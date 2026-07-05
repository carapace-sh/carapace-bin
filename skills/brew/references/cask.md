# Cask DSL

A cask is a Homebrew package definition that installs pre-compiled binaries — typically macOS GUI applications, fonts, and plugins. Casks use a distinct DSL from formulae and install into the Caskroom rather than the Cellar.

## Cask vs Formula

| Aspect | Formula | Cask |
|--------|---------|------|
| **Purpose** | CLI tools, libraries, dev utilities | GUI applications, desktop software, fonts, plugins |
| **Installation** | Builds from source (downloads, compiles, links) | Downloads and moves pre-built binaries |
| **Install location** | `Cellar/<formula>/<version>` | `Caskroom/<cask>/<version>` |
| **DSL** | Formula DSL (source-based: `def install`, `depends_on`) | Cask DSL (declarative: `version`, `sha256`, `app`, `pkg`) |
| **Default target** | `bin/` in prefix | `/Applications` (for `.app` files) |
| **Uninstall** | Remove Cellar directory | May require explicit `uninstall` stanza for `pkg`/`installer` |

> **Source of truth**: <https://docs.brew.sh/Cask-Cookbook>. For building from source, see [formula.md](formula.md).

## Cask Structure

A cask is a Ruby block beginning with a header line and enclosed in `do … end`. The DSL is **declarative** — order doesn't matter (except inside procedural blocks like `postflight`).

```ruby
cask "anybar" do
  version "0.2.3"
  sha256 "c87dbc6aff5411676a471e84905d69c671b62b93b1210bd95c9d776d087de95c"
  url "https://github.com/tonsky/AnyBar/releases/download/#{version}/AnyBar-#{version}.zip"
  name "AnyBar"
  desc "Menu bar status indicator"
  homepage "https://github.com/tonsky/AnyBar"
  app "AnyBar.app"
end
```

The cask name (token) on the header line must match the cask filename (without `.rb`).

## Stanza Order (Canonical)

```
arch
os
version
sha256
language
url
name
desc
homepage
livecheck
no_autobump!
deprecate!
disable!
auto_updates
conflicts_with
depends_on
container
rename
suite
app
pkg
installer
binary
manpage
bash_completion
fish_completion
zsh_completion
colorpicker
dictionary
font
input_method
internet_plugin
keyboard_layout
prefpane
qlplugin
mdimporter
screen_saver
service
audio_unit_plugin
vst_plugin
vst3_plugin
artifact, target:
stage_only
preflight
preflight_steps
postflight
postflight_steps
uninstall_preflight
uninstall_preflight_steps
uninstall_postflight
uninstall_postflight_steps
uninstall
zap
caveats
```

## Required Stanzas

| Stanza | Multiple? | Value |
|--------|-----------|-------|
| `version` | No | Application version string, or `:latest` |
| `sha256` | No | SHA-256 checksum, or `:no_check` |
| `url` | No | URL to the `.dmg`/`.zip`/`.tgz`/etc. |
| `name` | Yes | Full proper name defined by the vendor |
| `desc` | No | One-line description (≤80 chars) |
| `homepage` | No | Application homepage |
| `livecheck` | No | Ruby block describing how to find updates |
| `depends_on` | Yes | List of dependencies and requirements |
| `zap` | Yes | Additional procedures for more complete uninstall |

## Artifact Stanzas (At Least One Required)

At least one artifact stanza must be present (unless `stage_only` is used):

| Stanza | Purpose | Default Target |
|--------|---------|---------------|
| `app` | `.app` bundle | `/Applications` |
| `suite` | Directory of apps | `/Applications` |
| `pkg` | `.pkg` installer | macOS installer |
| `installer` | Executable installer | — |
| `binary` | Binary symlinked into prefix | `$(brew --prefix)/bin` |
| `manpage` | Man page | man page directories |
| `bash_completion` | Bash completion | `etc/bash_completion.d` |
| `fish_completion` | fish completion | `vendor_completions.d` |
| `zsh_completion` | Zsh completion | `site-functions` |
| `colorpicker` | ColorPicker plugin | `~/Library/ColorPickers` |
| `dictionary` | Dictionary | `~/Library/Dictionaries` |
| `font` | Font | `~/Library/Fonts` |
| `input_method` | Input Method | `~/Library/Input Methods` |
| `internet_plugin` | Internet Plugin | `~/Library/Internet Plug-Ins` |
| `keyboard_layout` | Keyboard Layout | `~/Library/Keyboard Layouts` |
| `prefpane` | Preference Pane | `~/Library/PreferencePanes` |
| `qlplugin` | Quick Look Plugin | `~/Library/QuickLook` |
| `mdimporter` | Spotlight Importer | `~/Library/Spotlight` |
| `screen_saver` | Screen Saver | `~/Library/Screen Savers` |
| `service` | Service | `~/Library/Services` |
| `audio_unit_plugin` | Audio Unit plugin | `~/Library/Audio/Plug-Ins/Components` |
| `vst_plugin` | VST plugin | `~/Library/Audio/Plug-Ins/VST` |
| `vst3_plugin` | VST3 plugin | `~/Library/Audio/Plug-Ins/VST3` |
| `artifact` | Arbitrary file (requires `target:`) | specified target |
| `stage_only` | No installable artifact (download/staging only) | — |

All artifact stanzas support a `target:` key for custom destination paths.

### `generate_completions_from_executable`

Runs an installed executable at install time to produce shell completion scripts:

```ruby
generate_completions_from_executable "#{bin}/foo", "--generate-completions",
  shells: [:bash, :zsh, :fish],
  base_name: "foo"
```

## Key Stanzas in Detail

### `version`

Accepts a string interpolated in other stanzas. Supports helper methods: `major`, `minor`, `patch`, `major_minor`, `major_minor_patch`, `minor_patch`, `csv.first`, `csv.second`, `dots_to_hyphens`, `no_dots`, etc.

Special value `:latest` — used when URL has no version info; requires `sha256 :no_check`.

### `sha256`

SHA-256 checksum of the downloaded file. Special value `:no_check` disables checking. Supports architecture-specific values:

```ruby
sha256 arm:   "abc123...",
       intel: "def456..."
```

### `url`

Download URL. HTTPS preferred. Additional parameters:

| Parameter | Purpose |
|-----------|---------|
| `verified:` | "Verified" URL label for non-obvious domains |
| `using:` | `:post`, `:homebrew_curl` (only legal values for cask URLs) |
| `cookies:` | HTTP cookies |
| `referer:` | HTTP referer |
| `header:` | HTTP headers |
| `user_agent:` | User agent string or `:fake` |
| `data:` | POST data |

### `uninstall`

Required for `pkg` and `installer` casks. Describes how to uninstall:

| Key | Purpose |
|-----|---------|
| `early_script:` | Runs early (special cases) |
| `launchctl:` | launchd job IDs to remove |
| `quit:` | Bundle IDs of running apps to quit |
| `signal:` | Unix signal + bundle ID arrays (when `quit:` doesn't work) |
| `login_item:` | Login item names to remove |
| `kext:` | Kernel extension bundle IDs to unload |
| `script:` | Uninstall script path (run via sudo) |
| `pkgutil:` | Bundle IDs matching installed packages (most useful) |
| `delete:` | Absolute paths to remove (last resort) |
| `trash:` | Paths to move to Trash |
| `rmdir:` | Directories to remove if empty |

### `zap`

More complete uninstall including user files and shared resources. Only performed with `brew uninstall --zap`. Uses same directives as `uninstall` (`trash:` preferred over `delete:`).

If no additional files exist, use: `# No zap stanza required`

Can be auto-generated with `brew generate-zap`.

### `conflicts_with`

Declares conflicts preventing installation. Primary key: `cask:` (another cask token).

### `depends_on`

| Key | Accepts |
|-----|---------|
| `cask:` | Cask tokens |
| `formula:` | Formula names |
| `macos:` | Symbol (`:big_sur`), comparison string (`">= :big_sur"`), or array |
| `maximum_macos:` | Newest compatible macOS |
| `arch:` | `:arm64`, `:x86_64`, `:intel` |
| `linux:` | Linux-only |
| `java:` | Java stub |

### `caveats`

Displays cask-specific information at install/info time.

**String form** — evaluated at compile time; supports interpolation: `token`, `version`, `homepage`, `caskroom_path`, `staged_path`.

**Block form** — deferred to install time; can reference `@cask`.

**Mini-DSL methods:** `path_environment_variable`, `zsh_path_helper`, `depends_on_java`, `requires_rosetta`, `logout`, `reboot`, `files_in_usr_local`, `kext`, `unsigned_accessibility`, `license`, `free_license`.

### `auto_updates`

Boolean. Asserts that the cask artifacts auto-update themselves. When `true`, `brew upgrade` skips this cask unless `--greedy` is passed.

### `container`

| Form | Purpose |
|------|---------|
| `container nested:` | Relative path to an inner container (e.g., `.dmg` inside `.tar`) |
| `container type:` | Override autodetect: `:air`, `:bz2`, `:cab`, `:dmg`, `:generic_unar`, `:gzip`, `:otf`, `:pkg`, `:rar`, `:seven_zip`, `:sit`, `:tar`, `:ttf`, `:xar`, `:zip`, `:naked` |

### `language`

Matches ISO 639-1 language codes, ISO 15924 script codes, and ISO 3166-1 Alpha 2 regional identifiers. US English should always be the default:

```ruby
language "en", default: true do
  "en-US"
end

language "zh" do
  "zh-CN"
end
```

Install with: `brew install firefox --language=it`

### `livecheck`

Automatically fetches latest version from changelogs, release notes, appcasts. Must contain a `url` (string or `:url`/`:homepage`).

### `no_autobump!`

Excludes a cask from autobump. Requires `because:` parameter.

### `deprecate!` / `disable!`

| Stanza | Behavior |
|--------|----------|
| `deprecate!` | Still installable but prints a warning |
| `disable!` | Cannot be installed/upgraded, prints an error |

Parameters: `date:` (YYYY-MM-DD), `because:` (string or preset symbol), `replacement:`, `replacement_formula:`, `replacement_cask:`.

A `disable!` with a future date is automatically deprecated until that date.

### `arch` / `os`

Define architecture/OS-specific strings for substitution:

```ruby
arch arm: "aarch64", intel: "x86-64"
os macos: "darwin", linux: "tux"
```

### `rename`

Renames files with unpredictable names (e.g., timestamps in paths):

```ruby
rename "foobar-*.pkg", "foobar.pkg"
```

### `*flight` Stanzas

| Stanza | Type | When |
|--------|------|------|
| `preflight` | Ruby block | Before installation |
| `preflight_steps` | Declarative | Before installation (JSON API compatible) |
| `postflight` | Ruby block | After installation |
| `postflight_steps` | Declarative | After installation (JSON API compatible) |
| `uninstall_preflight` | Ruby block | Before uninstall |
| `uninstall_preflight_steps` | Declarative | Before uninstall |
| `uninstall_postflight` | Ruby block | After uninstall |
| `uninstall_postflight_steps` | Declarative | After uninstall |

A cask may define either the Ruby flight block **or** the matching steps block, not both.

**Steps available:** `mkdir`, `mkdir_p`, `touch`, `move`/`mv`, `move_children`, `symlink`/`ln_s`, `ln_sf`, `write`.

**Mini-DSL:** `set_ownership(paths)`, `set_permissions(paths, permissions_str)`.

## Cask Installation Process

1. **Download** — The file at `url` is downloaded (via `curl` by default, or `:svn`/`:git`/`:post`)
2. **Verification** — SHA-256 checksum validated (unless `:no_check`)
3. **Extraction** — Container extracted (`.dmg`, `.zip`, `.tar`, etc.) — autodetected or overridden
4. **Staging** — Files staged in Caskroom at `$(brew --caskroom)/<token>/<version>/`
5. **Preflight** — `preflight` Ruby block or `preflight_steps` run
6. **Artifact Installation** — Artifacts installed (apps moved, binaries symlinked, `.pkg` installed, etc.)
7. **Postflight** — `postflight` Ruby block or `postflight_steps` run
8. **Completion generation** — If `generate_completions_from_executable` present

### Sandboxing

- `app` artifacts: moved via Homebrew-managed operations (sandboxed)
- `pkg` artifacts: installed by macOS `/usr/sbin/installer` (not sandboxed)
- `installer script:` runs without sandbox
- `generate_completions_from_executable`: sandboxed (where available)

## Global Cask Directory Flags

These flags apply to `brew install --cask`, `brew reinstall --cask`, and `brew upgrade --cask`:

| Flag | Default |
|------|---------|
| `--appdir` | `/Applications` |
| `--appimagedir` | `~/Applications` |
| `--fontdir` | `~/Library/Fonts` |
| `--prefpanedir` | `~/Library/PreferencePanes` |
| `--qlplugindir` | `~/Library/QuickLook` |
| `--mdimporterdir` | `~/Library/Spotlight` |
| `--colorpickerdir` | `~/Library/ColorPickers` |
| `--dictionarydir` | `~/Library/Dictionaries` |
| `--servicedir` | `~/Library/Services` |
| `--input-methoddir` | `~/Library/Input Methods` |
| `--internet-plugindir` | `~/Library/Internet Plug-Ins` |
| `--audio-unit-plugindir` | `~/Library/Audio/Plug-Ins/Components` |
| `--vst-plugindir` | `~/Library/Audio/Plug-Ins/VST` |
| `--vst3-plugindir` | `~/Library/Audio/Plug-Ins/VST3` |
| `--screen-saverdir` | `~/Library/Screen Savers` |
| `--keyboard-layoutdir` | `/Library/Keyboard Layouts` |
| `--language` | System language |

These can also be set globally via `HOMEBREW_CASK_OPTS`.

## Cask::DSL and Cask::Installer

The `Cask::DSL` class is the Ruby class representing the DSL. Key constants:

- `ORDINARY_ARTIFACT_CLASSES` — 27 artifact classes
- `ACTIVATABLE_ARTIFACT_CLASSES` — all minus `StageOnly`
- `DSL_METHODS` — set of all recognized DSL method symbols

Uses `set_unique_stanza` to enforce "may only appear once" constraints. Tracks `*_set_in_block` booleans for `on_<system>` block handling.

The `Cask::Installer` handles the installation process. The `Cask::CaskLoader` loads casks from taps or the API.

## References

- [Cask Cookbook](https://docs.brew.sh/Cask-Cookbook) — official cask DSL reference
- For formulae: [formula.md](formula.md)
- For taps: [taps.md](taps.md)
- For environment variables: [environment.md](environment.md)
