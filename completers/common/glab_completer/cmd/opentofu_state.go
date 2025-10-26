package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_stateCmd = &cobra.Command{
	Use:   "state <command> [flags]",
	Short: "Work with the OpenTofu or Terraform states.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_stateCmd).Standalone()

	opentofuCmd.AddCommand(opentofu_stateCmd)
}
