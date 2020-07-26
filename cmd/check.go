package cmd

import (
	"fmt"

	"github.com/particledecay/kartographer/pkg/helm"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate a chart",
	Long:  `Run a basic lint for a chart directory and validate version`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Error("You must supply a directory to check")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		basedir := args[1]
		linter := helm.Validate(basedir)

		for msg := range linter.Messages {
			fmt.Printf(msg)
		}
	},
}
