// Package cmd /*
package cmd

import (
	"com.tim.go/oclogin/core"
	"fmt"
	"github.com/spf13/cobra"
)

// list represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available openshift environments in your config",
	Long: `List available openshift environments in your config. For example:

oclogin list.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.List()
	},
}

// list represents the list command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login one openshift environments in your config",
	Long: `Login one available openshift environments in your config. For example:

oclogin login.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Login()
	},
}

// list represents the list command
var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Switch openshift config context by id",
	Long: `Switch openshift config context by id. For example:

oclogin context <id>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			core.UpdateContext(args[0])
		} else {
			fmt.Println("Only id parameter is allowed, e.g. oclogin context <id>")
		}
	},
}

// init represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the script environment, and create folder with an sample config file",
	Long: `Init the script environment, and create folder with an sample config file. For example:

oclogin init.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Init()
	},
}


func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(contextCmd)
	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
