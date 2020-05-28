// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
)

// deleteCategoryCmd represents the delete command
var deleteCategoryCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Category by its ID",
	Long: `Delete a category by its mongoDB Unique identifier.

	If no category is found for the ID, it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &categorypb.DeleteCategoryReq{
			Id: id,
		}
		// We only return true upon success for other cases an error is thrown
		// We can thus omit the response variable for now and just print something to console
		_, err = categoryClient.DeleteCategory(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Printf("Succesfully deleted the category with id %s\n", id)
		return nil
	},
}

func init() {
	deleteCategoryCmd.Flags().StringP("id", "i", "", "The id of the category")
	deleteCategoryCmd.MarkFlagRequired("id")
	categoryCmd.AddCommand(deleteCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
