package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_export_locationCmd = &cobra.Command{
	Use:   "location",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_export_locationCmd).Standalone()

	share_exportCmd.AddCommand(share_export_locationCmd)
}
