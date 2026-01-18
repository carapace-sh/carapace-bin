package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:     "ps",
	Short:   "Enumerate running sandboxes",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(psCmd).Standalone()

	psCmd.Flags().String("columns", "", "What information to show")
	psCmd.Flags().BoolP("help", "h", false, "Show help options")
	psCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	psCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(psCmd)

	carapace.Gen(psCmd).FlagCompletion(carapace.ActionMap{
		"columns": columns(flatpak.ActionProcessColums()),
	})
}
