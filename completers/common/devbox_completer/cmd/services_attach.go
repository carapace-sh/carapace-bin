package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running process-compose for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_attachCmd).Standalone()

	servicesCmd.AddCommand(services_attachCmd)
}
