package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listSessionsCmd = &cobra.Command{
	Use:     "list-sessions",
	Short:   "List active sessions",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSessionsCmd).Standalone()

	listSessionsCmd.Flags().BoolP("help", "h", false, "Print help")
	listSessionsCmd.Flags().BoolP("no-formatting", "n", false, "Do not add colors and formatting to the list (useful for parsing)")
	listSessionsCmd.Flags().BoolP("reverse", "r", false, "List the sessions in reverse order (default is ascending order)")
	listSessionsCmd.Flags().BoolP("short", "s", false, "Print just the session name")
	rootCmd.AddCommand(listSessionsCmd)
}
