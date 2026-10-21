package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshForEachCmd = &cobra.Command{
	Use:   "ssh-for-each",
	Short: "execute a command on each machine via SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshForEachCmd).Standalone()

	sshForEachCmd.Flags().Bool("all", false, "Run ssh-for-each for all deployments")
	sshForEachCmd.Flags().StringSlice("exclude", nil, "Run command on all except the specified machines")
	sshForEachCmd.Flags().StringSlice("include", nil, "Run command only on the specified machines")
	sshForEachCmd.Flags().BoolP("parallel", "p", false, "Run in parallel")

	rootCmd.AddCommand(sshForEachCmd)
}
