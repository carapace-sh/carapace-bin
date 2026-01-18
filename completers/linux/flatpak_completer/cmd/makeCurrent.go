package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var makeCurrentCmd = &cobra.Command{
	Use:     "make-current",
	Short:   "Make branch of application current",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeCurrentCmd).Standalone()

	makeCurrentCmd.Flags().String("arch", "", "Arch to make current for")
	makeCurrentCmd.Flags().BoolP("help", "h", false, "Show help options")
	makeCurrentCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	makeCurrentCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	makeCurrentCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	makeCurrentCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	makeCurrentCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(makeCurrentCmd)

	carapace.Gen(makeCurrentCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "installation": carapace.ActionValues(),
	})
}
