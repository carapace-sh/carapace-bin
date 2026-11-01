package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var libcCmd = &cobra.Command{
	Use:   "libc",
	Short: "Display native libc paths file or validate one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(libcCmd).Standalone()

	rootCmd.AddCommand(libcCmd)

	libcCmd.Flags().String("color", "", "Enable or disable colored error messages")
	libcCmd.Flags().BoolP("help", "h", false, "Print help")
	libcCmd.Flags().Bool("includes", false, "Print the libc include directories for the target")
	libcCmd.Flags().String("target", "", "Target triple (e.g. x86_64-linux-gnu)")

	carapace.Gen(libcCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"target": zig.ActionTargets(),
	})

	carapace.Gen(libcCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
