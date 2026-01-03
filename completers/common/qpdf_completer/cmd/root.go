package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "qpdf",
	Short: "PDF transformation software",
	Long:  "https://qpdf.readthedocs.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("accessibility", "", "restrict document accessibility")
	rootCmd.Flags().StringSlice("add-attachment", nil, "start add attachment options")
	rootCmd.Flags().Bool("allow-insecure", false, "allow empty owner passwords")
	rootCmd.Flags().Bool("allow-weak-crypto", false, "allow insecure cryptographic algorithms")
	rootCmd.Flags().String("annotate", "", "restrict document annotation")
	rootCmd.Flags().String("assemble", "", "restrict document assembly")
	rootCmd.Flags().String("bits", "", "specify encryption key length")
	rootCmd.Flags().Bool("check", false, "partially check whether PDF is valid")
	rootCmd.Flags().Bool("check-linearization", false, "check linearization tables")
	rootCmd.Flags().Bool("cleartext-metadata", false, "don't encrypt metadata")
	rootCmd.Flags().Bool("coalesce-contents", false, "combine content streams")
	rootCmd.Flags().String("collate", "", "collate with --pages")
	rootCmd.Flags().Bool("completion-bash", false, "enable bash completion")
	rootCmd.Flags().Bool("completion-zsh", false, "enable zsh completion")
	rootCmd.Flags().String("compress-streams", "", "compress uncompressed streams")
	rootCmd.Flags().String("compression-level", "", "set compression level for flate")
	rootCmd.Flags().String("copy-attachments-from", "", "start copy attachment options")
	rootCmd.Flags().String("copy-encryption", "", "copy another file's encryption details")
	rootCmd.Flags().Bool("copyright", false, "show copyright information")
	rootCmd.Flags().StringSlice("creationdate", nil, "set attachment's creation date")
	rootCmd.Flags().String("decode-level", "", "control which streams to uncompress")
	rootCmd.Flags().Bool("decrypt", false, "remove encryption from input file")
	rootCmd.Flags().StringSlice("description", nil, "set attachment's description")
	rootCmd.Flags().Bool("deterministic-id", false, "generate ID deterministically")
	rootCmd.Flags().Bool("empty", false, "use empty file as input")
	rootCmd.Flags().Bool("encrypt", false, "start encryption options")
	rootCmd.Flags().String("encryption-file-password", "", "supply password for --copy-encryption")
	rootCmd.Flags().Bool("externalize-inline-images", false, "convert inline to regular images")
	rootCmd.Flags().String("extract", "", "restrict text/graphic extraction")
	rootCmd.Flags().String("file", "", "source for pages")
	rootCmd.Flags().StringSlice("filename", nil, "set attachment's displayed filename")
	rootCmd.Flags().Bool("filtered-stream-data", false, "show filtered stream data")
	rootCmd.Flags().String("flatten-annotations", "", "push annotations into content")
	rootCmd.Flags().Bool("flatten-rotation", false, "remove rotation from page dictionary")
	rootCmd.Flags().Bool("force-R5", false, "forces use of deprecated R=5 encryption")
	rootCmd.Flags().Bool("force-V4", false, "forces use of V=4 encryption handler")
	rootCmd.Flags().Bool("force-version", false, "set output PDF version")
	rootCmd.Flags().String("form", "", "restrict form filling")
	rootCmd.Flags().String("from", "", "source pages for underlay/overlay")
	rootCmd.Flags().Bool("generate-appearances", false, "generate appearances for form fields")
	rootCmd.Flags().String("help", "", "provide help")
	rootCmd.Flags().Bool("ignore-xref-streams", false, "use xref tables rather than streams")
	rootCmd.Flags().String("ii-min-bytes", "", "set minimum size for --externalize-inline-images")
	rootCmd.Flags().Bool("is-encrypted", false, "silently test whether a file is encrypted")
	rootCmd.Flags().String("job-json-file", "", "job JSON file")
	rootCmd.Flags().Bool("job-json-help", false, "show format of job JSON")
	rootCmd.Flags().String("jpeg-quality", "", "set jpeg quality level for jpeg")
	rootCmd.Flags().String("json", "", "show file in JSON format")
	rootCmd.Flags().String("json-help", "", "show format of JSON output")
	rootCmd.Flags().Bool("json-input", false, "input file is qpdf JSON")
	rootCmd.Flags().StringSlice("json-key", nil, "limit which keys are in JSON output")
	rootCmd.Flags().StringSlice("json-object", nil, "limit which objects are in JSON")
	rootCmd.Flags().String("json-output", "", "apply defaults for JSON serialization")
	rootCmd.Flags().String("json-stream-data", "", "how to handle streams in json output")
	rootCmd.Flags().String("json-stream-prefix", "", "prefix for json stream data files")
	rootCmd.Flags().String("keep-files-open", "", "manage keeping multiple files open")
	rootCmd.Flags().String("keep-files-open-threshold", "", "set threshold for --keep-files-open")
	rootCmd.Flags().Bool("keep-inline-images", false, "exclude inline images from optimization")
	rootCmd.Flags().StringSlice("key", nil, "specify attachment key")
	rootCmd.Flags().Bool("linearize", false, "linearize (web-optimize) output")
	rootCmd.Flags().String("linearize-pass1", "", "save pass 1 of linearization")
	rootCmd.Flags().Bool("list-attachments", false, "list embedded files")
	rootCmd.Flags().StringSlice("mimetype", nil, "attachment mime type, e.g. application/pdf")
	rootCmd.Flags().String("min-version", "", "set minimum PDF version")
	rootCmd.Flags().StringSlice("moddate", nil, "set attachment's modification date")
	rootCmd.Flags().String("modify", "", "restrict document modification")
	rootCmd.Flags().String("modify-other", "", "restrict other modifications")
	rootCmd.Flags().Bool("newline-before-endstream", false, "force a newline before endstream")
	rootCmd.Flags().Bool("no-original-object-ids", false, "omit original object IDs in qdf")
	rootCmd.Flags().Bool("no-warn", false, "suppress printing of warning messages")
	rootCmd.Flags().String("normalize-content", "", "fix newlines in content streams")
	rootCmd.Flags().String("object-streams", "", "control use of object streams")
	rootCmd.Flags().String("oi-min-area", "", "minimum area for --optimize-images")
	rootCmd.Flags().String("oi-min-height", "", "minimum height for --optimize-images")
	rootCmd.Flags().String("oi-min-width", "", "minimum width for --optimize-images")
	rootCmd.Flags().Bool("optimize-images", false, "use efficient compression for images")
	rootCmd.Flags().String("overlay", "", "begin overlay options")
	rootCmd.Flags().String("owner-password", "", "specify owner password")
	rootCmd.Flags().Bool("pages", false, "show number of pages")
	rootCmd.Flags().String("password", "", "silently test a file's password")
	rootCmd.Flags().String("password-file", "", "read password from a file")
	rootCmd.Flags().Bool("password-is-hex-key", false, "provide hex-encoded encryption key")
	rootCmd.Flags().String("password-mode", "", "tweak how qpdf encodes passwords")
	rootCmd.Flags().String("prefix", "", "key prefix for copying attachments")
	rootCmd.Flags().Bool("preserve-unreferenced", false, "preserve unreferenced objects")
	rootCmd.Flags().Bool("preserve-unreferenced-resources", false, "use --remove-unreferenced-resources=no")
	rootCmd.Flags().String("print", "", "restrict printing")
	rootCmd.Flags().Bool("progress", false, "show progress when writing")
	rootCmd.Flags().Bool("qdf", false, "enable viewing PDF code in a text editor")
	rootCmd.Flags().String("range", "", "page range")
	rootCmd.Flags().Bool("raw-stream-data", false, "show raw stream data")
	rootCmd.Flags().Bool("recompress-flate", false, "uncompress and recompress flate")
	rootCmd.Flags().String("remove-attachment", "", "remove an embedded file")
	rootCmd.Flags().Bool("remove-info", false, "remove file information")
	rootCmd.Flags().Bool("remove-metadata", false, "remove metadata")
	rootCmd.Flags().Bool("remove-page-labels", false, "remove explicit page numbers")
	rootCmd.Flags().Bool("remove-restrictions", false, "remove security restrictions from input file")
	rootCmd.Flags().Bool("remove-structure", false, "remove metadata")
	rootCmd.Flags().Bool("remove-unreferenced-resources", false, "remove unreferenced page resources")
	rootCmd.Flags().String("repeat", "", "overlay/underlay pages to repeat")
	rootCmd.Flags().Count("replace", "replace attachment with same key")
	rootCmd.Flags().Bool("replace-input", false, "overwrite input with output")
	rootCmd.Flags().Bool("report-memory-usage", false, "best effort report of memory usage")
	rootCmd.Flags().Bool("requires-password", false, "silently test a file's password")
	rootCmd.Flags().String("rotate", "", "rotate pages")
	rootCmd.Flags().String("set-page-labels", "", "number pages for the entire document")
	rootCmd.Flags().String("show-attachment", "", "export an embedded file")
	rootCmd.Flags().Bool("show-crypto", false, "show available crypto providers")
	rootCmd.Flags().Bool("show-encryption", false, "information about encrypted files")
	rootCmd.Flags().Bool("show-encryption-key", false, "show key with --show-encryption")
	rootCmd.Flags().Bool("show-linearization", false, "show linearization hint tables")
	rootCmd.Flags().Bool("show-npages", false, "show number of pages")
	rootCmd.Flags().String("show-object", "", "show contents of an object")
	rootCmd.Flags().Bool("show-pages", false, "display page dictionary information")
	rootCmd.Flags().Bool("show-xref", false, "show cross reference data")
	rootCmd.Flags().String("split-pages", "", "write pages to separate files")
	rootCmd.Flags().Bool("static-aes-iv", false, "use a fixed AES vector")
	rootCmd.Flags().Bool("static-id", false, "use a fixed document ID")
	rootCmd.Flags().String("stream-data", "", "show filtered stream data")
	rootCmd.Flags().Bool("suppress-password-recovery", false, "don't try different password encodings")
	rootCmd.Flags().Bool("suppress-recovery", false, "suppress error recovery")
	rootCmd.Flags().Bool("test-json-schema", false, "test generated json against schema")
	rootCmd.Flags().String("to", "", "destination pages for underlay/overlay")
	rootCmd.Flags().Bool("underlay", false, "begin underlay options")
	rootCmd.Flags().String("update-from-json", "", "update a PDF from qpdf JSON")
	rootCmd.Flags().String("use-aes", "", "use AES with 128-bit encryption")
	rootCmd.Flags().String("user-password", "", "specify user password")
	rootCmd.Flags().Bool("verbose", false, "print additional information")
	rootCmd.Flags().Bool("version", false, "show qpdf version")
	rootCmd.Flags().Bool("warning-exit-0", false, "exit 0 even with warnings")
	rootCmd.Flags().Bool("with-images", false, "include image details with --show-pages")
	rootCmd.Flags().Bool("zopfli", false, "indicate whether zopfli is enabled and active")

	rootCmd.Flag("accessibility").NoOptDefVal = " "
	rootCmd.Flag("annotate").NoOptDefVal = " "
	rootCmd.Flag("assemble").NoOptDefVal = " "
	rootCmd.Flag("collate").NoOptDefVal = " "
	rootCmd.Flag("compress-streams").NoOptDefVal = " "
	rootCmd.Flag("extract").NoOptDefVal = " "
	rootCmd.Flag("form").NoOptDefVal = " "
	rootCmd.Flag("from").NoOptDefVal = " "
	rootCmd.Flag("help").NoOptDefVal = " "
	rootCmd.Flag("json").NoOptDefVal = " "
	rootCmd.Flag("json-help").NoOptDefVal = " "
	rootCmd.Flag("json-output").NoOptDefVal = " "
	rootCmd.Flag("keep-files-open").NoOptDefVal = " "
	rootCmd.Flag("modify-other").NoOptDefVal = " "
	rootCmd.Flag("normalize-content").NoOptDefVal = " "
	rootCmd.Flag("split-pages").NoOptDefVal = " "
	rootCmd.Flag("use-aes").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"accessibility":         carapace.ActionValues("y", "n"),
		"add-attachment":        carapace.ActionFiles(),
		"annotate":              carapace.ActionValues("y", "n"),
		"assemble":              carapace.ActionValues("y", "n"),
		"bits":                  carapace.ActionValues("48", "128", "256"),
		"compress-streams":      carapace.ActionValues("y", "n"),
		"copy-attachments-from": carapace.ActionFiles(),
		"copy-encryption":       carapace.ActionFiles(),
		"extract":               carapace.ActionValues("y", "n"),
		"file":                  carapace.ActionFiles(),
		"flatten-annotations": carapace.ActionValuesDescribed(
			"all", "include all annotations that are not marked invisible or hidden",
			"print", "only include annotations that should appear when the page is printed",
			"screen", "omit annotations that should not appear on the screen",
		),
		"form": carapace.ActionValues("y", "n"),
		"help": carapace.Batch(
			carapace.ActionValuesDescribed(

				"all", "all available help",
				"add-attachment", "attach (embed) files",
				"advanced-control", "tweak qpdf's behavior",
				"attachments", "work with embedded files",
				"completion", "shell completion",
				"copy-attachments", "copy attachments from another file",
				"encryption", "create encrypted files",
				"exit-status", "meanings of qpdf's exit codes",
				"general", "general options",
				"help", "information about qpdf",
				"inspection", "inspect PDF files",
				"json", "JSON output for PDF information",
				"modification", "change parts of the PDF",
				"overlay-underlay", "overlay/underlay pages from other PDF files",
				"page-ranges", "page range syntax",
				"page-selection", "select pages from one or more files",
				"pdf-dates", "PDF date format",
				"testing", "options for testing or debugging",
				"transformation", "make structural PDF changes",
				"usage", "basic invocation",
			),
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				values := []string{}

				rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
					values = append(values, "--"+f.Name, f.Usage)
				})

				return carapace.ActionValuesDescribed(values...)
			}),
		).ToA(),
		"job-json-file":    carapace.ActionFiles(),
		"json-stream-data": carapace.ActionValues("none", "inline", "file"),
		"keep-files-open":  carapace.ActionValues("y", "n"),
		"linearize-pass1":  carapace.ActionFiles(),
		"modify": carapace.ActionValuesDescribed(
			"all", "allow full document modification",
			"annotate", "form permissions plus commenting and modifying forms",
			"assembly", "allow document assembly only",
			"form", "assembly permissions plus filling in form fields and signing",
			"none", "allow no modifications",
		),
		"modify-other":      carapace.ActionValues("y", "n"),
		"normalize-content": carapace.ActionValues("y", "n"),
		"object-streams": carapace.ActionValuesDescribed(
			"preserve", "preserve original object streams, if any",
			"disable", "create output files with no object streams",
			"generate", "create object streams, and compress objects when possible",
		),
		"overlay":       carapace.ActionFiles(),
		"pages":         carapace.ActionFiles(),
		"password-file": carapace.ActionFiles(),
		"password-mode": carapace.ActionValues("auto", "unicode", "bytes", "hex-bytes"),
		"print": carapace.ActionValuesDescribed(
			"none", "disallow printing",
			"low", "allow low-resolution printing only",
			"full", "allow full printing",
		),
		"remove-unreferenced-resources": carapace.ActionValues("auto", "yes", "no").StyleF(style.ForKeyword),
		"stream-data": carapace.ActionValuesDescribed(
			"compress", "recompress stream data when possible",
			"preserve", "leave all stream data as is",
			"uncompress", "uncompress stream data compressed with generalized filters when possible",
		),
		"underlay":         carapace.ActionFiles(),
		"update-from-json": carapace.ActionFiles(),
		"use-aes":          carapace.ActionValues("y", "n"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
