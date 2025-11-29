package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_upCmd = &cobra.Command{
	Use:   "up [service]...",
	Short: "Starts process manager with specified services. If no services are listed, starts the process manager with all the services in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_upCmd).Standalone()

	services_upCmd.Flags().BoolP("background", "b", false, "run service in background")
	services_upCmd.Flags().StringSlice("pcflags", nil, "pass flags directly to process compose")
	services_upCmd.Flags().StringP("pcport", "p", "", "specify the port for process-compose to use. You can also set the pcport by exporting DEVBOX_PC_PORT_NUM")
	services_upCmd.Flags().String("process-compose-file", "", "path to process compose file or directory containing process compose-file.yaml|yml. Default is directory containing devbox.json")
	servicesCmd.AddCommand(services_upCmd)

	// TODO flag completion
	carapace.Gen(services_upCmd).FlagCompletion(carapace.ActionMap{
		"process-compose-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
