package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pkgidCmd = &cobra.Command{
	Use:   "pkgid",
	Short: "Print a fully qualified package specification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgidCmd).Standalone()

	pkgidCmd.Flags().BoolP("help", "h", false, "Print help")
	pkgidCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	pkgidCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	pkgidCmd.Flags().StringP("package", "p", "", "Argument to get the package ID specifier for")
	rootCmd.AddCommand(pkgidCmd)

	carapace.Gen(pkgidCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(pkgidCmd, true),
	})
}
