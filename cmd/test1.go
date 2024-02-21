/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/graytalent/test1"
)

// test1Cmd represents the test1 command
var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var res string
		var err error

		res, err = test1.TextInput("h1", "h2", "h3")
		if err != nil {
			panic(err)
		}
		println(res)
		res, err = test1.TextInput("h1", "h2", "h3")
		if err != nil {
			panic(err)
		}
		println(res)
		res, err = test1.TextInput("h1", "h2", "h3")
		if err != nil {
			panic(err)
		}
		println(res)
		res, err = test1.TextInput("h1", "h2", "h3")
		if err != nil {
			panic(err)
		}
		println(res)
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
