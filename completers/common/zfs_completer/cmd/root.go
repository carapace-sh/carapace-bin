package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "zfs", Short: "configure ZFS datasets", Long: "https://openzfs.github.io/openzfs-docs/man/v2.4/8/zfs.8.html", Run: func(cmd *cobra.Command, args []string) {}}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "print help")
	rootCmd.Flags().BoolP("version", "V", false, "print version")

	for _, spec := range []string{
		"allow delegate permissions", "bookmark create bookmark", "change-key change encryption key", "clone clone a snapshot", "create create a dataset", "destroy destroy datasets/snapshots", "diff show differences", "get get properties", "groupspace list group space", "hold hold snapshots", "holds list holds", "inherit inherit properties", "jail attach dataset to jail", "list list datasets", "load-key load encryption keys", "mount mount datasets", "program run ZFS channel program", "promote promote clone", "project project space accounting", "projectspace list project space", "receive receive snapshot stream", "redact create redaction bookmark", "release release snapshot holds", "rename rename dataset", "rollback roll back to snapshot", "send send snapshot stream", "set set properties", "share share datasets", "snapshot create snapshot", "unallow remove delegated permissions", "unjail detach dataset from jail", "unload-key unload encryption keys", "unmount unmount datasets", "upgrade upgrade filesystem version", "userspace list user space", "wait wait for activity",
	} {
		addCommand(spec)
	}

	carapace.Gen(rootCmd).PositionalCompletion(actionSubcommands())
}

func addCommand(spec string) {
	parts := strings.SplitN(spec, " ", 2)
	cmd := &cobra.Command{Use: parts[0], Short: parts[1], Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().BoolP("recursive", "r", false, "operate recursively")
	cmd.Flags().BoolP("verbose", "v", false, "verbose output")
	cmd.Flags().BoolP("parseable", "p", false, "parseable output")
	cmd.Flags().StringP("output", "o", "", "properties/fields to display")
	cmd.Flags().StringP("type", "t", "", "dataset type")
	cmd.Flags().StringP("property", "s", "", "property source")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output":   actionProperties().UniqueList(","),
		"type":     carapace.ActionValues("filesystem", "snapshot", "volume", "bookmark", "all").UniqueList(","),
		"property": carapace.ActionValues("local", "default", "inherited", "temporary", "received", "none").UniqueList(","),
	})
	carapace.Gen(cmd).PositionalAnyCompletion(actionDatasets())
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

func actionDatasets() carapace.Action {
	return carapace.ActionExecCommand("zfs", "list", "-H", "-o", "name", "-t", "all")(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}

func actionProperties() carapace.Action {
	return carapace.ActionValues("name", "type", "creation", "used", "available", "referenced", "compressratio", "mounted", "quota", "reservation", "recordsize", "mountpoint", "sharenfs", "checksum", "compression", "atime", "devices", "exec", "setuid", "readonly", "zoned", "snapdir", "aclmode", "aclinherit", "createtxg", "canmount", "xattr", "copies", "version", "utf8only", "normalization", "casesensitivity", "vscan", "nbmand", "sharesmb", "refquota", "refreservation", "guid", "primarycache", "secondarycache", "usedbysnapshots", "usedbydataset", "usedbychildren", "usedbyrefreservation", "logbias", "dedup", "mlslabel", "sync", "dnodesize", "refcompressratio", "written", "logicalused", "logicalreferenced", "volmode", "filesystem_limit", "snapshot_limit", "filesystem_count", "snapshot_count", "snapdev", "acltype", "context", "fscontext", "defcontext", "rootcontext", "relatime", "redundant_metadata", "overlay", "encryption", "keylocation", "keyformat", "pbkdf2iters", "special_small_blocks")
}
