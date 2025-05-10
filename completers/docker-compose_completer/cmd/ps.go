package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps [OPTIONS] [SERVICE...]",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(psCmd).Standalone()

	psCmd.Flags().BoolP("all", "a", false, "Show all stopped containers (including those created by the run command)")
	psCmd.Flags().String("filter", "", "Filter services by a property (supported filters: status)")
	psCmd.Flags().String("format", "", "Format output using a custom template:")
	psCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	psCmd.Flags().Bool("orphans", false, "Include orphaned services (not declared by project)")
	psCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	psCmd.Flags().Bool("services", false, "Display services")
	psCmd.Flags().StringSlice("status", nil, "Filter services by status. Values: [paused | restarting | removing | running | dead | created | exited]")
	rootCmd.AddCommand(psCmd)

	carapace.Gen(psCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionValues("status"),
		"format": carapace.ActionValues("table", "json"),
		"status": carapace.ActionValues("paused", "restarting", "removing", "running", "dead", "created", "exited").StyleF(style.ForKeyword),
	})

	carapace.Gen(psCmd).PositionalAnyCompletion(
		action.ActionServices(psCmd).FilterArgs(),
	)
}
