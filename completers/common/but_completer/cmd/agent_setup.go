package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure GitButler skills and workflow instructions for coding agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_setupCmd).Standalone()

	agent_setupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agent_setupCmd.Flags().Bool("print", false, "Print the default generated steering text without prompting or modifying files")
	agentCmd.AddCommand(agent_setupCmd)
}
