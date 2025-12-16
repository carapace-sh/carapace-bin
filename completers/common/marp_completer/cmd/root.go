package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "marp",
	Short: "A CLI interface for Marp and Marpit based converters",
	Long:  "https://github.com/marp-team/marp-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow-local-files", false, "Allow to access local files from Markdown")
	rootCmd.Flags().String("author", "", "Define author of the slide deck")
	rootCmd.Flags().Bool("bespoke.osc", false, "Use on-screen controller")
	rootCmd.Flags().Bool("bespoke.progress", false, "Use progress bar")
	rootCmd.Flags().Bool("bespoke.transition", false, "Use transitions")
	rootCmd.Flags().StringP("config", "c", "", "Specify path to a configuration file")
	rootCmd.Flags().String("config-file", "", "Specify path to a configuration file")
	rootCmd.Flags().String("description", "", "Define description of the slide deck")
	rootCmd.Flags().String("engine", "", "Select Marpit based engine by module name or path")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().Bool("html", false, "Enable or disable HTML tags")
	rootCmd.Flags().String("image", "", "Convert the first slide page into an image file")
	rootCmd.Flags().String("image-scale", "", "The scale factor for rendered images")
	rootCmd.Flags().String("images", "", "Convert slide deck into multiple image files")
	rootCmd.Flags().StringP("input-dir", "I", "", "The base directory to find markdown and theme CSS")
	rootCmd.Flags().String("jpeg-quality", "", "Set JPEG image quality")
	rootCmd.Flags().String("keywords", "", "Define comma-separated keywords for the slide deck")
	rootCmd.Flags().Bool("no-config", false, "Prevent looking up for a configuration file")
	rootCmd.Flags().Bool("no-config-file", false, "Prevent looking up for a configuration file")
	rootCmd.Flags().Bool("notes", false, "Convert slide deck notes into a text file")
	rootCmd.Flags().String("og-image", "", "Define Open Graph image URL")
	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.Flags().Bool("pdf", false, "Convert slide deck into PDF")
	rootCmd.Flags().Bool("pdf-notes", false, "Add presenter notes to PDF as annotations")
	rootCmd.Flags().Bool("pdf-outlines", false, "Add outlines (bookmarks) to PDF")
	rootCmd.Flags().Bool("pdf-outlines.headings", false, "Make outlines from Markdown headings")
	rootCmd.Flags().Bool("pdf-outlines.pages", false, "Make outlines from slide pages")
	rootCmd.Flags().Bool("pptx", false, "Convert slide deck into PowerPoint document")
	rootCmd.Flags().BoolP("preview", "p", false, "Open preview window")
	rootCmd.Flags().BoolP("server", "s", false, "Enable server mode")
	rootCmd.Flags().String("template", "", "Choose template")
	rootCmd.Flags().String("theme", "", "Override theme by name or CSS file")
	rootCmd.Flags().String("theme-set", "", "Path to additional theme CSS files")
	rootCmd.Flags().String("title", "", "Define title of the slide deck")
	rootCmd.Flags().String("url", "", "Define canonical URL")
	rootCmd.Flags().BoolP("version", "v", false, "Show versions")
	rootCmd.Flags().BoolP("watch", "w", false, "Watch input markdowns for changes")

	rootCmd.MarkFlagsMutuallyExclusive("config", "config-file")
	rootCmd.MarkFlagsMutuallyExclusive("no-config", "no-config-file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":      carapace.ActionFiles(),
		"config-file": carapace.ActionFiles(),
		"engine": carapace.Batch(
			carapace.ActionValues("@marp-team/marpit"),
			carapace.ActionFiles(),
		).ToA(),
		"image": carapace.ActionValues("png", "jpeg").StyleF(style.ForExtension),
		"image-scale": carapace.ActionValuesDescribed(
			"1", "default",
			"2", "for PPTX conversion",
		),
		"images":    carapace.ActionValues("png", "jpeg").StyleF(style.ForExtension),
		"input-dir": carapace.ActionDirectories(),
		"output":    carapace.ActionFiles(),
		"template":  carapace.ActionValues("bare", "bespoke"),
		"theme": carapace.Batch(
			carapace.ActionValues("default", "gaia", "uncover"),
			carapace.ActionFiles(),
		).ToA(),
		"theme-set": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("input-dir").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
