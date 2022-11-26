/*
Copyright © 2022 David Sosby <dsosby@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a days test",
	Long:  `Run a days test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Want me to run the latest day?")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
