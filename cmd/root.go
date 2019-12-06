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
  "context"
  "fmt"
  "google.golang.org/grpc"
  "os"
  "github.com/spf13/cobra"
  "strconv"
  "time"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
  pb "github.com/msvens/grpcgo/service"

)


var cfgFile string
var address string

type GrpcClient struct {
  Conn *grpc.ClientConn
  Ctx context.Context
  Cancel context.CancelFunc
  UC pb.UserServiceClient
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "grpcgo",
  Short: "Test application for a grpc client",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

func (c *GrpcClient) Close() {
  c.Conn.Close()
  c.Cancel()

}

func ConnClient() *GrpcClient {
  conn, err :=  grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
  c := pb.NewUserServiceClient(conn)
  return &GrpcClient{Conn:conn,Ctx:ctx, Cancel:cancel, UC:c}
}

var intArgFunc = func(com *cobra.Command, args []string) error {
  _, err := strconv.ParseInt(args[0], 0, 64)
  if len(args) != 1 {
    return fmt.Errorf("The command requires 1 argument")
  }
  if err != nil {
    return fmt.Errorf("Could not parse user id: %s", args[0])
  }
  return nil
}

func strToId(arg string) int64 {
  id,_ := strconv.ParseInt(arg, 0, 64)
  return id
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grpcgo.yaml)")
  rootCmd.PersistentFlags().StringVar(&address, "address", "localhost:50051", "address to grpc server")
  rootCmd.PersistentFlags().Int64("name", 0, "some")

  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".grpcgo" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".grpcgo")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

