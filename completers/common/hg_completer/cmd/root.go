package cmd

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "hg",
	Short: "Mercurial distributed SCM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

const (
	group_repository_creation = iota
	group_remote_repository_management
	group_change_creation
	group_change_manipulation
	group_change_organization
	group_file_content_management
	group_change_navigation
	group_working_directory_management
	group_change_import_export
	group_repository_maintenance
	group_help
	group_alias
	group_extension
)

var groups = []*cobra.Group{
	{ID: "creation", Title: "Repository creation"},
	{ID: "remote", Title: "Remote repository management"},
	{ID: "change-creation", Title: "Change creation"},
	{ID: "change-manipulation", Title: "Change manipulation"},
	{ID: "change-organization", Title: "Change organization"},
	{ID: "content", Title: "File content management"},
	{ID: "navigation", Title: "Change navigation"},
	{ID: "working-dir", Title: "Working directory management"},
	{ID: "import-export", Title: "Change import/export"},
	{ID: "maintenance", Title: "Repository maintenance"},
	{ID: "help", Title: "Help"},
	{ID: "alias", Title: "Alias commands"},
	{ID: "extension", Title: "Extension commands"},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(groups...)

	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("color", "", "when to colorize (boolean, always, auto, never, or debug)")
	rootCmd.Flags().String("config", "", "set/override config option (use 'section.name=value')")
	rootCmd.Flags().String("config-file", "", "load config file to set/override config options")
	rootCmd.Flags().String("cwd", "", "change working directory")
	rootCmd.Flags().Bool("debug", false, "enable debugging output")
	rootCmd.Flags().Bool("debugger", false, "start debugger")
	rootCmd.Flags().String("encoding", "", "set the charset encoding")
	rootCmd.Flags().String("encodingmode", "", "set the charset encoding mode")
	rootCmd.Flags().BoolP("help", "h", false, "display help and exit")
	rootCmd.Flags().Bool("hidden", false, "consider hidden changesets")
	rootCmd.Flags().BoolP("noninteractive", "y", false, "do not prompt, automatically pick the first choice for all prompts")
	rootCmd.Flags().String("pager", "", "when to paginate (boolean, always, auto, or never)")
	rootCmd.Flags().Bool("profile", false, "print command execution profile")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	rootCmd.Flags().StringP("repository", "R", "", "repository root directory or name of overlay bundle file")
	rootCmd.Flags().Bool("time", false, "time how long the command takes")
	rootCmd.Flags().Bool("traceback", false, "always print a traceback on exception")
	rootCmd.Flags().BoolP("verbose", "v", false, "enable additional output")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("true", "false", "always", "auto", "never", "debug"),
		"config": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return hg.ActionConfigKeys()
			default:
				return carapace.ActionValues()
			}
		}),
		"config-file": carapace.ActionFiles(),
		"cwd":         carapace.ActionDirectories(),
		"pager":       carapace.ActionValues("true", "false", "always", "auto", "never"),
		"repository":  carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("cwd").Value.String())
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if c, _, _ := rootCmd.Find(args); c == rootCmd && len(args) > 0 {
			addAliasCompletion(args)
			addExtensionCommands(args)
		}
	})
}

func addAliasCompletion(args []string) {
	cmd := &cobra.Command{}
	cmd.FParseErrWhitelist.UnknownFlags = true
	cmd.Flags().String("cwd", "", "")
	cmd.Flags().StringP("repository", "R", "", "")
	cmd.ParseFlags(args[:len(args)-1])

	aliases, err := hg.Aliases(cmd.Flag("cwd").Value.String(), cmd.Flag("repository").Value.String())
	if err != nil {
		carapace.LOG.Println(err.Error())
		return
	}

	for key, value := range aliases {
		if _, _, err := rootCmd.Find([]string{key}); err == nil {
			continue // don't clobber existing commands
		}

		aliasCmd := &cobra.Command{
			Use:                key,
			Short:              fmt.Sprintf("alias for '%s'", value),
			GroupID:            groups[group_alias].ID,
			DisableFlagParsing: true,
			Run:                func(cmd *cobra.Command, args []string) {},
		}

		rootCmd.AddCommand(aliasCmd)

		switch {
		case strings.HasPrefix(value, "!"): // shell alias
			tokens, err := shlex.Split(strings.TrimPrefix(value, "!"))
			if err != nil {
				carapace.LOG.Println("failed to parse shell alias: " + err.Error())
				continue
			}
			carapace.Gen(aliasCmd).PositionalAnyCompletion(
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Args = append(tokens.CurrentPipeline().Words().Strings(), c.Args...)
					return bridge.ActionCarapaceBin().Invoke(c).ToA()
				}),
			)

		default: // mercurial alias
			tokens, err := shlex.Split(value)
			if err != nil {
				carapace.LOG.Println("failed to parse alias: " + err.Error())
				continue
			}
			carapace.Gen(aliasCmd).PositionalAnyCompletion(
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Args = append(tokens.Words().Strings(), c.Args...)
					return bridge.ActionCarapaceBin("hg").Invoke(c).ToA()
				}),
			)
		}
	}
}

func addExtensionCommands(args []string) {
	hgArgs := append(globalArgs(args), "debugcomplete")
	if output, err := (carapace.Context{}).Command("hg", hgArgs...).Output(); err != nil {
		carapace.LOG.Println(err.Error())
		return
	} else {
		for _, name := range strings.Fields(string(output)) {
			if _, _, err := rootCmd.Find([]string{name}); err == nil {
				continue // don't clobber existing commands
			}

			extensionCmd := &cobra.Command{
				Use:                name,
				Short:              "extension",
				GroupID:            groups[group_extension].ID,
				DisableFlagParsing: true,
				Run:                func(cmd *cobra.Command, args []string) {},
			}

			rootCmd.AddCommand(extensionCmd)
		}
	}
}

func globalArgs(args []string) []string {
	cmd := &cobra.Command{}
	cmd.FParseErrWhitelist.UnknownFlags = true
	cmd.Flags().String("cwd", "", "")
	cmd.Flags().StringP("repository", "R", "", "")
	cmd.ParseFlags(args[:len(args)-1])

	globalArgs := []string{}
	if dir := cmd.Flag("cwd").Value.String(); dir != "" {
		globalArgs = append(globalArgs, "--cwd", dir)
	}
	if repo := cmd.Flag("repository").Value.String(); repo != "" {
		globalArgs = append(globalArgs, "--repository", repo)
	}
	return globalArgs
}
