/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	pb "github.com/msvens/grpcgo/service"
)

var getFunc = func(cmd *cobra.Command, args []string) {
	id := strToId(args[0])
	c := ConnClient()
	req := pb.GetUserRequest{Id: id}
	resp, err := c.UC.GetUser(c.Ctx, &req)
	if err != nil {
		fmt.Println("could not get get user: ",err)
	} else {
		fmt.Println(resp.User)
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a user",
	Long: `Given a user id return the given user or none if the user did not exist`,
	Args: intArgFunc,
	Run: getFunc,
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
