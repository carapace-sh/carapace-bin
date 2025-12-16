package cmd

import (
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chmod",
	Short: "change file mode bits",
	Long:  "https://www.man7.org/linux/man-pages/man1/chmod.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	// TODO ensure patching still works
	for i, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") || len(arg) < 2 {
			continue // skip non-flags
		}
		if arg == "--" {
			break // no need for patching at dash positions
		}
		switch arg[1] {
		case 'X', 'r', 's', 't', 'w', 'x':
			switch i {
			case len(os.Args) - 1:
				os.Args[i] = "--" // fake dash completion
				os.Args = append(os.Args, arg)
			default:
				os.Args[i] = "a" + arg // patch with the implicit (a)ll to avoid parsing issues
			}
		}
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "if a command line argument is a symbolic link to a directory, traverse it")
	rootCmd.Flags().BoolS("L", "L", false, "traverse every symbolic link to a directory encountered")
	rootCmd.Flags().BoolS("P", "P", false, "do not traverse any symbolic links")
	rootCmd.Flags().BoolP("changes", "c", false, "like verbose but report only when a change is made")
	rootCmd.Flags().Bool("dereference", false, "affect the referent of each symbolic link, rather than the symbolic link itself")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("no-dereference", "h", false, "affect each symbolic link, rather than the referent")
	rootCmd.Flags().Bool("no-preserve-root", false, "do not treat '/' specially (the default)")
	rootCmd.Flags().Bool("preserve-root", false, "fail to operate recursively on '/'")
	rootCmd.Flags().Bool("quiet", false, "suppress most error messages")
	rootCmd.Flags().BoolP("recursive", "R", false, "change files and directories recursively")
	rootCmd.Flags().String("reference", "", "use RFILE's mode instead of MODE values")
	rootCmd.Flags().BoolP("silent", "f", false, "suppress most error messages")
	rootCmd.Flags().BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("reference").Changed {
				return carapace.ActionFiles()
			}
			return fs.ActionFileModes()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionPositional(rootCmd),
	)
}
