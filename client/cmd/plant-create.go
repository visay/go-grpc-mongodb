// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new plant",
	Long: `Create a new plant on the server through gRPC.

	A plant requires a Name, Category and Description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		category_id, err := cmd.Flags().GetString("category_id")
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			return err
		}
		plant := &plantpb.Plant{
			Name:       name,
			CategoryId: category_id,
			Desc:       desc,
		}
		res, err := plantClient.CreatePlant(
			context.TODO(),
			&plantpb.CreatePlantReq{
				Plant: plant,
			},
		)
		if err != nil {
			log.Fatalf("Error creating plant:\n\n%s", err)
		}
		fmt.Printf("Plant created: %s\n", res.Plant.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Add a name")
	createCmd.Flags().StringP("category_id", "c", "", "The category for the plant")
	createCmd.Flags().StringP("desc", "d", "", "The description for the plant")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("category_id")
	createCmd.MarkFlagRequired("desc")
	rootCmd.AddCommand(createCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
