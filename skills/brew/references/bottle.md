# Bottles

Bottles are pre-built binary packages that Homebrew "pours" into the Cellar instead of building from source. This document covers the bottle DSL, the `brew bottle` command, pouring, relocation, and related commands.

## How Bottles Work

When running `brew install`, Homebrew checks if a bottle exists for the user's specific macOS version and CPU architecture. If so, it "pours" the bottle (extracts the archive) into the Cellar. Otherwise, it builds from source.

```
brew install <formula>
  → Check for matching bottle tag
    → Bottle exists? → Pour (extract) into Cellar → Link
    → No bottle?    → Build from source → Link
```

### Bottle Naming Convention

```
formula--version.os.bottle.tar.gz
```

Examples:

```
foo--0.1.ventura.bottle.tar.gz
qt--6.5.1.ventura.bottle.tar.gz
foo--0.1.arm64_ventura.bottle.tar.gz
```

Bottle tags encode OS and architecture: `ventura`, `arm64_ventura`, `sonoma`, `x86_64_linux`, etc.

### Bottle Hosting

Bottles for `homebrew/core` are hosted on GitHub Container Registry at `https://ghcr.io/v2/homebrew/core`. The bottle domain can be changed with `HOMEBREW_BOTTLE_DOMAIN`.

## Bottle DSL in Formulae

The `bottle` block in a formula defines pre-built binary packages. It is managed by `BottleSpecification`.

```ruby
bottle do
  root_url "https://ghcr.io/v2/homebrew/core"
  sha256 cellar: :any_skip_relocation, arm64_sonoma:   "abc123..."
  sha256 cellar: :any_skip_relocation, arm64_ventura:   "def456..."
  sha256 cellar: :any_skip_relocation, sonoma:         "ghi789..."
  sha256 cellar: :any_skip_relocation, ventura:        "jkl012..."
  sha256 cellar: :any_skip_relocation, x86_64_linux:   "mno345..."
end
```

### `cellar` Values

| Value | Meaning |
|-------|---------|
| `:any` | Bottle can be used regardless of Cellar location |
| `:any_skip_relocation` | Bottle can be used anywhere; no relocation needed |
| _(path)_ | Bottle must be poured into a specific Cellar path |

## Pouring and Relocation

When a bottle is poured into a non-default prefix, Homebrew performs **binary relocation** to fix hardcoded paths:

- Replaces the build-time `HOMEBREW_PREFIX` with the user's actual prefix
- Updates `install_name`/`rpath` entries in Mach-O binaries (macOS)
- Updates `RUNPATH` entries in ELF binaries (Linux)

The `--skip-relocation` flag to `brew bottle` marks bottles as not needing relocation (`:any_skip_relocation`).

## `brew bottle` Command

Generates a bottle from a formula built with `--build-bottle`.

```bash
brew install --build-bottle foo
brew bottle foo
```

| Flag | Purpose |
|------|---------|
| `--skip-relocation` | Mark bottle as not needing relocation |
| `--force-core-tap` | Bottle even if not from core tap |
| `--no-rebuild` | Don't add rebuild version |
| `--keep-old` | Keep old bottle block entries |
| `--json` | Write JSON metadata |
| `--merge` | Merge JSON from multiple platforms |
| `--write` | Write bottle block to formula |
| `--no-commit` | Don't commit changes |
| `--only-json-tab` | Only write JSON tab (no INSTALL_RECEIPT) |
| `--no-all-checks` | Don't create an `all` bottle or stop a no-change upload (requires `--merge`) |
| `--root-url <url>` | Override root URL |
| `--root-url-using <strategy>` | Download strategy for root URL |

### Workflow

```
brew install --build-bottle <formula>    # Build with bottle preparation
brew bottle <formula>                     # Generate bottle + JSON
# ... on multiple platforms ...
brew bottle --merge --write *.json        # Merge and write bottle block
```

## `brew unbottled`

Shows formulae that lack bottles for a given tag, and their dependents.

| Flag | Purpose |
|------|---------|
| `--tag <tag>` | Specific bottle tag |
| `--dependents` | Show dependents of unbottled formulae |
| `--total` | Show total counts |
| `--lost` | Print `homebrew/core` commits where bottles were lost in the last week |

## `brew verify`

Verifies bottle build provenance using GitHub attestations (requires `gh` CLI).

| Flag | Purpose |
|------|---------|
| `--os <os>` / `--arch <arch>` | Target platform |
| `--bottle-tag <tag>` | Specific bottle tag |
| `--deps` | Verify all dependencies too |
| `-f, --force` | Force re-verification |
| `-j, --json` | JSON output |

Related environment variables:

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_VERIFY_ATTESTATIONS` | Enable verification |
| `HOMEBREW_NO_VERIFY_ATTESTATIONS` | Disable verification |

## Related Environment Variables

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_BOTTLE_DOMAIN` | `https://ghcr.io/v2/homebrew/core` | Bottle download mirror URL |
| `HOMEBREW_ARTIFACT_DOMAIN` | — | Prefix all download URLs (mirror) |
| `HOMEBREW_ARTIFACT_DOMAIN_NO_FALLBACK` | — | Error instead of fallback when artifact domain fails |
| `HOMEBREW_GITHUB_PACKAGES_TOKEN` | — | GitHub Packages auth token |
| `HOMEBREW_GITHUB_PACKAGES_USER` | — | GitHub Packages username |
| `HOMEBREW_DOCKER_REGISTRY_TOKEN` | — | Docker registry bearer token |
| `HOMEBREW_DOCKER_REGISTRY_BASIC_AUTH_TOKEN` | — | Docker registry basic auth |

## References

- [Bottle documentation](https://docs.brew.sh/Bottles)
- For formula DSL: [formula.md](formula.md)
- For the Cellar: [cellar.md](cellar.md)
- For environment variables: [environment.md](environment.md)
- For developer commands: [developer.md](developer.md)
