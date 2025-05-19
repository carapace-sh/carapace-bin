package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var organizations_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an organization",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createCmd).Standalone()

	organizations_createCmd.Flags().StringP("description", "d", "", "")
	organizations_createCmd.Flags().StringP("location", "L", "", "")
	organizations_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	organizations_createCmd.Flags().StringP("name", "n", "", "")
	organizations_createCmd.Flags().Bool("repo-admins-can-change-team-access", false, "")
	organizations_createCmd.Flags().StringP("visibility", "v", "", "")
	organizations_createCmd.Flags().StringP("website", "w", "", "")
	organizationsCmd.AddCommand(organizations_createCmd)

	// TODO completion
	carapace.Gen(organizations_createCmd).FlagCompletion(carapace.ActionMap{
		"login": tea.ActionLogins(),
	})
}
