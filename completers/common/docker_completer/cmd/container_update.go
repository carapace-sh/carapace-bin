package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var container_updateCmd = &cobra.Command{
	Use:   "update [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Update configuration of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_updateCmd).Standalone()

	container_updateCmd.Flags().Uint16("blkio-weight", 0, "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)")
	container_updateCmd.Flags().Int64("cpu-period", 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	container_updateCmd.Flags().Int64("cpu-quota", 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	container_updateCmd.Flags().Int64("cpu-rt-period", 0, "Limit the CPU real-time period in microseconds")
	container_updateCmd.Flags().Int64("cpu-rt-runtime", 0, "Limit the CPU real-time runtime in microseconds")
	container_updateCmd.Flags().Int64P("cpu-shares", "c", 0, "CPU shares (relative weight)")
	container_updateCmd.Flags().String("cpus", "", "Number of CPUs")
	container_updateCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_updateCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	container_updateCmd.Flags().String("kernel-memory", "", "Kernel memory limit (deprecated)")
	container_updateCmd.Flags().StringP("memory", "m", "", "Memory limit")
	container_updateCmd.Flags().String("memory-reservation", "", "Memory soft limit")
	container_updateCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: -1 to enable unlimited swap")
	container_updateCmd.Flags().Int64("pids-limit", 0, "Tune container pids limit (set -1 for unlimited)")
	container_updateCmd.Flags().String("restart", "", "Restart policy to apply when a container exits")
	containerCmd.AddCommand(container_updateCmd)

	carapace.Gen(container_updateCmd).FlagCompletion(carapace.ActionMap{
		"restart": carapace.ActionValues("always", "no", "on-failure", "unless-stopped").StyleF(style.ForKeyword),
	})

	carapace.Gen(container_updateCmd).PositionalAnyCompletion(docker.ActionContainers())
}
