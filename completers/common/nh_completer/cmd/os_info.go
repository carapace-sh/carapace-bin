package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "List available generations from profile path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_infoCmd).Standalone()

	os_infoCmd.Flags().String("fields", "", "Comma-delimited list of field(s) to display")
	os_infoCmd.Flags().StringP("profile", "P", "/nix/var/nix/profiles/system", "Path to Nix profiles directory")

	carapace.Gen(os_infoCmd).FlagCompletion(carapace.ActionMap{
		"fields":  carapace.ActionValues("id", "date", "nver", "kernel", "confRev", "spec", "size").UniqueList(","),
		"profile": carapace.ActionFiles(),
	})

	osCmd.AddCommand(os_infoCmd)
}
