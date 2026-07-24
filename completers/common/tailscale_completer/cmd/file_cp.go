package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy file(s) to a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_cpCmd).Standalone()

	file_cpCmd.Flags().String("name", "", "alternate filename to use, especially useful when file is stdin (-)")
	file_cpCmd.Flags().Bool("targets", false, "list possible file cp targets")
	file_cpCmd.Flags().String("update-interval", "", "how often to repaint the progress line")
	file_cpCmd.Flags().Bool("verbose", false, "verbose output")
	fileCmd.AddCommand(file_cpCmd)

	carapace.Gen(file_cpCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
