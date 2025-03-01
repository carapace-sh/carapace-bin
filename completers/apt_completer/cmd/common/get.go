package common

import "github.com/spf13/cobra"

func AddGetFlags(cmd *cobra.Command) {
	// TODO documentation is lacking and descriptions are likely wrong
	cmd.Flags().Bool("allow-change-held-packages", false, "allow changing held packages")
	cmd.Flags().Bool("allow-downgrades", false, "allow downgrades")
	cmd.Flags().Bool("allow-insecure-repositories", false, "allow insecure repositories")
	cmd.Flags().Bool("allow-remove-essential", false, "allow removal of essential packages")
	cmd.Flags().Bool("allow-unauthenticated", false, "allow unauthenticated packages")
	cmd.Flags().Bool("arch-only", false, "only process architecture-dependent build-dependencies")
	cmd.Flags().Bool("assume-no", false, "automatic \"no\" to all prompts")
	cmd.Flags().BoolP("assume-yes", "y", false, "automatic \"yes\" to all prompts")
	cmd.Flags().Bool("download", false, "download the given binary package into the current directory")
	cmd.Flags().BoolP("download-only", "d", false, "do not unpack source package")
	cmd.Flags().Bool("fix-missing", false, "fix missing packages")
	cmd.Flags().Bool("fix-policy", false, "fix policies")
	cmd.Flags().Bool("ignore-hold", false, "ignore package holds")
	cmd.Flags().BoolP("ignore-missing", "m", false, "ignore missing packages")
	cmd.Flags().Bool("install-recommends", false, "install recommended packages")
	cmd.Flags().Bool("install-suggests", false, "install suggested packages")
	cmd.Flags().Bool("no-install-recommends", false, "do not install recommended packages")
	cmd.Flags().Bool("no-install-suggests", false, "do not install suggested packages")
	cmd.Flags().Bool("only-upgrade", false, "only updgrade packages")
	cmd.Flags().Bool("print-uris", false, "print file URIs")
	cmd.Flags().Bool("remove", false, "remove packages")
	cmd.Flags().BoolP("show-upgraded", "u", false, "show upgraded packages")
	cmd.Flags().Bool("trivial-only", false, "only perform operations that are 'trivial'")
	cmd.Flags().Bool("upgrade", false, "upgrade and install new dependencies")

	cmd.MarkFlagsMutuallyExclusive("assume-no", "assume-yes")
	cmd.MarkFlagsMutuallyExclusive("no-install-recommends", "install-recommends")
	cmd.MarkFlagsMutuallyExclusive("no-install-suggests", "install-suggests")
}
