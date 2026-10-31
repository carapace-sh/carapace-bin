package action

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

const defaultConfigFile = ".pre-commit-config.yaml"

// ActionHookTypes completes the git hook types supported by pre-commit
func ActionHookTypes() carapace.Action {
	return carapace.ActionValues(
		"commit-msg",
		"post-checkout",
		"post-commit",
		"post-merge",
		"post-rewrite",
		"pre-commit",
		"pre-merge-commit",
		"pre-push",
		"pre-rebase",
		"prepare-commit-msg",
	).Tag("hook types")
}

// ActionStages completes the hook stages pre-commit can run in
func ActionStages() carapace.Action {
	return carapace.ActionValues(
		"commit-msg",
		"manual",
		"post-checkout",
		"post-commit",
		"post-merge",
		"post-rewrite",
		"pre-commit",
		"pre-merge-commit",
		"pre-push",
		"pre-rebase",
		"prepare-commit-msg",
	).Tag("hook stages")
}

// ActionCommands completes pre-commit subcommands
func ActionCommands() carapace.Action {
	return carapace.ActionValues(
		"autoupdate",
		"clean",
		"gc",
		"hazmat",
		"help",
		"hook-impl",
		"init-templatedir",
		"install",
		"install-hooks",
		"migrate-config",
		"run",
		"sample-config",
		"try-repo",
		"uninstall",
		"validate-config",
		"validate-manifest",
	).Tag("commands")
}

// ActionColors completes the values of the --color flag
func ActionColors() carapace.Action {
	return carapace.ActionValues("always", "auto", "never").Tag("colors")
}

type config struct {
	Repos []struct {
		Hooks []struct {
			ID string `yaml:"id"`
		} `yaml:"hooks"`
	} `yaml:"repos"`
}

// ActionHookIDs completes hook ids of the configured pre-commit config
func ActionHookIDs(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		configFile := defaultConfigFile
		if flag := cmd.Flag("config"); flag != nil && flag.Value.String() != "" {
			configFile = flag.Value.String()
		}

		content, err := os.ReadFile(configFile)
		if err != nil {
			return carapace.ActionValues()
		}

		var cfg config
		if err := yaml.Unmarshal(content, &cfg); err != nil {
			return carapace.ActionValues()
		}

		ids := make([]string, 0)
		for _, repo := range cfg.Repos {
			for _, hook := range repo.Hooks {
				ids = append(ids, hook.ID)
			}
		}
		return carapace.ActionValues(ids...).Unique()
	}).Tag("hook ids")
}
