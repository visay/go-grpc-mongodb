// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
)

// listCategoryCmd represents the list command
var listCategoryCmd = &cobra.Command{
	Use:   "list",
	Short: "List all categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create the request (this can be inline below too)
		req := &categorypb.ListCategoriesReq{}
		// Call ListCategories that returns a stream
		stream, err := categoryClient.ListCategories(context.Background(), req)
		// Check for errors
		if err != nil {
			return err
		}
		// Start iterating
		for {
			// stream.Recv returns a pointer to a ListCategoryRes at the current iteration
			res, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well use the generated getter to print the category message
			fmt.Println(res.GetCategory())
		}
		return nil
	},
}

func init() {
	categoryCmd.AddCommand(listCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
