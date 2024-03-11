package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] CONTEXT",
	Short: "Create a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_createCmd).Standalone()

	context_createCmd.Flags().String("description", "", "Description of the context")
	context_createCmd.Flags().String("docker", "", "set the docker endpoint")
	context_createCmd.Flags().String("from", "", "create context from a named context")
	contextCmd.AddCommand(context_createCmd)
}
