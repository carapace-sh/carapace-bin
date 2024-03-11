package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var mfa_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Get a list of registered MFA devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mfa_lsCmd).Standalone()

	mfa_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	mfa_lsCmd.Flags().Bool("no-verbose", false, "Print more information about MFA devices")
	mfa_lsCmd.Flags().BoolP("verbose", "v", false, "Print more information about MFA devices")
	mfa_lsCmd.Flag("no-verbose").Hidden = true
	mfaCmd.AddCommand(mfa_lsCmd)

	carapace.Gen(mfa_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
