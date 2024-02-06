package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util/embed"
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
