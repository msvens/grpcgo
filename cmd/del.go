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
	"github.com/msvens/grpcgo/service"
)



var delFunc = func(cmd *cobra.Command, args []string) {
	id := strToId(args[0])
	c := ConnClient()
	req := service.DeleteUserRequest{Id: id}
	resp, err := c.UC.DeleteUser(c.Ctx, &req)
	if err != nil {
		fmt.Println("could not delete user: ",err)
	} else {
		fmt.Println("No of users deleted: ", resp.Count)
	}
}

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del ID",
	Short: "delete user",
	Long: `Delete a user from the database`,
	Args: intArgFunc,
	Run: delFunc,
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
