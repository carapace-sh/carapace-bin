package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var orbit_setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Guided setup for Orbit: verify access, install the skill, install the local CLI. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_setupCmd).Standalone()

	orbit_setupCmd.Flags().BoolP("global", "g", false, "Install the Orbit skill at user scope (`~/.agents/skills/`).")
	orbit_setupCmd.Flags().String("hostname", "", "GitLab hostname to verify. Defaults to the current repository's host or `gitlab.com`.")
	orbit_setupCmd.Flags().String("path", "", "Install the Orbit skill to the directory at `<path>`.")
	orbit_setupCmd.Flags().Bool("skip-local", false, "Skip the local CLI binary install step.")
	orbit_setupCmd.Flags().Bool("skip-skill", false, "Skip the agent-skill install step.")
	orbit_setupCmd.Flags().Bool("upgrade", false, "Re-fetch the skill and update the local CLI binary in place.")
	orbit_setupCmd.Flags().BoolP("yes", "y", false, "Skip every confirmation prompt.")
	orbitCmd.AddCommand(orbit_setupCmd)

	carapace.Gen(orbit_setupCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"path":     carapace.ActionDirectories(),
	})
}
