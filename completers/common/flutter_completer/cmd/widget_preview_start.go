package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var widget_preview_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the widget preview environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(widget_preview_startCmd).Standalone()

	widget_preview_startCmd.Flags().Bool("devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	widget_preview_startCmd.Flags().String("devtools-server-address", "", "When this value is provided, the Flutter tool will not spin up a new DevTools server instance, and will instead use the one provided at the given address. Ignored if \"--no-devtools\" is specified.")
	widget_preview_startCmd.Flags().Bool("disable-dtd-service-uuid", false, "Disables the addition of a UUID to the widget preview DTD service and stream.")
	widget_preview_startCmd.Flags().String("dtd-url", "", "The address of an existing Dart Tooling Daemon instance to be used by the Flutter CLI.")
	widget_preview_startCmd.Flags().Bool("headless", false, "Launches Chrome in headless mode for testing.")
	widget_preview_startCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	widget_preview_startCmd.Flags().Bool("launch-previewer", false, "Launches the widget preview environment.")
	widget_preview_startCmd.Flags().Bool("legacy-preview-detection", false, "Enables the legacy preview detection mechanism that uses package:analyzer instead of LSP.")
	widget_preview_startCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	widget_preview_startCmd.Flags().Bool("no-devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	widget_preview_startCmd.Flags().Bool("no-disable-dtd-service-uuid", false, "Disables the addition of a UUID to the widget preview DTD service and stream.")
	widget_preview_startCmd.Flags().Bool("no-headless", false, "Launches Chrome in headless mode for testing.")
	widget_preview_startCmd.Flags().Bool("no-launch-previewer", false, "Launches the widget preview environment.")
	widget_preview_startCmd.Flags().Bool("no-legacy-preview-detection", false, "Enables the legacy preview detection mechanism that uses package:analyzer instead of LSP.")
	widget_preview_startCmd.Flags().Bool("no-offline", false, "When \"flutter pub get\" is run by the create command, this indicates whether to run it in offline mode or not. In offline mode, it will need to have all dependencies already available in the pub cache to succeed.")
	widget_preview_startCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" after the project has been created.")
	widget_preview_startCmd.Flags().Bool("no-web-server", false, "Serve the widget preview environment using the web-server device instead of the browser.")
	widget_preview_startCmd.Flags().Bool("offline", false, "When \"flutter pub get\" is run by the create command, this indicates whether to run it in offline mode or not. In offline mode, it will need to have all dependencies already available in the pub cache to succeed.")
	widget_preview_startCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" after the project has been created.")
	widget_preview_startCmd.Flags().String("scaffold-output-dir", "", "Generated the widget preview environment scaffolding at a given location for testing purposes.")
	widget_preview_startCmd.Flags().Bool("web-server", false, "Serve the widget preview environment using the web-server device instead of the browser.")
	widget_previewCmd.AddCommand(widget_preview_startCmd)
}
