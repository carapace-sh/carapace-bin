package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var lsFilesCmd = &cobra.Command{
	Use:     "ls-files",
	Short:   "Show information about files in the index and the working tree",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(lsFilesCmd).Standalone()

	lsFilesCmd.Flags().String("abbrev", "", "use <n> digits to display object names")
	lsFilesCmd.Flags().BoolP("cached", "c", false, "show cached files in the output (default)")
	lsFilesCmd.Flags().Bool("debug", false, "show debugging data")
	lsFilesCmd.Flags().Bool("deduplicate", false, "suppress duplicate entries")
	lsFilesCmd.Flags().BoolP("deleted", "d", false, "show deleted files in the output")
	lsFilesCmd.Flags().Bool("directory", false, "show 'other' directories' names only")
	lsFilesCmd.Flags().Bool("empty-directory", false, "don't show empty directories")
	lsFilesCmd.Flags().Bool("eol", false, "show line endings of files")
	lsFilesCmd.Flags().Bool("error-unmatch", false, "if any <file> is not in the index, treat this as an error")
	lsFilesCmd.Flags().StringP("exclude", "x", "", "skip files matching pattern")
	lsFilesCmd.Flags().StringP("exclude-from", "X", "", "read exclude patterns from <file>")
	lsFilesCmd.Flags().String("exclude-per-directory", "", "read additional per-directory exclude patterns in <file>")
	lsFilesCmd.Flags().Bool("exclude-standard", false, "add the standard git exclusions")
	lsFilesCmd.Flags().BoolS("f", "f", false, "use lowercase letters for 'fsmonitor clean' files")
	lsFilesCmd.Flags().String("format", "", "format to use for the output")
	lsFilesCmd.Flags().Bool("full-name", false, "make the output relative to the project top directory")
	lsFilesCmd.Flags().BoolP("ignored", "i", false, "show ignored files in the output")
	lsFilesCmd.Flags().BoolP("killed", "k", false, "show files on the filesystem that need to be removed")
	lsFilesCmd.Flags().BoolP("modified", "m", false, "show modified files in the output")
	lsFilesCmd.Flags().BoolP("others", "o", false, "show other files in the output")
	lsFilesCmd.Flags().Bool("recurse-submodules", false, "recurse through submodules")
	lsFilesCmd.Flags().Bool("resolve-undo", false, "show resolve-undo information")
	lsFilesCmd.Flags().Bool("sparse", false, "show sparse directories in the presence of a sparse index")
	lsFilesCmd.Flags().BoolP("stage", "s", false, "show staged contents' object name in the output")
	lsFilesCmd.Flags().BoolS("t", "t", false, "identify the file status with tags")
	lsFilesCmd.Flags().BoolP("unmerged", "u", false, "show unmerged files in the output")
	lsFilesCmd.Flags().BoolS("v", "v", false, "use lowercase letters for 'assume unchanged' files")
	lsFilesCmd.Flags().String("with-tree", "", "pretend that paths removed since <tree-ish> are still present")
	lsFilesCmd.Flags().BoolS("z", "z", false, "separate paths with the NUL character")
	rootCmd.AddCommand(lsFilesCmd)

	carapace.Gen(lsFilesCmd).FlagCompletion(carapace.ActionMap{
		"exclude-from":          carapace.ActionFiles(),
		"exclude-per-directory": carapace.ActionFiles(),
		"with-tree":             git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(lsFilesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{"ls-files"}
			lsFilesCmd.Flags().Visit(func(f *pflag.Flag) {
				switch f.Name {
				case
					"cached",
					"deleted",
					"directory",
					"empty-directory",
					"error-unmatch",
					"exclude-standard",
					"ignored",
					"killed",
					"modified",
					"others",
					"recurse-submodules",
					"resolve-undo",
					"sparse",
					"unmerged":
					args = append(args, "--"+f.Name)

				case
					"exclude",
					"exclude-from",
					"exclude-per-directory",
					"with-tree":
					args = append(args, "--"+f.Name, f.Value.String())

				}
			})

			if dir := filepath.Dir(c.Value); dir != "." {
				args = append(args, dir)
			}

			return carapace.ActionExecCommand("git", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...).MultiParts("/").StyleF(style.ForPathExt)
			})
		}),
	)
}
