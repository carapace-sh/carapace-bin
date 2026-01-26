package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman-key",
	Short: "Manage pacman's list of trusted keys",
	Long:  "https://archlinux.org/pacman/pacman-key.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("add", "a", false, "Add the specified keys (empty for stdin)")
	rootCmd.Flags().String("config", "", "Use an alternate config file")
	rootCmd.Flags().BoolP("delete", "d", false, "Remove the specified keyids")
	rootCmd.Flags().Bool("edit-key", false, "Present a menu for key management task on keyids")
	rootCmd.Flags().BoolP("export", "e", false, "Export the specified or all keyids")
	rootCmd.Flags().BoolP("finger", "f", false, "List fingerprint for specified or all keyids")
	rootCmd.Flags().String("gpgdir", "", "Set an alternate directory for GnuPG")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().Bool("import", false, "Imports pubring.gpg from dir(s)")
	rootCmd.Flags().Bool("import-trustdb", false, "Imports ownertrust values from trustdb.gpg in dir(s)")
	rootCmd.Flags().Bool("init", false, "Ensure the keyring is properly initialized")
	rootCmd.Flags().String("keyserver", "", "Specify a keyserver to use if necessary")
	rootCmd.Flags().BoolP("list-keys", "l", false, "List the specified or all keys")
	rootCmd.Flags().Bool("list-sigs", false, "List keys and their signatures")
	rootCmd.Flags().Bool("lsign-key", false, "Locally sign the specified keyid")
	rootCmd.Flags().Bool("populate", false, "Reload the default keys from the (given) keyrings in '/usr/share/pacman/keyrings'")
	rootCmd.Flags().String("populate-from", "", "Set an alternate directory for --populate (instead of '/usr/share/pacman/keyrings')")
	rootCmd.Flags().BoolP("recv-keys", "r", false, "Fetch the specified keyids")
	rootCmd.Flags().Bool("refresh-keys", false, "Update specified or all keys from a keyserver")
	rootCmd.Flags().BoolP("updatedb", "u", false, "Update the trustdb of pacman")
	rootCmd.Flags().Bool("verbose", false, "Show extra information")
	rootCmd.Flags().BoolP("verify", "v", false, "Verify the file(s) specified by the signature(s)")
	rootCmd.Flags().BoolP("version", "V", false, "Show program version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":        carapace.ActionFiles(),
		"gpgdir":        carapace.ActionDirectories(),
		"keyserver":     carapace.ActionValues(), // TODO
		"populate-from": carapace.ActionDirectories(),
	})
}
