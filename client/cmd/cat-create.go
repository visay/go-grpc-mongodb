// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
)

// createCategoryCmd represents the create command
var createCategoryCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category",
	Long: `Create a new category on the server through gRPC.

	A category requires a Name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		category := &categorypb.Category{
			Name: name,
		}
		res, err := categoryClient.CreateCategory(
			context.TODO(),
			&categorypb.CreateCategoryReq{
				Category: category,
			},
		)
		if err != nil {
			log.Fatalf("Error creating category:\n\n%s", err)
		}
		fmt.Printf("Category created: %s\n", res.Category.Id)
		return nil
	},
}

func init() {
	createCategoryCmd.Flags().StringP("name", "n", "", "Add a name")
	createCategoryCmd.MarkFlagRequired("name")
	categoryCmd.AddCommand(createCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
