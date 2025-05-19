package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var trust_inspectCmd = &cobra.Command{
	Use:   "inspect IMAGE[:TAG] [IMAGE[:TAG]...]",
	Short: "Return low-level information about keys and signatures",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_inspectCmd).Standalone()

	trust_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	trustCmd.AddCommand(trust_inspectCmd)

	carapace.Gen(trust_inspectCmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
}
