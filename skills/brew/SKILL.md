---
name: brew
description: >
  Use when working with Homebrew (`brew`) — the package manager for macOS and Linux.
  Covers the command structure, formula and cask DSLs, taps, bottles, the Cellar, services,
  bundle/Brewfile, environment variables, installation internals, and developer commands.
  Triggers on: "brew", "homebrew", "brew install", "brew uninstall", "brew update",
  "brew upgrade", "brew search", "brew tap", "brew cask", "brew bottle", "brew services",
  "brew bundle", "brew cleanup", "brew doctor", "brew audit", "brew livecheck",
  "brew test-bot", "brew bump-formula-pr", "Brewfile", "HOMEBREW_PREFIX", "HOMEBREW_CELLAR",
  "HOMEBREW_NO_INSTALL_FROM_API", "HOMEBREW_CASK_OPTS", "HOMEBREW_BOTTLE_DOMAIN", "Formula",
  "Cask", "Cellar", "Caskroom", "keg", "rack", "tap", "bottle", "superenv", "Formulary",
  "FormulaInstaller", "shellenv", "restart_service", "keg-only", "brew deps", "brew uses",
  "brew leaves", "brew outdated", "brew pin", "brew fetch", "brew create", "brew edit",
  "livecheck", "auto_updates", "depends_on".
user-invocable: true
---

# Homebrew (`brew`) In-Depth Reference

Comprehensive reference for [Homebrew](https://brew.sh) — the missing package manager for macOS (and Linux). Covers the full CLI surface, the formula and cask DSLs, the installation pipeline, taps and bottles, services, bundle, environment variables, and developer tooling.

## Data Flow

```
brew install <formula>
  → Formulary loads formula (API JSON or tap .rb file)
    → FormulaInstaller
      → Fetch phase (bottle download or source download)
        → Install phase (pour bottle OR build from source via superenv)
          → Clean phase (remove build-only files)
            → Finish phase (link keg into prefix, run post-install)
              → Tab written (INSTALL_RECEIPT.json)
```

```
brew install --cask <cask>
  → CaskLoader loads cask (.rb from tap or API)
    → Cask::Installer
      → Download (curl)
        → Verify (sha256)
          → Extract container (dmg/zip/tar)
            → Stage in Caskroom
              → Preflight → install artifacts → Postflight
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| install, uninstall, reinstall, upgrade, update, search, info, list, ls, fetch, formula, cask, formulae, casks, --formula, --cask, --build-from-source, --force-bottle, --HEAD, --dry-run, --ignore-dependencies, --only-dependencies, --cc, --keep-tmp, --overwrite, --adopt, --zap, --require-sha, --binaries, brew command, brew commands, brew help, brew home, --version, --cache, --cellar, --caskroom, --prefix, --repository, --env, shellenv, brew which-formula, brew source, brew docs, global options, --debug, --verbose, --quiet | [references/commands.md](references/commands.md) |
| formula, Formula class, formula DSL, def install, depends_on, resource, url, sha256, version, head, license, desc, homepage, bottle block, keg_only, conflicts_with, uses_from_macos, option, deprecated_option, patch, livecheck, service do, caveats, post_install, test do, revision, version_scheme, fails_with, std_cmake_args, std_configure_args, std_go_args, std_cargo_args, std_meson_args, std_pip_args, ENV, superenv, on_macos, on_linux, on_arm, on_intel, on_system, prefix, bin, lib, include, share, libexec, opt_prefix, FormulaInstaller, Formulary, formula loading, path loading, API loading, tap loading | [references/formula.md](references/formula.md) |
| cask, Cask, cask DSL, cask stanza, version, sha256, url, app, pkg, installer, binary, manpage, font, zap, uninstall, caveats, auto_updates, conflicts_with, depends_on, container, livecheck, no_autobump, deprecate, disable, language, rename, suite, stage_only, artifact, preflight, postflight, preflight_steps, postflight_steps, generate_completions_from_executable, Cask::DSL, Cask::Installer, CaskLoader, Caskroom, --appdir, --fontdir, --language, --require-sha, cask vs formula | [references/cask.md](references/cask.md) |
| tap, untap, tap-info, tap-new, homebrew/core, homebrew/cask, CoreTap, CoreCaskTap, third-party tap, homebrew- prefix, tap directory structure, Formula/, Casks/, cmd/, Aliases/, API mode, git clone mode, brew tap --repair, --custom-remote, tap trust, brew trust, brew untrust, HOMEBREW_NO_INSTALL_FROM_API, formula resolution order, fully qualified name | [references/taps.md](references/taps.md) |
| bottle, BottleSpecification, bottle block, bottle tag, pour, relocate, relocation, --build-bottle, brew bottle, --skip-relocation, --force-core-tap, --no-rebuild, --keep-old, --json, --merge, --write, --no-commit, --root-url, ghcr.io, GitHub Container Registry, HOMEBREW_BOTTLE_DOMAIN, bottle naming convention, brew unbottled, --tag, --dependents | [references/bottle.md](references/bottle.md) |
| services, brew services, service do, service block, run, run_type, interval, cron, keep_alive, launch_only_once, require_root, environment_variables, working_dir, log_path, error_log_path, restart_delay, sockets, nice, process_type, throttle_interval, macos_legacy_timers, name, launchctl, launchd, systemctl, systemd, LaunchAgents, LaunchDaemons, start, stop, restart, run, kill, cleanup, --all, --file, --json, --keep, --no-wait, --max-wait, sudo, std_service_path_env | [references/services.md](references/services.md) |
| bundle, Brewfile, brew bundle, brew bundle install, brew bundle dump, brew bundle cleanup, brew bundle check, brew bundle exec, brew bundle sh, brew bundle list, brew bundle add, brew bundle remove, brew bundle edit, brew bundle env, brew bundle upgrade, tap entry, brew entry, cask entry, mas entry, vscode entry, go entry, cargo entry, uv entry, flatpak entry, winget entry, krew entry, npm entry, whalebrew entry, cask_args, restart_service, start_service, link, conflicts_with, postinstall, version_file, greedy, trusted, HOMEBREW_BUNDLE_FILE, HOMEBREW_BUNDLE_JOBS | [references/bundle.md](references/bundle.md) |
| environment variable, HOMEBREW_PREFIX, HOMEBREW_CELLAR, HOMEBREW_REPOSITORY, HOMEBREW_CACHE, HOMEBREW_TEMP, HOMEBREW_NO_INSTALL_FROM_API, HOMEBREW_NO_AUTO_UPDATE, HOMEBREW_AUTO_UPDATE_SECS, HOMEBREW_API_DOMAIN, HOMEBREW_BOTTLE_DOMAIN, HOMEBREW_ARTIFACT_DOMAIN, HOMEBREW_CASK_OPTS, HOMEBREW_NO_ANALYTICS, HOMEBREW_DEBUG, HOMEBREW_DEVELOPER, HOMEBREW_COLOR, HOMEBREW_EDITOR, HOMEBREW_BROWSER, HOMEBREW_GITHUB_API_TOKEN, HOMEBREW_FORCE_BREWED_CURL, HOMEBREW_FORCE_BREWED_GIT, HOMEBREW_SORBET_RUNTIME, HOMEBREW_NO_INSECURE_REDIRECT, HOMEBREW_CLEANUP_MAX_AGE_DAYS, HOMEBREW_DOWNLOAD_CONCURRENCY, HOMEBREW_FORBIDDEN_TAPS, HOMEBREW_FORBIDDEN_FORMULAE, HOMEBREW_ALLOWED_TAPS, brew.env, brew config, brew doctor, configuration | [references/environment.md](references/environment.md) |
| cleanup, autoremove, prune, brew cleanup, --prune, --scrub, --prune-prefix, brew autoremove, --dry-run, brew doctor, --list-checks, --audit-debug, brew missing, brew migrate, brew pin, brew unpin, brew outdated, --json, --greedy, --fetch-HEAD, brew leaves, --installed-on-request, --installed-as-dependency, brew tab, brew linkage, --test, --reverse, --cached, INSTALL_RECEIPT, tab metadata | [references/maintenance.md](references/maintenance.md) |
| deps, uses, brew deps, brew uses, --topological, --direct, --tree, --graph, --dot, --annotate, --union, --include-build, --include-optional, --include-test, --include-implicit, --skip-recommended, --installed, --missing, --for-each, --HEAD, --os, --arch, --full-name, recursive dependency graph, reverse dependency | [references/dependency.md](references/dependency.md) |
| audit, style, test, tests, test-bot, typecheck, tc, create, edit, cat, extract, livecheck, lc, bump, bump-formula-pr, bump-cask-pr, bump-revision, bump-unversioned-casks, bump-compatibility-version, readall, unpack, irb, ruby, sh, prof, rubocop, vendor-gems, update-python-resources, update-perl-resources, generate-man-completions, generate-zap, tap-new, debugger, gist-logs, developer mode, HOMEBREW_DEVELOPER, Sorbet, RuboCop, --strict, --online, --git, --new, --fix | [references/developer.md](references/developer.md) |
| Cellar, keg, rack, keg-only, opt prefix, opt_prefix, prefix, --prefix, --unbrewed, --installed, symlink, link, unlink, brew link, brew unlink, --overwrite, --force, --HEAD, --dry-run, pour, relocation, staged path, Caskroom, directory layout, HOMEBREW_PREFIX, HOMEBREW_CELLAR | [references/cellar.md](references/cellar.md) |

## Quick Guide

- **How do I install, uninstall, or upgrade a package?** → [references/commands.md](references/commands.md)
- **How do I write or modify a formula?** → [references/formula.md](references/formula.md)
- **How do I write or modify a cask?** → [references/cask.md](references/cask.md)
- **How do I manage taps?** → [references/taps.md](references/taps.md)
- **How do bottles work and how do I create one?** → [references/bottle.md](references/bottle.md)
- **How do I manage background services?** → [references/services.md](references/services.md)
- **How do I use Brewfiles and `brew bundle`?** → [references/bundle.md](references/bundle.md)
- **What environment variables does brew recognize?** → [references/environment.md](references/environment.md)
- **How do I clean up, diagnose, or troubleshoot?** → [references/maintenance.md](references/maintenance.md)
- **How do I inspect dependencies?** → [references/dependency.md](references/dependency.md)
- **What developer commands are available (audit, test, style, livecheck)?** → [references/developer.md](references/developer.md)
- **How does the Cellar and keg linking work?** → [references/cellar.md](references/cellar.md)
- **How do I search for packages?** → [references/commands.md](references/commands.md)
- **How do I pin packages or check what's outdated?** → [references/maintenance.md](references/maintenance.md)
- **What's the difference between a formula and a cask?** → [references/formula.md](references/formula.md) and [references/cask.md](references/cask.md)
- **How do I set up brew in my shell?** → [references/commands.md](references/commands.md) (`shellenv`)
- **How do I create a PR to update a formula or cask?** → [references/developer.md](references/developer.md)
- **What are the global cask directory flags (--appdir, --fontdir, etc.)?** → [references/cask.md](references/cask.md)

## Cross-Project References

- For **Ruby** language fundamentals (classes, blocks, DSLs, gems, Bundler), consult Ruby's official documentation — brew's formula and cask DSLs are Ruby-based.
- For **Git** internals (cloning, remotes, branches, PRs), use the **git** skill.
- For **carapace** completion integration (the `brew` completer in carapace-bin, specs, actions, macros), see the **carapace** skill.
- For **launchd** and **systemd** internals (plist format, unit files, service management beyond brew's wrapper), consult Apple's and systemd's official documentation — this skill covers only brew's `services` abstraction layer.
