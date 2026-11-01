package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duo_cli_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the GitLab Duo CLI binary to the latest compatible version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duo_cli_updateCmd).Standalone()

	duo_cli_updateCmd.Flags().BoolP("yes", "y", false, "Skip the download prompt when the binary is not installed.")
	duo_cliCmd.AddCommand(duo_cli_updateCmd)
}
