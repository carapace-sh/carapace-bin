// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap___serviceExecCmd = &cobra.Command{
	Use:    "__service-exec",
	Short:  "Run a user service that carries an environment (Windows, internal)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap___serviceExecCmd).Standalone()

	bootstrap___serviceExecCmd.Flags().String("digest", "", "The digest of the launch the task was registered with")
	bootstrap___serviceExecCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap___serviceExecCmd.Flags().String("launch", "", "The stored launch to run it from")
	bootstrap___serviceExecCmd.MarkFlagRequired("digest")
	bootstrap___serviceExecCmd.MarkFlagRequired("launch")
	bootstrapCmd.AddCommand(bootstrap___serviceExecCmd)
}
