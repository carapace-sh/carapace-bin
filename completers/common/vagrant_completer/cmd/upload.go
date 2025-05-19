package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload to machine via communicator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().BoolP("compress", "c", false, "Use gzip compression for upload")
	uploadCmd.Flags().StringP("compression-type", "C", "", "Type of compression to use (tgz, zip)")
	uploadCmd.Flags().BoolP("temporary", "t", false, "Upload source to temporary directory")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).FlagCompletion(carapace.ActionMap{
		"compression-type": carapace.ActionValues("tgz", "zip"),
	})

	carapace.Gen(uploadCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionValues(),
		vagrant.ActionMachines(),
	)
}
