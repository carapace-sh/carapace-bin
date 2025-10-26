package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_initCmd = &cobra.Command{
	Use:   "init <state> [flags]",
	Short: "Initialize OpenTofu or Terraform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_initCmd).Standalone()

	opentofu_initCmd.Flags().StringP("binary", "b", "", "Name or path of the OpenTofu or Terraform binary to use for the initialization.")
	opentofu_initCmd.Flags().StringP("directory", "d", "", "Directory of the OpenTofu or Terraform project to initialize.")
	opentofuCmd.AddCommand(opentofu_initCmd)

	carapace.Gen(opentofu_initCmd).FlagCompletion(carapace.ActionMap{
		"binary": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"directory": carapace.ActionDirectories(),
	})
}
