package cmd

import (
	"fmt"
	"strings"

	"github.com/google/shlex"
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "the stupid content tracker",
	Long:  "https://git-scm.com/",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().StringS("C", "C", "", "run as if git was started in given path")
	rootCmd.Flags().Bool("bare", false, "use $PWD as repository")
	rootCmd.Flags().StringS("c", "c", "", "pass configuration parameter to command")
	rootCmd.Flags().String("exec-path", "", "path containing core git-programs")
	rootCmd.Flags().String("git-dir", "", "path to repository")
	rootCmd.Flags().Bool("help", false, "display help message")
	rootCmd.Flags().Bool("html-path", false, "display path to HTML documentation and exit")
	rootCmd.Flags().Bool("info-path", false, "print the path where the info files are installed and exit")
	rootCmd.Flags().String("list-cmds", "", "list commands")
	rootCmd.Flags().Bool("literal-pathspecs", false, "treat pathspecs literally, rather than as glob patterns")
	rootCmd.Flags().Bool("man-path", false, "print the manpath for the man pages for this version of Git and exit")
	rootCmd.Flags().String("namespace", "", "set the Git namespace")
	rootCmd.Flags().BoolP("no-pager", "P", false, "don't pipe git output into a pager")
	rootCmd.Flags().Bool("no-replace-objects", false, "do not use replacement refs to replace git objects")
	rootCmd.Flags().BoolP("paginate", "p", false, "pipe output into a pager")
	rootCmd.Flags().Bool("version", false, "display version information")
	rootCmd.Flags().String("work-tree", "", "path to working tree")

	rootCmd.Flag("list-cmds").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
		"c": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionConfigs()
			default:
				return carapace.ActionValues()
			}
		}),
		"exec-path": carapace.ActionDirectories(),
		"git-dir":   carapace.ActionDirectories(),
		"list-cmds": carapace.ActionValues("main", "others", "alias", "nohelpers"),
		"work-tree": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		if f := rootCmd.Flag("C"); f != flag && f.Changed {
			return action.Chdir(f.Value.String())
		}
		return action
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		addAliasCompletion(args)
		addOtherCommands()
	})
}

func addAliasCompletion(args []string) {
	// pass through args related to config
	rootCmd.ParseFlags(args)
	gitArgs := []string{}
	if f := rootCmd.Flag("C"); f.Changed {
		gitArgs = append(gitArgs, "-C", f.Value.String())
	}
	if f := rootCmd.Flag("git-dir"); f.Changed {
		gitArgs = append(gitArgs, "--git-dir", f.Value.String())
	}

	if aliases, err := git.Aliases(gitArgs); err == nil {
		for key, value := range aliases {
			// don't clobber existing commands
			if _, _, err := rootCmd.Find([]string{key}); err == nil {
				continue
			}

			aliasCmd := &cobra.Command{
				Use:   key,
				Short: fmt.Sprintf("alias for '%s'", value),
				// disable flag parsing so that we can forward them together with Args
				DisableFlagParsing: true,
			}

			rootCmd.AddCommand(aliasCmd)

			// aliases beginning with ! are arbitrary shell commands so don't add completion
			if strings.HasPrefix(value, "!") {
				continue
			}

			args, err := shlex.Split(value)
			if err == nil {
				carapace.Gen(aliasCmd).PositionalAnyCompletion(
					carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						c.Args = append(args, c.Args...)
						return bridge.ActionCarapaceBin("git").Invoke(c).ToA()
					}),
				)
			}
		}
	}
}

func addOtherCommands() {
	if output, err := (carapace.Context{}).Command("git", "--list-cmds=others").Output(); err == nil {
		lines := strings.Split(string(output), "\n")

		for _, name := range lines[:len(lines)-1] {
			pluginCmd := &cobra.Command{
				Use:                name,
				Short:              "",
				Run:                func(cmd *cobra.Command, args []string) {},
				DisableFlagParsing: true,
			}

			carapace.Gen(pluginCmd).PositionalAnyCompletion(
				bridge.ActionCarapaceBin("git-" + name),
			)

			rootCmd.AddCommand(pluginCmd)
		}
	}
}
