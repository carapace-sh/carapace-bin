package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the tsh client and Proxy server versions for the current context.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("client", false, "Show the client version only (no server required).")
	versionCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	versionCmd.Flags().Bool("no-client", false, "Show the client version only (no server required).")
	versionCmd.Flag("no-client").Hidden = true
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
