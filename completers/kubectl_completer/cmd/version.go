package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the client and server version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("client", false, "If true, shows client version only (no server required).")
	versionCmd.Flags().StringP("output", "o", "", "One of 'yaml' or 'json'.")

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("yaml", "json"),
	})
}
