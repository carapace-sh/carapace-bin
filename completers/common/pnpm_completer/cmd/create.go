package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a project from a `create-*` starter kit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringSlice("allow-build", nil, "Package names allowed to run lifecycle (build) scripts during the install. May be repeated")
	createCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-tagged optional dependencies the install should keep. Repeat or comma-separate for multiple")
	createCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	createCmd.Flags().StringSlice("libc", nil, "libc families (`glibc`, `musl`) whose platform-tagged optional dependencies the install should keep")
	createCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-tagged optional dependencies the install should keep")
	createCmd.Flags().BoolP("shell-mode", "c", false, "Run the command inside of a shell. Uses `/bin/sh` on UNIX and `cmd.exe` on Windows")
	rootCmd.AddCommand(createCmd)
}
