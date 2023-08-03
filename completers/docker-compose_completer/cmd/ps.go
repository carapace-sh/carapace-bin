package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
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
	psCmd.Flags().String("filter", "", "Filter services by a property (supported filters: status).")
	psCmd.Flags().String("format", "table", "Format the output. Values: [table | json]")
	psCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	psCmd.Flags().Bool("services", false, "Display services")
	psCmd.Flags().StringArray("status", []string{}, "Filter services by status. Values: [paused | restarting | removing | running | dead | created | exited]")
	rootCmd.AddCommand(psCmd)

	carapace.Gen(psCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
		"status": carapace.ActionValues("paused", "restarting", "removing", "running", "dead", "created", "exited"),
	})

	carapace.Gen(psCmd).PositionalAnyCompletion(
		action.ActionServices(psCmd).FilterArgs(),
	)
}
