---
name: man-docs
description: >
  Use when writing or updating man page documentation (YAML cmd/ specs) for carapace-bin completers.
  Covers the documentation workflow, standards, helper tools (update, split-to, spec-diff, man-to-md),
  documentation.command and documentation.flag guidelines, and validation.
  Triggers on: "man docs", "cmd docs", "command documentation", "flag documentation",
  "YAML man pages", "spec documentation", "documentation.command", "documentation.flag",
  "man/cmd", "split-to", "spec-diff", "man-to-md", "update".
user-invocable: true
---

# Man Page Documentation for carapace-bin

Guide for writing and maintaining command documentation in `man/cmd/` YAML specs within carapace-bin.

## Architecture

Command documentation lives in `man/cmd/<completer>/` as split YAML specs — one file per subcommand:

```
man/cmd/
  git/
    git.yaml              # root command
    git.log.yaml           # subcommand
    git.remote.add.yaml   # nested subcommand
  docker/
    docker.yaml
    docker.container.run.yaml
```

The spec structure (flags, subcommands, descriptions) is generated from the completer. The `documentation:` section is the enrichment layer — prose that goes beyond the one-line description.

## Tools

### Update docs for a completer (recommended)

```bash
# One-step update: generate spec, split with doc preservation, show diff
carapace <completer> spec | carapace-man update - man/cmd/<completer>

# Or with a pre-generated spec file
carapace-man update /tmp/<completer>-spec.yaml man/cmd/<completer>

# Dry run to preview changes without writing files
carapace-man update --dry-run /tmp/<completer>-spec.yaml man/cmd/<completer>
```

The `update` command combines spec-diff and split-to into one step. It:
1. Reports current diff (new/removed subcommands, changed flags, missing docs)
2. Splits the spec into per-subcommand files, **preserving existing `documentation:` entries**
3. Creates new files for new subcommands

### Split spec into per-subcommand files

```bash
# Generate fresh spec from the completer
carapace <completer> spec > /tmp/<completer>-spec.yaml

# Split into per-subcommand files in the target directory
carapace-man split-to /tmp/<completer>-spec.yaml man/cmd/<completer>
```

`split-to` overwrites structural fields (name, description, flags, commands) and **preserves existing `documentation:` entries** from the target files. New files get no documentation.

### Check for drift

```bash
# Compare fresh spec against committed man/cmd/<completer>/ files
carapace <completer> spec > /tmp/<completer>-spec.yaml
carapace-man spec-diff /tmp/<completer>-spec.yaml man/cmd/<completer>
```

Reports: new/removed subcommands, changed flag descriptions, missing documentation, `[AI]` prefixed entries.

### Research source material

Priority order for sourcing documentation:

1. **Official upstream docs** — the project's own documentation site (e.g. docs.docker.com, docs.brew.sh, cli.github.com/manual). Closest to authoritative; rephrase for consistency but stay faithful to the original meaning and technical accuracy.
2. **System man pages** — via `man-to-md`. These are often derived from the upstream source and vetted by distributions.
3. **`<command> --help`** — last resort. Help output is often incomplete or auto-generated.

When rephrasing official docs:
- Preserve technical accuracy — enum values, default behaviors, and interaction effects must match the source
- Adjust tone and length to fit the skill's conventions (imperative mood, concise structure)
- Merge multiple paragraphs into a coherent summary where the full detail isn't needed
- Keep close to the original phrasing where it's already clear and concise — don't rewrite for the sake of rewriting

```bash
# Convert a system man page to markdown
carapace-man man-to-md docker-container-run
carapace-man man-to-md git-log --section 1

# Falls back to <command> --help if no man page exists
carapace-man man-to-md mycli
```

## Documentation Standards

### `documentation.command`

- **≤300 words** (complex commands with many flags or subtle behaviors may need more; the pager handles long content)
- First sentence: what the command does (imperative mood: "Create and start a new container")
- Remaining: key behavior, common usage patterns, important caveats
- Source from official upstream docs first, then man pages, then `--help`. Stay close to official wording for correctness; rephrase for consistency only
- Use UID cross-references where appropriate: `[git log](cmd://git/log)`, `[git-config](man://git-config/1)`

### `documentation.flag`

- **≤200 words** (concise when possible, but complex flags need full explanation)
- Only add when the spec's short description is insufficient
- Typical cases needing flag docs:
  - Complex value types (enum values, patterns, formats)
  - Non-obvious side effects or interactions with other flags
  - Flags where short description is missing or terse
  - Flags with subtle default behavior that differs from user expectation
- Source from official upstream docs first, then man pages, then `--help`. Stay close to official wording for correctness; rephrase for consistency only
- Most flags should have **no** flag doc entry
- Use the flag name without leading dashes as the key (e.g. `restart` not `--restart`)

### Examples

Good `documentation.command`:
```yaml
documentation:
  command: >-
    Create and start a new container from an image. Configures the
    container's filesystem, network, and resource constraints, then
    executes the specified command. Use `docker create` to create a
    container without starting it.
```

Good `documentation.flag` (only because short description is terse):
```yaml
flags:
  --sig-proxy: Proxy received signals to the process
documentation:
  flag:
    sig-proxy: >
      When set (default), signals received by `docker run` are forwarded
      to the container process. Disable with `--sig-proxy=false` to
      handle signals only on the client side.
```

Most flags need **no** `documentation.flag` entry:
```yaml
flags:
  --detach: Run container in background and print container ID
  --interactive: Keep STDIN open even if not attached
  --tty: Allocate a pseudo-TTY
# Short descriptions are clear — no flag docs needed
```

## Workflow

For each completer:

1. **Update**: Run `carapace-man update` to refresh spec structure while preserving docs
2. **Audit**: Check the diff output from `update` (or run `spec-diff` separately) to see what needs docs
3. **Research**: For each subcommand needing `documentation.command`:
   - Check the project's official documentation site first (most authoritative)
   - Run `man-to-md <command>` for system man pages
   - If neither available, run `<command> --help`
   - Stay close to official wording; rephrase only for consistency and length
4. **Write**: Fill in `documentation.command` entries — concise, accurate, grounded in real docs
5. **Flag docs**: Only add `documentation.flag` entries where the short description is insufficient
6. **Validate**: Re-read edited files, check word counts, verify no `[AI]` prefixes remain
7. **Preserve**: Never modify git, jj, or but docs — these are high-quality and should not be changed

## High-Quality Docs (Do Not Modify)

These completers have high-quality documentation from official sources and must not be changed:
- **git**: sourced from official Git man pages
- **jj**: auto-generated from Jujutsu CLI help output
- **but**: auto-generated from GitButler CLI help output

## YAML Format

Each file follows the carapace command spec schema:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: subcommand-name
description: one-line description from completer
flags:
    --flag-a: short description
    --flag-b=: short description with value
documentation:
  command: |
    Extended description of the command.
    Can span multiple lines with markdown.
  flag:
    flag-b: Extended description for flag-b only when needed.
```

The `documentation` section is optional. Only include it when there's something meaningful to add beyond the short descriptions already in `flags`.

## Cross-References

Use UID-style links in documentation text:
- `[git log](cmd://git/log)` — link to another subcommand's docs
- `[git-config](man://git-config/1)` — link to a system man page
- `[HTTP GET](http://method/GET)` — link to domain docs

## Validation Checklist

- [ ] `documentation.command` entries ≤ 300 words
- [ ] `documentation.flag` entries ≤ 200 words and only where needed
- [ ] YAML parses without errors
- [ ] No changes to git/, jj/, but/ docs
