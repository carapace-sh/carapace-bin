package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKey_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an SSH key's name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKey_updateCmd).Standalone()
	compute_sshKey_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `FingerPrint`")
	compute_sshKey_updateCmd.Flags().String("key-name", "", "Key name (required)")
	compute_sshKey_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_sshKeyCmd.AddCommand(compute_sshKey_updateCmd)

	carapace.Gen(compute_sshKey_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "FingerPrint").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
