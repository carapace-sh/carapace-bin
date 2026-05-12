package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newConfigCommand() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Modify the Baloo configuration",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	configCmd.AddCommand(
		newConfigAddCommand(),
		newConfigHelpCommand(),
		newConfigListCommand("list", "Show the value of a config parameter"),
		newConfigRemoveCommand("remove", "Remove a value from a config parameter"),
		newConfigRemoveCommand("rm", "Remove a value from a config parameter"),
		newConfigSetCommand(),
		newConfigListCommand("show", "Show the value of a config parameter"),
		newConfigListCommand("ls", "Show the value of a config parameter"),
	)

	carapace.Gen(configCmd).PositionalCompletion(
		actionConfigCommands(),
	)

	return configCmd
}

func newConfigAddCommand() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a value to config parameter",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(addCmd).PositionalCompletion(
		actionModifiableConfigParameters(),
		actionConfigValue(),
	)

	return addCmd
}

func newConfigHelpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "Display the config help menu",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}

func newConfigListCommand(use string, short string) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(listCmd).PositionalCompletion(
		actionListableConfigParameters(),
	)

	return listCmd
}

func newConfigRemoveCommand(use string, short string) *cobra.Command {
	removeCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(removeCmd).PositionalCompletion(
		actionModifiableConfigParameters(),
		actionConfigValue(),
	)

	return removeCmd
}

func newConfigSetCommand() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set the value of a config parameter",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(setCmd).PositionalCompletion(
		actionSettableConfigParameters(),
		actionBooleanValues(),
	)

	return setCmd
}

func actionBooleanValues() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "enable the option",
		"yes", "enable the option",
		"y", "enable the option",
		"1", "enable the option",
		"false", "disable the option",
		"no", "disable the option",
		"n", "disable the option",
		"0", "disable the option",
	)
}

func actionConfigCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"add", "Add a value to config parameter",
		"rm", "Remove a value from a config parameter",
		"remove", "Remove a value from a config parameter",
		"list", "Show the value of a config parameter",
		"ls", "Show the value of a config parameter",
		"show", "Show the value of a config parameter",
		"set", "Set the value of a config parameter",
		"help", "Display the config help menu",
	)
}

func actionConfigValue() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) < 1 {
			return carapace.ActionValues()
		}

		switch c.Args[0] {
		case "includeFolders", "excludeFolders":
			return carapace.ActionDirectories()
		default:
			return carapace.ActionValues()
		}
	})
}

func actionListableConfigParameters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"hidden", "Controls if Baloo indexes hidden files and folders",
		"contentIndexing", "Controls if Baloo indexes file content",
		"includeFolders", "The list of folders which Baloo indexes",
		"excludeFolders", "The list of folders which Baloo will never index",
		"excludeFilters", "The list of filters which are used to exclude files",
		"excludeMimetypes", "The list of mimetypes which are used to exclude files",
	)
}

func actionModifiableConfigParameters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"includeFolders", "The list of folders which Baloo indexes",
		"excludeFolders", "The list of folders which Baloo will never index",
		"excludeFilters", "The list of filters which are used to exclude files",
		"excludeMimetypes", "The list of mimetypes which are used to exclude files",
	)
}

func actionSettableConfigParameters() carapace.Action {
	return carapace.ActionValuesDescribed(
		"hidden", "Controls if Baloo indexes hidden files and folders",
		"contentIndexing", "Controls if Baloo indexes file content",
	)
}
