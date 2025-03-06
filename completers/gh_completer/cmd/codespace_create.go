package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_createCmd).Standalone()

	codespace_createCmd.Flags().StringP("branch", "b", "", "Repository branch")
	codespace_createCmd.Flags().Bool("default-permissions", false, "Do not prompt to accept additional permissions requested by the codespace")
	codespace_createCmd.Flags().String("devcontainer-path", "", "Path to the devcontainer.json file to use when creating codespace")
	codespace_createCmd.Flags().StringP("display-name", "d", "", "Display name for the codespace (48 characters or less)")
	codespace_createCmd.Flags().String("idle-timeout", "", "Allowed inactivity before codespace is stopped, e.g. \"10m\", \"1h\"")
	codespace_createCmd.Flags().StringP("location", "l", "", "Location: {EastUs|SouthEastAsia|WestEurope|WestUs2} (determined automatically if not provided)")
	codespace_createCmd.Flags().StringP("machine", "m", "", "Hardware specifications for the VM")
	codespace_createCmd.Flags().StringP("repo", "R", "", "Repository name with owner: user/repo")
	codespace_createCmd.Flags().StringP("repo-deprecated", "r", "", "(Deprecated) Shorthand for --repo")
	codespace_createCmd.Flags().String("retention-period", "", "Allowed time after shutting down before the codespace is automatically deleted (maximum 30 days), e.g. \"1h\", \"72h\"")
	codespace_createCmd.Flags().BoolP("status", "s", false, "Show status of post-create command and dotfiles")
	codespace_createCmd.Flags().BoolP("web", "w", false, "Create codespace from browser, cannot be used with --display-name, --idle-timeout, or --retention-period")
	codespace_createCmd.Flag("repo-deprecated").Hidden = true
	codespaceCmd.AddCommand(codespace_createCmd)

	// TODO devcontainer-path
	carapace.Gen(codespace_createCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(codespace_createCmd),
		"devcontainer-path": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			path := filepath.Dir(c.Value)
			path = strings.TrimPrefix(path, "/")
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return action.ActionContents(codespace_createCmd, path, codespace_createCmd.Flag("branch").Value.String())
			})
		}),
		"location": carapace.ActionValues("EastUs", "SouthEastAsia", "WestEurope", "WestUs2"),
		"machine":  action.ActionCodespaceMachines(),
		"repo":     action.ActionOwnerRepositories(codespace_createCmd),
	})
}
