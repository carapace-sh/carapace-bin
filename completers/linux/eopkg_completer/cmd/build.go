package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Build eopkg packages",
	Aliases: []string{"bi"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("fetch", false, "Break build after fetching the source archive")
	buildCmd.Flags().Bool("unpack", false, "Break build after unpacking the source archive")
	buildCmd.Flags().Bool("setup", false, "Break build after running the setup action")
	buildCmd.Flags().Bool("build", false, "Break build after running the build action")
	buildCmd.Flags().Bool("check", false, "Break build after running the check action")
	buildCmd.Flags().Bool("install", false, "Break build after running the install action")
	buildCmd.Flags().Bool("package", false, "Create the eopkg package")
	buildCmd.Flags().Bool("ignore-build-no", false, "Do not take build no into account")
	buildCmd.Flags().BoolP("quiet", "q", false, "Run quietly")
	buildCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	buildCmd.Flags().StringP("output-dir", "O", "", "Output directory for produced packages")
	buildCmd.Flags().Bool("ignore-action-errors", false, "Bypass errors from ActionsAPI")
	buildCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	buildCmd.Flags().Bool("ignore-check", false, "Skip check action")
	buildCmd.Flags().Bool("create-static", false, "Create a statically linked package")
	buildCmd.Flags().String("package-format", "", "Create the package using the given format")
	buildCmd.Flags().Bool("use-quilt", false, "Use quilt to apply patches")
	buildCmd.Flags().Bool("ignore-sandbox", false, "Do not use build sandbox")

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
