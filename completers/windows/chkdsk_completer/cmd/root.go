package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chkdsk",
	Short: "check a disk and display a status report for disk errors",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/chkdsk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "fix errors on the disk")
	rootCmd.Flags().BoolP("r", "r", false, "locate bad sectors and recover readable information")
	rootCmd.Flags().BoolP("v", "v", false, "display cleanup messages")
	rootCmd.Flags().BoolP("x", "x", false, "force the volume to dismount first")
}
