package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:     "query [flags] [OPERATION]",
	Short:   "Send API queries to a dagger engine",
	GroupID: "exec",
	Aliases: []string{"q"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().String("doc", "", "Read query from file (defaults to reading from stdin)")
	queryCmd.PersistentFlags().Bool("focus", false, "Only show output for focused commands")
	queryCmd.PersistentFlags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	queryCmd.Flags().StringSlice("var", nil, "List of query variables, in key=value format")
	queryCmd.Flags().String("var-json", "", "Query variables in JSON format (overrides --var)")
	rootCmd.AddCommand(queryCmd)

	carapace.Gen(queryCmd).FlagCompletion(carapace.ActionMap{
		"doc": carapace.ActionFiles(),
		"mod": dagger.ActionMods(),
	})

	// TODO positional
}
