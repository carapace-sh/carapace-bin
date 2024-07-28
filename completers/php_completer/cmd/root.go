package cmd

import (
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "php",
	Short: "PHP Command Line Interface",
	Long:  "https://www.php.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("B", "B", "", "Run PHP <begin_code> before processing input lines")
	rootCmd.Flags().StringS("E", "E", "", "Run PHP <end_code> after processing all input lines")
	rootCmd.Flags().StringS("F", "F", "", "Parse and execute <file> for every input line")
	rootCmd.Flags().BoolS("H", "H", false, "Hide any passed arguments from external tools")
	rootCmd.Flags().StringS("R", "R", "", "Run PHP <code> for every input line")
	rootCmd.Flags().BoolS("S", "S", false, "Run with built-in web server")
	rootCmd.Flags().BoolS("a", "a", false, "Run as interactive shell (requires readline extension)")
	rootCmd.Flags().BoolS("c", "c", false, "Look for php.ini file in this directory")
	rootCmd.Flags().StringS("d", "d", "", "Define INI entry foo with value 'bar'")
	rootCmd.Flags().BoolS("e", "e", false, "Generate extended information for debugger/profiler")
	rootCmd.Flags().StringS("f", "f", "", "Parse and execute <file>")
	rootCmd.Flags().BoolS("h", "h", false, "This help")
	rootCmd.Flags().BoolS("i", "i", false, "PHP information")
	rootCmd.Flags().Bool("ini", false, "Show configuration file names")
	rootCmd.Flags().BoolS("l", "l", false, "Syntax check only (lint)")
	rootCmd.Flags().BoolS("m", "m", false, "Show compiled in modules")
	rootCmd.Flags().BoolS("n", "n", false, "No configuration (ini) files will be used")
	rootCmd.Flags().StringS("r", "r", "", "Run PHP <code> without using script tags <?..?>")
	rootCmd.Flags().String("rc", "", "Show information about class <name>")
	rootCmd.Flags().String("re", "", "Show information about extension <name>")
	rootCmd.Flags().String("rf", "", "Show information about function <name>")
	rootCmd.Flags().String("ri", "", "Show configuration for extension <name>")
	rootCmd.Flags().String("rz", "", "Show information about Zend extension <name>")
	rootCmd.Flags().BoolS("s", "s", false, "Output HTML syntax highlighted source")
	rootCmd.Flags().StringS("t", "t", "", "Specify document root <docroot> for built-in web server")
	rootCmd.Flags().BoolS("v", "v", false, "Version number")
	rootCmd.Flags().BoolS("w", "w", false, "Output source with stripped comments and whitespace")
	rootCmd.Flags().StringS("z", "z", "", "Load Zend extension <file>")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionFiles(),
		"S": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return net.ActionPorts()
			}
		}),
		"c": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
		"t": carapace.ActionDirectories(),
		"z": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("S").Changed || rootCmd.Flag("r").Changed {
				return carapace.ActionValues()
			}

			shift := 1
			var file string
			switch {
			case rootCmd.Flag("F").Changed:
				file = rootCmd.Flag("F").Value.String()
				shift = 0
			case rootCmd.Flag("f").Changed:
				file = rootCmd.Flag("f").Value.String()
				shift = 0
			case len(c.Args) == 0:
				return carapace.ActionFiles()
			default:
				file = c.Args[0]
			}

			switch filepath.Base(file) {
			case "artisan": // assume laravel artisan
				return bridge.ActionCarapaceBin("artisan").Shift(shift).Chdir(filepath.Dir(file))
			default:
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionPositional(rootCmd),
	)
}
