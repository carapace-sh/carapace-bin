package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var dlxCmd = &cobra.Command{
	Use:   "dlx",
	Short: "Run a package in a temporary environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlxCmd).Standalone()

	dlxCmd.Flags().StringSlice("allow-build", nil, "Package names allowed to run lifecycle (build) scripts during the dlx install. May be repeated")
	dlxCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-tagged optional dependencies the dlx install should keep. Repeat or comma-separate for multiple")
	dlxCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	dlxCmd.Flags().StringSlice("libc", nil, "libc families (`glibc`, `musl`) whose platform-tagged optional dependencies the dlx install should keep")
	dlxCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-tagged optional dependencies the dlx install should keep")
	dlxCmd.Flags().StringSlice("package", nil, "The package to install before running the command. May be repeated. When omitted, the command name is the package")
	dlxCmd.Flags().BoolP("shell-mode", "c", false, "Run the command inside of a shell. Uses `/bin/sh` on UNIX and `cmd.exe` on Windows")
	rootCmd.AddCommand(dlxCmd)

	carapace.Gen(dlxCmd).PositionalCompletion(
		npm.ActionPackageSearch(""),
	)

	carapace.Gen(dlxCmd).FlagCompletion(carapace.ActionMap{
		"cpu":  carapace.ActionValues("arm", "arm64", "ia32", "loong64", "mips", "mipsel", "ppc64", "riscv64", "s390", "s390x", "x64"),
		"libc": carapace.ActionValues("glibc", "musl"),
		"os":   carapace.ActionValues("aix", "android", "darwin", "freebsd", "linux", "openbsd", "sunos", "win32"),
	})
}
