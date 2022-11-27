/*
Copyright © 2022 David Sosby <dsosby@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
  "os/exec"
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
    // TODO Look at files to figure out which is latest -- can just sort I think
		day := 1

		if len(args) == 1 {
			argday, err := strconv.Atoi(args[0])

			if err != nil {
				log.Fatal(err)
			}

			day = argday
		}

    packageName := fmt.Sprintf("day%02d/day.go", day)
    runCmd := exec.Command("go", "run", packageName)
    runCmd.Stdout = os.Stdout

    isSecond, _ := cmd.Flags().GetBool("second")

    if isSecond {
      runCmd.Args = append(runCmd.Args, "--second")
    }

    if runErr := runCmd.Run(); runErr != nil {
      panic(runErr)
    }
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().BoolP("second", "s", false, "-s to run Problem2")
}
