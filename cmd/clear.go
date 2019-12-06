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

var clearFunc = func(cmd *cobra.Command, args []string) {
	c := ConnClient()
	resp, err := c.UC.DeleteAll(c.Ctx, &service.DeleteAllRequest{})
	if err != nil {
		fmt.Println("Could not delete all users: ", err)
	} else {
		fmt.Println("Number of users deleted: ", resp.Count)
	}
}

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove all users",
	Long: `Delete all users from the database and report how many users deleted`,
	Args: cobra.ExactArgs(0),
	Run: clearFunc,
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
