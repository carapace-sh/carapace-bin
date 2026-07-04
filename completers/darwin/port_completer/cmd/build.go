package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Compute checksums of distribution files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean files used for building a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Run configure process for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var destrootCmd = &cobra.Command{
	Use:   "destroot",
	Short: "Install a port to a temporary directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract distribution files for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch distribution files for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Apply required patches to a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(checksumCmd).Standalone()
	rootCmd.AddCommand(checksumCmd)

	carapace.Gen(cleanCmd).Standalone()
	cleanCmd.Flags().Bool("all", false, "Remove all build files (dist, archive, logs, work)")
	cleanCmd.Flags().Bool("archive", false, "Remove temporary archives")
	cleanCmd.Flags().Bool("dist", false, "Remove downloaded files")
	cleanCmd.Flags().Bool("logs", false, "Remove log files")
	cleanCmd.Flags().Bool("work", false, "Remove the work directory (default)")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(configureCmd).Standalone()
	rootCmd.AddCommand(configureCmd)

	carapace.Gen(destrootCmd).Standalone()
	rootCmd.AddCommand(destrootCmd)

	carapace.Gen(extractCmd).Standalone()
	rootCmd.AddCommand(extractCmd)

	carapace.Gen(fetchCmd).Standalone()
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(patchCmd).Standalone()
	rootCmd.AddCommand(patchCmd)

	carapace.Gen(testCmd).Standalone()
	rootCmd.AddCommand(testCmd)
}
