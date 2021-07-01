package cmd

import (
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Retrieve the ssh identity key path of the specified node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	sshKeyCmd.Flags().StringP("node", "n", "", "The node to get ssh-key path. Defaults to the primary control plane.")
	rootCmd.AddCommand(sshKeyCmd)
}
