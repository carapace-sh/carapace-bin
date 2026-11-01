package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugchangedfilesCmd = &cobra.Command{
	Use:    "debugchangedfiles",
	Short:  "REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugchangedfilesCmd).Standalone()

	debugchangedfilesCmd.Flags().Bool("compute", false, "compute information instead of reading it from storage")
	rootCmd.AddCommand(debugchangedfilesCmd)
}
