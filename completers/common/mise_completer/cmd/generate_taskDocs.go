// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_taskDocsCmd = &cobra.Command{
	Use:   "task-docs",
	Short: "Generate Markdown documentation for project tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_taskDocsCmd).Standalone()

	generate_taskDocsCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_taskDocsCmd.Flags().BoolP("index", "I", false, "Write only an index of tasks, intended for use with `--multi`")
	generate_taskDocsCmd.Flags().BoolP("inject", "i", false, "Replace the marked section of the --output file")
	generate_taskDocsCmd.Flags().BoolP("multi", "m", false, "Render each task as a separate document; requires `--output` to be a directory")
	generate_taskDocsCmd.Flags().StringP("output", "o", "", "Write the generated docs to a file or directory")
	generate_taskDocsCmd.Flags().StringP("root", "r", "", "Has no effect; tasks always come from the current directory")
	generate_taskDocsCmd.Flags().StringP("style", "s", "simple", "Has no effect; both styles produce the same output")
	generate_taskDocsCmd.Flag("root").Hidden = true
	generate_taskDocsCmd.Flag("style").Hidden = true
	generateCmd.AddCommand(generate_taskDocsCmd)

	carapace.Gen(generate_taskDocsCmd).FlagCompletion(carapace.ActionMap{
		"root":  carapace.ActionDirectories(),
		"style": carapace.ActionValues("simple", "detailed"),
	})
}
