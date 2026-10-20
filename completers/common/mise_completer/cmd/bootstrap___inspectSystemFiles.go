package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap___inspectSystemFilesCmd = &cobra.Command{
	Use:    "__inspect-system-files",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap___inspectSystemFilesCmd).Standalone()

	bootstrap___inspectSystemFilesCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap___inspectSystemFilesCmd)
}
