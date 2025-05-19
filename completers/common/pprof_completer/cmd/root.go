package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pprof",
	Short: "pprof is a tool for visualization and analysis of profiling data",
	Long:  "https://github.com/google/pprof",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("add_comment", "add_comment", "", "Free-form annotation to add to the profile")
	rootCmd.Flags().BoolS("addresses", "addresses", false, "Aggregate at the address level")
	rootCmd.Flags().BoolS("alloc_objects", "alloc_objects", false, "Same as -sample_index=alloc_objects")
	rootCmd.Flags().BoolS("alloc_space", "alloc_space", false, "Same as -sample_index=alloc_space")
	rootCmd.Flags().StringS("base", "base", "", "Source of base profile for profile subtraction")
	rootCmd.Flags().StringS("buildid", "buildid", "", "Override build id for main binary")
	rootCmd.Flags().BoolS("call_tree", "call_tree", false, "Create a context-sensitive call tree")
	rootCmd.Flags().BoolS("callgrind", "callgrind", false, "Outputs a graph in callgrind format")
	rootCmd.Flags().BoolS("comments", "comments", false, "Output all profile comments")
	rootCmd.Flags().BoolS("compact_labels", "compact_labels", false, "Show minimal headers")
	rootCmd.Flags().BoolS("contentions", "contentions", false, "Same as -sample_index=contentions")
	rootCmd.Flags().BoolS("cum", "cum", false, "Sort entries based on cumulative weight")
	rootCmd.Flags().StringS("diff_base", "diff_base", "", "Source of base profile for comparison")
	rootCmd.Flags().BoolS("disasm", "disasm", false, "Output assembly listings annotated with samples")
	rootCmd.Flags().StringS("divide_by", "divide_by", "", "Ratio to divide all samples before visualization")
	rootCmd.Flags().BoolS("dot", "dot", false, "Outputs a graph in DOT format")
	rootCmd.Flags().BoolS("drop_negative", "drop_negative", false, "Ignore negative differences")
	rootCmd.Flags().StringS("edgefraction", "edgefraction", "", "Hide edges below <f>*total")
	rootCmd.Flags().BoolS("eog", "eog", false, "Visualize graph through eog")
	rootCmd.Flags().BoolS("evince", "evince", false, "Visualize graph through evince")
	rootCmd.Flags().BoolS("filefunctions", "filefunctions", false, "Aggregate at the function level")
	rootCmd.Flags().BoolS("files", "files", false, "Aggregate at the file level")
	rootCmd.Flags().BoolS("flat", "flat", false, "Sort entries based on own weight")
	rootCmd.Flags().StringS("focus", "focus", "", "Restricts to samples going through a node matching regexp")
	rootCmd.Flags().BoolS("functions", "functions", false, "Aggregate at the function level")
	rootCmd.Flags().BoolS("gif", "gif", false, "Outputs a graph image in GIF format")
	rootCmd.Flags().BoolS("gv", "gv", false, "Visualize graph through gv")
	rootCmd.Flags().StringS("hide", "hide", "", "Skips nodes matching regexp")
	rootCmd.Flags().StringS("http", "http", "", "Provide web interface at host:port")
	rootCmd.Flags().StringS("ignore", "ignore", "", "Skips paths going through any nodes matching regexp")
	rootCmd.Flags().BoolS("intel_syntax", "intel_syntax", false, "Show assembly in Intel syntax")
	rootCmd.Flags().BoolS("inuse_objects", "inuse_objects", false, "Same as -sample_index=inuse_objects")
	rootCmd.Flags().BoolS("inuse_space", "inuse_space", false, "Same as -sample_index=inuse_space")
	rootCmd.Flags().BoolS("kcachegrind", "kcachegrind", false, "Visualize report in KCachegrind")
	rootCmd.Flags().BoolS("lines", "lines", false, "Aggregate at the source code line level")
	rootCmd.Flags().BoolS("list", "list", false, "Output annotated source for functions matching regexp")
	rootCmd.Flags().BoolS("mean", "mean", false, "Average sample value over first value (count)")
	rootCmd.Flags().BoolS("mean_delay", "mean_delay", false, "Same as -mean -sample_index=delay")
	rootCmd.Flags().BoolS("no_browser", "no_browser", false, "Skip opening a browser for the interactive web UI")
	rootCmd.Flags().StringS("nodecount", "nodecount", "", "Max number of nodes to show")
	rootCmd.Flags().StringS("nodefraction", "nodefraction", "", "Hide nodes below <f>*total")
	rootCmd.Flags().BoolS("noinlines", "noinlines", false, "Ignore inlines")
	rootCmd.Flags().BoolS("normalize", "normalize", false, "Scales profile based on the base profile")
	rootCmd.Flags().StringS("output", "output", "", "Output filename for file-based outputs")
	rootCmd.Flags().BoolS("pdf", "pdf", false, "Outputs a graph in PDF format")
	rootCmd.Flags().StringS("peek", "peek", "", "Output callers/callees of functions matching regexp")
	rootCmd.Flags().BoolS("png", "png", false, "Outputs a graph image in PNG format")
	rootCmd.Flags().BoolS("proto", "proto", false, "Outputs the profile in compressed protobuf format")
	rootCmd.Flags().StringS("prune_from", "prune_from", "", "Drops any functions below the matched frame")
	rootCmd.Flags().BoolS("ps", "ps", false, "Outputs a graph in PS format")
	rootCmd.Flags().BoolS("raw", "raw", false, "Outputs a text representation of the raw profile")
	rootCmd.Flags().BoolS("relative_percentages", "relative_percentages", false, "Show percentages relative to focused subgraph")
	rootCmd.Flags().StringS("sample_index", "sample_index", "", "Sample value to report (0-based index or name)")
	rootCmd.Flags().StringS("seconds", "seconds", "", "Duration for time-based profile collection")
	rootCmd.Flags().StringS("show", "show", "", "Only show nodes matching regexp")
	rootCmd.Flags().StringS("show_from", "show_from", "", "Drops functions above the highest matched frame")
	rootCmd.Flags().StringS("source_path", "source_path", "", "Search path for source files")
	rootCmd.Flags().BoolS("svg", "svg", false, "Outputs a graph in SVG format")
	rootCmd.Flags().StringS("symbolize", "symbolize", "", "Controls source of symbol information")
	rootCmd.Flags().StringS("tagfocus", "tagfocus", "", "Restricts to samples with tags in range or matched by regexp")
	rootCmd.Flags().StringS("taghide", "taghide", "", "Skip tags matching this regexp")
	rootCmd.Flags().StringS("tagignore", "tagignore", "", "Discard samples with tags in range or matched by regexp")
	rootCmd.Flags().StringS("tagleaf", "tagleaf", "", "Adds pseudo stack frames for labels key/value pairs at the callstack leaf")
	rootCmd.Flags().StringS("tagroot", "tagroot", "", "Adds pseudo stack frames for labels key/value pairs at the callstack root")
	rootCmd.Flags().BoolS("tags", "tags", false, "Outputs all tags in the profile")
	rootCmd.Flags().StringS("tagshow", "tagshow", "", "Only consider tags matching this regexp")
	rootCmd.Flags().BoolS("text", "text", false, "Outputs top entries in text form")
	rootCmd.Flags().StringS("timeout", "timeout", "", "Timeout in seconds for profile collection")
	rootCmd.Flags().StringS("tls_ca", "tls_ca", "", "TLS CA certs file for fetching profile and symbols")
	rootCmd.Flags().StringS("tls_cert", "tls_cert", "", "TLS client certificate file for fetching profile and symbols")
	rootCmd.Flags().StringS("tls_key", "tls_key", "", "TLS private key file for fetching profile and symbols")
	rootCmd.Flags().StringS("tools", "tools", "", "Search path for object tools")
	rootCmd.Flags().BoolS("top", "top", false, "Outputs top entries in text form")
	rootCmd.Flags().BoolS("topproto", "topproto", false, "Outputs top entries in compressed protobuf format")
	rootCmd.Flags().BoolS("total_delay", "total_delay", false, "Same as -sample_index=delay")
	rootCmd.Flags().BoolS("traces", "traces", false, "Outputs all profile samples in text form")
	rootCmd.Flags().BoolS("tree", "tree", false, "Outputs a text rendering of call graph")
	rootCmd.Flags().BoolS("trim", "trim", false, "Honor nodefraction/edgefraction/nodecount defaults")
	rootCmd.Flags().StringS("trim_path", "trim_path", "", "Path to trim from source paths before search")
	rootCmd.Flags().StringS("unit", "unit", "", "Measurement units to display")
	rootCmd.Flags().BoolS("web", "web", false, "Visualize graph through web browser")
	rootCmd.Flags().BoolS("weblist", "weblist", false, "Display annotated source in a web browser")

	rootCmd.Flag("symbolize").NoOptDefVal = " "
	rootCmd.Flag("sample_index").NoOptDefVal = " "

	rootCmd.MarkFlagsMutuallyExclusive(
		"callgrind",
		"comments",
		"disasm",
		"dot",
		"eog",
		"evince",
		"gif",
		"gv",
		"kcachegrind",
		"list",
		"pdf",
		"peek",
		"png",
		"proto",
		"ps",
		"raw",
		"svg",
		"tags",
		"text",
		"top",
		"topproto",
		"traces",
		"tree",
		"web",
		"weblist",
	)
	rootCmd.MarkFlagsMutuallyExclusive(
		"functions",
		"filefunctions",
		"files",
		"lines",
		"addresses",
	)
	rootCmd.MarkFlagsMutuallyExclusive(
		"cum",
		"flat",
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"base":      carapace.ActionFiles(),
		"diff_base": carapace.ActionFiles(),
		"http": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"output": carapace.ActionFiles(),
		"sample_index": carapace.ActionValues(
			"inuse_space",
			"inuse_objects",
			"alloc_space",
			"alloc_objects",
			"delay",
			"contentions",
			"delay",
		),
		"source_path": carapace.ActionDirectories(),
		"symbolize": carapace.ActionValuesDescribed(
			"none", "Do not attempt symbolization",
			"local", "Examine only local binaries",
			"fastlocal", "Only get function names from local binaries",
			"remote", "Do not examine local binaries",
			"force", "Force re-symbolization",
		),
		"tls_ca":   carapace.ActionFiles(),
		"tls_cert": carapace.ActionFiles(),
		"tls_key":  carapace.ActionFiles(),
		"tools":    carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
