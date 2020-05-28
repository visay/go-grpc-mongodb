// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
)

// deleteCmd represents the read command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Plant by its ID",
	Long: `Delete a plant by its mongoDB Unique identifier.

	If no plant is found for the ID, it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &plantpb.DeletePlantReq{
			Id: id,
		}
		// We only return true upon success for other cases an error is thrown
		// We can thus omit the response variable for now and just print something to console
		_, err = plantClient.DeletePlant(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Printf("Succesfully deleted the plant with id %s\n", id)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("id", "i", "", "The id of the plant")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
