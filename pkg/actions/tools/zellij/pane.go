package zellij

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
)

type paneEntry struct {
	ID           int     `json:"id"`
	IsPlugin     bool    `json:"is_plugin"`
	IsFocused    bool    `json:"is_focused"`
	IsFloating   bool    `json:"is_floating"`
	IsFullscreen bool    `json:"is_fullscreen"`
	Title        string  `json:"title"`
	Exited       bool    `json:"exited"`
	ExitStatus   *int    `json:"exit_status"`
	IsHeld       bool    `json:"is_held"`
	IsSelectable bool    `json:"is_selectable"`
	TabID        int     `json:"tab_id"`
	TabPosition  int     `json:"tab_position"`
	TabName      string  `json:"tab_name"`
	PaneCommand  *string `json:"pane_command"`
	PaneCwd      *string `json:"pane_cwd"`
	PluginURL    *string `json:"plugin_url"`
	TerminalCmd  *string `json:"terminal_command"`
}

func actionPanes(f func(panes []paneEntry) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("zellij", "action", "list-panes", "--json")(func(output []byte) carapace.Action {
		var panes []paneEntry
		if err := json.Unmarshal(output, &panes); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(panes)
	})
}

// ActionPanes completes pane IDs in the current session
//
//	terminal_1 (bash — Tab #1)
//	plugin_2 (status-bar — Tab #1)
func ActionPanes() carapace.Action {
	return actionPanes(func(panes []paneEntry) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			id := paneIDString(p)
			if id == "" {
				continue
			}
			vals = append(vals, id, paneDescription(p))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no panes in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("panes")
}

// ActionSelectablePanes completes selectable pane IDs (excludes non-selectable plugin panes like status-bar)
//
//	terminal_1 (bash — Tab #1)
//	terminal_3 (vim — editor)
func ActionSelectablePanes() carapace.Action {
	return actionPanes(func(panes []paneEntry) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			if !p.IsSelectable {
				continue
			}
			id := paneIDString(p)
			if id == "" {
				continue
			}
			vals = append(vals, id, paneDescription(p))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no selectable panes in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("panes")
}

// ActionPaneIds completes bare numeric pane IDs
//
//	1 (bash — Tab #1)
//	2 (status-bar — Tab #1)
func ActionPaneIds() carapace.Action {
	return actionPanes(func(panes []paneEntry) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			id := paneIDString(p)
			if id == "" {
				continue
			}
			vals = append(vals, strconv.Itoa(p.ID), paneDescription(p))
		}
		if len(vals) == 0 {
			return carapace.ActionMessage("no panes in current session")
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("panes")
}

// paneIDString returns the full pane ID string (e.g. "terminal_1", "plugin_2")
func paneIDString(p paneEntry) string {
	if p.IsPlugin {
		return "plugin_" + strconv.Itoa(p.ID)
	}
	return "terminal_" + strconv.Itoa(p.ID)
}

// paneDescription builds a human-readable description for a pane
func paneDescription(p paneEntry) string {
	parts := make([]string, 0, 4)

	if p.PaneCommand != nil && *p.PaneCommand != "" {
		parts = append(parts, *p.PaneCommand)
	} else if p.PluginURL != nil && *p.PluginURL != "" {
		parts = append(parts, *p.PluginURL)
	} else if p.Title != "" {
		parts = append(parts, p.Title)
	}

	if p.TabName != "" {
		parts = append(parts, p.TabName)
	}

	if p.IsFloating {
		parts = append(parts, "floating")
	}
	if p.IsFullscreen {
		parts = append(parts, "fullscreen")
	}
	if p.Exited {
		parts = append(parts, "exited")
	}

	return strings.Join(parts, " — ")
}
