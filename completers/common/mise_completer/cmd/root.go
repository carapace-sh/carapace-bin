package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type commandSpec struct {
	path        []string
	description string
	aliases     []string
	flags       []flagSpec
}

type flagSpec struct {
	spec        string
	description string
	value       bool
	array       bool
	count       bool
	persistent  bool
	choices     []string
}

var rootCmd = &cobra.Command{
	Use:   "mise",
	Short: "Dev tools, env vars, and tasks in one CLI",
	Long:  "https://mise.jdx.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	commands := map[string]*cobra.Command{
		"": rootCmd,
	}
	addFlags(rootCmd, rootFlagSpecs)

	for _, spec := range commandSpecs {
		cmd := &cobra.Command{
			Use:     spec.path[len(spec.path)-1],
			Short:   firstLine(spec.description),
			Aliases: spec.aliases,
			Run:     func(cmd *cobra.Command, args []string) {},
		}
		addFlags(cmd, spec.flags)
		addPositionalCompletion(strings.Join(spec.path, " "), cmd)

		commands[strings.Join(spec.path, " ")] = cmd
		commands[strings.Join(spec.path[:len(spec.path)-1], " ")].AddCommand(cmd)
	}

	addPositionalCompletion("", rootCmd)
}

func firstLine(s string) string {
	if before, _, ok := strings.Cut(s, `\n`); ok {
		return before
	}
	if before, _, ok := strings.Cut(s, "\n"); ok {
		return before
	}
	return s
}

func addFlags(cmd *cobra.Command, flags []flagSpec) {
	for _, spec := range flags {
		name, shorthand := parseFlagSpec(spec.spec)
		if name == "" {
			continue
		}

		flagSet := cmd.Flags()
		if spec.persistent {
			flagSet = cmd.PersistentFlags()
		}
		addFlag(flagSet, name, shorthand, spec)

		if completion, ok := flagCompletion(name, spec); ok {
			carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
				name: completion,
			})
		}
	}
}

func parseFlagSpec(spec string) (name string, shorthand string) {
	for _, field := range strings.Fields(spec) {
		switch {
		case strings.HasPrefix(field, "--"):
			name = strings.TrimPrefix(field, "--")
		case strings.HasPrefix(field, "-"):
			shorthand = strings.TrimPrefix(field, "-")
		}
	}
	if name == "" {
		name = shorthand
	}
	return
}

func addFlag(flagSet *pflag.FlagSet, name string, shorthand string, spec flagSpec) {
	switch {
	case spec.count:
		if shorthand != "" && shorthand != name {
			flagSet.CountP(name, shorthand, spec.description)
		} else {
			flagSet.Count(name, spec.description)
		}
	case spec.value && spec.array:
		if shorthand != "" && shorthand != name {
			flagSet.StringArrayP(name, shorthand, nil, spec.description)
		} else {
			flagSet.StringArray(name, nil, spec.description)
		}
	case spec.value:
		if shorthand != "" && shorthand != name {
			flagSet.StringP(name, shorthand, "", spec.description)
		} else {
			flagSet.String(name, "", spec.description)
		}
	default:
		if shorthand != "" && shorthand != name {
			flagSet.BoolP(name, shorthand, false, spec.description)
		} else {
			flagSet.Bool(name, false, spec.description)
		}
	}
}

func flagCompletion(name string, spec flagSpec) (carapace.Action, bool) {
	if len(spec.choices) > 0 {
		return carapace.ActionValues(spec.choices...), true
	}
	if !spec.value {
		return carapace.ActionValues(), false
	}

	switch name {
	case "backend":
		return action.ActionBackends(), true
	case "cd", "dir", "localized-dir", "prefix", "root":
		return carapace.ActionDirectories(), true
	case "config", "env-file", "file", "manifest-path", "output", "tool-versions", "write":
		return carapace.ActionFiles(), true
	case "plugin":
		return action.ActionPlugins(), true
	case "shell":
		return action.ActionShells(), true
	case "task":
		return action.ActionTasks(), true
	case "tool":
		return action.ActionToolVersions(), true
	default:
		return carapace.ActionValues(), false
	}
}

func addPositionalCompletion(path string, cmd *cobra.Command) {
	switch path {
	case "":
		carapace.Gen(cmd).PositionalCompletion(action.ActionTasks())
	case "activate", "completion", "env":
		carapace.Gen(cmd).PositionalCompletion(action.ActionShells())
	case "config get", "config set":
		carapace.Gen(cmd).PositionalCompletion(action.ActionConfigKeys())
	case "en":
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionDirectories())
	case "exec", "install", "install-into", "latest", "link", "ls-remote", "shell", "use":
		carapace.Gen(cmd).PositionalAnyCompletion(action.ActionToolVersions().FilterArgs())
	case "ls", "outdated", "prune", "uninstall", "upgrade", "where":
		carapace.Gen(cmd).PositionalAnyCompletion(action.ActionInstalledToolVersions().FilterArgs())
	case "plugins install":
		carapace.Gen(cmd).PositionalCompletion(action.ActionRemotePlugins())
	case "plugins link":
		carapace.Gen(cmd).PositionalCompletion(action.ActionPlugins(), carapace.ActionDirectories())
	case "plugins uninstall", "plugins update":
		carapace.Gen(cmd).PositionalAnyCompletion(action.ActionPlugins().FilterArgs())
	case "registry", "search", "tool":
		carapace.Gen(cmd).PositionalCompletion(action.ActionRegistryTools())
	case "settings add", "settings get", "settings ls", "settings set", "settings unset":
		carapace.Gen(cmd).PositionalCompletion(action.ActionSettings())
	case "shell-alias get", "shell-alias set", "shell-alias unset":
		carapace.Gen(cmd).PositionalCompletion(action.ActionShellAliases())
	case "tasks", "tasks edit", "tasks info":
		carapace.Gen(cmd).PositionalCompletion(action.ActionTasks())
	case "run", "tasks deps", "tasks run", "tasks validate", "watch":
		carapace.Gen(cmd).PositionalAnyCompletion(action.ActionTasks().FilterArgs())
	case "token forgejo", "token github", "token gitlab":
		carapace.Gen(cmd).PositionalCompletion(action.ActionHosts())
	case "trust", "untrust":
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionFiles(".toml"))
	case "which":
		carapace.Gen(cmd).PositionalCompletion(carapace.ActionExecutables())
	default:
		if strings.HasSuffix(path, " config") {
			carapace.Gen(cmd).PositionalCompletion(carapace.ActionFiles(".toml"))
		}
	}
}
