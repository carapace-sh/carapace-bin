package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Retrieve funding information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fundCmd).Standalone()
	fundCmd.Flags().String("browser", "", "browser to use")
	fundCmd.Flags().Bool("json", false, "output as json")
	fundCmd.Flags().Bool("unicode", false, "use unicode characters in output")
	fundCmd.Flags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	fundCmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")
	fundCmd.Flags().Int("which", 1, "index of funding source to open")

	rootCmd.AddCommand(fundCmd)

	carapace.Gen(fundCmd).PositionalCompletion(
		action.ActionDependencies(), // TODO only list dependencies listed by `npm fund`
	)
}
