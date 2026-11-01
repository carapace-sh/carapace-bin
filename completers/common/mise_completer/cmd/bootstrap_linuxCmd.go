package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_linuxCmd = &cobra.Command{
	Use:   "linux",
	Short: "Manage systemd user units from `[bootstrap.linux.systemd.units]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_linuxCmd).Standalone()

	bootstrap_linuxCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_linuxCmd)
}
