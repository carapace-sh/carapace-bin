package wezterm

import (
	"encoding/json"
	"strconv"

	"github.com/carapace-sh/carapace"
)

type pane struct {
	WindowID  int    `json:"window_id"`
	TabID     int    `json:"tab_id"`
	PaneID    int    `json:"pane_id"`
	Workspace string `json:"workspace"`
	Size      struct {
		Rows        int `json:"rows"`
		Cols        int `json:"cols"`
		PixelWidth  int `json:"pixel_width"`
		PixelHeight int `json:"pixel_height"`
		Dpi         int `json:"dpi"`
	} `json:"size"`
	Title            string `json:"title"`
	Cwd              string `json:"cwd"`
	CursorX          int    `json:"cursor_x"`
	CursorY          int    `json:"cursor_y"`
	CursorShape      string `json:"cursor_shape"`
	CursorVisibility string `json:"cursor_visibility"`
	LeftCol          int    `json:"left_col"`
	TopRow           int    `json:"top_row"`
	TabTitle         string `json:"tab_title"`
	WindowTitle      string `json:"window_title"`
	IsActive         bool   `json:"is_active"`
	IsZoomed         bool   `json:"is_zoomed"`
	TtyName          string `json:"tty_name"`
}

func actionPanes(f func(panes []pane) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("wezterm", "cli", "list", "--format", "json")(func(output []byte) carapace.Action {
		var panes []pane
		if err := json.Unmarshal(output, &panes); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(panes)
	})
}

// ActionPanes completes panes
// 0 (zsh)
// 1 (fish)
func ActionPanes() carapace.Action {
	return actionPanes(func(panes []pane) carapace.Action {
		vals := make([]string, 0)
		for _, p := range panes {
			vals = append(vals, strconv.Itoa(p.PaneID), p.Title)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionPaneDirections completes pane directions
//
//	Up
//	Down
func ActionPaneDirections() carapace.Action {
	return carapace.ActionValues("Up", "Down", "Left", "Right", "Next", "Prev")
}
