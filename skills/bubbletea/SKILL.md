---
name: bubbletea
description: >
  Use when working with Bubble Tea (charmbracelet/bubbletea) — the Elm Architecture TUI
  framework for Go. Covers the Model/Update/View cycle, commands, messages, the View struct,
  renderer, input handling, program lifecycle, testing, component composition, integration
  with the Charm ecosystem (lipgloss, bubbles, huh, glamour), and production design patterns.
  Triggers on: "bubbletea", "bubble tea", "tea.Model", "tea.Cmd", "tea.Msg", "tea.Program",
  "tea.View", "tea.KeyPressMsg", "tea.WindowSizeMsg", "tea.MouseMsg", "tea.Batch",
  "tea.Sequence", "tea.Quit", "Elm Architecture", "charmbracelet", "Bubble Tea v2",
  "tea.NewProgram", "tea.Send", "tea.Tick", "tea.Every", "cursedRenderer",
  "ultraviolet", "ScreenBuffer", "ProgramContext", "compositor".
user-invocable: true
---

# Bubble Tea In-Depth Reference

Comprehensive reference for the Bubble Tea TUI framework (charmbracelet/bubbletea v2), covering the Elm Architecture in Go, the declarative View system, command/message patterns, renderer internals, input handling, and Charm ecosystem integration.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| Model, Init, Update, View, Elm Architecture, tea.Model interface, tea.View struct, NewView, View.Content, AltScreen, Cursor, ProgressBar, WindowTitle, MouseMode, KeyboardEnhancements, ReportFocus, declarative view, view composition, view diffing, viewEquals | [references/model-view.md](references/model-view.md) |
| Cmd, tea.Cmd, Batch, Sequence, Tick, Every, Quit, Suspend, Interrupt, RequestWindowSize, Println, Printf, ExecProcess, Exec, ExecCallback, BatchMsg, sequenceMsg, compactCmds, command dispatch, goroutine, async I/O | [references/commands.md](references/commands.md) |
| Msg, tea.Msg, KeyMsg, KeyPressMsg, KeyReleaseMsg, Key struct, KeyMod, Key constants, MouseMsg, MouseClickMsg, MouseReleaseMsg, MouseWheelMsg, MouseMotionMsg, Mouse struct, MouseButton, WindowSizeMsg, FocusMsg, BlurMsg, QuitMsg, InterruptMsg, SuspendMsg, ResumeMsg, ColorProfileMsg, EnvMsg, PasteMsg, CapabilityMsg, ModeReportMsg, KeyboardEnhancementsMsg, ClipboardMsg, custom message, type switch, message filtering, WithFilter | [references/messages.md](references/messages.md) |
| Program, tea.NewProgram, Program.Run, Program.Send, Program.Quit, Program.Kill, Program.Wait, ProgramOption, WithContext, WithOutput, WithInput, WithFilter, WithFPS, WithColorProfile, WithWindowSize, WithoutSignals, WithoutRenderer, WithoutSignalHandler, WithoutCatchPanics, WithEnvironment, event loop, startup sequence, shutdown, graceful exit, ErrProgramKilled, ErrInterrupted, ReleaseTerminal, RestoreTerminal, exec subprocess, LogToFile, TEA_TRACE | [references/program.md](references/program.md) |
| renderer, cursedRenderer, nilRenderer, TerminalRenderer, ScreenBuffer, ultraviolet, ANSI escape, synchronized output, mode 2026, grapheme width, mode 2027, color profile, diffing, flush, ticker, FPS, hard tabs, backspace, cursor optimization, cell buffer, StyledString | [references/renderer.md](references/renderer.md) |
| input, raw mode, term.MakeRaw, cancelreader, TerminalReader, translateInputEvent, keyboard protocol, Kitty keyboard, modifyOtherKeys, key sequence, bracketed paste, PasteMsg, PasteStartMsg, PasteEndMsg, SIGWINCH, resize, signal, SIGINT, SIGTERM, SIGTSTP, suspend, resume, focus reporting, terminal state, initTerminal, restoreTerminalState, platform-specific, Unix, Windows, BSD | [references/input-signals.md](references/input-signals.md) |
| testing, WithInput, WithOutput, WithoutSignals, WithoutRenderer, WithContext, bytes.Buffer, test model, p.Send, p.Quit, p.Kill, ErrProgramKilled, ErrInterrupted, errors.Is, test pattern, unit test, integration test, context cancellation | [references/testing.md](references/testing.md) |
| composition, component, sub-model, bubbles, textinput, textarea, viewport, spinner, list, table, help, paginator, progress, key.Binding, key.Matches, KeyMap, Focus, Blur, focus management, lipgloss, layout, JoinVertical, JoinHorizontal, glamour, huh, form, custom component, value type, reassignment pattern | [references/composition.md](references/composition.md) |
| v1, v2, migration, upgrade, View() string, View() tea.View, KeyMsg struct, KeyPressMsg, msg.Type, msg.Code, msg.Runes, msg.Text, msg.Alt, msg.Mod, EnterAltScreen, AltScreen, EnableMouseCellMotion, MouseMode, SetWindowTitle, WindowTitle, HideCursor, ShowCursor, Cursor, p.Start, p.Run, Sequentially, Sequence, charm.land, module path, import path | [references/v1-v2-migration.md](references/v1-v2-migration.md) |
| architecture, design pattern, best practice, Common struct, state machine, Draw/View split, ScreenBuffer, pubsub, program.Send, external events, bridge goroutine, fan-in, render cache, versioned cache, freeze, throttling, WithFilter, animation, dialog overlay, grace period, interface-based items, builder pattern, golden file, version bump contract, render hit counting, key map, dynamic help, theme, style struct, icon constants, status bar, TTL, double-click, ProgramContext, adaptive color, compositor layers, rebind, user-configurable keys, view-conditional help, section interface, base model embedding, input controller, mode switching, task callback, explicit sync, feature methods | [references/patterns.md](references/patterns.md) |

## Quick Guide

- **How do I create a Bubble Tea program?** → [references/program.md](references/program.md)
- **How does the Model/Update/View cycle work?** → [references/model-view.md](references/model-view.md)
- **What fields does the View struct have?** → [references/model-view.md](references/model-view.md)
- **How do I use commands for async I/O?** → [references/commands.md](references/commands.md)
- **What message types does Bubble Tea define?** → [references/messages.md](references/messages.md)
- **How do I handle keyboard input?** → [references/messages.md](references/messages.md) and [references/input-signals.md](references/input-signals.md)
- **How do I handle mouse events?** → [references/messages.md](references/messages.md)
- **How does the renderer work internally?** → [references/renderer.md](references/renderer.md)
- **How do I enter alt screen?** → [references/model-view.md](references/model-view.md)
- **How do I run a subprocess from within a program?** → [references/program.md](references/program.md)
- **How do I test a Bubble Tea program?** → [references/testing.md](references/testing.md)
- **How do I compose multiple components?** → [references/composition.md](references/composition.md)
- **How do I build a custom component?** → [references/composition.md](references/composition.md)
- **How do I manage focus between components?** → [references/composition.md](references/composition.md)
- **How do I filter or block messages?** → [references/messages.md](references/messages.md) and [references/program.md](references/program.md)
- **How do I handle terminal resize?** → [references/messages.md](references/messages.md) and [references/input-signals.md](references/input-signals.md)
- **How do I suspend and resume the program?** → [references/input-signals.md](references/input-signals.md) and [references/program.md](references/program.md)
- **How do I migrate from v1 to v2?** → [references/v1-v2-migration.md](references/v1-v2-migration.md)
- **What changed between v1 and v2?** → [references/v1-v2-migration.md](references/v1-v2-migration.md)
- **How does the event loop process messages?** → [references/program.md](references/program.md)
- **How do I debug a Bubble Tea program?** → [references/program.md](references/program.md) (LogToFile, TEA_TRACE)
- **How does the renderer diff views?** → [references/renderer.md](references/renderer.md)
- **How does input parsing work (Kitty protocol, etc.)?** → [references/input-signals.md](references/input-signals.md)
- **How do I structure a large Bubble Tea application?** → [references/patterns.md](references/patterns.md)
- **How do I bridge external events (goroutines, channels) into the update loop?** → [references/patterns.md](references/patterns.md)
- **How do I optimize rendering performance (caching, freezing, throttling)?** → [references/patterns.md](references/patterns.md)
- **How do I test complex UIs without running a tea.Program?** → [references/patterns.md](references/patterns.md)
- **How do I manage dialogs, overlays, and focus in a multi-region UI?** → [references/patterns.md](references/patterns.md)
- **How do I organize styles, themes, and key bindings in a large app?** → [references/patterns.md](references/patterns.md)
- **How do I support light and dark terminal backgrounds?** → [references/patterns.md](references/patterns.md)
- **How do I use compositor layers for overlays?** → [references/patterns.md](references/patterns.md)
- **How do I propagate mutable context across components safely?** → [references/patterns.md](references/patterns.md)
- **How do I let users rebind keys from a config file?** → [references/patterns.md](references/patterns.md)
- **How do I build a section-based UI with shared base model?** → [references/patterns.md](references/patterns.md)
- **How do I implement an input controller with mode switching?** → [references/patterns.md](references/patterns.md)
