package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocryptfs",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("aessiv", "aessiv", false, "Use AES-SIV encryption (with -init)")
	rootCmd.Flags().BoolS("allow_other", "allow_other", false, "Allow other users to access the mount")
	rootCmd.Flags().BoolS("config", "config", false, "Custom path to config file")
	rootCmd.Flags().BoolS("ctlsock", "ctlsock", false, "Create control socket at location")
	rootCmd.Flags().BoolS("extpass", "extpass", false, "Call external program to prompt for the password")
	rootCmd.Flags().BoolS("fg", "fg", false, "Stay in the foreground")
	rootCmd.Flags().BoolS("fsck", "fsck", false, "Check filesystem integrity")
	rootCmd.Flags().BoolS("fusedebug", "fusedebug", false, "Debug FUSE calls")
	rootCmd.Flags().BoolN("help", "h", false, "Short help text")
	rootCmd.Flags().BoolS("hh", "hh", false, "Long help text with all options")
	rootCmd.Flags().BoolN("idle", "i", false, "Unmount automatically after specified idle duration")
	rootCmd.Flags().BoolS("info", "info", false, "Display information about encrypted directory")
	rootCmd.Flags().BoolS("init", "init", false, "Initialize encrypted directory")
	rootCmd.Flags().BoolS("masterkey", "masterkey", false, "Mount with explicit master key instead of password")
	rootCmd.Flags().BoolS("nonempty", "nonempty", false, "Allow mounting over non-empty directory")
	rootCmd.Flags().BoolS("nosyslog", "nosyslog", false, "Do not redirect log messages to syslog")
	rootCmd.Flags().BoolS("passfile", "passfile", false, "Read password from plain text file(s)")
	rootCmd.Flags().BoolS("passwd", "passwd", false, "Change password")
	rootCmd.Flags().BoolS("plaintextnames", "plaintextnames", false, "Do not encrypt file names (with -init)")
	rootCmd.Flags().BoolN("quiet", "q", false, "Silence informational messages")
	rootCmd.Flags().BoolS("reverse", "reverse", false, "Enable reverse mode")
	rootCmd.Flags().BoolS("ro", "ro", false, "Mount read-only")
	rootCmd.Flags().BoolS("speed", "speed", false, "Run crypto speed test")
	rootCmd.Flags().BoolS("version", "version", false, "Print version information")
}
