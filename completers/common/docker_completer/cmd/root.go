package cmd

import (
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker [OPTIONS] COMMAND [ARG...]",
	Short: "A self-sufficient runtime for containers",
	Long:  "https://docs.docker.com/compose/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "common", Title: "Common Commands"},
		&cobra.Group{ID: "management", Title: "Management Commands"},
		&cobra.Group{ID: "swarm", Title: "Swarm Commands"},
		&cobra.Group{ID: "legacy", Title: "Legacy Commands"},
	)

	rootCmd.Flags().String("config", "~/.docker", "Location of client config files")
	rootCmd.Flags().StringP("context", "c", "", "Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with \"docker context use\")")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket(s) to connect to")
	rootCmd.Flags().StringP("log-level", "l", "info", "Set the logging level (\"debug\", \"info\", \"warn\", \"error\", \"fatal\")")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "~/.docker/ca.pem", "Trust certs signed only by this CA")
	rootCmd.Flags().String("tlscert", "~/.docker/cert.pem", "Path to TLS certificate file")
	rootCmd.Flags().String("tlskey", "~/.docker/key.pem", "Path to TLS key file")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":    carapace.ActionFiles(),
		"log-level": carapace.ActionValues("debug", "info", "warn", "error", "fatal").StyleF(style.ForLogLevel),
		"tlscacert": carapace.ActionFiles(),
		"tlscert":   carapace.ActionFiles(),
		"tlskey":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		r := regexp.MustCompile(`^  (?P<plugin>[^ ]+)[*] +(?P<description>.*)$`)
		if output, err := (carapace.Context{}).Command("docker", "--help").Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); len(matches) > 2 {
					name := matches[1]
					description := matches[2]
					pluginCmd := &cobra.Command{
						Use:                name,
						Short:              description,
						GroupID:            "management",
						Run:                func(cmd *cobra.Command, args []string) {},
						DisableFlagParsing: true,
					}

					if d := completers.Description("docker-" + name); d != "" {
						pluginCmd.Short = d
					}

					carapace.Gen(pluginCmd).PositionalAnyCompletion(
						bridge.ActionCarapaceBin("docker-" + name),
					)

					rootCmd.AddCommand(pluginCmd)
				}
			}
		}

		addRootAlias("run", "common", false, container_runCmd)

		hideLegacyCommands := os.Getenv("DOCKER_HIDE_LEGACY_COMMANDS") == "1"
		addRootAlias("attach", "legacy", hideLegacyCommands, container_attachCmd)
		addRootAlias("commit", "legacy", hideLegacyCommands, container_commitCmd)
		addRootAlias("cp", "legacy", hideLegacyCommands, container_cpCmd)
		addRootAlias("create", "legacy", hideLegacyCommands, container_createCmd)
		addRootAlias("diff", "legacy", hideLegacyCommands, container_diffCmd)
		addRootAlias("events", "legacy", hideLegacyCommands, system_eventsCmd)
		addRootAlias("export", "legacy", hideLegacyCommands, container_exportCmd)
		addRootAlias("history", "legacy", hideLegacyCommands, image_historyCmd)
		addRootAlias("import", "legacy", hideLegacyCommands, image_importCmd)
		addRootAlias("kill", "legacy", hideLegacyCommands, container_killCmd)
		addRootAlias("load", "legacy", hideLegacyCommands, image_loadCmd)
		addRootAlias("logs", "legacy", hideLegacyCommands, container_logsCmd)
		addRootAlias("pause", "legacy", hideLegacyCommands, container_pauseCmd)
		addRootAlias("port", "legacy", hideLegacyCommands, container_portCmd)
		addRootAlias("rename", "legacy", hideLegacyCommands, container_renameCmd)
		addRootAlias("restart", "legacy", hideLegacyCommands, container_restartCmd)
		addRootAlias("rm", "legacy", hideLegacyCommands, container_rmCmd)
		addRootAlias("rmi", "legacy", hideLegacyCommands, image_rmCmd)
		addRootAlias("save", "legacy", hideLegacyCommands, image_saveCmd)
		addRootAlias("start", "legacy", hideLegacyCommands, container_startCmd)
		addRootAlias("stats", "legacy", hideLegacyCommands, container_statsCmd)
		addRootAlias("stop", "legacy", hideLegacyCommands, container_stopCmd)
		addRootAlias("tag", "legacy", hideLegacyCommands, image_tagCmd)
		addRootAlias("top", "legacy", hideLegacyCommands, container_topCmd)
		addRootAlias("unpause", "legacy", hideLegacyCommands, container_unpauseCmd)
		addRootAlias("update", "legacy", hideLegacyCommands, container_updateCmd)
		addRootAlias("wait", "legacy", hideLegacyCommands, container_waitCmd)
	})
}

func addRootAlias(use string, groupID string, hidden bool, cmd *cobra.Command) {
	aliasCmd := &cobra.Command{
		Use:                use,
		Short:              cmd.Short,
		GroupID:            groupID,
		Hidden:             hidden,
		Run:                func(cmd *cobra.Command, args []string) {},
		DisableFlagParsing: true,
	}

	carapace.Gen(aliasCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			command := []string{cmd.Name()}
			for cmd.HasParent() {
				cmd = cmd.Parent()
				command = append(command, cmd.Name())
			}
			slices.Reverse(command)
			return bridge.ActionCarapaceBin(command...)
		}),
	)

	rootCmd.AddCommand(aliasCmd)
}
