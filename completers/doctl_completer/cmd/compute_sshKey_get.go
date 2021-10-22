package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKey_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about an SSH key on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKey_getCmd).Standalone()
	compute_sshKey_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `FingerPrint`")
	compute_sshKey_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_sshKeyCmd.AddCommand(compute_sshKey_getCmd)

	carapace.Gen(compute_sshKey_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "FingerPrint").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
