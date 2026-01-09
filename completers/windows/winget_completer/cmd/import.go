package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Installs all the packages in a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().Bool("accept-package-agreements", false, "Accept all license agreements for packages")
	importCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	importCmd.Flags().Bool("ignore-unavailable", false, "Ignore unavailable packages")
	importCmd.Flags().Bool("ignore-versions", false, "Ignore package versions in import file")
	importCmd.Flags().StringP("import-file", "i", "", "File describing the packages to install")
	importCmd.Flags().Bool("no-upgrade", false, "Skips upgrade if an installed version already exists")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"import-file": carapace.ActionFiles(),
	})
}
