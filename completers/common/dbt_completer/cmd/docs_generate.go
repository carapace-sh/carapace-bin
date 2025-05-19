package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docs_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docs_generateCmd).Standalone()

	docs_generateCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	docs_generateCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	docs_generateCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	docs_generateCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	docs_generateCmd.Flags().Bool("no-compile", false, "Do not run \"dbt compile\" as part of docs generation")
	docs_generateCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	docs_generateCmd.Flags().Bool("no-favor-state", false, "If defer is set, expect standard defer behaviour")
	docs_generateCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	docs_generateCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	docs_generateCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	docs_generateCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	docs_generateCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	docs_generateCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	addProjectDirFlag(docs_generateCmd)
	addSelectionFlags(docs_generateCmd)
	addModelsFlag(docs_generateCmd)
	addProfileFlag(docs_generateCmd)
	docsCmd.AddCommand(docs_generateCmd)

	carapace.Gen(docs_generateCmd).FlagCompletion(carapace.ActionMap{
		"log-path":    carapace.ActionFiles(),
		"target-path": carapace.ActionDirectories(),
	})
}
