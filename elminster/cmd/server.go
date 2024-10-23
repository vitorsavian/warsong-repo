/*
Copyright © 2024 Vitor Savian
*/
package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	environment "github.com/vitorsavian/warsong-repo/elminster/pkg/infra/env"
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

		env, _ := cmd.Flags().GetString("env")

		environment.SetEnv(env)
		config, err := tcp.ConfigServer()
		if err != nil {
			logrus.Fatal("teste")
		}

		config.Listen()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().String("env", "dev", "Environment that will run the app")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
