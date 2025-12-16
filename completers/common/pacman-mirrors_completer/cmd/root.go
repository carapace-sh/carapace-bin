package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman-mirrors",
	Short: "generate pacman mirrorlist",
	Long:  "https://wiki.manjaro.org/index.php?title=Pacman-mirrors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("S", "S", "", "API: Replace branch in configuration")
	rootCmd.Flags().BoolP("api", "a", false, "TODO")
	rootCmd.Flags().Bool("continent", false, "Use continent from geolocation")
	rootCmd.Flags().StringP("country", "c", "", "Comma separated list of countries, from which mirrors will be used")
	rootCmd.Flags().BoolP("country-config", "lc", false, "lists configured mirror countries")
	rootCmd.Flags().Bool("country-list", false, "List all available countries")
	rootCmd.Flags().BoolP("default", "d", false, "INTERACTIVE: Load default mirror file")
	rootCmd.Flags().StringP("fasttrack", "f", "", "Generate mirrorlist with a number of up-to-date mirrors")
	rootCmd.Flags().BoolS("fc", "fc", false, "TODO")
	rootCmd.Flags().BoolS("g", "g", false, "Create mirror list from active pool.")
	rootCmd.Flags().Bool("geoip", false, "Get current country using geolocation")
	rootCmd.Flags().BoolP("get-branch", "G", false, "Return branch from configuration")
	rootCmd.Flags().BoolP("help", "h", false, "")
	rootCmd.Flags().BoolP("interactive", "i", false, "Generate custom mirrorlist")
	rootCmd.Flags().String("interval", "", "Max. number of hours since last sync")
	rootCmd.Flags().BoolP("list", "l", false, "List all available countries")
	rootCmd.Flags().StringP("method", "m", "", "Generation method")
	rootCmd.Flags().Bool("no-color", false, "Disable color")
	rootCmd.Flags().BoolP("no-mirrorlist", "n", false, "Use to skip generation of mirrorlist")
	rootCmd.Flags().BoolP("no-status", "s", false, "Ignore mirror branch status")
	rootCmd.Flags().StringP("prefix", "p", "", "API: Set prefix to : $mnt | /mnt/install")
	rootCmd.Flags().StringP("proto", "P", "", "API: Replace protocols in configuration")
	rootCmd.Flags().String("protocols", "", "API: Replace protocols in configuration")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode - less verbose output")
	rootCmd.Flags().BoolP("re-branch", "R", false, "API: Replace branch in mirrorlist")
	rootCmd.Flags().StringP("set-branch", "B", "", "API: Replace branch in configuration")
	rootCmd.Flags().Bool("status", false, "Status for the current mirror list.")
	rootCmd.Flags().StringP("timeout", "t", "", "Maximum waiting time for server response")
	rootCmd.Flags().StringP("url", "U", "", "API: Replace mirror url in mirrorlist")
	rootCmd.Flags().BoolP("version", "v", false, "Print the pacman-mirrors version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"S":          carapace.ActionValues("stable", "testing", "unstable"),
		"country":    pacman.ActionCountries().UniqueList(","),
		"method":     carapace.ActionValues("rank", "random"),
		"proto":      carapace.ActionValues("all", "http", "https", "ftp", "ftps").UniqueList(","),
		"protocols":  carapace.ActionValues("all", "http", "https", "ftp", "ftps").UniqueList(","),
		"set-branch": carapace.ActionValues("stable", "testing", "unstable"),
	})
}
