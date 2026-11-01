package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List sessions",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().Bool("json", false, "output sessions as JSON")
	lsCmd.Flags().BoolP("quiet", "q", false, "print IDs only, one per line")
	lsCmd.Flags().Bool("short", true, "show short IDs (--short=false shows full IDs)")

	rootCmd.AddCommand(lsCmd)
}
