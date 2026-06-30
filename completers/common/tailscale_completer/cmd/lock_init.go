package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_initCmd).Standalone()

	lock_initCmd.Flags().Bool("confirm", false, "do not prompt for confirmation")
	lock_initCmd.Flags().Bool("gen-disablement-for-support", false, "generate and transmit a disablement secret for Tailscale support")
	lock_initCmd.Flags().String("gen-disablements", "", "number of disablement secrets to generate")
	lockCmd.AddCommand(lock_initCmd)
}
