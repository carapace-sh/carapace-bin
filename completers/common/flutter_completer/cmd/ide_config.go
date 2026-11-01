package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ide_configCmd = &cobra.Command{
	Use:   "ide-config",
	Short: "Configure the IDE for use in the Flutter tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ide_configCmd).Standalone()

	ide_configCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	ide_configCmd.Flags().Bool("no-overwrite", false, "When performing operations, overwrite existing files.")
	ide_configCmd.Flags().Bool("no-with-root-module", false, "Also create module that corresponds to the root of Flutter tree. This makes the entire Flutter tree browsable and searchable in IDE. Without this flag, only the child modules will be visible in IDE.")
	ide_configCmd.Flags().Bool("overwrite", false, "When performing operations, overwrite existing files.")
	ide_configCmd.Flags().Bool("update-templates", false, "Update the templates in the template directory from the current configuration files. This is the opposite of what ide-config usually does. Will search the flutter tree for *.iml files and copy any missing ones into the template directory. If \"--overwrite\" is also specified, it will update any out-of-date files, and remove any deleted files from the template directory.")
	ide_configCmd.Flags().Bool("with-root-module", false, "Also create module that corresponds to the root of Flutter tree. This makes the entire Flutter tree browsable and searchable in IDE. Without this flag, only the child modules will be visible in IDE.")
	rootCmd.AddCommand(ide_configCmd)
}
