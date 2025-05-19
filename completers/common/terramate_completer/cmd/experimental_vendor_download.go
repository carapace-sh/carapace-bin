package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_vendor_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads a Terraform module and stores it on the project vendor dir",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_vendor_downloadCmd).Standalone()

	experimental_vendor_downloadCmd.Flags().StringP("dir", "d", "", "dir to vendor downloaded project")
	experimental_vendorCmd.AddCommand(experimental_vendor_downloadCmd)

	carapace.Gen(experimental_vendor_downloadCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
