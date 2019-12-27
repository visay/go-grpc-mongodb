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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new plant",
	Long: `Create a new plant on the server through gRPC.

	A plant requires a Name, Group and Description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		group, err := cmd.Flags().GetString("group")
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			return err
		}
		plant := &plantpb.Plant{
			Name:  name,
			Group: group,
			Desc:  desc,
		}
		res, err := client.CreatePlant(
			context.TODO(),
			&plantpb.CreatePlantReq{
				Plant: plant,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Plant created: %s\n", res.Plant.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Add a name")
	createCmd.Flags().StringP("group", "g", "", "A group for the plant")
	createCmd.Flags().StringP("desc", "d", "", "The description for the plant")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("group")
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
