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
	agent_setupCmd.Flags().Bool("stub", false, "Install discovery stubs that point agents at but skill instead of the full skill files. Team-internal while the approach is evaluated")
	agent_setupCmd.Flag("stub").Hidden = true
	agentCmd.AddCommand(agent_setupCmd)
}
