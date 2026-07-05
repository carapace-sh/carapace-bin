# The Cellar and Keg Linking

The Cellar is Homebrew's installation directory. This document covers the Cellar structure, kegs, racks, the opt prefix, linking/unlinking, and keg-only formulae.

## Directory Structure

```
HOMEBREW_PREFIX/                    (e.g. /opt/homebrew)
├── Cellar/                         ← The Cellar: all installed formula versions
│   ├── foo/
│   │   ├── 0.1/                    ← A keg (one version)
│   │   │   ├── bin/
│   │   │   ├── lib/
│   │   │   ├── include/
│   │   │   └── INSTALL_RECEIPT.json   ← The tab
│   │   └── 0.2/                    ← Another keg (newer version)
│   │       ├── bin/
│   │       └── ...
│   └── bar/
│       └── 1.0/
│           └── ...
├── Caskroom/                       ← The Caskroom: all installed casks
│   └── firefox/
│       └── 120.0/
├── opt/                            ← Symlinks to active kegs
│   ├── foo -> ../Cellar/foo/0.2
│   └── bar -> ../Cellar/bar/1.0
├── bin/                            ← Symlinks from kegs into prefix
│   ├── foo -> ../opt/foo/bin/foo
│   └── bar -> ../opt/bar/bin/bar
├── lib/
├── include/
├── share/
├── etc/                            ← Non-versioned config (not in Cellar)
└── var/                            ← Non-versioned runtime data (not in Cellar)
```

## Terminology

| Term | Definition | Example |
|------|-----------|---------|
| **Prefix** | Homebrew's installation root | `/opt/homebrew` (Apple Silicon), `/usr/local` (Intel Mac), `/home/linuxbrew/.linuxbrew` (Linux) |
| **Cellar** | Directory containing all named racks | `Cellar/` |
| **Rack** | Directory containing one or more versioned kegs for a formula | `Cellar/foo` |
| **Keg** | A specific versioned installation directory | `Cellar/foo/0.1` |
| **Opt prefix** | Symlink to the active version of a keg | `opt/foo` → `Cellar/foo/0.2` |
| **Tab** | Metadata JSON file in a keg | `Cellar/foo/0.1/INSTALL_RECEIPT.json` |
| **Caskroom** | Directory containing all installed casks | `Caskroom/` |

## The Prefix

The prefix is determined by the installation:

| Platform | Default Prefix |
|----------|---------------|
| Apple Silicon (ARM Macs) | `/opt/homebrew` |
| Intel Macs | `/usr/local` |
| Linux | `/home/linuxbrew/.linuxbrew` |

Query with `brew --prefix`. The prefix is set in the environment via `brew shellenv`:

```sh
eval "$(/opt/homebrew/bin/brew shellenv)"
```

This exports `HOMEBREW_PREFIX`, `HOMEBREW_CELLAR`, and `HOMEBREW_REPOSITORY`.

## Keg Linking

When a formula is installed (or `brew link` is run), Homebrew symlinks files from the keg into the prefix:

```
Cellar/foo/0.2/bin/foo  →  opt/foo/bin/foo  →  bin/foo
```

The `opt/` directory provides a stable path that always points to the active version, regardless of which version is currently linked. This is important for:

- Configuring build tools to find libraries: `pkg-config --cflags foo` uses `opt/foo/lib/pkgconfig`
- Setting `PATH` entries that survive upgrades: `export PATH="$HOMEBREW_PREFIX/opt/foo/bin:$PATH"`
- Referencing formulae in other formulae: `depends_on "foo"` uses `opt/foo`

### `brew link` / `brew unlink`

| Command | Purpose |
|---------|---------|
| `brew link <formula>` | Symlink keg into prefix (done automatically on install) |
| `brew unlink <formula>` | Remove symlinks (temporarily disable) |
| `brew link --overwrite <formula>` | Overwrite conflicting files |
| `brew link --force <formula>` | Force-link a keg-only formula |
| `brew link --dry-run <formula>` | Show what would be linked |

See [maintenance.md](maintenance.md) for full flag reference.

## Keg-Only Formulae

A **keg-only** formula is not symlinked into the prefix. Its files remain only in the Cellar, accessible via the `opt/` path.

```ruby
class Openjdk < Formula
  keg_only "it shadows macOS' built-in java"
  # ...
end
```

### Why Keg-Only?

Formulae are marked keg-only when:

1. **Conflicts with macOS built-ins** — e.g. `openssl`, `sqlite`, `zlib`, `python` — to avoid shadowing system versions
2. **Conflicts with other formulae** — multiple versions of the same library (e.g. `postgresql@14`, `postgresql@16`)
3. **Post-install scripts** — formulae that need special handling before being on `PATH`

### Using Keg-Only Formulae

Keg-only formulae are accessible via their opt prefix:

```bash
# Add to PATH manually
export PATH="$HOMEBREW_PREFIX/opt/openjdk/bin:$PATH"

# Or invoke directly
"$(brew --prefix)/opt/openjdk/bin/java" -version

# In other formulae, keg-only deps are automatically added to -I and -L flags by superenv
```

When another formula `depends_on` a keg-only formula, Homebrew's `superenv` automatically adds the keg-only formula's `include/` and `lib/` to compiler flags (`-I` and `-L`). This is a key reason keg-only formulae work correctly as dependencies without being on `PATH`.

## Multiple Versions

A rack can contain multiple kegs (versions):

```
Cellar/postgresql@14
  14.10/
  14.11/      ← after upgrade, old version kept until cleanup
```

- `brew upgrade` installs the new version but keeps the old one
- `brew cleanup` removes old versions
- `brew pin` prevents a specific version from being upgraded
- `brew switch` (removed in 4.0) was used to switch between versions; now use versioned formulae (`postgresql@14`, `postgresql@16`)

## The Tab (INSTALL_RECEIPT.json)

Each keg contains a tab recording installation metadata:

```json
{
  "homebrew_version": "4.2.0",
  "used_options": [],
  "unused_options": ["--with-test"],
  "built_as_bottle": false,
  "poured_from_bottle": true,
  "loaded_from_api": true,
  "installed_as_dependency": false,
  "installed_on_request": true,
  "changed_files": [],
  "time": 1700000000.0,
  "source_modified_time": 1699900000.0,
  "compiler": "clang",
  "stdlib": "libc++",
  "aliases": [],
  "runtime_dependencies": [...],
  "source": {
    "path": "/opt/homebrew/Library/Taps/homebrew/homebrew-core/Formula/f/foo.rb",
    "tap": "homebrew/core",
    "spec": "stable",
    "versions": { "stable": "0.2", "head": null }
  }
}
```

The tab is used by:
- `brew autoremove` — checks `installed_as_dependency` to determine if a formula is still needed
- `brew leaves` — filters by `installed_on_request`
- `brew reinstall` — reproduces original `used_options`
- `brew list --poured-from-bottle` / `--built-from-source`

## Path Queries

| Command | Output |
|---------|--------|
| `brew --prefix` | `/opt/homebrew` |
| `brew --prefix foo` | `/opt/homebrew/opt/foo` (or Cellar path) |
| `brew --cellar` | `/opt/homebrew/Cellar` |
| `brew --cellar foo` | `/opt/homebrew/Cellar/foo` |
| `brew --caskroom` | `/opt/homebrew/Caskroom` |
| `brew --repository` | `/opt/homebrew` (Git repo location) |
| `brew --cache` | `~/Library/Caches/Homebrew` (macOS) |
| `brew --cache foo` | Cache path for foo's download |

## References

- [Formula Cookbook](https://docs.brew.sh/Formula-Cookbook) — directory variables, keg-only
- For formula DSL: [formula.md](formula.md)
- For bottles: [bottle.md](bottle.md)
- For maintenance (link/unlink/cleanup): [maintenance.md](maintenance.md)
- For environment variables: [environment.md](environment.md)
