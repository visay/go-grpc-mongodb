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
	plantpb "github.com/visay/go-grpc-mongodb/proto"
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
		group, err := cmd.Flags().GetString("group")
		desc, err := cmd.Flags().GetString("desc")

		// Create an UpdatePlantRequest
		req := &plantpb.UpdatePlantReq{
			Plant: &plantpb.Plant{
				Id:    id,
				Name:  name,
				Group: group,
				Desc:  desc,
			},
		}

		res, err := client.UpdatePlant(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "The id of the plant")
	updateCmd.Flags().StringP("name", "n", "", "Name of the plant")
	updateCmd.Flags().StringP("group", "g", "", "A group for the plant")
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
