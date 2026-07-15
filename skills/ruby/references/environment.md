# Environment Variables

Environment variables that affect Ruby, RubyGems, and Bundler — interpreter behavior, load path, gem locations, and Bundler configuration.

> **Source of truth**: `ruby --help`, `gem env`, `bundle config`. For **CLI flags**, see [cli.md](cli.md). For **gem paths**, see [gems-bundler.md](gems-bundler.md).

## Ruby Interpreter

### RUBYOPT

Extra options appended to the command line for every `ruby` invocation:

```bash
export RUBYOPT="-w"
ruby script.rb  # equivalent to: ruby -w script.rb
```

Common uses: `-w` (warnings), `-I/path` (load path), `-rfoo` (require).

### RUBYLIB

Prepends directories to `$LOAD_PATH`:

```bash
export RUBYLIB=/home/user/mylib
ruby -e 'puts $LOAD_PATH.first'  # => /home/user/mylib
```

### RUBY_ENGINE, RUBY_VERSION, etc.

These are constants (not env vars), but commonly referenced:

| Constant | Value |
|----------|-------|
| `RUBY_VERSION` | e.g. `"3.2.0"` |
| `RUBY_ENGINE` | `"ruby"` (MRI), `"jruby"`, `"truffleruby"`, `"mruby"` |
| `RUBY_PLATFORM` | e.g. `"x86_64-linux"` |
| `RUBY_PATCHLEVEL` | patch level (e.g. `0`) |
| `RUBY_RELEASE_DATE` | release date string |
| `RUBY_DESCRIPTION` | full version string |
| `RUBY_REVISION` | git revision |

Check at runtime:

```ruby
if RUBY_ENGINE == "jruby"
  # JRuby-specific code
end
```

### Encoding-Related

| Variable | Effect |
|----------|--------|
| `LANG` | Sets locale, affects `Encoding.default_external` |
| `LC_ALL` | Overrides `LANG` for all locale categories |
| `LC_CTYPE` | Locale for character encoding |

Ruby reads `LANG`/`LC_ALL`/`LC_CTYPE` to set `Encoding.default_external` at startup. Override with `-E` flag.

### RubyVM::FrozenStringLiteral

There's no env var for this — use the magic comment or `--enable-frozen-string-literal`:

```bash
ruby --enable-frozen-string-literal script.rb
ruby --disable-frozen-string-literal script.rb
```

## RubyGems

### GEM_HOME

Where gems are installed:

```bash
export GEM_HOME=/opt/gems
gem install foo    # installs to /opt/gems/gems/foo-1.0/
```

### GEM_PATH

Where RubyGems looks for installed gems (colon-separated):

```bash
export GEM_PATH=/opt/gems:/home/user/.gem/ruby/3.2.0
```

Default: `GEM_HOME` plus the user's `.gem` directory.

### GEM_ROOT

The Ruby's built-in gem directory (read-only, comes with Ruby).

### GEM_RC

Path to `.gemrc` config file (default `~/.gemrc`):

```yaml
# ~/.gemrc
---
:sources:
  - https://rubygems.org
:install: --user-install
gem: "--no-document"
```

### GEM_SPEC_CACHE

Cache directory for gemspec data (default: `GEM_HOME/specifications`).

### Common Gem Env Vars

| Variable | Effect |
|----------|--------|
| `GEM_HOME` | Install location |
| `GEM_PATH` | Search path |
| `GEM_SPEC_CACHE` | Spec cache dir |
| `GEM_RC` | Config file path |
| `RUBY_GEMSPEC` | (rarely used) |

Check with `gem env`:

```bash
$ gem env home
/usr/lib/ruby/gems/3.2.0

$ gem env path
/usr/lib/ruby/gems/3.2.0:/home/user/.gem/ruby/3.2.0

$ gem env gemdir
/usr/lib/ruby/gems/3.2.0
```

## Bundler

### BUNDLE_GEMFILE

Which Gemfile to use:

```bash
BUNDLE_GEMFILE=Gemfile.alt bundle exec rake
```

Default: `Gemfile` in the current directory (searched upward).

### BUNDLE_PATH

Where to install gems for `bundle install`:

```bash
bundle config set --local path 'vendor/bundle'
# or
export BUNDLE_PATH=vendor/bundle
bundle install
```

### BUNDLE_FROZEN

Equivalent to `--frozen` — prevents modifications to `Gemfile.lock`:

```bash
BUNDLE_FROZEN=true bundle install
```

### BUNDLE_WITHOUT

Skip groups:

```bash
BUNDLE_WITHOUT=test:development bundle install
```

### BUNDLE_BIN

Where binstubs are installed (default: `bin/`).

### BUNDLE_JOBS

Parallel jobs for installation:

```bash
BUNDLE_JOBS=4 bundle install
```

### BUNDLE_IGNORE_CONFIG

Ignore Bundler config files (only use env vars):

```bash
BUNDLE_IGNORE_CONFIG=true bundle install
```

### BUNDLE_CACHE_PATH

Where `bundle package` caches gems (default: `vendor/cache`).

### BUNDLER_VERSION

Force a specific Bundler version:

```bash
BUNDLER_VERSION=2.4.0 bundle _2.4.0_ install
```

The `bundle _<version>_` syntax also works (passed as arg).

### Bundler Config (Persistent)

Bundler reads config from:
1. Local: `.bundle/config` (per-project)
2. Env: `BUNDLE_*` variables
3. Global: `~/.bundle/config`

```bash
bundle config set --global path vendor/bundle  # global
bundle config set --local path vendor/bundle   # per-project
bundle config                                 # show all
bundle config path                           # show one
bundle config unset path                     # remove
```

## Load Path and $LOAD_PATH

`$LOAD_PATH` (alias `$:`) is an array of directories Ruby searches when you `require`:

```ruby
$LOAD_PATH
# => ["/usr/lib/ruby/gems/3.2.0/gems/foo-1.0/lib",
#     "/usr/lib/ruby/3.2.0",
#     "/usr/lib/ruby/3.2.0/x86_64-linux",
#     ...]
```

Modify at runtime:

```ruby
$LOAD_PATH.unshift("/my/lib")
require "my_file"  # looks in /my/lib first
```

`-I` flag and `RUBYLIB` env var prepend to `$LOAD_PATH`.

## ARGV and ENV

```ruby
ARGV  # command-line args (Array)
ENV   # environment (Hash-like, read-only)
ENV["PATH"]
ENV.fetch("HOME", "/tmp")
ENV.to_h  # full hash
```

`ENV` is a pseudo-Hash. Keys and values are Strings. Modifying `ENV` with `ENV[key] = val` calls the OS `setenv`.

## Edge Cases and Known Issues

- **`RUBYOPT` is parsed early** — errors in options will prevent Ruby from starting
- **`GEM_HOME` must exist** — RubyGems may create it, but missing parent dirs can cause errors
- **`BUNDLE_FROZEN` + `Gemfile.lock` out of date** — fails immediately; run `bundle install` without `BUNDLE_FROZEN` to update the lockfile
- **`ENV` is process-wide** — modifying it affects all threads and child processes
- **`LANG=C` sets ASCII encoding** — use `LANG=C.UTF-8` or `LANG=en_US.UTF-8` for UTF-8 default
- **`GEM_PATH` order matters** — earlier paths shadow later; useful for overriding a system gem with a local one
- **Bundler env vars override config files** — `BUNDLE_PATH=foo` takes precedence over `.bundle/config`

## Related

- [cli.md](cli.md) — `ruby` flags, `-I`, `-E`
- [gems-bundler.md](gems-bundler.md) — `gem env`, `bundle config`
- [io-encoding.md](io-encoding.md) — `Encoding.default_external` and locale
