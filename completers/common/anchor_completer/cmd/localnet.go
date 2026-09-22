package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var localnetCmd = &cobra.Command{
	Use:   "localnet",
	Short: "Localnet commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localnetCmd).Standalone()

	localnetCmd.Flags().StringSliceP("env", "e", nil, "Environment variables to pass into the docker container")
	localnetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	localnetCmd.Flags().Bool("ignore-keys", false, "Skip checking for program ID mismatch between keypair and declare_id")
	localnetCmd.Flags().Bool("skip-build", false, "Flag to skip building the program in the workspace, use this to save time when running test and the program code is not altered")
	localnetCmd.Flags().Bool("skip-deploy", false, "Use this flag if you want to run tests against previously deployed programs")
	localnetCmd.Flags().Bool("skip-lint", false, "True if the build should not fail even if there are no \"CHECK\" comments where normally required")
	localnetCmd.Flags().String("validator", "surfpool", "Validator type to use for local testing")
	rootCmd.AddCommand(localnetCmd)
}
