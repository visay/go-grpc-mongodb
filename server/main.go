// Copyright © 2020 Visay Keo <keo@visay.info>

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"reflect"

	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// -----------------------------------------------------------------------------

type PlantServiceServer struct {
}

func (s *PlantServiceServer) ReadPlant(ctx context.Context, req *plantpb.ReadPlantReq) (*plantpb.ReadPlantRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := plantdb.FindOne(ctx, bson.M{"_id": oid})

	// Create an empty PlantItem to write our decode result to
	data := PlantItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find plant with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadPlantRes type
	response := &plantpb.ReadPlantRes{
		Plant: &plantpb.Plant{
			Id:         oid.Hex(),
			Name:       data.Name,
			CategoryId: data.CategoryID,
			Desc:       data.Desc,
		},
	}
	return response, nil
}

func (s *PlantServiceServer) CreatePlant(ctx context.Context, req *plantpb.CreatePlantReq) (*plantpb.CreatePlantRes, error) {
	// Get the protobuf plant type from the protobuf request type
	// Essentially doing req.Plant to access the struct with a nil check
	plant := req.GetPlant()

	// convert string id (from proto) to mongoDB ObjectId
	oidCat, err := primitive.ObjectIDFromHex(plant.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	resultCat := categorydb.FindOne(ctx, bson.M{"_id": oidCat})
	// Create an empty CategoryItem to write our decode result to
	dataCat := CategoryItem{}
	// decode and write to data
	if err := resultCat.Decode(&dataCat); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find category with Object Id %s: %v", plant.GetCategoryId(), err))
	}

	// Now we have to convert this into a PlantItem type to convert into BSON
	data := PlantItem{
		// ID:       primitive.NilObjectID,
		Name:       plant.GetName(),
		CategoryID: plant.GetCategoryId(),
		Desc:       plant.GetDesc(),
	}

	// Insert the data into the database
	// *InsertOneResult contains the oid
	result, err := plantdb.InsertOne(mongoCtx, data)
	// check error
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to plant
	oid := result.InsertedID.(primitive.ObjectID)
	plant.Id = oid.Hex()
	// return the plant in a CreatePlantRes type
	return &plantpb.CreatePlantRes{Plant: plant}, nil
}

func (s *PlantServiceServer) UpdatePlant(ctx context.Context, req *plantpb.UpdatePlantReq) (*plantpb.UpdatePlantRes, error) {
	// Get the plant data from the request
	plant := req.GetPlant()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(plant.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied plant id to a MongoDB ObjectId: %v", err),
		)
	}

	// convert string id (from proto) to mongoDB ObjectId
	oidCat, err := primitive.ObjectIDFromHex(plant.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	resultCat := categorydb.FindOne(ctx, bson.M{"_id": oidCat})
	// Create an empty CategoryItem to write our decode result to
	dataCat := CategoryItem{}
	// decode and write to data
	if err := resultCat.Decode(&dataCat); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find category with Object Id %s: %v", plant.GetCategoryId(), err))
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"name":        plant.GetName(),
		"category_id": plant.GetCategoryId(),
		"desc":        plant.GetDesc(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := plantdb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := PlantItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find plant with supplied ID: %v", err),
		)
	}
	return &plantpb.UpdatePlantRes{
		Plant: &plantpb.Plant{
			Id:         decoded.ID.Hex(),
			Name:       decoded.Name,
			CategoryId: decoded.CategoryID,
			Desc:       decoded.Desc,
		},
	}, nil
}

func (s *PlantServiceServer) DeletePlant(ctx context.Context, req *plantpb.DeletePlantReq) (*plantpb.DeletePlantRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = plantdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete plant with id %s: %v", req.GetId(), err))
	}
	return &plantpb.DeletePlantRes{
		Success: true,
	}, nil
}

func (s *PlantServiceServer) ListPlants(req *plantpb.ListPlantsReq, stream plantpb.PlantService_ListPlantsServer) error {
	// Initiate a PlantItem type to write decoded data to
	data := &PlantItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := plantdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send plant over stream
		stream.Send(&plantpb.ListPlantsRes{
			Plant: &plantpb.Plant{
				Id:         data.ID.Hex(),
				Name:       data.Name,
				CategoryId: data.CategoryID,
				Desc:       data.Desc,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type PlantItem struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	CategoryID string             `bson:"category_id"`
	Desc       string             `bson:"desc"`
}

// -----------------------------------------------------------------------------

type CategoryServiceServer struct {
}

func (s *CategoryServiceServer) ReadCategory(ctx context.Context, req *categorypb.ReadCategoryReq) (*categorypb.ReadCategoryRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := categorydb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty CategoryItem to write our decode result to
	data := CategoryItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find category with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadCategoryRes type
	response := &categorypb.ReadCategoryRes{
		Category: &categorypb.Category{
			Id:   oid.Hex(),
			Name: data.Name,
		},
	}
	return response, nil
}

func (s *CategoryServiceServer) CreateCategory(ctx context.Context, req *categorypb.CreateCategoryReq) (*categorypb.CreateCategoryRes, error) {
	// Get the protobuf category type from the protobuf request type
	// Essentially doing req.Category to access the struct with a nil check
	category := req.GetCategory()
	// Now we have to convert this into a CategoryItem type to convert into BSON
	data := CategoryItem{
		// ID:       primitive.NilObjectID,
		Name: category.GetName(),
	}

	// Insert the data into the database
	// *InsertOneResult contains the oid
	result, err := categorydb.InsertOne(mongoCtx, data)
	// check error
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to category
	oid := result.InsertedID.(primitive.ObjectID)
	category.Id = oid.Hex()
	// return the category in a CreateCategoryRes type
	return &categorypb.CreateCategoryRes{Category: category}, nil
}

func (s *CategoryServiceServer) UpdateCategory(ctx context.Context, req *categorypb.UpdateCategoryReq) (*categorypb.UpdateCategoryRes, error) {
	// Get the category data from the request
	category := req.GetCategory()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(category.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied category id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"name": category.GetName(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := categorydb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := CategoryItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find category with supplied ID: %v", err),
		)
	}
	return &categorypb.UpdateCategoryRes{
		Category: &categorypb.Category{
			Id:   decoded.ID.Hex(),
			Name: decoded.Name,
		},
	}, nil
}

func (s *CategoryServiceServer) DeleteCategory(ctx context.Context, req *categorypb.DeleteCategoryReq) (*categorypb.DeleteCategoryRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// TODO: check if category is being used

	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = categorydb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete category with id %s: %v", req.GetId(), err))
	}
	return &categorypb.DeleteCategoryRes{
		Success: true,
	}, nil
}

func (s *CategoryServiceServer) ListCategories(req *categorypb.ListCategoriesReq, stream categorypb.CategoryService_ListCategoriesServer) error {
	// Initiate a CategoryItem type to write decoded data to
	data := &CategoryItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := categorydb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send category over stream
		stream.Send(&categorypb.ListCategoriesRes{
			Category: &categorypb.Category{
				Id:   data.ID.Hex(),
				Name: data.Name,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type CategoryItem struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

// -----------------------------------------------------------------------------

var db *mongo.Client
var plantdb *mongo.Collection
var categorydb *mongo.Collection
var mongoCtx context.Context

func main() {

	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// just the filename & line number:
	// log.SetFlags(log.Lshortfile)
	// Or add timestamps and pipe file name and line number to it:
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port :50051...")

	// 50051 is the default port for gRPC
	// Ideally we'd use 0.0.0.0 instead of localhost as well
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// slice of gRPC options
	// Here we can configure things like TLS
	opts := []grpc.ServerOption{}
	// var s *grpc.Server
	s := grpc.NewServer(opts...)
	// var plantsrv *PlantServiceServer
	plantsrv := &PlantServiceServer{}
	// var categorysrv *CategoryServiceServer
	categorysrv := &CategoryServiceServer{}

	plantpb.RegisterPlantServiceServer(s, plantsrv)
	categorypb.RegisterCategoryServiceServer(s, categorysrv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	plantdb = db.Database("plant_db").Collection("plant_collection")
	categorydb = db.Database("plant_db").Collection("category_collection")

	// Declare an index model object to pass to CreateOne()
	// db.members.createIndex( { "SOME_FIELD": 1 }, { unique: true } )
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1, // index in ascending order
		}, Options: options.Index().SetUnique(true),
	}

	// Create an Index using the CreateOne() method
	plantIndex, err := plantdb.Indexes().CreateOne(mongoCtx, indexModel)

	// Check if the CreateOne() method returned any errors
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		// API call returns string of the index name
		fmt.Println("Created plant index:", plantIndex)
		fmt.Println("Field type:", reflect.TypeOf(plantIndex), "\n")
	}

	// Create an Index using the CreateOne() method
	categoryIndex, err := categorydb.Indexes().CreateOne(mongoCtx, indexModel)

	// Check if the CreateOne() method returned any errors
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		// API call returns string of the index name
		fmt.Println("Created category index:", categoryIndex)
		fmt.Println("Field type:", reflect.TypeOf(categoryIndex), "\n")
	}

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

	// Bad way to stop the server
	// if err := s.Serve(listener); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
	// Right way to stop the server using a SHUTDOWN HOOK

	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed
	// And our main routine keeps running
	// If the main routine were to shutdown so would the child routine that is Serving the server
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
