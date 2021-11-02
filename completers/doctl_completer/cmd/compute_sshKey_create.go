package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKey_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new SSH key on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKey_createCmd).Standalone()
	compute_sshKey_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `FingerPrint`")
	compute_sshKey_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_sshKey_createCmd.Flags().String("public-key", "", "Key contents (required)")
	compute_sshKeyCmd.AddCommand(compute_sshKey_createCmd)

	carapace.Gen(compute_sshKey_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "FingerPrint").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
