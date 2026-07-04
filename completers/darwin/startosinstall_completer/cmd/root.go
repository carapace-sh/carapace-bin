package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "startosinstall",
	Short: "Start macOS OS installation",
	Long:  "https://support.apple.com/en-us/101896",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("agreetolicense", false, "agree to the license agreement")
	rootCmd.Flags().Bool("allowremoval", false, "force removal of files on the root of the system disk")
	rootCmd.Flags().String("bridgeos-pkg", "", "supply a custom BridgeOS package")
	rootCmd.Flags().Bool("converttoapfs", false, "convert the existing filesystem to APFS (deprecated)")
	rootCmd.Flags().Bool("eraseinstall", false, "erase all volumes in the APFS container and install to a new volume")
	rootCmd.Flags().Bool("forcequitapps", false, "forcefully quit applications on restart")
	rootCmd.Flags().String("installpackage", "", "path to a package to install after the OS installation")
	rootCmd.Flags().Bool("license", false, "print the user license agreement only")
	rootCmd.Flags().String("newvolumename", "", "name of the volume to be created with --eraseinstall")
	rootCmd.Flags().Bool("nointeraction", false, "run without prompts")
	rootCmd.Flags().String("pidtosignal", "", "send SIGUSR1 to this PID upon completion of the prepare phase")
	rootCmd.Flags().Bool("preservecontainer", false, "preserve other volumes in the APFS container when using --eraseinstall")
	rootCmd.Flags().Int("rebootdelay", 0, "delay the reboot after preparation completes, in seconds (max: 300)")
	rootCmd.Flags().Bool("usage", false, "display all parameters available")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bridgeos-pkg":   carapace.ActionFiles(),
		"installpackage": carapace.ActionFiles(),
	})
}
