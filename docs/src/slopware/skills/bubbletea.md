# bubbletea

Bubble Tea — the Elm Architecture TUI framework for Go (charmbracelet/bubbletea).

| reference                              | description                                          |
|----------------------------------------|------------------------------------------------------|
| [model-view][ref-model-view]           | Model/Init/Update/View cycle, View struct, alt screen |
| [commands][ref-commands]               | Cmd, Batch, Sequence, Tick, Every, async I/O         |
| [messages][ref-messages]               | KeyMsg, MouseMsg, WindowSizeMsg, custom messages     |
| [program][ref-program]                 | Program, Run, Send, Quit, Wait, ProgramOptions        |
| [renderer][ref-renderer]               | cursedRenderer, ScreenBuffer, ultraviolet, diffing   |
| [input-signals][ref-input-signals]     | raw mode, key sequences, bracketed paste, signals    |
| [testing][ref-testing]                 | WithInput, WithoutRenderer, test patterns             |
| [composition][ref-composition]         | sub-models, bubbles, lipgloss, glamour, focus         |
| [v1-v2-migration][ref-v1-v2-migration] | View() string vs View() tea.View, KeyMsg changes      |
| [patterns][ref-patterns]               | architecture, state machine, pubsub, render cache     |

[ref-model-view]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/model-view.md
[ref-commands]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/commands.md
[ref-messages]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/messages.md
[ref-program]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/program.md
[ref-renderer]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/renderer.md
[ref-input-signals]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/input-signals.md
[ref-testing]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/testing.md
[ref-composition]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/composition.md
[ref-v1-v2-migration]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/v1-v2-migration.md
[ref-patterns]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/bubbletea/references/patterns.md