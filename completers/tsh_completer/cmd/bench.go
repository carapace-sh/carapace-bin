package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:    "bench",
	Short:  "Run Teleport benchmark tests.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchCmd).Standalone()

	benchCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	benchCmd.Flags().String("duration", "", "Test duration")
	benchCmd.Flags().Bool("export", false, "Export the latency profile")
	benchCmd.Flags().Bool("no-export", false, "Export the latency profile")
	benchCmd.Flags().String("path", "", "Directory to save the latency profile to, default path is the current directory")
	benchCmd.Flags().String("rate", "", "Requests per second rate")
	benchCmd.Flags().String("scale", "", "Value scale in which to scale the recorded values")
	benchCmd.Flags().String("ticks", "", "Ticks per half distance")
	benchCmd.Flag("no-export").Hidden = true
	rootCmd.AddCommand(benchCmd)

	carapace.Gen(benchCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"path":    carapace.ActionDirectories(),
	})
}
