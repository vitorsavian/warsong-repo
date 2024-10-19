/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/grpc/proto"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3432))
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}

		fmt.Printf("server listening at %v\n", lis.Addr())

		var opts []gRPC.ServerOption

		s := gRPC.NewServer(opts...)

		proto.RegisterCharacterServer(s, grpc.NewServer())

		reflection.Register(s)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
