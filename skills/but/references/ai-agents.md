# AI Agents and Skills

GitButler's AI features: agent setup, skills system, `--ai` commit message generation, and parallel agent workflows.

> **Source of truth**: <https://docs.gitbutler.com/ai-agents/overview> and <https://docs.gitbutler.com/ai-agents/tuning-agent-behavior>

## `but agent` — AI Agent Setup

Parent command for AI coding agent configuration.

### `but agent setup`

Interactive wizard for tuning version control behavior for AI coding agents. Easier to set up than the documentation-based approach.

```bash
but agent setup              # Interactive wizard
but agent setup --print      # Print default steering text without modifying files
```

| Flag | Description |
|------|-------------|
| `--print` | Print the default generated steering text without prompting or modifying files |

#### What the Wizard Does

1. Installs the GitButler skill (for Codex, Claude Code, Cursor, GitHub Copilot, Windsurf/Devin, OpenCode, or Agent Skills)
2. Writes workflow steering instructions into supported agent instruction files (globally, locally in repo, or both)
3. Runs `but setup` for the repository when GitButler needs workspace mode

Before writing anything, the wizard shows the exact paths, files, and generated text.

The generated instructions tell the agent to **use `but` commands instead of** Git checkout/stash/add/commit/rebase operations.

### Tuning Agent Behavior (Optional Policies)

The agent wizard can write standing instructions including:

| Policy | Description |
|--------|-------------|
| Amend local fixes | Fold follow-up fixes into the right unpublished commits |
| Checkpoint commits | Create savepoints after each completed turn |
| Stacked PRs | Stack dependent work and create PRs with `but pr` (not `gh`) |
| Custom branch names | e.g. `feature/<ticket-id>-<short-description>` |
| Custom commit messages | e.g. Conventional Commits (`feat:`, `fix:`) |
| "Ship it" publish | A short phrase that commits, pushes, and opens/updates a PR |
| Auto-update from main | Run `but pull --check` / `but pull` automatically |
| Draft PRs by default | Create PRs as drafts unless told ready |
| Recovery snapshots | Run `but oplog snapshot` before large history edits |
| Split unrelated hunks | Split files by hunk when they contain unrelated changes |

## `but skill` — Manage AI Agent Skills

Parent command for skill management.

### `but skill install`

Install GitButler skill files for coding agents. Interactive by default.

```bash
but skill install                           # Interactive
but skill install --global                  # Global install (home directory)
but skill install --path .agents/skills     # Custom path
but skill install --detect                  # Auto-detect existing installation
```

| Flag | Short | Description |
|------|-------|-------------|
| `--detect` | `-d` | Auto-detect where to install by finding existing installation |
| `--global` | `-g` | Install globally instead of in the current repository |
| `--path` | `-p` | Custom installation path (relative to repo root or absolute) |

Interactive prompts:
1. **Scope**: Current repository or global home directory
2. **Format**: Agent Skills / `.agents`, Claude Code, OpenCode, Codex, GitHub Copilot, Cursor, Windsurf

Positional argument (optional): file path for skill file (completed relative to git worktree).

### `but skill check`

Check if installed GitButler skills are up to date with the CLI version.

```bash
but skill check                    # Check all installations
but skill check --update           # Auto-update outdated skills
but skill check --global           # Check only global installations
but skill check --local            # Check only local installations
```

| Flag | Short | Description |
|------|-------|-------------|
| `--global` | `-g` | Only check global installations (home directory) |
| `--local` | `-l` | Only check local installations (current repository) |
| `--update` | `-u` | Automatically update any outdated skills found |

## `--ai` Flag

Several commands support the `--ai` flag for AI-generated content:

| Command | AI Usage |
|---------|----------|
| `but commit --ai` | Generate commit message from changes being committed |
| `but commit --ai="instructions"` | Generate with specific user guidance |
| `but squash --ai` | Generate combined commit message when squashing |
| `but squash --ai="instructions"` | Generate with specific user guidance |
| `but branch show --ai` | Generate AI summary of branch changes |

The `--ai` flag uses `NoOptDefVal = " "`, so `--ai` by itself generates a message, or `--ai="your instructions"` provides guidance (equals sign required for value).

## AI Provider Configuration

AI providers are configured via `but config ai`. See [config.md](config.md) for details.

Supported providers:
- **Anthropic** (Claude) — `but config ai anthropic`
- **OpenAI** — `but config ai openai`
- **OpenRouter** — `but config ai openrouter`
- **Ollama** (local) — `but config ai ollama`
- **LM Studio** (local) — `but config ai lmstudio`

Since 0.20.4, commit-message and branch-name generation works for non-OpenAI providers (Anthropic, Ollama, and others).

## Parallel Agents

Multiple coding agents can work in **one GitButler workspace** without separate worktrees:

- Each agent commits changes to its own GitButler branch
- They share one working directory, dependency install, and runtime state
- If a dependency arises, stack one branch on another
- If work is independent, move it to a separate branch

### When NOT to Use Parallel Agents

- Competing attempts at the same task
- Incompatible checkout states
- Isolated runtimes that conflict
- Use git worktrees for those cases instead

## What Agents Can Do With GitButler

- **Parallel branches** in one workspace without worktrees
- **Selected files and hunks** assigned to specific branches
- **Stacked branches** for dependent work
- **History edits** as direct commands (no interactive rebase needed)
- **Review and recovery** via `but oplog` and `but undo`
- Multi-committing, amending, and splitting complete quicker and use fewer tokens (since 0.20.4)

## Related

- [config.md](config.md) — AI provider configuration
- [commits.md](commits.md) — `--ai` flag on commit and squash
- [branches.md](branches.md) — `--ai` flag on branch show
