package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var incident_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List project incidents.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_listCmd).Standalone()

	incident_listCmd.Flags().BoolP("all", "A", false, "Get all incidents.")
	incident_listCmd.Flags().StringP("assignee", "a", "", "Filter incident by assignee <username>.")
	incident_listCmd.Flags().String("author", "", "Filter incident by author <username>.")
	incident_listCmd.Flags().BoolP("closed", "c", false, "Get only closed incidents.")
	incident_listCmd.Flags().BoolP("confidential", "C", false, "Filter by confidential incidents.")
	incident_listCmd.Flags().StringP("epic", "e", "", "List issues belonging to a given epic (requires --group, no pagination support).")
	incident_listCmd.PersistentFlags().StringP("group", "g", "", "Select a group or subgroup. Ignored if a repo argument is set.")
	incident_listCmd.Flags().String("in", "", "search in: title, description.")
	incident_listCmd.Flags().StringSliceP("label", "l", nil, "Filter incident by label <name>.")
	incident_listCmd.Flags().StringP("milestone", "m", "", "Filter incident by milestone <id>.")
	incident_listCmd.Flags().BoolP("mine", "M", false, "Filter only incidents assigned to me.")
	incident_listCmd.Flags().String("not-assignee", "", "Filter incident by not being assigned to <username>.")
	incident_listCmd.Flags().String("not-author", "", "Filter incident by not being by author(s) <username>.")
	incident_listCmd.Flags().StringSlice("not-label", nil, "Filter incident by lack of label <name>.")
	incident_listCmd.Flags().BoolP("opened", "o", false, "Get only open incidents.")
	incident_listCmd.Flags().String("order", "", "Order incident by <field>. Order options: created_at, updated_at, priority, due_date, relative_position, label_priority, milestone_due, popularity, weight.")
	incident_listCmd.Flags().StringP("output", "O", "", "Options: 'text' or 'json'.")
	incident_listCmd.Flags().StringP("output-format", "F", "", "Options: 'details', 'ids', 'urls'.")
	incident_listCmd.Flags().StringP("page", "p", "", "Page number.")
	incident_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	incident_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	incident_listCmd.Flags().String("search", "", "Search <string> in the fields defined by '--in'.")
	incident_listCmd.Flags().String("sort", "", "Return incident sorted in asc or desc order.")
	incident_listCmd.Flag("mine").Hidden = true
	incident_listCmd.Flag("opened").Hidden = true
	incidentCmd.AddCommand(incident_listCmd)

	// TODO complete epic
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
		"order":         carapace.ActionValues("created_at", "updated_at", "priority", "due_date", "relative_position", "label_priority", "milestone_due", "popularity", "weight"),
		"output-format": carapace.ActionValues("details", "ids", "urls"),
		"repo":          action.ActionRepo(incident_listCmd),
		"sort":          carapace.ActionValues("asc", "desc"),
	})
}
