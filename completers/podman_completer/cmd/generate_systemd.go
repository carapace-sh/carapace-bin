package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_systemdCmd = &cobra.Command{
	Use:   "systemd [options] {CONTAINER|POD}",
	Short: "[DEPRECATED] Generate systemd units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_systemdCmd).Standalone()

	generate_systemdCmd.Flags().StringSlice("after", []string{}, "Add dependencies order to the generated unit file")
	generate_systemdCmd.Flags().String("container-prefix", "", "Systemd unit name prefix for containers")
	generate_systemdCmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables to the systemd unit files")
	generate_systemdCmd.Flags().BoolP("files", "f", false, "Generate .service files instead of printing to stdout")
	generate_systemdCmd.Flags().String("format", "", "Print the created units in specified format (json)")
	generate_systemdCmd.Flags().BoolP("name", "n", false, "Use container/pod names instead of IDs")
	generate_systemdCmd.Flags().Bool("new", false, "Create a new container or pod instead of starting an existing one")
	generate_systemdCmd.Flags().Bool("no-header", false, "Skip header generation")
	generate_systemdCmd.Flags().String("pod-prefix", "", "Systemd unit name prefix for pods")
	generate_systemdCmd.Flags().StringSlice("requires", []string{}, "Similar to wants, but declares stronger requirement dependencies")
	generate_systemdCmd.Flags().String("restart-policy", "", "Systemd restart-policy")
	generate_systemdCmd.Flags().String("restart-sec", "", "Systemd restart-sec")
	generate_systemdCmd.Flags().String("separator", "", "Systemd unit name separator between name/id and prefix")
	generate_systemdCmd.Flags().String("start-timeout", "", "Start timeout override")
	generate_systemdCmd.Flags().String("stop-timeout", "", "Stop timeout override")
	generate_systemdCmd.Flags().Bool("template", false, "Make it a template file and use %i and %I specifiers. Working only for containers")
	generate_systemdCmd.Flags().StringP("time", "t", "", "Backwards alias for --stop-timeout")
	generate_systemdCmd.Flags().StringSlice("wants", []string{}, "Add (weak) requirement dependencies to the generated unit file")
	generate_systemdCmd.Flag("time").Hidden = true
	generateCmd.AddCommand(generate_systemdCmd)
}
