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

var addFunc = func(cmd *cobra.Command, args []string) {
	c := ConnClient()
	defer c.Close()
	req := pb.AddUserRequest{Email:args[0],Name:args[1]}
	resp, err := c.UC.AddUser(c.Ctx, &req)
	if err != nil {
		fmt.Println("could not add user: ", err)
	} else {
		fmt.Println("user added with id: ", resp.Id)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Args: cobra.ExactArgs(2),
	Use:   "add EMAIL NAME",
	Short: "add user",
	Long: `Adds a user to the database. If the email is already in use an error will be thrown`,
	Run: addFunc,
}

func init() {
	rootCmd.AddCommand(addCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
