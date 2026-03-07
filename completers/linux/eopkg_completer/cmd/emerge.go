package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emergeCmd = &cobra.Command{
	Use:     "emerge",
	Short:   "Build and install packages from source",
	Aliases: []string{"em"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emergeCmd).Standalone()

	emergeCmd.Flags().Bool("ignore-build-no", false, "Do not take build no into account")
	emergeCmd.Flags().BoolP("quiet", "q", false, "Run quietly")
	emergeCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	emergeCmd.Flags().StringP("output-dir", "O", "", "Output directory for produced packages")
	emergeCmd.Flags().Bool("ignore-action-errors", false, "Bypass errors from ActionsAPI")
	emergeCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	emergeCmd.Flags().Bool("ignore-check", false, "Skip check action")
	emergeCmd.Flags().Bool("create-static", false, "Create a statically linked package")
	emergeCmd.Flags().String("package-format", "", "Create the package using the given format")
	emergeCmd.Flags().Bool("use-quilt", false, "Use quilt to apply patches")
	emergeCmd.Flags().Bool("ignore-sandbox", false, "Do not use build sandbox")
	emergeCmd.Flags().StringP("component", "c", "", "Emerge component's packages")
	emergeCmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	emergeCmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	emergeCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")

	rootCmd.AddCommand(emergeCmd)
}
