package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "dpkg",
	Short:              "package manager for Debian",
	Long:               "https://linux.die.net/man/1/dpkg",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			cmd := &cobra.Command{}
			carapace.Gen(cmd).Standalone()
			addRootFlags(cmd)
			return carapace.ActionExecute(cmd)
		}), // TODO merge with dpkg-deb
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO use embed.SubcommandsAsFlags()
			switch c.Args[0] {
			case "-i", "--install":
				return carapace.ActionExecute(installCmd).Shift(1)
			case "--unpack":
				return carapace.ActionExecute(unpackCmd).Shift(1)
			case "-A", "--record-avail":
				return carapace.ActionExecute(recordAvailCmd).Shift(1)
			case "--configure":
				return carapace.ActionExecute(configureCmd).Shift(1)
			case "--triggers-only":
				return carapace.ActionExecute(triggersOnlyCmd).Shift(1)
			case "--purge":
				return carapace.ActionExecute(purgeCmd).Shift(1)
			case "-V", "--verify":
				return carapace.ActionExecute(verifyCmd).Shift(1)
			case "--update-avail":
				return carapace.ActionExecute(updateAvailCmd).Shift(1)
			case "--merge-avail":
				return carapace.ActionExecute(mergeAvailCmd).Shift(1)
			case "-s", "--status":
				return carapace.ActionExecute(statusCmd).Shift(1)
			case "-p", "--print-avail":
				return carapace.ActionExecute(printAvailCmd).Shift(1)
			case "-L", "--listfiles":
				return carapace.ActionExecute(listfilesCmd).Shift(1)
			case "-C", "--audit":
				return carapace.ActionExecute(auditCmd).Shift(1)
			case "--compare-versions":
				return carapace.ActionExecute(compareVersionsCmd).Shift(1)
			default:
				return carapace.ActionValues()
			}
		}), // TODO merge with dpkg-deb
	)
}

func addRootFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("add-architecture", false, "Add <arch> to the list of architectures")
	cmd.Flags().Bool("assert-help", false, "Show help on assertions")
	cmd.Flags().Bool("assert-long-filenames", false, "long filenames in .deb archives")
	cmd.Flags().Bool("assert-multi-arch", false, "multi-arch fields and semantics")
	cmd.Flags().Bool("assert-multi-conrep", false, "multiple Conflicts and Replaces")
	cmd.Flags().Bool("assert-protected-field", false, "the Protected field")
	cmd.Flags().Bool("assert-support-predepends", false, "the Pre-Depends field")
	cmd.Flags().Bool("assert-versioned-provides", false, "versioned relationships in the Provides field")
	cmd.Flags().Bool("assert-working-epoch", false, "epochs in versions")
	cmd.Flags().BoolP("audit", "C", false, "Check for broken package(s)")
	cmd.Flags().Bool("clear-avail", false, "Erase existing available info")
	cmd.Flags().Bool("clear-selections", false, "Deselect every non-essential package")
	cmd.Flags().Bool("compare-versions", false, "Compare version numbers - see below")
	cmd.Flags().Bool("configure", false, "Reconfigure an unpacked package")
	cmd.Flags().StringP("debug", "Dh", "", "Show help on debugging")
	cmd.Flags().Bool("force-help", false, "Show help on forcing")
	cmd.Flags().Bool("forget-old-unavail", false, "Forget uninstalled unavailable pkgs")
	cmd.Flags().Bool("get-selections", false, "Get list of selections to stdout")
	cmd.Flags().BoolP("help", "?", false, "Show this help message")
	cmd.Flags().BoolP("install", "i", false, "Install the package")
	cmd.Flags().BoolP("list", "l", false, "List packages concisely")
	cmd.Flags().BoolP("listfiles", "L", false, "List files 'owned' by package(s)")
	cmd.Flags().Bool("merge-avail", false, "Merge with info from file")
	cmd.Flags().Bool("predep-package", false, "Print pre-dependencies to unpack")
	cmd.Flags().Bool("print-architecture", false, "Print dpkg architecture")
	cmd.Flags().BoolP("print-avail", "p", false, "Display available version details")
	cmd.Flags().Bool("print-foreign-architectures", false, "Print allowed foreign architectures")
	cmd.Flags().BoolP("purge", "P", false, "Remove an installed Package including conffiles")
	cmd.Flags().BoolP("record-avail", "A", false, "Update which packages are available with information from the package")
	cmd.Flags().BoolP("remove", "r", false, "Remove an installed package")
	cmd.Flags().Bool("remove-architeture", false, "Remove <arch> from the list of architectures")
	cmd.Flags().BoolP("search", "S", false, "Find package(s) owning file(s)")
	cmd.Flags().Bool("set-selections", false, "Set package selections from stdin")
	cmd.Flags().BoolP("status", "s", false, "Display package status details")
	cmd.Flags().Bool("triggers-only", false, "Processes only triggers")
	cmd.Flags().Bool("unpack", false, "Unpack the package ")
	cmd.Flags().Bool("update-avail", false, "Replace available packages info")
	cmd.Flags().Bool("validate-archname", false, "Validate arch name")
	cmd.Flags().Bool("validate-pkgname", false, "Validate package name")
	cmd.Flags().Bool("validate-trigname", false, "Validate trigger name")
	cmd.Flags().Bool("validate-version", false, "Validate version")
	cmd.Flags().BoolP("verify", "V", false, "Verify the integrity of package(s)")
	cmd.Flags().Bool("version", false, "Show the version")
	cmd.Flags().Bool("yet-to-unpack", false, "Print packages selected for installation")

	cmd.Flag("debug").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"debug": carapace.ActionValues("help"),
	})
}
