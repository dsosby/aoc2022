/*
Copyright © 2022 David Sosby <dsosby@gmail.com>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/dsosby/aoc2022/pkg"
	"github.com/spf13/cobra"
)

// newdayCmd represents the newday command
var newdayCmd = &cobra.Command{
	Use:   "newday",
	Short: "Initializes a new day package",
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    day, err := strconv.Atoi(args[0])

    if err != nil {
      log.Fatal(err)
    }

    util.CreateDayPackage(day)
	},
}

func init() {
	rootCmd.AddCommand(newdayCmd)
}
