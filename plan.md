# Adding missing completers (batch 1)

Source of the shortlist: a diff of the ~1000 completion files shipped by fish
(`share/completions/`) against `completers/*/*_completer` here, minus the 76
commands already served by known bridges (`cmd/carapace/cmd/completers/bridges.go`).
627 fish-supported commands had no native completer; the entries below are the
subset worth implementing in this batch.

fish files are only used as a popularity/surface-size signal. Every flag,
subcommand and argument is verified against the source code or documentation of
the current upstream release, not copied from fish.

## Status

| Command | Framework | Source used | Status |
|---|---|---|---|
| `kind` | cobra | kubernetes-sigs/kind v0.33.0 | done (scraped) |
| `skaffold` | cobra | GoogleContainerTools/skaffold v2.25.0 | done (scraped) |
| `ruff` | clap | astral-sh/ruff 0.17.3 | done (scraped) |
| `pre-commit` | argparse | pre-commit/pre-commit v4.7.0 | done (scraped) |
| `iwctl` | hand-written | iwd 3.12 (`client/*.c`) | done |
| `psql` | hand-written | PostgreSQL 18 (`psql --help`, `\?` help) | done |

Scrapers live in `../scrape` on branch `feat/completers-batch-1`; generated
specs were captured under `gh-pages/{command}/{version}.yaml` and are excluded
from git there by design.

## What each completer covers

### kind (scraped, cobra)

Full command tree from the release. Completion added for `--config`,
`--kubeconfig`, `--image`/`--base-image` (docker images), `--type`
(`url|file|release|ci|source`), `--arch` (`amd64|arm64`, the only architectures
`pkg/build/nodeimage` accepts), cluster names via `kind get clusters`, node
names via `kind get nodes --name <cluster>`, log output directories, and image
archives (`*.tar`, `*.tar.gz`).

### skaffold (scraped, cobra)

Full command tree from the release. A shared helper adds completion for the
flags most commands carry: config/kubeconfig/log files, kube contexts, `--profile`
(parsed from the skaffold config), `--mute-logs`, `--sync-remote-cache` and
`--verbosity`. The helper only registers completions for flags the command
actually declares, because inherited flags must be registered on the parent.

### ruff (scraped, clap)

Full command tree from the release. Rule selectors (`--select`, `--ignore`,
`--fixable`, `--unfixable`, and the `extend-*` variants) complete linter prefixes
and rule codes with rule names as descriptions; `ruff rule` completes rule codes
and `ruff config` completes configuration keys, both generated from the scraped
spec (970 rule codes, 1029 selectors, 215 config options). `--config` completes
TOML files, `--output-format`, `--target-version`, `--color` and
`generate-shell-completion` use the documented values.

### pre-commit (scraped, argparse)

Full command tree from the release, including `hazmat` subcommands. `run` and
`try-repo` complete hook ids from the configured `.pre-commit-config.yaml`
(honoring `--config`), `--hook-stage`/`--hook-type` complete the stages and hook
types from `clientlib`, and file/directory arguments are completed where upstream
documents them.

### iwctl (hand-written, iwd 3.12)

The command model follows `client/command.c`: a command family is followed by an
optional entity (`<wlan>`/`<phy>`) and the command itself. Families, commands and
their descriptions come from the `struct command`/`struct command_family` tables
in `client/*.c`. Writable `set-property` properties and their values come from the
`proxy_interface_property` tables (`is_read_write` plus the option lists), so
`adapter` only offers `Powered`, `device` offers `Mode`/`Powered`, and
`known-networks` offers `AutoConnect`. Device entities use the shared wifi device
action, adapter entities are read from `/sys/class/ieee80211`.

### psql (hand-written, PostgreSQL 18)

Flags, backslash commands, `\pset` options/values and special variables come from
`src/bin/psql/help.c` and `src/bin/psql/command.c`. Backslash commands are
completed for `-c/--command` (the only argv position where they occur), `\pset`/
`--pset` values via multiparts, and `--set` values for the documented enumerations
(`ECHO`, `COMP_KEYWORD_CASE`, `HISTCONTROL`, `ON_ERROR_ROLLBACK`, `VERBOSITY`).

## Deliberate limitations

- `iwctl`: known-network names and dynamic station/BSS data need D-Bus, so they
  are not completed; `known-networks <name>` accepts free text.
- `psql`: database names, table patterns and `\d*` patterns require a live server,
  so positional arguments are not completed.
- `skaffold`: `--namespace` has no local source of truth and is left alone.

## Scraper notes

- `pre-commit` builds its argparse parser inside `main()`, so `scrapers/pre-commit/scrape.py`
  intercepts parser construction and aborts before parsing.
- The argparse converter emits a phantom `- name: ""` command for the root parser.
  `scrapers/pre-commit/clean.py` removes that block so root flags survive without
  producing an empty subcommand during codegen.
- The release optimizer (`cmd/carapace-generate/pkg/completer/optimize.go`) assumes
  every `.go` file directly in `cmd/` has an `init()` function, so helper files in
  `cmd/` are avoided; helpers live in `cmd/action/` or are added to `root.go`.

## Follow-ups

Deferred to a later batch (recorded so the research is not lost): `virsh`, `svn`,
`gem`/`bundle`, `mysql`/`mysqldump`, `btrfs`, `zypper`/`rpm`, `nmap`/`ncat`,
`meson`/`ninja`, `firewall-cmd`, `mdadm`, `nvme`,
`hostnamectl`/`timedatectl`/`loginctl`/`machinectl`/`busctl`, `gcc`/`clang`,
`strace`, `valgrind`, plus native versions of the bridge-only `podman` and
`kustomize`.
