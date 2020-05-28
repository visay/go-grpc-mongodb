// Copyright © 2019 Visay Keo <keo@visay.info>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
)

// updateCmd represents the update command
var UpdateCategoryCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a Category by its ID",
	Long: `Update a category by its mongoDB Unique identifier.

	If no category is found for the ID, it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		name, err := cmd.Flags().GetString("name")

		// Create an UpdateCategoryRequest
		req := &categorypb.UpdateCategoryReq{
			Category: &categorypb.Category{
				Id:   id,
				Name: name,
			},
		}

		res, err := categoryClient.UpdateCategory(context.Background(), req)
		if err != nil {
			log.Fatalf("Error updating category:\n\n%s", err)
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	UpdateCategoryCmd.Flags().StringP("id", "i", "", "The id of the category")
	UpdateCategoryCmd.Flags().StringP("name", "n", "", "Name of the category")
	UpdateCategoryCmd.MarkFlagRequired("id")
	UpdateCategoryCmd.MarkFlagRequired("name")
	categoryCmd.AddCommand(UpdateCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
