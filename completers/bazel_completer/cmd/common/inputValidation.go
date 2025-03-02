package common

import "github.com/spf13/cobra"

func AddInputValidationFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("check_bzl_visibility", false, "If disabled, .bzl load visibility errors are demoted to warnings")
	cmd.Flags().Bool("no-check_bzl_visibility", false, "If disabled, .bzl load visibility errors are demoted to warnings")
}
