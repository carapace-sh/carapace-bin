# RubyGems and Bundler

Ruby's package system — `gem` (RubyGems), `Gem::Specification` (gemspec), `Gemfile`, `Bundler`, and the relationship between them.

> **Source of truth**: <https://guides.rubygems.org/>, <https://bundler.io/man/>. For **environment variables** like `GEM_HOME`, see [environment.md](environment.md).

## RubyGems (`gem`)

RubyGems is the package manager bundled with Ruby. It installs gems from <https://rubygems.org>.

### Common Commands

| Command | Description |
|---------|-------------|
| `gem install <name>` | Install a gem |
| `gem install <name> -v 1.2.3` | Install specific version |
| `gem install <name> --version "< 2.0"` | Install with version constraint |
| `gem install <name> --pre` | Install pre-release |
| `gem uninstall <name>` | Uninstall |
| `gem update <name>` | Update to latest |
| `gem update` | Update all gems |
| `gem list` | List installed gems |
| `gem list <name>` | List gems matching pattern |
| `gem search <name> --remote` | Search remote gems |
| `gem sources` | List gem sources |
| `gem sources -a <url>` | Add a source |
| `gem env` | Show gem environment |
| `gem env home` | Show GEM_HOME |
| `gem which <lib>` | Find where a gem's file is |
| `gem dependency <name>` | Show dependencies |
| `gem open <name>` | Open gem source in editor |
| `gem cleanup` | Remove old versions |
| `gem specification <name>` | Show gemspec |
| `gem build` | Build .gem from .gemspec |
| `gem push <file.gem>` | Push to rubygems.org |
| `gem owner <name>` | Manage gem owners |
| `gem yank <name> -v <ver>` | Unpush a version (marks as uninstallable) |

### Gem Paths

```bash
$ gem env
RubyGems Environment:
  - RUBYGEMS VERSION: 3.4.0
  - RUBY VERSION: 3.2.0
  - INSTALLATION DIRECTORY: /usr/lib/ruby/gems/3.2.0  # GEM_HOME
  - GEM PATHS:
     - /usr/lib/ruby/gems/3.2.0
     - /home/user/.gem/ruby/3.2.0  # user-installed fallback
```

Gems are installed to `GEM_HOME/gems/<name>-<version>/`. Executables go to `GEM_HOME/bin/` (or `GEM_HOME/bin` on the `PATH`).

## Gem::Specification (gemspec)

A gem is defined by a `.gemspec` file (or block in a `Rakefile`):

```ruby
Gem::Specification.new do |spec|
  spec.name          = "mygem"
  spec.version       = "1.0.0"
  spec.summary       = "A useful gem"
  spec.description   = "A longer description..."
  spec.authors       = ["Alice"]
  spec.email         = ["alice@example.com"]
  spec.homepage      = "https://example.com"
  spec.license       = "MIT"
  spec.required_ruby_version = ">= 2.7"

  spec.files = Dir["lib/**/*.rb"]
  spec.bindir = "exe"
  spec.executables = ["mygem"]

  spec.add_dependency "json", "~> 2.0"
  spec.add_dependency "thor", ">= 1.0"

  spec.add_development_dependency "rspec", "~> 3.0"
  spec.add_development_dependency "rake", "~> 13.0"
end
```

### Key Fields

| Field | Purpose |
|-------|---------|
| `name`, `version` | Identity (SemVer) |
| `summary`, `description` | Short/long description |
| `authors`, `email` | Authorship |
| `license` / `licenses` | SPDX identifier(s) |
| `homepage` | Project URL |
| `files` | Files included in the gem |
| `bindir`, `executables` | Binary scripts |
| `require_paths` | Default: `["lib"]` — where `require` looks |
| `required_ruby_version` | Ruby version constraint |
| `add_dependency` | Runtime dependency |
| `add_development_dependency` | Dev-only dependency (not installed when gem is a dep) |

### Version Constraints

| Syntax | Meaning |
|--------|---------|
| `"1.2.3"` | Exactly 1.2.3 |
| `"~> 1.0"` | >= 1.0, < 2.0 |
| `"~> 1.2.3"` | >= 1.2.3, < 1.3.0 |
| `">= 1.0"` | >= 1.0 |
| `">= 1.0", "< 2.0"` | >= 1.0, < 2.0 (multiple) |
| `">= 1.2.3", "< 2.0"` | Range |

`~>` (pessimistic / twiddle-wakka) is the most common — it allows patch updates within the same minor version.

## Bundler

Bundler ensures the right gem versions are used across all environments. It's the standard for Ruby projects.

### Gemfile

```ruby
source "https://rubygems.org"

gem "rails", "7.0.0"
gem "puma", "~> 5.0"
gem "pg", ">= 1.0"

# Groups
gem "rspec", group: :test
gem "pry", group: :development

# Multiple groups
gem "byebug", groups: [:development, :test]

# Git source
gem "foo", git: "https://github.com/foo/foo.git", branch: "main"

# Path source (local)
gem "bar", path: "../bar"

# Platform-conditional
gem "tzinfo", platforms: [:ruby, :mingw, :mswin]
gem "tzinfo-data", platforms: [:x64_mingw, :mswin]

# Version from method call
ruby "3.2.0"  # require Ruby version
```

### Gemfile.lock

`bundle install` resolves all dependencies and writes `Gemfile.lock`:

```
GEM
  remote: https://rubygems.org/
  specs:
    actionpack (7.0.0)
      activesupport (= 7.0.0)
      ...
    puma (5.6.4)

PLATFORMS
  ruby

DEPENDENCIES
  puma (~> 5.0)
  rails (= 7.0.0)

RUBY VERSION
   ruby 3.2.0p0

BUNDLED WITH
   2.4.0
```

Commit `Gemfile.lock` for applications. For gems, it's optional (downstream resolves their own).

### Common Bundler Commands

| Command | Description |
|---------|-------------|
| `bundle install` | Install all gems from Gemfile |
| `bundle install --without test` | Skip group |
| `bundle install --deployment` | Deploy mode (frozen, no updates) |
| `bundle update <gem>` | Update a specific gem |
| `bundle update` | Update all gems |
| `bundle exec <cmd>` | Run a command in the bundle context |
| `bundle config` | Show/set config |
| `bundle config set --local path 'vendor/bundle'` | Install to local path |
| `bundle add <gem>` | Add a gem to Gemfile |
| `bundle remove <gem>` | Remove a gem from Gemfile |
| `bundle outdated` | Show gems with newer versions |
| `bundle lock` | Update Gemfile.lock without installing |
| `bundle check` | Verify all deps installed |
| `bundle clean` | Remove unused gems |
| `bundle open <gem>` | Open gem source |
| `bundle binstubs <gem>` | Generate executable stubs in `bin/` |
| `bundle info <gem>` | Show gem info |
| `bundle show` | Show all gems |
| `bundle list` | List all gems |
| `bundle console` | Start IRB with bundle loaded |
| `bundle viz` | Visualize dependency graph |
| `bundle lock --add-platform x86_64-linux` | Add a platform to lockfile |

### bundle exec

```bash
bundle exec rspec
bundle exec rails server
```

`bundle exec` activates only the gems listed in `Gemfile.lock`, ensuring the right versions are used. Without it, RubyGems may activate the latest installed version.

Mechanism: Bundler manipulates `$LOAD_PATH` and RubyGems' activate hook to load exactly the locked versions.

### BUNDLE_GEMFILE

```bash
BUNDLE_GEMFILE=Gemfile.alt bundle exec rake
```

Sets which Gemfile to use. Default is `Gemfile` in the current directory.

### Groups

```ruby
gem "pry", group: :development
gem "rspec-rails", group: [:development, :test]

# Install without a group
# bundle install --without development test
```

Groups let you skip dev/test gems in production. Common groups: `:default` (implicit), `:development`, `:test`, `:production`.

### Path and Git Sources

```ruby
# Local path (for development across gems)
gem "mylib", path: "../mylib"

# Git
gem "foo", git: "https://github.com/foo/foo.git"
gem "foo", git: "https://github.com/foo/foo.git", branch: "main"
gem "foo", git: "https://github.com/foo/foo.git", tag: "v1.0.0"
gem "foo", git: "https://github.com/foo/foo.git", ref: "abc123"

# GitHub shorthand
gem "foo", github: "foo/foo", branch: "main"
```

Path sources are useful in a monorepo or for local development. Git sources lock to a specific commit in `Gemfile.lock`.

## Gemfile vs Gemspec

| Aspect | Gemfile | gemspec |
|--------|---------|---------|
| Used by | Applications (Bundler) | Gems (RubyGems) |
| Purpose | Lock exact versions for the app | Declare what the gem requires |
| Lives in | Project root | Project root, named `<name>.gemspec` |
| Consumed by | Bundler | RubyGems + Bundler |
| Lockfile | `Gemfile.lock` | None (downstream resolves) |

A **gem project** typically has both:
- A `.gemspec` declaring dependencies (consumable by anyone who installs the gem)
- A `Gemfile` that references the gemspec:

```ruby
# Gemfile
source "https://rubygems.org"

gemspec
# or with groups:
# gemspec development_group: :development
```

This lets Bundler install the gem's own dependencies for development, using the gemspec as the source of truth.

## Edge Cases and Known Issues

- **`bundle exec` vs `bin/rails`** — binstubs (`bin/rails`) already have the bundle loaded, so `bundle exec` is redundant when using them
- **Gem conflict resolution** — Bundler fails if two gems require incompatible versions of a transitive dep. Resolve by updating or pinning.
- **Platform-specific lockfile entries** — `x86_64-linux` is different from `x86_64-darwin`. Use `bundle lock --add-platform` to support CI on different OSes.
- **`Gemfile.lock` should not be edited** — regenerate with `bundle install` or `bundle lock`
- **`gemspec` `files` must include everything** — if you forget a file, it won't be in the `.gem`; use `spec.files = Dir["{lib,exe}/**/*"]`
- **`required_ruby_version`** — set this to prevent install on unsupported Ruby versions
- **`BUNDLE_FROZEN=true`** — equivalent to `--frozen`, prevents any changes to `Gemfile.lock`
- **`BUNDLE_PATH`** — sets the install location for `bundle install --path`
- **`bundle outdated` only checks rubygems.org** — for git sources, use `bundle outdated <gem>` or check manually

## Related

- [cli.md](cli.md) — `ruby` interpreter flags, `rake`, `rdoc`, `irb`
- [environment.md](environment.md) — `GEM_HOME`, `GEM_PATH`, `BUNDLE_GEMFILE`, `BUNDLE_PATH`
- [stdlib.md](stdlib.md) — `rubygems` integration, `bundler/inline` for single-file scripts
