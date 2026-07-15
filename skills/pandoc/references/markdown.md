# Pandoc Markdown

Pandoc's Markdown dialect: syntax, extensions, and Markdown variants. Pandoc Markdown is an extended version of the original Markdown spec, designed for conversion to multiple output formats.

> **Source of truth**: <https://pandoc.org/MANUAL.html#pandocs-markdown>. For the extension system and Markdown variants, see [formats.md](formats.md).

## Philosophy

Pandoc Markdown is designed for documents that may be converted to multiple output formats. It discourages raw HTML in favor of portable syntax, emphasizes readability, and provides extensions for structure (tables, definition lists, footnotes, citations) not found in original Markdown.

## Paragraphs and Line Breaks

One or more lines of text followed by a blank line form a paragraph. A hard line break is created by two or more trailing spaces.

```markdown
This is a paragraph
with a soft line break.

This is another paragraph.
```

**Extensions:**

| Extension | Effect |
|-----------|--------|
| `escaped_line_breaks` | Backslash + newline = hard line break |
| `hard_line_breaks` | All newlines are hard breaks |
| `ignore_line_breaks` | All newlines are soft (ignored) |
| `east_asian_line_breaks` | Newlines between East Asian characters are soft |

## Headings

### ATX Style

```markdown
# Heading 1
## Heading 2
### Heading 3
```

### Setext Style

```markdown
Heading 1
=========

Heading 2
---------
```

**Extensions:**

| Extension | Effect |
|-----------|--------|
| `blank_before_header` | Require blank line before heading |
| `space_in_atx_header` | Require space after `#` |
| `header_attributes` | `{#id .class key=val}` on headings |
| `implicit_header_references` | Headings create reference link definitions |

Class `unnumbered` prevents numbering; `unlisted` excludes from TOC:

```markdown
## Unnumbered Section { .unnumbered }

## Unlisted Section { .unlisted }
```

## Block Quotations

```markdown
> This is a block quote.
> It can span multiple lines.

> Nested
>> blockquote
```

Lazy form (continuation without `>`) is allowed for single-paragraph quotes.

**Extension:** `blank_before_blockquote` requires a blank line before.

## Verbatim (Code) Blocks

### Indented Code Blocks

Four spaces or one tab:

```markdown
    This is a code block.
    Multiple lines.
```

**Extension:** `indented_code_classes` — set default classes with `--indented-code-classes`.

### Fenced Code Blocks

```markdown
~~~
code block
~~~

```python
def hello():
    print("Hello")
```
```

**Extensions:**

| Extension | Effect |
|-----------|--------|
| `fenced_code_blocks` | Enable `~~~` fences |
| `backtick_code_blocks` | Enable `` ``` `` fences |
| `fenced_code_attributes` | `{#id .class key=val}` on fenced blocks |
| `raw_attribute` | `` ```{=format} `` for explicit raw content |

## Line Blocks

```markdown
| First line
| Second line
|     Third line with indent
```

Preserves line breaks and leading spaces. **Extension:** `line_blocks`.

## Lists

### Bullet Lists

```markdown
- item 1
- item 2
  - nested item
```

Markers: `*`, `+`, `-`.

### Ordered Lists

```markdown
1. first
2. second
3. third
```

**Extensions:**

| Extension | Effect |
|-----------|--------|
| `startnum` | Preserve starting number and marker type |
| `fancy_lists` | Letters, roman numerals, `#`, parentheses |

```markdown
i.  first
ii. second

(1) parenthesized
```

### Task Lists

```markdown
- [ ] todo
- [x] done
```

**Extension:** `task_lists`.

### Definition Lists

```markdown
Term
:   definition
:   another definition
```

**Extension:** `definition_lists`.

### Example Lists

```markdown
(@) First example
(@good) Labeled example

Reference to @good.
```

**Extension:** `example_lists`.

### Ending a List

Use an HTML comment to cut off a list:

```markdown
- item 1
- item 2

<!-- -->

Not part of the list.
```

## Horizontal Rules

Three or more `*`, `-`, or `_`:

```markdown
* * *
---
___
```

## Tables

### Pipe Tables

```markdown
| Name  | Age | City     |
|-------|-----|----------|
| Alice | 30  | NYC      |
| Bob   | 25  | Boston   |
```

Alignment via colons:

```markdown
| Left | Center | Right |
|:-----|:------:|------:|
| a    | b      | c     |
```

**Extension:** `pipe_tables`.

### Grid Tables

```markdown
+---------+---------+
| Cell 1  | Cell 2  |
+=========+=========+
| Cell 3  | Cell 4  |
+---------+---------+
```

Supports cell spanning, multi-row headers/footers. **Extension:** `grid_tables`.

### Simple Tables

```markdown
  Name   Age
  ----  ---
  Alice  30
  Bob    25
```

**Extension:** `simple_tables`.

### Multiline Tables

Multi-row headers/rows with blank lines between rows. **Extension:** `multiline_tables`.

### Table Captions

```markdown
Table: Caption text

: Caption text (alternative form)
```

**Extension:** `table_captions`.

## Metadata Blocks

### Pandoc Title Block

```markdown
% Title
% Author One; Author Two
% Date
```

**Extension:** `pandoc_title_block`.

### YAML Metadata Block

```markdown
---
title: My Document
author: Jane Doe
date: 2024-01-15
keywords: [pandoc, markdown]
---
```

**Extension:** `yaml_metadata_block`.

## Backslash Escapes

Any ASCII punctuation character preceded by a backslash is literal.

**Extension:** `all_symbols_escapable` — extends to all Unicode punctuation.

```markdown
This is not a \*bold\* text.
```

Backslash-escaped space = nonbreaking space. Backslash-escaped newline = hard line break.

## Inline Formatting

| Syntax | Result | Extension |
|--------|--------|-----------|
| `*emph*` / `_emph_` | *emphasis* | (default) |
| `**strong**` / `__strong__` | **strong** | (default) |
| `~~strikeout~~` | ~~strikeout~~ | `strikeout` |
| `^superscript^` | superscript | `superscript` |
| `~subscript~` | subscript | `subscript` |
| `` `verbatim` `` | `code` | (default) |
| `[text]{.underline}` | underline | (default) |
| `[text]{.smallcaps}` | smallcaps | (default) |
| `[text]{.mark}` | highlight | `mark` |

**Extension:** `intraword_underscores` — `_` inside words is not emphasis.

### Inline Code Attributes

```markdown
`code`{.python #mycode}
```

**Extension:** `inline_code_attributes`.

## Math

```markdown
Inline math: $E = mc^2$

Display math:

$$
\int_0^\infty e^{-x^2} dx = \frac{\sqrt{\pi}}{2}
$$
```

**Extension:** `tex_math_dollars`. Other math input extensions: `tex_math_gfm`, `tex_math_single_backslash`, `tex_math_double_backslash`.

Math is rendered differently per output format: LaTeX (native), HTML (MathJax/KaTeX/MathML), docx (OMML), etc.

## Raw HTML and TeX

| Extension | Effect |
|-----------|--------|
| `raw_html` | Raw HTML passed through in HTML-based formats |
| `raw_tex` | Raw LaTeX/ConTeXt preserved in LaTeX/ConTeXt output |
| `markdown_in_html_blocks` | Markdown inside HTML block tags is parsed |
| `native_divs` | `<div>` as native Div blocks |
| `native_spans` | `<span>` as native Span inlines |
| `raw_attribute` | `` ```{=format} `` for explicit raw content |

```markdown
```{=html}
<div class="custom">Raw HTML</div>
```

```{=latex}
\textbf{Raw LaTeX}
```
```

## LaTeX Macros

**Extension:** `latex_macros` — parse and apply `\newcommand` definitions:

```markdown
\newcommand{\R}{\mathbb{R}}
\newcommand{\vec}[1]{\boldsymbol{#1}}

Let $\vec{x} \in \R^n$.
```

## Links

| Type | Syntax |
|------|--------|
| Automatic | `<https://example.com>` or `<user@example.com>` |
| Inline | `[text](url "title")` |
| Reference | `[text][label]` then `[label]: url "title"` |
| Shortcut | `[text]` (when `[text]: url` is defined) |
| Internal | `[Section](#identifier)` |

**Extensions:**

| Extension | Effect |
|-----------|--------|
| `shortcut_reference_links` | Omit second brackets |
| `spaced_reference_links` | Allow whitespace between components |
| `autolink_bare_uris` | Bare URLs become links |
| `link_attributes` | Attributes on links/images |

## Images

```markdown
![Alt text](image.png "Title")

![Caption](image.png){width=50% height=200px}
```

**Extension:** `implicit_figures` — image alone in a paragraph becomes a figure with the alt text as caption.

## Divs and Spans

### Fenced Divs

```markdown
::: {.warning}
This is a warning.
:::

::::: {.columns}
::: {.column}
Left
:::
::: {.column}
Right
:::
:::::
```

**Extension:** `fenced_divs`.

### Bracketed Spans

```markdown
[highlighted text]{.highlight}

[text with attributes]{#id .class key=value}
```

**Extension:** `bracketed_spans`.

## Footnotes

### Reference Footnotes

```markdown
Here is a footnote reference[^1].

[^1]: Here is the footnote text.

    It can span multiple paragraphs.
```

### Inline Footnotes

```markdown
Here is an inline footnote^[with text].
```

**Extensions:** `footnotes`, `inline_notes`.

## Citation Syntax

```markdown
Citations use @key syntax [@smith2020].

See [@jones2021, p. 42; @lee2019, chap. 3].

As @smith2020 argues...

[-@smith2020] shows only the date.
```

**Extension:** `citations` (enabled by default in pandoc Markdown).

For full citation processing details, see [citations.md](citations.md).

## Alerts

GitHub-style alerts (extension: `alerts`):

```markdown
::: warning
Warning text
:::

> [!NOTE]
> Note text
```

## Emoji

**Extension:** `emoji` — `:smile:` becomes the corresponding emoji character.

## Wikilinks

| Extension | Syntax | Link text |
|-----------|--------|-----------|
| `wikilinks_title_after_pipe` | `[[Target|Display Text]]` | Display Text |
| `wikilinks_title_before_pipe` | `[[Display Text|Target]]` | Display Text |

## Non-Default Extensions

Extensions not enabled by default in `markdown`:

| Extension | Effect |
|-----------|--------|
| `rebase_relative_paths` | Rewrite relative paths based on input file location |
| `mark` | `[text]{.mark}` highlighting |
| `attributes` | Enable attribute syntax |
| `old_dashes` | Old-style dash rules |
| `angle_brackets_escapable` | `<` is escapable |
| `lists_without_preceding_blankline` | Lists don't need blank line before |
| `four_space_rule` | Original Markdown indented code rule |
| `hard_line_breaks` | All newlines are hard breaks |
| `ignore_line_breaks` | All newlines are soft |
| `east_asian_line_breaks` | East Asian line break handling |
| `emoji` | `:emoji:` syntax |
| `tex_math_gfm` | GitHub math syntax |
| `tex_math_single_backslash` | `\(...\)`, `\[...\]` |
| `tex_math_double_backslash` | `\\(...\\)`, `\\[...\\]` |
| `markdown_attribute` | `markdown=` attribute on HTML |
| `mmd_title_block` | MultiMarkdown title block |
| `abbreviations` | Abbreviations file |
| `alerts` | GitHub-style alerts |
| `autolink_bare_uris` | Bare URLs become links |
| `mmd_link_attributes` | MultiMarkdown link attributes |
| `mmd_header_identifiers` | MultiMarkdown header IDs |
| `gutenberg` | Project Gutenberg style |
| `sourcepos` | Source position attributes |
| `short_subsuperscripts` | No braces needed for sub/sup |
| `wikilinks_title_after_pipe` | `[[Target\|Text]]` wikilinks |
| `wikilinks_title_before_pipe` | `[[Text\|Target]]` wikilinks |

## Related References

- [formats.md](formats.md) — Extension system, Markdown variants, `--list-extensions`
- [citations.md](citations.md) — Citation processing, CSL, bibliography
- [cli-options.md](cli-options.md) — Reader options, `--shift-heading-level-by`, etc.
- [architecture.md](architecture.md) — Markdown reader modules
