package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scopes_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List scopes at which user has assigned privileges.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scopes_lsCmd).Standalone()

	scopes_lsCmd.Flags().Bool("no-verbose", false, "Show table with details of per-scope privileges.")
	scopes_lsCmd.Flags().BoolP("verbose", "v", false, "Show table with details of per-scope privileges.")
	scopes_lsCmd.Flag("no-verbose").Hidden = true
	scopesCmd.AddCommand(scopes_lsCmd)
}
