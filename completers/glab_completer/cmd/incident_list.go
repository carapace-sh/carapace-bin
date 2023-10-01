package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List project incidents",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_listCmd).Standalone()

	incident_listCmd.Flags().BoolP("all", "A", false, "Get all incidents")
	incident_listCmd.Flags().StringP("assignee", "a", "", "Filter incident by assignee <username>")
	incident_listCmd.Flags().String("author", "", "Filter incident by author <username>")
	incident_listCmd.Flags().BoolP("closed", "c", false, "Get only closed incidents")
	incident_listCmd.Flags().BoolP("confidential", "C", false, "Filter by confidential incidents")
	incident_listCmd.PersistentFlags().StringP("group", "g", "", "Select a group/subgroup. This option is ignored if a repo argument is set.")
	incident_listCmd.Flags().String("in", "", "search in {title|description}")
	incident_listCmd.Flags().StringSliceP("label", "l", []string{}, "Filter incident by label <name>")
	incident_listCmd.Flags().StringP("milestone", "m", "", "Filter incident by milestone <id>")
	incident_listCmd.Flags().BoolP("mine", "M", false, "Filter only incidents assigned to me")
	incident_listCmd.Flags().StringSlice("not-assignee", []string{}, "Filter incident by not being assigneed to <username>")
	incident_listCmd.Flags().StringSlice("not-author", []string{}, "Filter by not being by author(s) <username>")
	incident_listCmd.Flags().StringSlice("not-label", []string{}, "Filter incident by lack of label <name>")
	incident_listCmd.Flags().BoolP("opened", "o", false, "Get only open incidents")
	incident_listCmd.Flags().StringP("output-format", "F", "", "One of 'details', 'ids', or 'urls'")
	incident_listCmd.Flags().StringP("page", "p", "", "Page number")
	incident_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	incident_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	incident_listCmd.Flags().String("search", "", "Search <string> in the fields defined by --in")
	incident_listCmd.Flag("mine").Hidden = true
	incident_listCmd.Flag("opened").Hidden = true
	incidentCmd.AddCommand(incident_listCmd)

	carapace.Gen(incident_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      action.ActionProjectMembers(incident_listCmd),
		"author":        action.ActionUsers(incident_listCmd),
		"group":         action.ActionGroups(incident_listCmd).UniqueList(","),
		"in":            carapace.ActionValues("title", "description"),
		"label":         action.ActionLabels(incident_listCmd).UniqueList(","),
		"milestone":     action.ActionMilestones(incident_listCmd),
		"not-assignee":  action.ActionProjectMembers(incident_listCmd).UniqueList(","),
		"not-author":    action.ActionUsers(incident_listCmd).UniqueList(","),
		"not-label":     action.ActionLabels(incident_listCmd).UniqueList(","),
		"output-format": carapace.ActionValues("details", "ids", "urls"),
		"repo":          action.ActionRepo(incident_listCmd),
	})
}
