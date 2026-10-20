// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packslip_pinsCmd = &cobra.Command{
	Use:     "pins",
	Short:   "List the signer mise has pinned for each packslip project",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packslip_pinsCmd).Standalone()

	packslip_pinsCmd.Flags().BoolP("help", "h", false, "Print help")
	packslip_pinsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	packslipCmd.AddCommand(packslip_pinsCmd)
}
