package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "zpool", Short: "configure ZFS storage pools", Long: "https://openzfs.github.io/openzfs-docs/man/v2.4/8/zpool.8.html", Run: func(cmd *cobra.Command, args []string) {}}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "print help")
	rootCmd.Flags().BoolP("version", "V", false, "print version")
	for _, spec := range []string{
		"add add vdevs to a pool", "attach attach a device", "checkpoint checkpoint pool", "clear clear device errors", "create create a pool", "destroy destroy a pool", "detach detach a device", "events list events", "export export pools", "get get properties", "history show command history", "import import pools", "initialize initialize devices", "iostat display I/O statistics", "labelclear remove ZFS labels", "list list pools", "offline take device offline", "online bring device online", "reguid generate new pool GUID", "remove remove devices", "reopen reopen pools", "replace replace a device", "resilver restart resilver", "scrub start/stop scrub", "set set properties", "split split a mirrored pool", "status show pool status", "sync sync pool", "trim trim devices", "upgrade upgrade pool version", "wait wait for activity",
	} {
		addCommand(spec)
	}
	carapace.Gen(rootCmd).PositionalCompletion(actionSubcommands())
}

func addCommand(spec string) {
	parts := strings.SplitN(spec, " ", 2)
	cmd := &cobra.Command{Use: parts[0], Short: parts[1], Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().BoolP("verbose", "v", false, "verbose output")
	cmd.Flags().BoolP("parseable", "p", false, "parseable output")
	cmd.Flags().StringP("output", "o", "", "properties/fields to display")
	cmd.Flags().StringP("interval", "T", "", "timestamp style")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output":   actionProperties().UniqueList(","),
		"interval": carapace.ActionValues("d", "u"),
	})
	carapace.Gen(cmd).PositionalAnyCompletion(actionPools())
	rootCmd.AddCommand(cmd)
}

func actionSubcommands() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{}
		for _, cmd := range rootCmd.Commands() {
			vals = append(vals, cmd.Name(), cmd.Short)
		}
		return carapace.ActionValuesDescribed(vals...).StyleF(style.ForKeyword)
	})
}

func actionPools() carapace.Action {
	return carapace.ActionExecCommand("zpool", "list", "-H", "-o", "name")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}

func actionProperties() carapace.Action {
	return carapace.ActionValues("name", "size", "capacity", "altroot", "health", "guid", "version", "bootfs", "delegation", "autoreplace", "cachefile", "failmode", "listsnapshots", "autoexpand", "dedupditto", "dedupratio", "free", "allocated", "readonly", "ashift", "comment", "expandsize", "freeing", "fragmentation", "leaked", "multihost", "checkpoint", "load_guid", "autotrim", "compatibility", "feature@async_destroy", "feature@empty_bpobj", "feature@lz4_compress", "feature@spacemap_histogram", "feature@enabled_txg", "feature@hole_birth", "feature@extensible_dataset", "feature@embedded_data", "feature@bookmarks", "feature@filesystem_limits", "feature@large_blocks", "feature@large_dnode", "feature@sha512", "feature@skein", "feature@edonr", "feature@userobj_accounting", "feature@encryption", "feature@project_quota", "feature@device_removal", "feature@obsolete_counts", "feature@zpool_checkpoint", "feature@spacemap_v2", "feature@allocation_classes", "feature@resilver_defer", "feature@bookmark_v2", "feature@redaction_bookmarks", "feature@redacted_datasets", "feature@bookmark_written", "feature@log_spacemap", "feature@livelist", "feature@device_rebuild", "feature@zstd_compress", "feature@draid")
}
