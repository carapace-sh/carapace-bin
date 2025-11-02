package but

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type entry struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Details   struct {
		Version   int    `json:"version"`
		Operation string `json:"operation"`
		Title     string `json:"title"`
		Body      any    `json:"body"`
		Trailers  []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"trailers"`
	} `json:"details"`
}

func (e entry) style() string {
	switch e.Details.Operation {
	case "CreateCommit": // "CREATE",
		return style.Green
	case "CreateBranch": // "BRANCH",
		return style.Magenta
	case "AmendCommit": // "AMEND",
		return style.Yellow
	case "UndoCommit": // "UNDO",
		return style.Red
	case "SquashCommit": // "SQUASH",
		return style.Default
	case "UpdateCommitMessage": // "REWORD",
		return style.Yellow
	case "MoveCommit": // "MOVE",
		return style.Cyan
	case "RestoreFromSnapshot": // "RESTORE",
		return style.Red
	case "ReorderCommit": // "REORDER",
		return style.Cyan
	case "InsertBlankCommit": // "INSERT",
		return style.Default
	case "MoveHunk": // "MOVE_HUNK",
		return style.Cyan
	case "ReorderBranches": // "REORDER_BRANCH",
		return style.Default
	case "UpdateWorkspaceBase": // "UPDATE_BASE",
		return style.Default
	case "UpdateBranchName": // "RENAME",
		return style.Default
	case "GenericBranchUpdate": // "BRANCH_UPDATE",
		return style.Default
	case "ApplyBranch": // "APPLY",
		return style.Default
	case "UnapplyBranch": // "UNAPPLY",
		return style.Default
	case "DeleteBranch": // "DELETE",
		return style.Default
	case "DiscardChanges": // "DISCARD",
		return style.Default
	default: // "OTHER",
		return style.Default
	}
}

// ActionOplogEntries completes oplog entries
//
//	b59e6ab3fa3d (2025-11-02 18:15:22 CreateCommit)
//	ba997064b522 (2025-11-02 18:40:58 CreateCommit)
func ActionOplogEntries() carapace.Action {
	return carapace.ActionExecCommand("but", "--json", "oplog")(func(output []byte) carapace.Action {
		var entries []entry
		if err := json.Unmarshal(output, &entries); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, e := range entries {
			t := time.Unix(int64(e.CreatedAt), 0).UTC() // TODO UTC intentional?
			vals = append(vals, e.ID[:12], fmt.Sprintf("%v %v", t.Format(time.DateTime), e.Details.Operation), e.style())
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
