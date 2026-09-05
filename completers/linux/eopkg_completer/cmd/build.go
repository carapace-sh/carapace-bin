package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build <path to pspec.xml>",
	Aliases: []string{"bi"},
	Short:   "build a package from the given pspec.xml source specification",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("build", false, "run the build stage")
	buildCmd.Flags().Bool("check", false, "run the check stage")
	buildCmd.Flags().Bool("create-static", false, "create a static package")
	buildCmd.Flags().Bool("fetch", false, "run the fetch stage")
	buildCmd.Flags().Bool("ignore-action-errors", false, "ignore errors from actions")
	buildCmd.Flags().Bool("ignore-build-no", false, "ignore build number errors")
	buildCmd.Flags().Bool("ignore-check", false, "ignore checking metadata for a valid distribution specifier")
	buildCmd.Flags().Bool("ignore-dependency", false, "do not validate build dependencies")
	buildCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	buildCmd.Flags().Bool("ignore-sandbox", false, "ignore build sandboxing")
	buildCmd.Flags().Bool("install", false, "run the install stage")
	buildCmd.Flags().StringP("output-dir", "O", "", "override the output directory for the resulting .eopkg")
	buildCmd.Flags().Bool("package", false, "run the package stage")
	buildCmd.Flags().StringP("package-format", "F", "", "override the eopkg internal format")
	buildCmd.Flags().Bool("quiet", false, "reduce build output verbosity")
	buildCmd.Flags().Bool("setup", false, "run the setup stage")
	buildCmd.Flags().Bool("unpack", false, "run the unpack stage")
	buildCmd.Flags().Bool("use-quilt", false, "use quilt to apply patches")

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.ActionFiles("xml"),
	)
}
