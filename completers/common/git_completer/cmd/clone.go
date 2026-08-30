package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "Clone a repository into a new directory",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().Bool("also-filter-submodules", false, "apply partial clone filter to submodules")
	cloneCmd.Flags().Bool("bare", false, "create a bare repository")
	cloneCmd.Flags().StringP("branch", "b", "", "checkout <branch> instead of the remote's HEAD")
	cloneCmd.Flags().String("bundle-uri", "", "fetch additional objects from this URI")
	cloneCmd.Flags().StringArrayP("config", "c", nil, "set config inside the new repository")
	cloneCmd.Flags().String("depth", "", "create a shallow clone of that depth")
	cloneCmd.Flags().Bool("dissociate", false, "use --reference only while cloning")
	cloneCmd.Flags().StringArray("filter", nil, "object filtering")
	cloneCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	cloneCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	cloneCmd.Flags().StringP("jobs", "j", "", "number of submodules cloned in parallel")
	cloneCmd.Flags().BoolP("local", "l", false, "to clone from a local repository")
	cloneCmd.Flags().Bool("mirror", false, "create a mirror repository (implies bare)")
	cloneCmd.Flags().Bool("naked", false, "create a bare repository")
	cloneCmd.Flags().BoolP("no-checkout", "n", false, "don't create a checkout")
	cloneCmd.Flags().Bool("no-hardlinks", false, "don't use local hardlinks, always copy")
	cloneCmd.Flags().Bool("no-local", false, "override --local, as if file:/// URL was given")
	cloneCmd.Flags().Bool("no-reject-shallow", false, "do not reject shallow repository")
	cloneCmd.Flags().Bool("no-remote-submodules", false, "do not use remote-tracking branch for submodules")
	cloneCmd.Flags().Bool("no-shallow-submodules", false, "do not clone submodules as shallow")
	cloneCmd.Flags().Bool("no-single-branch", false, "clone history leading up to each branch")
	cloneCmd.Flags().Bool("no-tags", false, "don't clone any tags, and make later fetches not to follow them")
	cloneCmd.Flags().StringP("origin", "o", "", "use <name> instead of 'origin' to track upstream")
	cloneCmd.Flags().Bool("progress", false, "force progress reporting")
	cloneCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	cloneCmd.Flags().String("recurse-submodules", "", "initialize submodules in the clone")
	cloneCmd.Flags().String("recursive", "", "alias of --recurse-submodules")
	cloneCmd.Flags().String("ref-format", "", "specify the reference format for the repository")
	cloneCmd.Flags().StringArray("reference", nil, "reference repository")
	cloneCmd.Flags().StringArray("reference-if-able", nil, "reference repository")
	cloneCmd.Flags().Bool("reject-shallow", false, "do not clone shallow repository")
	cloneCmd.Flags().Bool("remote-submodules", false, "any cloned submodules will use their remote-tracking branch")
	cloneCmd.Flags().String("revision", "", "clone a specific revision")
	cloneCmd.Flags().String("separate-git-dir", "", "separate git dir from working tree")
	cloneCmd.Flags().StringArray("server-option", nil, "option to transmit")
	cloneCmd.Flags().StringArray("shallow-exclude", nil, "deepen history of shallow clone, excluding rev")
	cloneCmd.Flags().String("shallow-since", "", "create a shallow clone since a specific time")
	cloneCmd.Flags().Bool("shallow-submodules", false, "any cloned submodules will be shallow")
	cloneCmd.Flags().BoolP("shared", "s", false, "setup as shared repository")
	cloneCmd.Flags().Bool("single-branch", false, "clone only one branch, HEAD or --branch")
	cloneCmd.Flags().Bool("sparse", false, "initialize sparse-checkout file to include only files at root")
	cloneCmd.Flags().Bool("tags", false, "clone tags")
	cloneCmd.Flags().String("template", "", "directory from which templates will be used")
	cloneCmd.Flags().StringP("upload-pack", "u", "", "path to git-upload-pack on the remote")
	cloneCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	cloneCmd.Flag("naked").Hidden = true
	cloneCmd.Flag("recurse-submodules").NoOptDefVal = "."
	cloneCmd.Flag("recursive").NoOptDefVal = "."

	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"config": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionConfigs().NoSpace()
			default:
				return git.ActionConfigValues(c.Parts[0])
			}
		}).UniqueList(" "),
		"filter":             git.ActionObjectFilters(),
		"recurse-submodules": carapace.ActionValues("yes", "on-demand").StyleF(style.ForKeyword),
		"recursive":          carapace.ActionValues("yes", "on-demand").StyleF(style.ForKeyword),
		"ref-format":         carapace.ActionValues("files", "flat", "tree"),
		"reference":          carapace.ActionDirectories(),
		"reference-if-able":  carapace.ActionDirectories(),
		"revision":           git.ActionRefs(git.RefOption{}.Default()),
		"separate-git-dir":   carapace.ActionFiles(),
		"server-option":      carapace.ActionValues(),
		"shallow-exclude":    git.ActionRefs(git.RefOption{}.Default()),
		"shallow-since":      time.ActionDate(),
		"template":           carapace.ActionDirectories(),
		"upload-pack":        carapace.ActionFiles(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		carapace.ActionDirectories(),
	)

	carapace.Gen(cloneCmd).DashAnyCompletion(
		carapace.ActionPositional(cloneCmd),
	)
}
