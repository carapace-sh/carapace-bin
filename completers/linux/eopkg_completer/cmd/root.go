package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/eopkg_completer/cmd/actions"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eopkg",
	Short: "Solus package manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.Flags().BoolP("yes-all", "y", false, "assume yes in all yes/no queries")
	rootCmd.Flags().StringP("destdir", "D", "", "change the system root for eopkg commands")
	rootCmd.Flags().StringP("username", "u", "", "username used for authentication")
	rootCmd.Flags().StringP("password", "p", "", "password used for authentication")

	rootCmd.AddCommand(
		command("install", "Install packages", true),
		command("it", "Install packages", true),
		command("remove", "Remove packages", false),
		command("rm", "Remove packages", false),
		command("upgrade", "Upgrade packages", false),
		command("up", "Upgrade packages", false),
		command("reinstall", "Reinstall packages", false),
		command("search", "Search packages", true),
		command("sr", "Search packages", true),
		command("info", "Display package information", true),
		command("list-installed", "List installed packages", false),
		command("li", "List installed packages", false),
		command("list-available", "List available packages", true),
		command("la", "List available packages", true),
		command("list-upgrades", "List available upgrades", false),
		command("lu", "List available upgrades", false),
		command("update-repo", "Update repository database", false),
		command("ur", "Update repository database", false),
		command("add-repo", "Add repository", false),
		command("ar", "Add repository", false),
		command("remove-repo", "Remove repository", false),
		command("rr", "Remove repository", false),
		command("list-repo", "List repositories", false),
		command("lr", "List repositories", false),
		command("history", "Show history", false),
		command("hs", "Show history", false),
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"destdir": carapace.ActionDirectories(),
	})
}

func command(use, short string, available bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()

	switch use {
	case "install", "it", "remove", "rm", "upgrade", "up", "reinstall", "search", "sr", "info":
		carapace.Gen(cmd).PositionalAnyCompletion(
			actions.ActionPackages(!available).FilterArgs(),
		)
	}

	return cmd
}
