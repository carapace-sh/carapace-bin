# Release Summary: v1.7.0 → master

**426 commits** | May 29 – Aug 14, 2026

---

## New Completers (~80 new)

### Cross-platform
`7z`, `adb`, `az`, `crush`, `drift`, `gh-stack`, `grit`, `herdr`, `hl`, `hunspell`, `img2pdf`, `ip`, `iqtree3`, `k3d`, `lore`, `maturin`, `pytest`, `rails`, `scc`, `tailscale`, `vercel`, `wg`, `zig`, `zpaq`

### Shell builtins
- **bash**: `alias`, `bg`, `bind`, `builtin`, `caller`, `cd`, `command`, `compgen`, `complete`, `compopt`, `declare`, `dirs`, `disown`, `echo`, `enable`, `eval`, `exec`, `export`, `fc`, `fg`, `getopts`, `hash`, `help`, `history`, `jobs`, `let`, `local`, `mapfile`, `popd`, `printf`, `pushd`, `read`, `readarray`, `readonly`, `set`, `shopt`, `source`, `test`, `trap`, `type`, `typeset`, `ulimit`, `umask`, `unalias`, `unset`, `wait`
- **zsh**: builtin completers
- **elvish**: builtin completers
- **fish**: builtin completers
- **cmd.exe** (Windows): `assoc`, `call`, `cd`, `chcp`, `color`, `copy`, `date`, `del`, `dir`, `echo`, `endlocal`, `exit`, `for`, `ftype`, `goto`, `if`, `mkdir`, `mklink`, `move`, `path`, `popd`, `prompt`, `pushd`, `rem`, `ren`, `rmdir`, `set`, `setlocal`, `shift`, `start`, `time`, `title`, `type`, `ver`, `vol`

### Darwin (macOS)
`airport`, `asr`, `base64`, `brctl`, `codesign`, `cupsfilter`, `cupsreject`, `defaults`, `dmesg`, `dscl`, `dstro`, `duti`, `hdiutil`, `ioreg`, `kextfind`, `kextload`, `kextstat`, `kextunload`, `kmutil`, `launchctl`, `log`, `mdfind`, `mdutil`, `mount`, `nettop`, `nfs4`, `nvram`, `osascript`, `otool`, `pbcopy`, `pbpaste`, `pmset`, `port`, `pstop`, `sba`, `screencapture`, `scselect`, `scutil`, `security`, `shortcuts`, `simctl`, `softwareupdate`, `spctl`, `sw_vers`, `sysctl`, `system_profiler`, `tmutil`, `umount`, `update_dyld_shared_cache`

### Linux
`dnf5`, `qemu-*` (full suite), `zfs`/`zpool`

---

## Major Updates

### MCP Server (`cmd/carapace/cmd/mcp/`)
- Refactored from a monolithic file into focused modules (`complete_command.go`, `complete_macro.go`, `list_macros.go`, `codegen.go`, `protocol.go`, `tools.go`, `util.go`).
- Added `codegen` tool for generating Go code from YAML specs.
- Added `list_macros` with optional `executable` parameter and version info.
- `complete_command` now supports `bridge` and `executable` parameters for cross-tool completion, and invokes bridge actions directly when `executable` is set.
- Added `generate-all` subcommand to scan completers directory once.

### Completer Updates
- **brew**: fully rewritten with many new subcommands, alias support, corrected flag types
- **pnpm**: filter-aware script completion via `carapace-pnpm` lexer, persistent flags, missing flag completions
- **npm**: updated to match current CLI, `package.json` key completion for `pkg get/set/delete`, positional completion for `help`/`init`/`exec`
- **npm/bun**: compatible `run` completers with graceful error suppression
- **pixi**: updated for v0.73.0–v0.76.0 (multiple rounds)
- **jj**: updated for v0.44.0, added missing tool completion
- **cargo**: updated for 0.98.0, switched from `cargo read-manifest` to `cargo metadata`, unique dependencies/features
- **rustup**: updated for 1.29.0
- **but**: updated for v0.20.3–v0.22.0
- **herdr**: updated for v0.7.5–v0.8.0
- **gh-stack**: updated for v0.1.0
- **ghostty**: updated for 1.3.x, man docs added
- **glab**: updated to v1.112.0
- **wt**: updated to v0.71.0
- **nmcli**: fully rewritten to match current manpage/source
- **yay**: root-level & pacman operation flags, missing flags for v13
- **paru**: updated to v2.1.0
- **git**: updated for Git 2.50–2.55, config-based hooks, dash completion on `--`, fixed flag descriptions, fix for translations in `ActionHeads`
- **git-extras**: new completers for all subcommands
- **git-bulk**: workspace action moved to `cmd/action` subpackage
- **gh**: `jq` filter completion added
- **jq**: uses `carapace-jq` for filter completion
- **grit**: filter completion
- **vercel**: full flag and positional completions added
- **bun**: missing flag completions, `run` suppression of `package.json` not found
- **node, nu, rust-analyzer, rustdoc, xdotool, ufw, xz, upx, tsc, gopls, xxhsum, idea, jq, pip, python, tldr/tealdeer, cura, gh-copilot, cargo-metadata, chdman**: missing flag completions added
- ~36 other completers had flag corrections (wrong types, missing flags, incorrect shorthand mappings)

### Shell Integration
- `CARAPACE_SHELL` environment variable added to all shell snippets.
- Elvish group description added.
- Legacy bash detection added.
- Nushell: fallback to default completions on empty.
- Fix: `DetermineShell` skip unnecessary call on init.
- Fix: `sed` forks replaced with native shell in completion snippets.
- Oil, tcsh, xonsh, powershell snippets updated.

### Misc
- Shell aliases now included in `ActionExecutables` completion.
- `ps.ActionKillSignals` defaults to `runtime.GOOS` on empty.
- Macro modifier chaining support in `carapace --macro`.
- `style.ForKeyword` applied to keyword-valued flag completions.
- `--all` flag added to `list` command for builtin/excluded completers.
- `carapace-magick` added to goreleaser.
- man documentation added for ghostty, git config keys.
- Static analysis fixes, `go fmt` pass.

---

## Documentation & Skills

- **New compound skills**: Bun, But, Cargo, Cobra, dnf5, Git, Go, HTTPie, iproute2, kubeadm, kubectl, Lore, Pandoc, QEMU, Rails, Ruby, rustup, winget
- **Carapace skill**: consolidated from `carapace-spec`, `carapace-scrape`, `carapace-integrate` into a single compound skill with 14 reference documents (action, spec, macro, env, scrape, integrate, setup, choice, mcp, man, convert, lexer, update, lexer).
- **Bubble Tea skill**: 12 reference documents.
- **Bash, fish, zsh, nushell, elvish, oil, xonsh, tcsh, powershell, hilbish, ion, cmd-clink**: compound skills written for each shell's completion system.
- `AGENTS.md` added to the repo root as a project-level onboarding document.
- Cross-references fixed in skill docs, new `carapace-pflag` APIs documented.

---

## carapace-sh Dependency Changes

| Dependency | v1.7.0 → master | Notes |
|---|---|---|
| `carapace` | `v1.11.6` → **`v1.15.1`** | Core library, major bump |
| `carapace-spec` | `v1.5.3` → **`v1.8.0`** | Spec engine, major bump |
| `carapace-bridge` | `v1.5.3` → **`v1.6.3`** | Bridge framework |
| `carapace-pflag` | `v1.1.0` → **`v1.3.0`** | pflag replacement (still under `replace` directive) |
| `carapace-jjlex` | `v0.0.7` → **`v0.1.9`** | Jujutsu lexer |
| `carapace-jq` | **new** → `v0.0.3` | jq filter completion |
| `carapace-pnpm` | **new** → `v0.0.2` | pnpm filter lexer |
| `carapace-selfupdate` | `v0.0.10` (unchanged) | |
| `carapace-shlex` | `v1.1.1` (unchanged) | |

---

## Contributors

Special thanks to external contributors: `@rfaulhaber`, `@kusutori`, `@Malix-Labs`, `@ahfoysal`, `@kaathewisegit`, `@ruben-arts`.