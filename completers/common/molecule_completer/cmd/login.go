package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [flags]",
	Short: "Login to one instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd)

	loginCmd.Flags().Bool("help", false, "Show help message and exit")
	loginCmd.Flags().StringP("host", "h", "", "Host to access")
	loginCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"host":          net.ActionHosts(),
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(loginCmd)
}
