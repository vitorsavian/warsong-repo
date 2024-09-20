/*
Copyright © 2024 Vitor Savian
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/tcp"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server for tcp connections",
	Long: `It will create a tcp server with the possibility to also
  create a character by the api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initianting tcp server")

		tcp.Server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
