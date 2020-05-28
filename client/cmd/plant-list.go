// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plants",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create the request (this can be inline below too)
		req := &plantpb.ListPlantsReq{}
		// Call ListPlants that returns a stream
		stream, err := plantClient.ListPlants(context.Background(), req)
		// Check for errors
		if err != nil {
			return err
		}
		// Start iterating
		for {
			// stream.Recv returns a pointer to a ListPlantRes at the current iteration
			res, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well use the generated getter to print the plant message
			fmt.Println(res.GetPlant())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
