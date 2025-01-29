package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_createCmd = &cobra.Command{
	Use:   "create [options] NAME DESTINATION",
	Short: "Record destination for the Podman service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_createCmd).Standalone()

	context_createCmd.Flags().String("default-stack-orchestrator", "", "Ignored.  Just for script compatibility")
	context_createCmd.Flags().String("description", "", "Ignored.  Just for script compatibility")
	context_createCmd.Flags().String("docker", "", "Description of the context")
	context_createCmd.Flags().String("from", "", "Ignored.  Just for script compatibility")
	context_createCmd.Flags().String("kubernetes", "", "Ignored.  Just for script compatibility")
	contextCmd.AddCommand(context_createCmd)
}
