package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mise",
	Short: "The front-end to your dev env",
	Long:  "https://mise.jdx.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringP("cd", "C", "", "Change to this directory before executing the command")
	rootCmd.PersistentFlags().String("config", "", "Use a specific config file")
	rootCmd.PersistentFlags().StringP("env", "E", "", "Use a mise environment")
	rootCmd.PersistentFlags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	rootCmd.PersistentFlags().StringP("log-level", "L", "", "Log level")
	rootCmd.PersistentFlags().Bool("raw", false, "Directly pipe stdin/stdout/stderr from plugin to user")
	rootCmd.PersistentFlags().CountP("quiet", "q", "Suppress output")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Show extra output")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Answer yes to all prompts")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cd":        carapace.ActionDirectories(),
		"config":    carapace.ActionFiles(".toml"),
		"env":       actionEnvironments(),
		"log-level": carapace.ActionValues("trace", "debug", "info", "warn", "error"),
	})

	addRootCommands()
}

func addRootCommands() {
	activateCmd := addCommand(rootCmd, "activate [SHELL]", "Initialize mise in the current shell session")
	carapace.Gen(activateCmd).PositionalCompletion(actionShells())

	addCommand(rootCmd, "bin-paths", "List paths to installed tool binaries")

	backendsCmd := addCommand(rootCmd, "backends", "Manage tool backends")
	addCommand(backendsCmd, "ls", "List enabled tool backends")

	cacheCmd := addCommand(rootCmd, "cache", "Manage mise cache")
	addCommand(cacheCmd, "clear", "Clear the mise cache")
	addCommand(cacheCmd, "path", "Print the mise cache path")
	addCommand(cacheCmd, "prune", "Prune unused cache entries")

	completionCmd := addCommand(rootCmd, "completion [SHELL]", "Generate shell completion script")
	carapace.Gen(completionCmd).PositionalCompletion(actionCompletionShells())

	configCmd := addCommand(rootCmd, "config", "Manage mise config")
	configGetCmd := addCommand(configCmd, "get [KEY]", "Get a config value")
	configLsCmd := addCommand(configCmd, "ls", "List config values")
	configSetCmd := addCommand(configCmd, "set KEY VALUE", "Set a config value")
	configLsCmd.Flags().Bool("json", false, "Output in JSON format")
	carapace.Gen(configGetCmd).PositionalCompletion(actionSettings())
	carapace.Gen(configSetCmd).PositionalCompletion(actionSettings())

	deactivateCmd := addCommand(rootCmd, "deactivate [SHELL]", "Disable mise for the current shell session")
	carapace.Gen(deactivateCmd).PositionalCompletion(actionShells())

	depsCmd := addCommand(rootCmd, "deps", "Manage tool dependencies")
	addCommand(depsCmd, "add TOOL", "Add a dependency")
	addCommand(depsCmd, "install", "Install dependencies")
	depsRemoveCmd := addCommand(depsCmd, "remove TOOL", "Remove a dependency")
	carapace.Gen(depsRemoveCmd).PositionalCompletion(actionTools())

	doctorCmd := addCommand(rootCmd, "doctor", "Check mise installation for problems")
	addCommand(doctorCmd, "path", "Inspect PATH configuration")

	editCmd := addCommand(rootCmd, "edit", "Edit mise config")
	editCmd.Flags().BoolP("global", "g", false, "Edit the global config file")
	editCmd.Flags().StringP("path", "p", "", "Path to a config file or directory")
	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(".toml"),
	})

	addCommand(rootCmd, "en", "Show environment variables as envrc")
	addCommand(rootCmd, "env", "Show the mise environment")

	execCmd := addCommand(rootCmd, "exec [TOOL@VERSION] -- COMMAND", "Execute a command with mise tools")
	execCmd.Flags().Bool("raw", false, "Directly pipe stdin/stdout/stderr from plugin to user")
	carapace.Gen(execCmd).PositionalCompletion(actionTools())
	carapace.Gen(execCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin().Split())

	addCommand(rootCmd, "fmt", "Format mise config files")

	generateCmd := addCommand(rootCmd, "generate", "Generate mise integration files")
	addCommand(generateCmd, "bootstrap", "Generate a bootstrap script")
	addCommand(generateCmd, "config", "Generate config")
	addCommand(generateCmd, "devcontainer", "Generate devcontainer config")
	addCommand(generateCmd, "git-pre-commit", "Generate git pre-commit config")
	addCommand(generateCmd, "github-action", "Generate GitHub Action config")
	addCommand(generateCmd, "task-docs", "Generate task documentation")
	addCommand(generateCmd, "task-stubs", "Generate task stubs")
	addCommand(generateCmd, "tool-stub", "Generate a tool stub")

	addCommand(rootCmd, "implode", "Remove mise and its data")

	installCmd := addCommand(rootCmd, "install [TOOL@VERSION]...", "Install tools")
	installCmd.Flags().BoolP("force", "f", false, "Force reinstall even if already installed")
	installCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	carapace.Gen(installCmd).PositionalAnyCompletion(actionTools())

	installIntoCmd := addCommand(rootCmd, "install-into DIR [TOOL@VERSION]...", "Install tools into a directory")
	carapace.Gen(installIntoCmd).PositionalCompletion(carapace.ActionDirectories())
	carapace.Gen(installIntoCmd).PositionalAnyCompletion(actionTools())

	latestCmd := addCommand(rootCmd, "latest TOOL", "Show the latest available tool version")
	carapace.Gen(latestCmd).PositionalCompletion(actionTools())

	linkCmd := addCommand(rootCmd, "link TOOL VERSION PATH", "Link a tool version to a local path")
	carapace.Gen(linkCmd).PositionalCompletion(actionTools())
	carapace.Gen(linkCmd).PositionalAnyCompletion(carapace.ActionDirectories())

	addCommand(rootCmd, "lock", "Update mise.lock")

	lsCmd := addCommand(rootCmd, "ls [TOOL]", "List installed tool versions")
	lsCmd.Flags().Bool("current", false, "Only show current versions")
	lsCmd.Flags().Bool("json", false, "Output in JSON format")
	carapace.Gen(lsCmd).PositionalCompletion(actionInstalledTools())

	lsRemoteCmd := addCommand(rootCmd, "ls-remote TOOL", "List remote versions for a tool")
	carapace.Gen(lsRemoteCmd).PositionalCompletion(actionTools())

	addCommand(rootCmd, "mcp", "Run mise MCP server")

	ociCmd := addCommand(rootCmd, "oci", "Manage mise OCI images")
	ociBuildCmd := addCommand(ociCmd, "build IMAGE", "Build a mise OCI image")
	addCommand(ociCmd, "push IMAGE", "Push a mise OCI image")
	addCommand(ociCmd, "run IMAGE", "Run a mise OCI image")
	carapace.Gen(ociBuildCmd).PositionalCompletion(carapace.ActionFiles())

	outdatedCmd := addCommand(rootCmd, "outdated", "Show outdated tools")
	carapace.Gen(outdatedCmd).PositionalAnyCompletion(actionInstalledTools())

	pluginsCmd := addCommand(rootCmd, "plugins", "Manage plugins")
	pluginsInstallCmd := addCommand(pluginsCmd, "install PLUGIN", "Install a plugin")
	pluginsLinkCmd := addCommand(pluginsCmd, "link NAME PATH", "Link a local plugin")
	addCommand(pluginsCmd, "ls", "List installed plugins")
	addCommand(pluginsCmd, "ls-remote", "List remote plugins")
	pluginsUninstallCmd := addCommand(pluginsCmd, "uninstall PLUGIN", "Uninstall a plugin")
	pluginsUpdateCmd := addCommand(pluginsCmd, "update [PLUGIN]", "Update plugins")
	carapace.Gen(pluginsInstallCmd).PositionalCompletion(actionRemotePlugins())
	carapace.Gen(pluginsLinkCmd).PositionalAnyCompletion(carapace.ActionDirectories())
	carapace.Gen(pluginsUninstallCmd).PositionalCompletion(actionPlugins())
	carapace.Gen(pluginsUpdateCmd).PositionalCompletion(actionPlugins())

	addCommand(rootCmd, "prune", "Delete unused tools")

	registryCmd := addCommand(rootCmd, "registry", "List tools in the registry")
	registryCmd.Flags().Bool("json", false, "Output in JSON format")

	addCommand(rootCmd, "reshim", "Rebuild shims")

	runCmd := addCommand(rootCmd, "run [TASK]", "Run task(s)")
	runCmd.Flags().BoolP("continue-on-error", "c", false, "Continue running tasks even if one fails")
	runCmd.Flags().BoolP("force", "f", false, "Force tasks to run even if outputs are up to date")
	runCmd.Flags().StringP("jobs", "j", "", "Number of tasks to run in parallel")
	runCmd.Flags().StringP("output", "o", "", "Change how task output is displayed")
	runCmd.Flags().BoolP("raw", "r", false, "Read/write directly to stdin/stdout/stderr instead of by line")
	runCmd.Flags().StringP("shell", "s", "", "Shell to use to run toml tasks")
	runCmd.Flags().BoolP("silent", "S", false, "Do not show any output except errors")
	runCmd.Flags().StringSliceP("tool", "t", nil, "Tool(s) to run in addition to configured tools")
	runCmd.Flags().Bool("no-deps", false, "Skip automatic dependency preparation")
	runCmd.Flags().Bool("skip-deps", false, "Run only the specified tasks skipping all dependencies")
	runCmd.Flags().Bool("skip-tools", false, "Skip installing tools before running tasks")
	runCmd.Flags().String("timeout", "", "Timeout for the task to complete")
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("prefix", "interleave", "replacing", "timed", "keep-order", "quiet", "silent"),
		"shell":  actionShells(),
		"tool":   actionTools(),
	})
	carapace.Gen(runCmd).PositionalAnyCompletion(actionTasks())

	addCommand(rootCmd, "search", "Search the mise registry")

	addCommand(rootCmd, "self-update", "Update mise itself")

	setCmd := addCommand(rootCmd, "set KEY VALUE", "Set an environment variable")
	carapace.Gen(setCmd).PositionalCompletion(actionEnvironmentVariables())

	settingsCmd := addCommand(rootCmd, "settings", "Manage settings")
	settingsAddCmd := addCommand(settingsCmd, "add KEY VALUE", "Add a setting value")
	settingsGetCmd := addCommand(settingsCmd, "get KEY", "Get a setting")
	addCommand(settingsCmd, "ls", "List settings")
	settingsSetCmd := addCommand(settingsCmd, "set KEY VALUE", "Set a setting")
	settingsUnsetCmd := addCommand(settingsCmd, "unset KEY", "Unset a setting")
	carapace.Gen(settingsAddCmd).PositionalCompletion(actionSettings())
	carapace.Gen(settingsGetCmd).PositionalCompletion(actionSettings())
	carapace.Gen(settingsSetCmd).PositionalCompletion(actionSettings())
	carapace.Gen(settingsUnsetCmd).PositionalCompletion(actionSettings())

	shellCmd := addCommand(rootCmd, "shell TOOL@VERSION", "Set a tool version in the current shell")
	shellCmd.Flags().Bool("unset", false, "Remove a shell-specific version")
	carapace.Gen(shellCmd).PositionalAnyCompletion(actionTools())

	shellAliasCmd := addCommand(rootCmd, "shell-alias", "Manage shell aliases")
	shellAliasGetCmd := addCommand(shellAliasCmd, "get NAME", "Get a shell alias")
	addCommand(shellAliasCmd, "ls", "List shell aliases")
	shellAliasSetCmd := addCommand(shellAliasCmd, "set NAME VALUE", "Set a shell alias")
	shellAliasUnsetCmd := addCommand(shellAliasCmd, "unset NAME", "Unset a shell alias")
	carapace.Gen(shellAliasGetCmd).PositionalCompletion(actionShellAliases())
	carapace.Gen(shellAliasSetCmd).PositionalCompletion(actionShellAliases())
	carapace.Gen(shellAliasUnsetCmd).PositionalCompletion(actionShellAliases())

	syncCmd := addCommand(rootCmd, "sync", "Sync tool versions with language-specific files")
	addCommand(syncCmd, "node", "Sync Node.js versions")
	addCommand(syncCmd, "python", "Sync Python versions")
	addCommand(syncCmd, "ruby", "Sync Ruby versions")

	tasksCmd := addCommand(rootCmd, "tasks", "Manage tasks")
	tasksAddCmd := addCommand(tasksCmd, "add NAME", "Add a task")
	tasksDepsCmd := addCommand(tasksCmd, "deps TASK", "Show task dependencies")
	tasksEditCmd := addCommand(tasksCmd, "edit TASK", "Edit a task")
	tasksInfoCmd := addCommand(tasksCmd, "info TASK", "Show task information")
	tasksLsCmd := addCommand(tasksCmd, "ls", "List tasks")
	tasksRunCmd := addCommand(tasksCmd, "run TASK", "Run a task")
	addCommand(tasksCmd, "validate", "Validate tasks")
	tasksLsCmd.Flags().Bool("json", false, "Output in JSON format")
	carapace.Gen(tasksAddCmd).PositionalCompletion(actionTasks())
	carapace.Gen(tasksDepsCmd).PositionalCompletion(actionTasks())
	carapace.Gen(tasksEditCmd).PositionalCompletion(actionTasks())
	carapace.Gen(tasksInfoCmd).PositionalCompletion(actionTasks())
	carapace.Gen(tasksRunCmd).PositionalAnyCompletion(actionTasks())

	testToolCmd := addCommand(rootCmd, "test-tool TOOL", "Test a tool")
	carapace.Gen(testToolCmd).PositionalCompletion(actionTools())

	tokenCmd := addCommand(rootCmd, "token", "Manage tokens")
	addCommand(tokenCmd, "forgejo", "Manage Forgejo token")
	addCommand(tokenCmd, "github", "Manage GitHub token")
	addCommand(tokenCmd, "gitlab", "Manage GitLab token")

	toolCmd := addCommand(rootCmd, "tool TOOL", "Show tool metadata")
	carapace.Gen(toolCmd).PositionalCompletion(actionTools())

	toolAliasCmd := addCommand(rootCmd, "tool-alias", "Manage tool aliases")
	toolAliasGetCmd := addCommand(toolAliasCmd, "get ALIAS", "Get a tool alias")
	addCommand(toolAliasCmd, "ls", "List tool aliases")
	toolAliasSetCmd := addCommand(toolAliasCmd, "set ALIAS TOOL", "Set a tool alias")
	toolAliasUnsetCmd := addCommand(toolAliasCmd, "unset ALIAS", "Unset a tool alias")
	carapace.Gen(toolAliasGetCmd).PositionalCompletion(actionToolAliases())
	carapace.Gen(toolAliasSetCmd).PositionalCompletion(actionToolAliases())
	carapace.Gen(toolAliasUnsetCmd).PositionalCompletion(actionToolAliases())

	toolStubCmd := addCommand(rootCmd, "tool-stub TOOL", "Generate a tool stub")
	carapace.Gen(toolStubCmd).PositionalCompletion(actionTools())

	trustCmd := addCommand(rootCmd, "trust [CONFIG]", "Trust a mise config file")
	carapace.Gen(trustCmd).PositionalCompletion(carapace.ActionFiles(".toml"))

	uninstallCmd := addCommand(rootCmd, "uninstall TOOL", "Uninstall tools")
	uninstallCmd.Flags().BoolP("all", "a", false, "Uninstall all versions")
	carapace.Gen(uninstallCmd).PositionalAnyCompletion(actionInstalledTools())

	unsetCmd := addCommand(rootCmd, "unset KEY", "Unset an environment variable")
	carapace.Gen(unsetCmd).PositionalCompletion(actionEnvironmentVariables())

	untrustCmd := addCommand(rootCmd, "untrust [CONFIG]", "Untrust a mise config file")
	carapace.Gen(untrustCmd).PositionalCompletion(carapace.ActionFiles(".toml"))

	unuseCmd := addCommand(rootCmd, "unuse TOOL", "Remove a tool from config")
	unuseCmd.Flags().BoolP("global", "g", false, "Use the global config file")
	unuseCmd.Flags().StringP("path", "p", "", "Path to a config file or directory")
	carapace.Gen(unuseCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(".toml"),
	})
	carapace.Gen(unuseCmd).PositionalAnyCompletion(actionInstalledTools())

	upgradeCmd := addCommand(rootCmd, "upgrade [TOOL]", "Upgrade tools")
	upgradeCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	carapace.Gen(upgradeCmd).PositionalAnyCompletion(actionInstalledTools())

	useCmd := addCommand(rootCmd, "use [TOOL@VERSION]...", "Install a tool and add it to mise.toml")
	useCmd.Flags().StringP("env", "e", "", "Create/modify an environment-specific config file")
	useCmd.Flags().BoolP("force", "f", false, "Force reinstall even if already installed")
	useCmd.Flags().BoolP("global", "g", false, "Use the global config file")
	useCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	useCmd.Flags().BoolP("dry-run", "n", false, "Show changes without making them")
	useCmd.Flags().StringP("path", "p", "", "Path to a config file or directory")
	useCmd.Flags().String("before", "", "Only install versions released before this date")
	useCmd.Flags().Bool("dry-run-code", false, "Exit with code 1 if there are changes to make")
	useCmd.Flags().Bool("fuzzy", false, "Save fuzzy version to config file")
	useCmd.Flags().Bool("pin", false, "Save exact version to config file")
	useCmd.Flags().Bool("raw", false, "Directly pipe stdin/stdout/stderr from plugin to user")
	useCmd.Flags().StringSlice("remove", nil, "Remove plugin(s) from config file")
	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"env":    actionEnvironments(),
		"path":   carapace.ActionFiles(".toml"),
		"remove": actionInstalledTools(),
	})
	carapace.Gen(useCmd).PositionalAnyCompletion(actionTools())

	addCommand(rootCmd, "version", "Print mise version")

	watchCmd := addCommand(rootCmd, "watch [TASK]", "Watch task(s)")
	carapace.Gen(watchCmd).PositionalAnyCompletion(actionTasks())

	whereCmd := addCommand(rootCmd, "where TOOL", "Show where a tool is installed")
	carapace.Gen(whereCmd).PositionalCompletion(actionInstalledTools())

	whichCmd := addCommand(rootCmd, "which BIN", "Show which binary mise will use")
	carapace.Gen(whichCmd).PositionalCompletion(bridge.ActionCarapaceBin().Split())
}

func addCommand(parent *cobra.Command, use string, short string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	parent.AddCommand(cmd)
	return cmd
}

func actionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh")
}

func actionCompletionShells() carapace.Action {
	return carapace.ActionValues("bash", "elvish", "fish", "nushell", "zsh")
}

func actionEnvironments() carapace.Action {
	return carapace.ActionValues("development", "dev", "test", "staging", "production", "prod", "local")
}

func actionTools() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("mise", "registry", "--complete")(func(output []byte, err error) carapace.Action {
			if err == nil {
				if vals := describedValuesFromColonLines(output); len(vals) > 0 {
					return carapace.ActionValuesDescribed(vals...)
				}
			}
			return carapace.ActionExecCommandE("mise", "registry", "--json")(func(output []byte, err error) carapace.Action {
				if err == nil {
					if vals := describedValuesFromJSON(output, []string{"name", "id", "short", "tool", "plugin"}, []string{"description", "desc", "backend"}); len(vals) > 0 {
						return carapace.ActionValuesDescribed(vals...)
					}
				}
				return commonTools()
			})
		})
	})
}

func actionInstalledTools() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("mise", "ls", "--json")(func(output []byte, err error) carapace.Action {
			if err == nil {
				if vals := describedValuesFromJSON(output, []string{"name", "tool", "plugin"}, []string{"version", "requested_version", "install_path"}); len(vals) > 0 {
					return carapace.ActionValuesDescribed(vals...)
				}
			}
			return actionTools()
		})
	})
}

func actionTasks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("mise", "tasks", "ls", "--complete")(func(output []byte, err error) carapace.Action {
			if err == nil {
				if vals := describedValuesFromColonLines(output); len(vals) > 0 {
					return carapace.ActionValuesDescribed(vals...).FilterArgs()
				}
			}
			return carapace.ActionExecCommandE("mise", "tasks", "ls", "--json")(func(output []byte, err error) carapace.Action {
				if err == nil {
					if vals := describedValuesFromJSON(output, []string{"name", "task"}, []string{"description", "desc", "source"}); len(vals) > 0 {
						return carapace.ActionValuesDescribed(vals...).FilterArgs()
					}
				}
				return carapace.ActionExecCommandE("mise", "tasks", "ls")(func(output []byte, err error) carapace.Action {
					if err == nil {
						if vals := describedValuesFromColumns(output); len(vals) > 0 {
							return carapace.ActionValuesDescribed(vals...).FilterArgs()
						}
					}
					return carapace.ActionValues()
				})
			})
		})
	})
}

func actionPlugins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("mise", "plugins", "ls")(func(output []byte, err error) carapace.Action {
			if err == nil {
				if vals := describedValuesFromColumns(output); len(vals) > 0 {
					return carapace.ActionValuesDescribed(vals...)
				}
			}
			return carapace.ActionValues()
		})
	})
}

func actionRemotePlugins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("mise", "plugins", "ls-remote")(func(output []byte, err error) carapace.Action {
			if err == nil {
				if vals := describedValuesFromColumns(output); len(vals) > 0 {
					return carapace.ActionValuesDescribed(vals...)
				}
			}
			return commonTools()
		})
	})
}

func actionSettings() carapace.Action {
	return carapace.ActionValues(
		"activate_aggressive",
		"always_keep_download",
		"always_keep_install",
		"auto_install",
		"cache_prune_age",
		"experimental",
		"fetch_remote_versions_cache",
		"go_default_packages_file",
		"jobs",
		"legacy_version_file",
		"lockfile",
		"node_default_packages_file",
		"not_found_auto_install",
		"paranoid",
		"pipx_uvx",
		"plugin_autoupdate_last_check_duration",
		"python_compile",
		"python_default_packages_file",
		"raw",
		"status",
		"task_output",
		"task_run_auto_install",
		"trusted_config_paths",
		"yes",
	)
}

func actionEnvironmentVariables() carapace.Action {
	return carapace.ActionExecCommandE("mise", "set")(func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}
		return carapace.ActionValuesDescribed(describedValuesFromColumns(output)...)
	})
}

func actionShellAliases() carapace.Action {
	return carapace.ActionExecCommandE("mise", "shell-alias", "ls")(func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}
		return carapace.ActionValuesDescribed(describedValuesFromColumns(output)...)
	})
}

func actionToolAliases() carapace.Action {
	return carapace.ActionExecCommandE("mise", "tool-alias", "ls")(func(output []byte, err error) carapace.Action {
		if err != nil {
			return carapace.ActionValues()
		}
		return carapace.ActionValuesDescribed(describedValuesFromColumns(output)...)
	})
}

func commonTools() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bun", "JavaScript runtime",
		"deno", "JavaScript runtime",
		"elixir", "Elixir language",
		"erlang", "Erlang runtime",
		"go", "Go language",
		"java", "Java runtime",
		"node", "Node.js runtime",
		"python", "Python language",
		"ruby", "Ruby language",
		"rust", "Rust language",
		"swift", "Swift language",
		"zig", "Zig language",
	)
}

func describedValuesFromJSON(output []byte, nameKeys []string, descriptionKeys []string) []string {
	var data any
	if err := json.Unmarshal(output, &data); err != nil {
		return nil
	}

	seen := make(map[string]bool)
	vals := make([]string, 0)
	for _, entry := range jsonEntries(data) {
		name := firstString(entry, nameKeys...)
		description := firstString(entry, descriptionKeys...)
		if name == "" {
			name = firstString(entry, "key")
		}
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		vals = append(vals, name, description)
	}
	sortPairs(vals)
	return vals
}

func jsonEntries(data any) []map[string]any {
	switch value := data.(type) {
	case []any:
		entries := make([]map[string]any, 0, len(value))
		for _, item := range value {
			if entry, ok := item.(map[string]any); ok {
				entries = append(entries, entry)
			}
		}
		return entries
	case map[string]any:
		entries := make([]map[string]any, 0, len(value))
		for key, item := range value {
			if entry, ok := item.(map[string]any); ok {
				if _, ok := entry["key"]; !ok {
					entry["key"] = key
				}
				entries = append(entries, entry)
				continue
			}
			entries = append(entries, map[string]any{"key": key, "description": fmt.Sprint(item)})
		}
		return entries
	default:
		return nil
	}
}

func firstString(entry map[string]any, keys ...string) string {
	for _, key := range keys {
		if value, ok := entry[key]; ok {
			switch typed := value.(type) {
			case string:
				return typed
			case float64, bool:
				return fmt.Sprint(typed)
			case []any:
				parts := make([]string, 0, len(typed))
				for _, item := range typed {
					parts = append(parts, fmt.Sprint(item))
				}
				return strings.Join(parts, ", ")
			}
		}
	}
	return ""
}

func describedValuesFromColumns(output []byte) []string {
	seen := make(map[string]bool)
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(strings.TrimSuffix(line, "\r"))
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		name := fields[0]
		if seen[name] || skipColumnName(name) {
			continue
		}
		seen[name] = true
		description := strings.TrimSpace(strings.TrimPrefix(line, name))
		vals = append(vals, name, description)
	}
	sortPairs(vals)
	return vals
}

func describedValuesFromColonLines(output []byte) []string {
	seen := make(map[string]bool)
	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(strings.TrimSuffix(line, "\r"))
		if line == "" {
			continue
		}
		name, description, ok := splitEscapedColon(line)
		if !ok || name == "" || seen[name] {
			continue
		}
		seen[name] = true
		vals = append(vals, name, description)
	}
	sortPairs(vals)
	return vals
}

func splitEscapedColon(line string) (string, string, bool) {
	escaped := false
	for i, r := range line {
		if escaped {
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}
		if r == ':' {
			return unescapeCompletionValue(line[:i]), unescapeCompletionValue(line[i+1:]), true
		}
	}
	return "", "", false
}

func unescapeCompletionValue(value string) string {
	return strings.ReplaceAll(value, `\:`, ":")
}

func skipColumnName(name string) bool {
	switch strings.ToLower(name) {
	case "name", "plugin", "task", "tool", "version", "source", "description":
		return true
	default:
		return strings.HasPrefix(name, "-")
	}
}

func sortPairs(vals []string) {
	pairs := make([][2]string, 0, len(vals)/2)
	for i := 0; i+1 < len(vals); i += 2 {
		pairs = append(pairs, [2]string{vals[i], vals[i+1]})
	}
	sort.SliceStable(pairs, func(i int, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	for i, pair := range pairs {
		vals[i*2] = pair[0]
		vals[i*2+1] = pair[1]
	}
}
