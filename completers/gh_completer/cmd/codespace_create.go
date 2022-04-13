package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_createCmd).Standalone()
	codespace_createCmd.Flags().StringP("branch", "b", "", "repository branch")
	codespace_createCmd.Flags().Bool("default-permissions", false, "do not prompt to accept additional permissions requested by the codespace")
	codespace_createCmd.Flags().Duration("idle-timeout", 0, "allowed inactivity before codespace is stopped, e.g. \"10m\", \"1h\"")
	codespace_createCmd.Flags().StringP("location", "l", "", "location: {EastUs|SouthEastAsia|WestEurope|WestUs2} (determined automatically if not provided)")
	codespace_createCmd.Flags().StringP("machine", "m", "", "hardware specifications for the VM")
	codespace_createCmd.Flags().StringP("repo", "r", "", "repository name with owner: user/repo")
	codespace_createCmd.Flags().BoolP("status", "s", false, "show status of post-create command and dotfiles")
	codespaceCmd.AddCommand(codespace_createCmd)

	carapace.Gen(codespace_createCmd).FlagCompletion(carapace.ActionMap{
		"branch":   action.ActionBranches(codespace_createCmd),
		"location": carapace.ActionValues("EastUs", "SouthEastAsia", "WestEurope", "WestUs2"),
		"machine":  action.ActionCodespaceMachines(),
		"repo":     action.ActionOwnerRepositories(codespace_createCmd),
	})
}
