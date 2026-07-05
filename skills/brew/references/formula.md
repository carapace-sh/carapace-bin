# Formula DSL

The formula is Homebrew's core package definition — a Ruby class that builds software from upstream sources. This document covers the `Formula` class, its DSL, the installation pipeline, and the directory layout.

## Core Terminology

| Term | Description | Example |
|------|-------------|---------|
| **Formula** | A package definition written in Ruby that builds from upstream sources | `Formula/f/foo.rb` |
| **Prefix** | The path in which Homebrew is installed | `/opt/homebrew` (Apple Silicon), `/usr/local` (Intel) |
| **Keg** | The installation destination directory of a given formula version | `Cellar/foo/0.1` |
| **Rack** | A directory containing one or more versioned kegs | `Cellar/foo` |
| **Keg-only** | A formula that is **not** symlinked into Homebrew's prefix | `openjdk` |
| **Opt prefix** | A symlink to the active version of a keg | `opt/foo` |
| **Cellar** | The directory containing all named racks | `Cellar/` |
| **Tap** | A Git repository of formulae, casks, and/or external commands | `Library/Taps/homebrew/homebrew-core` |
| **Bottle** | A pre-built keg poured into a rack instead of building from source | `foo--0.1.ventura.bottle.tar.gz` |
| **Tab** | Metadata JSON file about a keg | `Cellar/foo/0.1/INSTALL_RECEIPT.json` |

> **Source of truth**: <https://docs.brew.sh/Formula-Cookbook>. For casks (pre-built binary packages), see [cask.md](cask.md).

## Formula Structure

A formula is a Ruby class inheriting from `Formula`. The file defines a subclass with installation instructions and metadata.

```ruby
class Foo < Formula
  desc "Description of the package"
  homepage "https://example.com/"
  url "https://example.com/foo-0.1.tar.gz"
  sha256 "abc123def456..."
  license "BSD-2-Clause"

  depends_on "pkgconf" => :build
  depends_on "jpeg"
  depends_on "readline" => :recommended
  depends_on "gtk+" => :optional

  def install
    system "./configure", "--disable-debug",
                          "--disable-dependency-tracking",
                          "--prefix=#{prefix}"
    system "make", "install"
  end

  test do
    assert_match "OK", shell_output("#{bin}/foo --version")
  end
end
```

### Key DSL Elements

| Element | Purpose |
|---------|---------|
| `desc` | One-line description (shown in `brew info`, `brew search`) |
| `homepage` | Project homepage URL |
| `url` | Source download URL |
| `sha256` | Checksum of the source archive |
| `version` | Explicit version (if not inferable from URL) |
| `license` | SPDX license identifier (required for homebrew-core) |
| `head` | Alternate download for development/cutting-edge source (activated with `--HEAD`) |
| `revision` | Forced recompile number for non-version-bump changes |
| `version_scheme` | For version scheme changes (e.g. `13` → `1.0.0`) |
| `keg_only` | Marks a formula as keg-only (not symlinked into prefix) |
| `depends_on` | Specifies dependencies (build-time, test-time, optional, recommended) |
| `conflicts_with` | Declares conflicts with other formulae |
| `uses_from_macos` | Dependencies provided by macOS itself (Linux uses formula instead) |
| `resource` | Additional downloadable resources (e.g. Python dependencies) |
| `patch` | Patches to apply (embedded via `__END__` or external URLs) |
| `livecheck` | Controls version detection for `brew livecheck` |
| `service` | Defines launchd/systemd service definitions for `brew services` |
| `caveats` | Post-install informational messages to users |
| `post_install` | Post-installation setup steps |
| `option` | Optional build flags (not allowed in homebrew-core) |
| `fails_with` | Declares compiler versions known to cause build failures |
| `deprecated_option` | For renaming/removing old options |
| `bottle` | Pre-built binary package definitions (see [bottle.md](bottle.md)) |

### Dependency Types

The `depends_on` method accepts a hash to specify the dependency type:

| Symbol | Meaning |
|--------|---------|
| `:build` | Needed only at build time, not at runtime |
| `:test` | Needed only for running `brew test` |
| `:recommended` | Installed by default, but can be skipped with `--without-<dep>` |
| `:optional` | Not installed by default, enabled with `--with-<dep>` |
| _(none)_ | Runtime dependency (always installed) |

```ruby
depends_on "pkgconf" => :build    # build-time only
depends_on "openssl@3"            # runtime
depends_on "sqlite" => :recommended  # installed by default
depends_on "gui" => :optional     # opt-in
```

### Platform-Specific Blocks

Formulae can nest dependencies and logic inside platform blocks:

| Block | Condition |
|-------|-----------|
| `on_macos do ... end` | macOS-specific |
| `on_linux do ... end` | Linux-specific |
| `on_arm do ... end` | Apple Silicon-specific |
| `on_intel do ... end` | Intel-specific |
| `on_sequoia :or_newer do ... end` | macOS version-specific |
| `on_system :linux, macos: :sonoma_or_older do ... end` | Combined conditions |

```ruby
on_macos do
  depends_on "openjdk" => :build
end

on_linux do
  depends_on "openjdk@17" => :build
end
```

### Resources

A `resource` is an additional downloadable file used during the build (e.g. Python dependencies, test data):

```ruby
resource "homebrew-test" do
  url "https://example.com/test-data-1.0.tar.gz"
  sha256 "def456..."
end
```

Resources are accessed in the `install` method via `buildpath` and `staging` methods.

### Patches

Patches can be embedded or fetched:

```ruby
# External patch (fetched from URL)
patch do
  url "https://example.com/fix.patch"
  sha256 "abc123..."
end

# Inline patch (at end of file)
patch :DATA
__END__
diff --git a/foo.c b/foo.c
...
```

## The `install` Method

The `install` method contains the build instructions. It runs in the `staged_path` (the temporary build directory where the source tarball was extracted) with `superenv` active. Files are installed *into* the keg (`prefix`) during this method.

### Standard Build Arguments

Homebrew provides `std_*_args` methods for common build systems:

| Method | Build System |
|--------|-------------|
| `std_cmake_args` | CMake |
| `std_configure_args` | Autotools (`./configure`) |
| `std_cargo_args` | Cargo/Rust |
| `std_meson_args` | Meson |
| `std_go_args` | Go |
| `std_cabal_v2_args` | Haskell Cabal |
| `std_npm_args` | npm |
| `std_pip_args` | pip |
| `std_zig_args` | Zig |

```ruby
def install
  system "cmake", *std_cmake_args, "-S", ".", "-B", "build"
  system "cmake", "--build", "build"
  system "cmake", "--install", "build"
end
```

### The `system` Call

The `system` method runs a command with `superenv` environment. Arguments are passed as separate strings (not a single shell string) to avoid shell injection:

```ruby
# Good — arguments are separate
system "make", "install", "PREFIX=#{prefix}"

# Bad — shell interpolation
system "make install PREFIX=#{prefix}"
```

## Directory Location Variables

These methods return paths within the formula's keg or Homebrew's prefix:

| Variable | Path | Example |
|----------|------|---------|
| `prefix` | `#{HOMEBREW_PREFIX}/Cellar/#{name}/#{version}` | `/opt/homebrew/Cellar/foo/0.1` |
| `opt_prefix` | `#{HOMEBREW_PREFIX}/opt/#{name}` | `/opt/homebrew/opt/foo` |
| `bin` | `#{prefix}/bin` | `Cellar/foo/0.1/bin` |
| `lib` | `#{prefix}/lib` | `Cellar/foo/0.1/lib` |
| `include` | `#{prefix}/include` | `Cellar/foo/0.1/include` |
| `share` | `#{prefix}/share` | `Cellar/foo/0.1/share` |
| `libexec` | `#{prefix}/libexec` | `Cellar/foo/0.1/libexec` |
| `etc` | `#{HOMEBREW_PREFIX}/etc` | `/opt/homebrew/etc` |
| `var` | `#{HOMEBREW_PREFIX}/var` | `/opt/homebrew/var` |
| `buildpath` | Temporary build directory | — |
| `staged_path` | Temporary source build directory where the tarball is extracted | — |

The `opt_` variants (`opt_bin`, `opt_lib`, etc.) provide stable paths across version updates — they're symlinks to the active keg.

## Superenv

When building from source, Homebrew uses `superenv` — a "super environment" that:

1. **Isolates builds** by removing `/usr/local/bin` and non-essential `PATH` entries
2. **Removes bad flags** from compiler commands
3. **Injects keg-only dependencies** into `-I` and `-L` flags automatically

This ensures reproducible builds regardless of the user's system configuration.

### ENV Helper Methods

Within the `install` method, `ENV` provides helpers for manipulating the build environment:

| Method | Purpose |
|--------|---------|
| `ENV.cxx11` | Compile with C++11 features |
| `ENV.deparallelize` | Single-job compilation (can take a block) |
| `ENV.O0` / `ENV.O1` / `ENV.O3` | Set optimization levels (default: macOS `-Os`, Linux `-O2`) |
| `ENV.runtime_cpu_detection` | Handle runtime CPU detection |
| `ENV.append_to_cflags` | Add to `CFLAGS`, `CXXFLAGS`, `OBJCFLAGS`, `OBJCXXFLAGS` |
| `ENV.prepend_create_path` | Create and prepend a path to a list |
| `ENV.remove` | Remove a string from a variable |
| `ENV.delete` | Unset a variable |
| `ENV["VAR"] = "VALUE"` | Set directly |
| `with_env` | Temporarily set variables within a block |

## Messaging Methods

| Method | Purpose |
|--------|---------|
| `ohai` | General info message |
| `opoo` | Warning message |
| `odie` | Error message (exits immediately) |

## The `test` Block

The `test do ... end` block defines a test that `brew test` runs in a temporary directory. It should verify the installation works:

```ruby
test do
  assert_match "1.0", shell_output("#{bin}/foo --version")
  system "#{bin}/foo", "--help"
end
```

The test sandbox has the formula's `bin`, `lib`, and `include` on `PATH` and related env vars. Tests should be fast and not require network access (unless necessary).

## Caveats

The `caveats` method returns a string shown after installation and in `brew info`:

```ruby
def caveats
  <<~EOS
    To start the service:
      brew services start foo
  EOS
end
```

## Post-Install

The `post_install` method runs after the keg is installed and linked:

```ruby
def post_install
  system "#{bin}/foo", "--init"
end
```

## Formulary — Formula Loading

The `Formulary` module is the factory for loading formulae. It supports three loading strategies:

| Strategy | Description |
|----------|-------------|
| **Path loading** | Loads from a direct filesystem path by evaluating the Ruby DSL within a unique `FormulaNamespace` module |
| **API loading** | Loads formula metadata from the Homebrew JSON API (default in modern Homebrew) |
| **Tap loading** | Resolves a formula from a specific tap |

In API mode (default since Homebrew 4.0), formula metadata is fetched as JSON from `https://formulae.brew.sh/api` rather than requiring a full local Git clone. The Ruby formula file is still needed for building from source.

## FormulaInstaller — Installation Pipeline

The `FormulaInstaller` orchestrates the installation of a single formula through these phases:

| Phase | What Happens |
|-------|-------------|
| **Prelude** | Fetches dependencies, validates environment |
| **Fetch** | Downloads bottles or source code via a `DownloadQueue` |
| **Install** | Pours a bottle (extracts archive) OR executes the formula's `install` method |
| **Clean** | Removes build-only files and temporary directories |
| **Finish** | Links the keg into the prefix, runs `post_install`, writes tab |

### Installation Command Flow

```
brew install <formula>
  → Build Flag Validation (check for dev tools if building from source)
    → Formula Filtering (determine which formulae need installation)
      → Installer Creation (FormulaInstaller instances)
        → Fetch Phase (DownloadQueue — parallel downloads)
          → Install Execution (pour bottle or build from source)
            → Clean + Finish (link, post-install, tab)
```

The `brew install` command partitions input into formulae and casks, then processes each through separate paths.

## References

- [Formula Cookbook](https://docs.brew.sh/Formula-Cookbook) — official DSL reference
- For casks: [cask.md](cask.md)
- For bottles: [bottle.md](bottle.md)
- For the Cellar and keg linking: [cellar.md](cellar.md)
- For service blocks: [services.md](services.md)
- For taps: [taps.md](taps.md)
