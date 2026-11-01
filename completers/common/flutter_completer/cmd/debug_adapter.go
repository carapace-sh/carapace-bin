package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_adapterCmd = &cobra.Command{
	Use:     "debug-adapter",
	Short:   "Run a Debug Adapter Protocol (DAP) server to communicate with the Flutter tool",
	Aliases: []string{"debug_adapter"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_adapterCmd).Standalone()

	debug_adapterCmd.Flags().Bool("dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	debug_adapterCmd.Flags().Int("dds-port", 0, "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port. Specifying port 0 (the default) will find a random free port.")
	debug_adapterCmd.Flags().Bool("disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	debug_adapterCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	debug_adapterCmd.Flags().Bool("ipv6", false, "Binds to IPv6 localhost instead of IPv4 when the flutter tool forwards the host port to a device port.")
	debug_adapterCmd.Flags().Bool("no-dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	debug_adapterCmd.Flags().Bool("no-disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	debug_adapterCmd.Flags().Bool("no-test", false, "Whether to use the \"flutter test\" debug adapter to run tests and emit custom events for test progress/results.")
	debug_adapterCmd.Flags().Bool("test", false, "Whether to use the \"flutter test\" debug adapter to run tests and emit custom events for test progress/results.")
	rootCmd.AddCommand(debug_adapterCmd)
}
