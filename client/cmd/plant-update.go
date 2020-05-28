// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a Plant by its ID",
	Long: `Update a plant by its mongoDB Unique identifier.

	If no plant is found for the ID, it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		name, err := cmd.Flags().GetString("name")
		category_id, err := cmd.Flags().GetString("category_id")
		desc, err := cmd.Flags().GetString("desc")

		// Create an UpdatePlantRequest
		req := &plantpb.UpdatePlantReq{
			Plant: &plantpb.Plant{
				Id:         id,
				Name:       name,
				CategoryId: category_id,
				Desc:       desc,
			},
		}

		res, err := plantClient.UpdatePlant(context.Background(), req)
		if err != nil {
			log.Fatalf("Error updating plant:\n\n%s", err)
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "The id of the plant")
	updateCmd.Flags().StringP("name", "n", "", "Name of the plant")
	updateCmd.Flags().StringP("category_id", "c", "", "The category for the plant")
	updateCmd.Flags().StringP("desc", "d", "", "The description for the plant")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
