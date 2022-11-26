/*
Copyright © 2022 David Sosby <dsosby@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a days test",
	Long:  `Run a days test`,
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return nil
		case 1:
			_, err := strconv.Atoi(args[0])
			return err
		default:
			return errors.New("Too many day args")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		day := 1

		if len(args) == 1 {
			argday, err := strconv.Atoi(args[0])

			if err != nil {
				log.Fatal(err)
			}

			day = argday
		}

		fmt.Println("Running day", day)
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
