package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emergeCmd = &cobra.Command{
	Use:     "emerge <name>",
	Aliases: []string{"em"},
	Short:   "build and install a package from the legacy pspec.xml source format",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emergeCmd).Standalone()

	emergeCmd.Flags().StringP("component", "c", "", "restrict the operation to the given component")
	emergeCmd.Flags().Bool("create-static", false, "create a static package")
	emergeCmd.Flags().Bool("ignore-action-errors", false, "ignore errors from actions")
	emergeCmd.Flags().Bool("ignore-build-no", false, "ignore build number errors")
	emergeCmd.Flags().Bool("ignore-check", false, "ignore checking metadata for a valid distribution specifier")
	emergeCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	emergeCmd.Flags().Bool("ignore-dependency", false, "do not validate build dependencies")
	emergeCmd.Flags().Bool("ignore-file-conflicts", false, "allow the package to install even if it conflicts with another package's files")
	emergeCmd.Flags().Bool("ignore-package-conflicts", false, "forcibly install a package even though it is marked as conflicting")
	emergeCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	emergeCmd.Flags().Bool("ignore-sandbox", false, "ignore build sandboxing")
	emergeCmd.Flags().StringP("output-dir", "O", "", "override the output directory for the resulting .eopkg")
	emergeCmd.Flags().StringP("package-format", "F", "", "override the eopkg internal format")
	emergeCmd.Flags().Bool("quiet", false, "reduce build output verbosity")
	emergeCmd.Flags().Bool("use-quilt", false, "use quilt to apply patches")

	rootCmd.AddCommand(emergeCmd)
}
