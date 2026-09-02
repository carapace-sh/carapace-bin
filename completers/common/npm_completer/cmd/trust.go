package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:   "trust",
	Short: "Manage trusted publishing relationships between packages and CI/CD providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()
	trustCmd.Flags().Bool("allow-publish", false, "allow the trusted publisher to run npm publish")
	trustCmd.Flags().Bool("allow-stage-publish", false, "allow the trusted publisher to run npm stage")
	trustCmd.Flags().String("context-id", "", "CircleCI context UUID to match")
	trustCmd.Flags().Bool("dry-run", false, "report what would be done without making changes")
	trustCmd.Flags().String("env", "", "CI environment name")
	trustCmd.Flags().String("environment", "", "CI environment name")
	trustCmd.Flags().String("file", "", "name of workflow or pipeline file")
	trustCmd.Flags().Bool("json", false, "output as json")
	trustCmd.Flags().String("org-id", "", "CircleCI organization UUID")
	trustCmd.Flags().String("pipeline-definition-id", "", "CircleCI pipeline definition UUID")
	trustCmd.Flags().String("project", "", "GitLab project name")
	trustCmd.Flags().String("project-id", "", "CircleCI project UUID")
	trustCmd.Flags().String("repo", "", "repository name")
	trustCmd.Flags().String("repository", "", "repository name")
	trustCmd.Flags().BoolP("yes", "y", false, "auto-answer yes to any prompts")

	rootCmd.AddCommand(trustCmd)

	carapace.Gen(trustCmd).PositionalCompletion(
		carapace.ActionValues("circleci", "github", "gitlab", "list", "revoke"),
	)
}
