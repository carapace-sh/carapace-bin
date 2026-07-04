package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dscl",
	Short: "directory service command line utility",
	Long:  "https://keith.github.io/xcode-manpages/dscl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")
	rootCmd.Flags().BoolP("plist", "p", false, "Output in plist format")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().BoolP("raw", "raw", false, "Print raw data")
	rootCmd.Flags().BoolP("url", "url", false, "Print data as URL-encoded strings")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"localhost", "Local directory node",
			"/Local/Default", "Local default directory",
			"/Active Directory", "Active directory",
			"/LDAPv3", "LDAP v3 directory",
			"read", "Read a record",
			"create", "Create a record or attribute",
			"delete", "Delete a record or attribute",
			"list", "List records",
			"search", "Search for records",
			"append", "Append to an attribute",
			"merge", "Merge values into an attribute",
			"change", "Change a value in an attribute",
			"readall", "Read all records",
		),
	)
}
