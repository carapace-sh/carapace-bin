package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [OPTION…]",
	Short:   "List installed apps and/or runtimes",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("all", "a", false, "List all refs (including locale/debug)")
	listCmd.Flags().Bool("app", false, "List installed applications")
	listCmd.Flags().String("app-runtime", "", "List all applications using RUNTIME")
	listCmd.Flags().String("arch", "", "Arch to show")
	listCmd.Flags().String("columns", "", "What information to show")
	listCmd.Flags().BoolP("help", "h", false, "Show help options")
	listCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	listCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	listCmd.Flags().Bool("runtime", false, "List installed runtimes")
	listCmd.Flags().BoolP("show-details", "d", false, "Show extra information")
	listCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	listCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	listCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(listCmd)

	// TODO flags
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		// "app-runtime":  carapace.ActionValues(),
		"arch":    flatpak.ActionArches(),
		"columns": columns(flatpak.ActionApplicationColumns()),
		// "installation": carapace.ActionValues(),
	})
}
