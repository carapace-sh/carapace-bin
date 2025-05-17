package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brew",
	Short: "The missing package manager for macOS",
	Long:  "https://brew.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(
		&cobra.Group{ID: "main", Title: "main commands"},
		&cobra.Group{ID: "developer", Title: "developer commands"},
		&cobra.Group{ID: "external", Title: "external commands"},
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if _, ok := os.LookupEnv("HOMEBREW_DEVELOPER"); !ok {
			for _, c := range cmd.Commands() {
				if c.GroupID == "developer" {
					c.Hidden = true
				}
			}
		}
	})

	embed.SubcommandsAsFlags(rootCmd,
		flagCacheCmd,
		flagCaskroomCmd,
		flagCellarCmd,
		flagConfigCmd,
		flagEnvCmd,
		flagPrefixCmd,
		flagRepoCmd,
		flagRepositoryCmd,
	)
	embed.SubcommandsAsFlagsS(rootCmd, flagSCmd)

	// TODO external commands
}
