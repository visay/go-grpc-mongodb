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

	"github.com/spf13/cobra"
	categorypb "github.com/visay/go-grpc-mongodb/proto/category"
)

// deleteCmd represents the read command
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
