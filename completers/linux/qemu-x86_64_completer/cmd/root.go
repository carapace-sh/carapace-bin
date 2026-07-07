package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-x86_64_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-x86_64",
	Short: "QEMU linux-user emulator (x86_64)",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("0", "0", "", "forces target process argv[0] to be 'argv0'")
	rootCmd.Flags().StringS("B", "B", "", "set guest_base address to 'address'")
	rootCmd.Flags().StringS("D", "D", "", "write logs to 'logfile' (default stderr)")
	rootCmd.Flags().StringS("E", "E", "", "sets targets environment variable (var=value)")
	rootCmd.Flags().StringS("L", "L", "", "set the elf interpreter prefix to 'path'")
	rootCmd.Flags().StringS("R", "R", "", "reserve 'size' bytes for guest virtual address space")
	rootCmd.Flags().StringS("U", "U", "", "unsets targets environment variable")
	rootCmd.Flags().StringS("cpu", "cpu", "", "select CPU (-cpu help for list)")
	rootCmd.Flags().StringS("d", "d", "", "enable logging of specified items (use '-d help' for a list)")
	rootCmd.Flags().StringS("dfilter", "dfilter", "", "filter logging based on address range")
	rootCmd.Flags().StringS("g", "g", "", "wait gdb connection to 'port'")
	rootCmd.Flags().BoolS("help", "help", false, "print this help")
	rootCmd.Flags().BoolS("jitdump", "jitdump", false, "Generate a jit-${pid}.dump file for perf")
	rootCmd.Flags().BoolS("one-insn-per-tb", "one-insn-per-tb", false, "run with one guest instruction per emulated TB")
	rootCmd.Flags().BoolS("perfmap", "perfmap", false, "Generate a /tmp/perf-${pid}.map file for perf")
	rootCmd.Flags().String("plugin", "", "[file=]<file>[,<argname>=<argvalue>]")
	rootCmd.Flags().StringS("r", "r", "", "set qemu uname release string to 'uname'")
	rootCmd.Flags().StringS("s", "s", "", "set the stack size to 'size' bytes")
	rootCmd.Flags().String("seed", "", "Seed for pseudo-random number generator")
	rootCmd.Flags().Bool("strace", false, "log system calls")
	rootCmd.Flags().StringS("t", "t", "", "map target rt signals [tsig,tsig+n) to [hsig,hsig+n)")
	rootCmd.Flags().String("tb-size", "", "TCG translation block cache size")
	rootCmd.Flags().String("trace", "", "[[enable=]<pattern>][,events=<file>][,file=<file>]")
	rootCmd.Flags().Bool("version", false, "display version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionFiles(),
		"L": carapace.ActionDirectories(),
		"cpu": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					return action.ActionCpuModels(rootCmd).Suffix(",").NoSpace(',')
				})
			}
			return qemu.ActionCpuFeatures().Prefix("+").Suffix(",").NoSpace(',')
		}),
		"d":      carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action { return qemu.ActionDebugItems() }),
		"plugin": carapace.ActionFiles(".so"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
