// Copyright © 2020 Visay Keo <keo@visay.info>

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	plantpb "github.com/visay/go-grpc-mongodb/proto/plant"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Find a Plant by its ID",
	Long: `Find a plant by its mongoDB Unique identifier.

	If no plant is found for the ID, it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &plantpb.ReadPlantReq{
			Id: id,
		}
		res, err := plantClient.ReadPlant(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	readCmd.Flags().StringP("id", "i", "", "The id of the plant")
	readCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
