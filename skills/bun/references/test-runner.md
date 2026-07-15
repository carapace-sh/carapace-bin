# Bun Test Runner

Reference for Bun's built-in Jest-compatible test runner: `bun test`, the `bun:test` module, configuration, snapshots, coverage, mocks, and concurrent execution.

> **Source of truth**: <https://bun.sh/docs/test>. For bunfig.toml `[test]` section, see [config.md](config.md). For CLI flags, see [cli.md](cli.md).

## Overview

Bun ships with a fast, built-in, Jest-compatible test runner. Tests run in the Bun runtime. It supports TypeScript, JSX, lifecycle hooks, snapshot testing, UI and DOM testing, watch mode (`--watch`), and script pre-loading (`--preload`).

**Module**: `import { test, expect, ... } from "bun:test"`

## Running Tests

```bash
bun test                              # run all tests
bun test <filter> <filter> ...        # filter by file name
bun test ./test/specific-file.test.ts # specific file (prefix with ./ or /)
bun test --test-name-pattern addition # filter by test name (regex)
bun test -t addition                  # shorthand for --test-name-pattern
```

### Test File Discovery

The runner recursively searches the working directory for files matching:

- `*.test.{js|jsx|ts|tsx|mjs|cjs|mts|cts}`
- `*_test.{js|jsx|ts|tsx|mjs|cjs|mts|cts}`
- `*.spec.{js|jsx|ts|tsx|mjs|cjs|mts|cts}`
- `*_spec.{js|jsx|ts|tsx|mjs|cjs|mts|cts}`

### Execution Model

All tests run in a **single process**. Preload scripts are loaded first, then all tests run. If a test fails, the runner exits with non-zero.

## CLI Options

### Execution Control

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--timeout` | number | `5000` | Per-test timeout in milliseconds |
| `--rerun-each` | number | | Re-run each test file N times (surfaces flaky tests) |
| `--retry` | number | | Retry failed tests up to N times |
| `--concurrent` | boolean | `false` | Treat all tests as `test.concurrent()` |
| `--max-concurrency` | number | `20` | Max concurrent tests executing at once |
| `--randomize` | boolean | `false` | Run tests in random order |
| `--seed` | number | | Set random seed (implies `--randomize`) |
| `--bail` | number | `1` | Exit after N failures |

### Test Filtering

| Flag | Type | Description |
|------|------|-------------|
| `--todo` | boolean | Include tests marked with `test.todo()` |
| `--test-name-pattern` / `-t` | string | Run only tests whose name matches the regex |
| `--only-failures` | boolean | Only display failed tests in output |

### Reporting

| Flag | Type | Description |
|------|------|-------------|
| `--reporter` | string | Output format: `junit` (requires `--reporter-outfile`), `dots` |
| `--reporter-outfile` | string | Output file path for the reporter |
| `--dots` | boolean | Shorthand for `--reporter=dots` |

### Snapshots

| Flag | Type | Description |
|------|------|-------------|
| `--update-snapshots` / `-u` | boolean | Update snapshot files |

### Coverage

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--coverage` | boolean | | Generate a coverage profile |
| `--coverage-reporter` | string | `"text"` | Coverage format: `text` and/or `lcov` |
| `--coverage-dir` | string | `"coverage"` | Directory for coverage files |

### Other

| Flag | Type | Description |
|------|------|-------------|
| `--watch` | boolean | Watch for changes and re-run tests |
| `--preload` | string | Load a script before running tests |
| `--path-ignore-patterns` | string | Glob patterns to exclude from test discovery |

## bunfig.toml Configuration

```toml
[test]
root = "./__tests__"
preload = ["./setup.ts"]
pathIgnorePatterns = ["vendor/**", "fixtures/**"]
smol = true
coverage = true
coverageThreshold = 0.9
coverageReporter = ["text", "lcov"]
coverageDir = "coverage"
coverageSkipTestFiles = false
coveragePathIgnorePatterns = ["**/*.spec.ts"]
randomize = true
seed = 2444615283
rerunEach = 3
retry = 3
onlyFailures = true

[test.reporter]
dots = true
junit = "test-results.xml"
```

See [config.md](config.md) for the full `[test]` section reference.

## Test API

### Basic Tests

```ts
import { test, expect } from "bun:test";

test("1 + 1 = 2", () => {
  expect(1 + 1).toBe(2);
});
```

### Test Qualifiers

| API | Description |
|-----|-------------|
| `test(name, fn)` | Sequential test (default) |
| `test.concurrent(name, fn)` | Run concurrently |
| `test.serial(name, fn)` | Force sequential (even with `--concurrent`) |
| `test.todo(name)` | Placeholder test (skipped) |
| `test.failing(name, fn)` | Expected to fail, errors if it passes |
| `test.skip(name, fn)` | Skip this test |
| `test.each([...])(name, fn)` | Parameterized test (runs sequentially) |
| `test.concurrent.each([...])(name, fn)` | Parameterized concurrent |
| `test.failing.each([...])(name, fn)` | Chained qualifiers |

Qualifiers can be chained: `test.failing.each([1, 2, 3])("name %d", (input) => { ... })`.

### Per-Test Options

```ts
test("flaky test", { retry: 5, timeout: 10000 }, () => {
  // ...
});
```

Per-test `{ retry: N }` overrides the global `--retry` flag.

### `describe` Blocks

```ts
import { describe, test, expect } from "bun:test";

describe("math", () => {
  test("addition", () => {
    expect(1 + 1).toBe(2);
  });

  test("subtraction", () => {
    expect(2 - 1).toBe(1);
  });
});
```

## Lifecycle Hooks

| Hook | Description |
|------|-------------|
| `beforeAll(fn)` | Runs once before all tests |
| `beforeEach(fn)` | Runs before each test |
| `afterEach(fn)` | Runs after each test |
| `afterAll(fn)` | Runs once after all tests |

```ts
import { describe, beforeAll, beforeEach, afterEach, afterAll, test } from "bun:test";

beforeAll(async () => {
  // setup once
});

beforeEach(async () => {
  // setup before each test
});
```

Hooks can be defined in test files or in a separate file preloaded with `--preload ./setup.ts`.

## Expect Assertions

Bun implements the Jest-compatible `expect()` API:

### Common Matchers

| Matcher | Description |
|---------|-------------|
| `toBe(value)` | Strict equality (`===`) |
| `toEqual(value)` | Deep equality |
| `toBeNull()` / `toBeUndefined()` / `toBeDefined()` | Type checks |
| `toBeTruthy()` / `toBeFalsy()` | Boolean checks |
| `toBeGreaterThan(n)` / `toBeLessThan(n)` | Numeric comparisons |
| `toBeGreaterThanOrEqual(n)` / `toBeLessThanOrEqual(n)` | |
| `toBeCloseTo(n, digits)` | Float comparison |
| `toMatch(regex)` / `toMatch(string)` | String matching |
| `toContain(item)` | Array/string contains |
| `toHaveLength(n)` | Array/string length |
| `toThrow(error?)` | Expects function to throw |
| `InstanceOf(Class)` | Instance check |
| `toHaveProperty(path, value?)` | Object property check |

### Negation and Async

| Matcher | Description |
|---------|-------------|
| `.not` | Negate any matcher |
| `.resolves` | Unwrap resolved promise value |
| `.rejects` | Unwrap rejected promise reason |

### Snapshot Matchers

| Matcher | Description |
|---------|-------------|
| `toMatchSnapshot()` | Match against stored snapshot |
| `toMatchInlineSnapshot()` | Inline snapshot in test file |
| `toThrowErrorMatchingSnapshot()` | Snapshot error message |

### Mock Matchers

| Matcher | Description |
|---------|-------------|
| `toHaveBeenCalled()` | Mock was called |
| `toHaveBeenCalledTimes(n)` | Mock called exactly n times |
| `toHaveBeenCalledWith(...args)` | Mock called with specific args |

## Mocks

```ts
import { mock, jest } from "bun:test";

const fn = mock(() => 42);   // same as jest.fn(() => 42)
fn();
expect(fn).toHaveBeenCalled();
expect(fn).toHaveBeenCalledTimes(1);
```

`mock` and `jest.fn` are identical — both create mock functions.

## Concurrent Test Execution

- `--concurrent` flag: runs all tests in parallel within each file (unless marked `test.serial`).
- `--max-concurrency` flag: caps parallel tests (default: 20).
- `test.concurrent()`: mark individual tests to run concurrently.
- `test.serial()`: force tests to run sequentially, even when `--concurrent` is enabled.

## Snapshots

```ts
expect(result).toMatchSnapshot();
expect(result).toMatchInlineSnapshot();
expect(() => { throw new Error("oops") }).toThrowErrorMatchingSnapshot();
```

Update snapshots:

```bash
bun test --update-snapshots    # or -u
```

## Coverage

```bash
bun test --coverage
bun test --coverage --coverage-reporter lcov --coverage-dir ./coverage
```

### Coverage Thresholds

```toml
[test]
coverageThreshold = 0.9                          # 90% overall

# or per-metric:
coverageThreshold = { lines = 0.7, functions = 0.8, statements = 0.9 }
```

If the threshold is not met, `bun test` exits with a non-zero code.

## CI/CD Integration

### GitHub Actions

Auto-detected. Emits GitHub Actions annotations with no extra config. Install Bun via `oven-sh/setup-bun@v2`.

### JUnit XML Reports

```bash
bun test --reporter=junit --reporter-outfile=./bun.xml
```

### AI Agent Integration

Set environment variables for quieter output (failure details only, hide passing/skipped/todo):

| Variable | For |
|----------|-----|
| `CLAUDECODE=1` | Claude Code |
| `REPL_ID=1` | Replit |
| `AGENT=1` | Generic AI agent flag |

```bash
CLAUDECODE=1 bun test
```

## Randomization

```bash
bun test --randomize                  # run tests in random order
bun test --seed 123456                # reproduce a specific order (implies --randomize)
```

The seed is displayed in the summary output when `--randomize` is used.

## DOM Testing

Compatible with:

- [HappyDOM](https://github.com/capricorn86/happy-dom)
- [DOM Testing Library](https://testing-library.com/docs/dom-testing-library/intro/)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)

Install as dev dependencies and import in tests.

## References

- <https://bun.sh/docs/test>
- <https://bun.sh/docs/test/writing-tests>
- <https://bun.sh/docs/test/lifecycle>
- <https://bun.sh/docs/test/mocks>
- <https://bun.sh/docs/test/snapshots>
- <https://bun.sh/docs/test/dom>
