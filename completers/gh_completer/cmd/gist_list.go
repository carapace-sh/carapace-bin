package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gist_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your gists",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_listCmd).Standalone()

	gist_listCmd.Flags().String("filter", "", "Filter gists using a regular `expression`")
	gist_listCmd.Flags().Bool("include-content", false, "Include gists' file content when filtering")
	gist_listCmd.Flags().StringP("limit", "L", "", "Maximum number of gists to fetch")
	gist_listCmd.Flags().Bool("public", false, "Show only public gists")
	gist_listCmd.Flags().Bool("secret", false, "Show only secret gists")
	gistCmd.AddCommand(gist_listCmd)
}
