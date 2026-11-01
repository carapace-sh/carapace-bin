package common

import "github.com/spf13/cobra"

func AddSkillAgentFlags(cmd *cobra.Command) {
	cmd.Flags().String("agent", "", "limit to the agent NAME (repeatable)")
	cmd.Flags().Bool("dry-run", false, "report what would change without changing anything")
	cmd.Flags().Bool("project", false, "use the current repository instead of your account")
}
