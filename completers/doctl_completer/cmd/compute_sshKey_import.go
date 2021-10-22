package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKey_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import an SSH key from your computer to your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKey_importCmd).Standalone()
	compute_sshKey_importCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `FingerPrint`")
	compute_sshKey_importCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_sshKey_importCmd.Flags().String("public-key-file", "", "Public key file (required)")
	compute_sshKeyCmd.AddCommand(compute_sshKey_importCmd)

	carapace.Gen(compute_sshKey_importCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "FingerPrint").Invoke(c).Filter(c.Parts).ToA()
		}),
		"public-key-file": carapace.ActionFiles(),
	})
}
