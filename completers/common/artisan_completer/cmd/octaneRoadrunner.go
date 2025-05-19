package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneRoadrunnerCmd = &cobra.Command{
	Use:   "octane:roadrunner [--host [HOST]] [--port [PORT]] [--rpc-host [RPC-HOST]] [--rpc-port [RPC-PORT]] [--workers [WORKERS]] [--max-requests [MAX-REQUESTS]] [--rr-config [RR-CONFIG]] [--watch] [--poll] [--log-level [LOG-LEVEL]]",
	Short: "Start the Octane RoadRunner server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneRoadrunnerCmd).Standalone()

	octaneRoadrunnerCmd.Flags().String("host", "", "The IP address the server should bind to")
	octaneRoadrunnerCmd.Flags().String("log-level", "", "Log messages at or above the specified log level")
	octaneRoadrunnerCmd.Flags().String("max-requests", "", "The number of requests to process before reloading the server")
	octaneRoadrunnerCmd.Flags().Bool("poll", false, "Use file system polling while watching in order to watch files over a network")
	octaneRoadrunnerCmd.Flags().String("port", "", "The port the server should be available on")
	octaneRoadrunnerCmd.Flags().String("rpc-host", "", "The RPC IP address the server should bind to")
	octaneRoadrunnerCmd.Flags().String("rpc-port", "", "The RPC port the server should be available on")
	octaneRoadrunnerCmd.Flags().String("rr-config", "", "The path to the RoadRunner .rr.yaml file")
	octaneRoadrunnerCmd.Flags().Bool("watch", false, "Automatically reload the server when the application is modified")
	octaneRoadrunnerCmd.Flags().String("workers", "", "The number of workers that should be available to handle requests")
	rootCmd.AddCommand(octaneRoadrunnerCmd)
}
