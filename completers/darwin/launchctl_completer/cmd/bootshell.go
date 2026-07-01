package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootshellCmd = &cobra.Command{
	Use:   "bootshell",
	Short: "Brings the system up from single-user mode with a console shell",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(bootshellCmd).Standalone()
	rootCmd.AddCommand(bootshellCmd)
}
