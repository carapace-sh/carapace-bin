package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var clangFormatCmd = &cobra.Command{
	Use:   "clang-format",
	Short: "run clang-format on lines that differ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clangFormatCmd).Standalone()

	clangFormatCmd.Flags().String("binary", "", "path to clang-format")
	clangFormatCmd.Flags().String("commit", "", "default commit to use if none is specified")
	clangFormatCmd.Flags().Bool("diff", false, "print a diff instead of applying the changes")
	clangFormatCmd.Flags().String("extensions", "", "comma-separated list of file extensions to format,")
	clangFormatCmd.Flags().BoolP("force", "f", false, "allow changes to unstaged files")
	clangFormatCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	clangFormatCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	clangFormatCmd.Flags().BoolP("quiet", "q", false, "print less information")
	clangFormatCmd.Flags().String("style", "", "passed to clang-format")
	clangFormatCmd.Flags().BoolP("verbose", "v", false, "print extra information")
	rootCmd.AddCommand(clangFormatCmd)

	carapace.Gen(clangFormatCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if clangFormatCmd.Flags().ArgsLenAtDash() != -1 {
				return carapace.ActionFiles()
			}

			if diff := clangFormatCmd.Flag("diff").Changed; len(c.Args) == 0 || (len(c.Args) == 1 && diff) {
				return git.ActionRefs(git.RefOptionDefault)
			}
			return carapace.ActionFiles()
		}),
	)
}
