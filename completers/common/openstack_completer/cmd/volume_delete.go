package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete volume(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_deleteCmd).Standalone()

	volume_deleteCmd.Flags().Bool("cascade", false, "Remove any snapshots along with volume(s) (defaults to False)")
	volume_deleteCmd.Flags().Bool("force", false, "Attempt forced removal of volume(s), regardless of state (defaults to False)")
	volume_deleteCmd.Flags().Bool("purge", false, "==SUPPRESS==")
	volume_deleteCmd.Flags().Bool("remote", false, "Specify this parameter to unmanage a volume.")
	volumeCmd.AddCommand(volume_deleteCmd)
}
