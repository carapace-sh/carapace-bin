package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chooseFilesCmd = &cobra.Command{
	Use:   "choose-files",
	Short: "Choose files, fast",
}

func init() {
	rootCmd.AddCommand(chooseFilesCmd)
	carapace.Gen(chooseFilesCmd).Standalone()

	chooseFilesCmd.Flags().Bool("clear-cache", false, "Clear the caches used by this kitten.")
	chooseFilesCmd.Flags().String("config", "", "Specify a path to the configuration file(s) to use. All configuration files are merged onto the builtin choose-files.conf, overriding the builtin values. This option can be specified multiple times to read multiple configuration files in sequence, which are merged. Use the special value NONE to not load any config file.")
	chooseFilesCmd.Flags().Bool("display-title", false, "Show the window title at the top, useful when this kitten is used in an OS window without a title bar.")
	chooseFilesCmd.Flags().StringArray("file-filter", nil, "A list of filters to restrict the displayed files. Can be either mimetypes, or glob style patterns. Can be specified multiple times. The syntax is type:expression:Descriptive Name. For example: mime:image/png:Images and mime:image/gif:Images and glob:*.[tT][xX][Tt]:Text files.")
	chooseFilesCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	chooseFilesCmd.Flags().String("mode", "", "The type of object(s) to select")
	chooseFilesCmd.Flags().String("output-format", "", "The format in which to write the output. The text format is absolute paths separated by newlines. The shell format is quoted absolute paths separated by spaces, quoting is done only if needed. The shell-relative format is the same as shell except it returns paths relative to the starting directory.")
	chooseFilesCmd.Flags().StringArrayP("override", "o", nil, "Override individual configuration options, can be specified multiple times. Syntax: name=value.")
	chooseFilesCmd.Flags().String("suggested-save-file-name", "", "A suggested name when picking a save file.")
	chooseFilesCmd.Flags().String("suggested-save-file-path", "", "Path to an existing file to use as the save file.")
	chooseFilesCmd.Flags().String("title", "", "Window title to use for this chooser")
	chooseFilesCmd.Flags().String("write-output-to", "", "Path to a file to which the output is written in addition to STDOUT.")
	chooseFilesCmd.Flags().String("write-pid-to", "", "Path to a file to which to write the process ID (PID) of this process to.")

	carapace.Gen(chooseFilesCmd).FlagCompletion(carapace.ActionMap{
		"mode":                     carapace.ActionValues("file", "dir", "dirs", "files", "save-dir", "save-file", "save-files"),
		"suggested-save-file-path": carapace.ActionFiles(),
		"config":                   carapace.ActionFiles("~/.config/kitty"),
		"write-output-to":          carapace.ActionFiles(),
		"output-format":            carapace.ActionValues("text", "json", "shell", "shell-relative"),
		"write-pid-to":             carapace.ActionFiles(),
	})

	carapace.Gen(chooseFilesCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
