package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_upCmd = &cobra.Command{
	Use:   "up [service]...",
	Short: "Starts process manager with specified services. If no services are listed, starts the process manager with all the services in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_upCmd).Standalone()

	global_services_upCmd.Flags().BoolP("background", "b", false, "run service in background")
	global_services_upCmd.Flags().StringSlice("pcflags", nil, "pass flags directly to process compose")
	global_services_upCmd.Flags().StringP("pcport", "p", "", "specify the port for process-compose to use. You can also set the pcport by exporting DEVBOX_PC_PORT_NUM")
	global_services_upCmd.Flags().String("process-compose-file", "", "path to process compose file or directory containing process compose-file.yaml|yml. Default is directory containing devbox.json")
	global_servicesCmd.AddCommand(global_services_upCmd)

	// TODO complete flags
	carapace.Gen(global_services_upCmd).FlagCompletion(carapace.ActionMap{
		"process-compose-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
