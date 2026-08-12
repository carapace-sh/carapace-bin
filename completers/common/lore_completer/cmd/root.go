package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lore",
	Short: "next-generation, open source version control system",
	Long:  "https://lore.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("cache", false, "Cache fragment payloads fetched from remote in the local store")
	rootCmd.PersistentFlags().String("compress-limit", "", "Set maximum number of parallel compress operations")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Dry run mode, only report what would have been changed and perform no changes to local file system")
	rootCmd.PersistentFlags().String("file-count-limit", "", "Set maximum number of parallel files opened")
	rootCmd.PersistentFlags().String("file-size-limit", "", "Set maximum total size in bytes of parallel files opened")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Force the operation if possible")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().String("identity", "", "Use given identity")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Enable machine-readable json output")
	rootCmd.PersistentFlags().Bool("local", false, "Use local data")
	rootCmd.PersistentFlags().String("log-level", "", "Set the logging level")
	rootCmd.Flags().Bool("markdown-help", false, "")
	rootCmd.PersistentFlags().String("max-connections", "", "Set maximum number of parallel connections")
	rootCmd.PersistentFlags().String("max-threads", "", "Set maximum total number of threads Lore sizes its pools for")
	rootCmd.PersistentFlags().Bool("no-gc", false, "Prevent automatic incremental garbage collection for this command; it otherwise runs in the background on writes. `lore repository gc` always runs a full pass regardless")
	rootCmd.PersistentFlags().BoolP("no-pager", "P", false, "Disable pagination")
	rootCmd.PersistentFlags().Bool("nocompress", false, "Avoid using compression")
	rootCmd.PersistentFlags().Bool("non-interactive", false, "Disable interactive prompts (e.g., per-link commit messages)")
	rootCmd.PersistentFlags().Bool("offline", false, "Force offline mode")
	rootCmd.PersistentFlags().Bool("remote", false, "Use remote data")
	rootCmd.PersistentFlags().String("repository", "", "Use given path as repository path")
	rootCmd.PersistentFlags().String("search-limit", "", "Set maximum number of revisions to search when matching or finding revisions")
	rootCmd.PersistentFlags().Bool("search-nearest", false, "Set to search for nearest match when matching revisions")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "Suppress all output")
	rootCmd.PersistentFlags().Bool("sync-data", false, "Force sync data to storage media during flush")
	rootCmd.PersistentFlags().BoolP("time", "t", false, "Time execution of command")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flag("json").Hidden = true
	rootCmd.Flag("markdown-help").Hidden = true
	rootCmd.Flag("max-threads").Hidden = true
	rootCmd.Flag("nocompress").Hidden = true
	rootCmd.Flag("silent").Hidden = true
	rootCmd.Flag("time").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"compress-limit":   carapace.ActionValues(),
		"file-count-limit": carapace.ActionValues(),
		"file-size-limit":  carapace.ActionValues(),
		"identity":         carapace.ActionValues(),
		"log-level":        carapace.ActionValues("trace", "debug", "info", "warn", "error"),
		"max-connections":  carapace.ActionValues(),
		"max-threads":      carapace.ActionValues(),
		"repository":       carapace.ActionDirectories(),
		"search-limit":     carapace.ActionValues(),
	})
}
