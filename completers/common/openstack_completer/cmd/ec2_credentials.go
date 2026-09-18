package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_credentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_credentialsCmd).Standalone()

	ec2Cmd.AddCommand(ec2_credentialsCmd)
}
