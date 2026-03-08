package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnf",
	Short: "DNF5 is a package manager for RPM-based Linux distributions",
	Long:  "https://github.com/rpm-software-management/dnf5",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("assumeno", false, "Automatically answer no for all questions")
	rootCmd.Flags().BoolP("assumeyes", "y", false, "Automatically answer yes for all questions")
	rootCmd.Flags().Bool("best", false, "Try the best available package versions in transactions")
	rootCmd.Flags().BoolP("cacheonly", "C", false, "Use only cached data")
	rootCmd.Flags().String("color", "", "Control whether color is used in terminal output")
	rootCmd.Flags().String("comment", "", "Add a comment to the transaction history")
	rootCmd.Flags().StringP("config", "c", "", "Define configuration file location")
	rootCmd.Flags().Bool("debugsolver", false, "Dump additional data from solver for debugging")
	rootCmd.Flags().String("disable-plugin", "", "Disable specified libdnf5 library plugins")
	rootCmd.Flags().String("disable-repo", "", "Temporarily disable active repositories")
	rootCmd.Flags().Bool("dump-main-config", false, "Print main configuration values to stdout")
	rootCmd.Flags().String("dump-repo-config", "", "Print repository configuration values to stdout")
	rootCmd.Flags().Bool("dump-variables", false, "Print variable values to stdout")
	rootCmd.Flags().String("enable-plugin", "", "Enable specified libdnf5 library plugins")
	rootCmd.Flags().String("enable-repo", "", "Temporarily enable additional repositories")
	rootCmd.Flags().StringP("exclude", "x", "", "Exclude packages from the transaction")
	rootCmd.Flags().String("forcearch", "", "Force the use of a specific architecture")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().String("installroot", "", "Setup installroot path")
	rootCmd.Flags().Bool("no-best", false, "Do not limit the transaction to the best candidates only")
	rootCmd.Flags().Bool("no-docs", false, "Do not install any files marked as documentation")
	rootCmd.Flags().Bool("no-gpgchecks", false, "Skip checking OpenPGP signatures on packages")
	rootCmd.Flags().Bool("no-plugins", false, "Disable all libdnf5 plugins")
	rootCmd.Flags().BoolP("quiet", "q", false, "In combination with a non-interactive command, shows just the relevant content")
	rootCmd.Flags().Bool("refresh", false, "Force refreshing metadata before running the command")
	rootCmd.Flags().String("releasever", "", "Override the value of the distribution release")
	rootCmd.Flags().String("releasever-major", "", "Override the releasever_major variable")
	rootCmd.Flags().String("releasever-minor", "", "Override the releasever_minor variable")
	rootCmd.Flags().String("repo", "", "Enable just specified repositories")
	rootCmd.Flags().String("repofrompath", "", "Specify a repository to add to the repositories only for this run")
	rootCmd.Flags().String("setopt", "", "Override a configuration option")
	rootCmd.Flags().String("setvar", "", "Override a DNF5 variable value")
	rootCmd.Flags().Bool("show-new-leaves", false, "Show newly installed leaf packages")
	rootCmd.Flags().Bool("skip-file-locks", false, "Skip acquiring file locks")
	rootCmd.Flags().Bool("use-host-config", false, "Use configuration files from the host system")
	rootCmd.Flags().Bool("version", false, "Display the version of dnf5")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("always", "never", "auto"),
		"config":      carapace.ActionFiles(),
		"forcearch":   carapace.ActionValues("i386", "i486", "i586", "i686", "x86_64", "aarch64", "armv7l", "ppc64", "ppc64le", "s390x"),
		"installroot": carapace.ActionDirectories(),
	})
}
