package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "create a new repository in the given directory",
	GroupID: groups[group_repository_creation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	initCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	initCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"remotecmd": carapace.ActionExecutables(),
		"ssh":       carapace.ActionExecutables(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
