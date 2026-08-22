package zellij

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

type tabInfo struct {
	Position                     int     `json:"position"`
	Name                         string  `json:"name"`
	Active                       bool    `json:"active"`
	PanesToHide                  int     `json:"panes_to_hide"`
	IsFullscreenActive           bool    `json:"is_fullscreen_active"`
	IsSyncPanesActive            bool    `json:"is_sync_panes_active"`
	AreFloatingPanesVisible      bool    `json:"are_floating_panes_visible"`
	OtherFocusedClients          []int   `json:"other_focused_clients"`
	ActiveSwapLayoutName         *string `json:"active_swap_layout_name"`
	IsSwapLayoutDirty            bool    `json:"is_swap_layout_dirty"`
	ViewportRows                 int     `json:"viewport_rows"`
	ViewportColumns              int     `json:"viewport_columns"`
	DisplayAreaRows              int     `json:"display_area_rows"`
	DisplayAreaColumns           int     `json:"display_area_columns"`
	SelectableTiledPanesCount    int     `json:"selectable_tiled_panes_count"`
	SelectableFloatingPanesCount int     `json:"selectable_floating_panes_count"`
	TabID                        int     `json:"tab_id"`
	HasBellNotification          bool    `json:"has_bell_notification"`
	IsFlashingBell               bool    `json:"is_flashing_bell"`
}

func actionTabs(f func(tabs []tabInfo) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("zellij", "action", "list-tabs", "--json")(func(output []byte) carapace.Action {
		var tabs []tabInfo
		if err := json.Unmarshal(output, &tabs); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(tabs)
	})
}

// ActionTabs completes tab IDs in the current session
//
//	0 (Tab #1)
//	1 (editor)
func ActionTabs() carapace.Action {
	return actionTabs(func(tabs []tabInfo) carapace.Action {
		vals := make([]string, 0)
		for _, t := range tabs {
			vals = append(vals, strconv.Itoa(t.TabID), tabDescription(t))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no tabs in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("tabs")
}

// ActionTabPositions completes tab positions (0-indexed)
//
//	0 (Tab #1)
//	1 (editor)
func ActionTabPositions() carapace.Action {
	return actionTabs(func(tabs []tabInfo) carapace.Action {
		vals := make([]string, 0)
		for _, t := range tabs {
			vals = append(vals, strconv.Itoa(t.Position), tabDescription(t))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no tabs in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("tabs")
}

// ActionTabNames completes tab names
//
//	Tab #1
//	editor
func ActionTabNames() carapace.Action {
	return actionTabs(func(tabs []tabInfo) carapace.Action {
		vals := make([]string, 0)
		for _, t := range tabs {
			vals = append(vals, t.Name, tabDescription(t))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no tabs in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("tabs")
}

// tabDescription builds a human-readable description for a tab
func tabDescription(t tabInfo) string {
	parts := make([]string, 0, 3)
	parts = append(parts, t.Name)

	if t.Active {
		parts = append(parts, "active")
	}
	if t.IsFullscreenActive {
		parts = append(parts, "fullscreen")
	}
	if t.IsSyncPanesActive {
		parts = append(parts, "sync")
	}

	return strings.Join(parts, " — ")
}
