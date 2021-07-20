package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/micro_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "micro",
	Short: "A modern and intuitive terminal-based text editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("autoindent", "", "")
	rootCmd.Flags().String("autosave", "", "")
	rootCmd.Flags().String("autosu", "", "")
	rootCmd.Flags().String("backup", "", "")
	rootCmd.Flags().String("backupdir", "", "")
	rootCmd.Flags().String("basename", "", "")
	rootCmd.Flags().Bool("clean", false, "Cleans the configuration directory")
	rootCmd.Flags().String("clipboard", "", "")
	rootCmd.Flags().String("colorcolumn", "", "")
	rootCmd.Flags().String("colorscheme", "", "")
	rootCmd.Flags().String("config-dir", "", "Specify a custom location for the configuration directory")
	rootCmd.Flags().String("cursorline", "", "")
	rootCmd.Flags().Bool("debug", false, "Enable debug mode")
	rootCmd.Flags().String("diffgutter", "", "")
	rootCmd.Flags().String("divchars", "", "")
	rootCmd.Flags().String("divreverse", "", "")
	rootCmd.Flags().String("encoding", "", "")
	rootCmd.Flags().String("eofnewline", "", "")
	rootCmd.Flags().String("fastdirty", "", "")
	rootCmd.Flags().String("fileformat", "", "")
	rootCmd.Flags().String("filetype", "", "")
	rootCmd.Flags().String("ignorecase", "", "")
	rootCmd.Flags().String("incsearch", "", "")
	rootCmd.Flags().String("indentchar", "", "")
	rootCmd.Flags().String("infobar", "", "")
	rootCmd.Flags().String("keepautoindent", "", "")
	rootCmd.Flags().String("keymenu", "", "")
	rootCmd.Flags().String("matchbrace", "", "")
	rootCmd.Flags().String("mkparents", "", "")
	rootCmd.Flags().String("mouse", "", "")
	rootCmd.Flags().Bool("options", false, "Show all option help")
	rootCmd.Flags().String("parsecursor", "", "")
	rootCmd.Flags().String("paste", "", "")
	rootCmd.Flags().String("permbackup", "", "")
	rootCmd.Flags().String("plugin", "", "Manage plugins")
	rootCmd.Flags().String("pluginchannels", "", "")
	rootCmd.Flags().String("pluginrepos", "", "")
	rootCmd.Flags().String("readonly", "", "")
	rootCmd.Flags().String("relativeruler", "", "")
	rootCmd.Flags().String("rmtrailingws", "", "")
	rootCmd.Flags().String("ruler", "", "")
	rootCmd.Flags().String("savecursor", "", "")
	rootCmd.Flags().String("savehistory", "", "")
	rootCmd.Flags().String("saveundo", "", "")
	rootCmd.Flags().String("scrollbar", "", "")
	rootCmd.Flags().String("scrollmargin", "", "")
	rootCmd.Flags().String("scrollspeed", "", "")
	rootCmd.Flags().String("smartpaste", "", "")
	rootCmd.Flags().String("softwrap", "", "")
	rootCmd.Flags().String("splitbottom", "", "")
	rootCmd.Flags().String("splitright", "", "")
	rootCmd.Flags().String("statusformatl", "", "")
	rootCmd.Flags().String("statusformatr", "", "")
	rootCmd.Flags().String("statusline", "", "")
	rootCmd.Flags().String("sucmd", "", "")
	rootCmd.Flags().String("syntax", "", "")
	rootCmd.Flags().String("tabmovement", "", "")
	rootCmd.Flags().String("tabsize", "", "")
	rootCmd.Flags().String("tabstospaces", "", "")
	rootCmd.Flags().String("useprimary", "", "")
	rootCmd.Flags().Bool("version", false, "Show the version number and information")
	rootCmd.Flags().String("wordwrap", "", "")
	rootCmd.Flags().String("xterm", "", "")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir": carapace.ActionDirectories(),
		"plugin": carapace.ActionValuesDescribed(
			"available", "List available plugins",
			"install", "Install plugins",
			"list", "List installed plugins",
			"remove", "Remove plugins",
			"search", "Search for a plugin",
			"update", "Update plugins",
		),
		// TODO complete options
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := rootCmd.Flag("plugin"); flag.Changed {
				switch flag.Value.String() {
				case "available":
					return carapace.ActionValues()
				case "install":
					return action.ActionAvailablePlugins().Invoke(c).Filter(c.Args).ToA()
				case "list":
					return carapace.ActionValues()
				case "remove":
					return action.ActionInstalledPlugins().Invoke(c).Filter(c.Args).ToA()
				case "search":
					return carapace.ActionValues()
				case "update":
					return action.ActionInstalledPlugins().Invoke(c).Filter(c.Args).ToA()
				default:
					return carapace.ActionValues()
				}
			}
			return carapace.ActionFiles()
		}),
	)
}
