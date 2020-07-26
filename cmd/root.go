package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "karto",
	Short: "work with helm charts",
	Long:  `package helm charts with proper semantic versioning and push them to a repository.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Error().Msg("An action was not provided")
			return fmt.Errorf("An action is required")
		}
		return nil
	},
}

func init() {
	// flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display debug messages")
}

// Execute combines all of the available command functions
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msgf("Error during execution: %v", err)
	}
}
