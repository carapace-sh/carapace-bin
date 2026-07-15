# Templates

Pandoc's template system: syntax, variables, context construction, and per-format variable reference.

> **Source of truth**: <https://pandoc.org/MANUAL.html#templates>. For the internal template engine (doctemplates), see [architecture.md](architecture.md).

## Overview

Templates wrap the writer-generated document body in a complete document structure (headers, footers, formatting). They are only used in **standalone** mode (`-s`/`--standalone`), which sets the `writerTemplate` field in `WriterOptions`.

```bash
pandoc -s --template=my-template.html input.md -o output.html
pandoc -s -D html5   # Print the default HTML5 template
```

Default templates are stored in `data/templates/` (e.g., `default.latex`, `default.html5`, `default.revealjs`). The `getDefaultTemplate` function maps writer names to template files. Custom templates are specified with `--template`.

## Template Syntax

Templates use `$...$` or `${...}` delimiters.

### Comments

```text
$-- This is a comment, ignored to end of line
```

### Variable Interpolation

```text
$title$
$foo.bar.baz$    ${foo.bar.baz}
```

If a variable is not set, it renders as empty string.

### Conditionals

```text
$if(variable)$
  Content shown if variable is truthy.
$else$
  Content shown if variable is falsy.
$endif$
```

```text
$if(author)$
$elseif(editor)$
  Content shown if author is not set but editor is.
$endif$
```

### For Loops

```text
$for(author)$
  $author$
$endfor$
```

With separator:

```text
$for(author)$
  $author$
$sep$,
$endfor$
```

The anaphoric keyword `it` refers to the current loop item:

```text
$for(items)$
  $it.name$: $it.value$
$endfor$
```

### Partials

```text
${styles()}
${date:fancy()}
```

Partials can have arguments in parentheses and separators specified:

```text
${months[, ]}
```

### Indentation Control

The `^` directive controls indentation of interpolated content:

```text
^  $body$
```

### Breakable Spaces

```text
$~$...$~$
```

Content between `$~$` pairs has its whitespace preserved as breakable.

### Pipes

Pipes transform variable output:

| Pipe | Effect |
|------|--------|
| `/uppercase` | Convert to uppercase |
| `/lowercase` | Convert to lowercase |
| `/length` | Length of the value |
| `/reverse` | Reverse the value |
| `/first` | First element |
| `/last` | Last element |
| `/rest` | All but first |
| `/allbutlast` | All but last |
| `/chomp` | Remove trailing newlines |
| `/nowrap` | Disable line wrapping |
| `/alpha` | Alphabetic numbering |
| `/roman` | Roman numeral |
| `/left` `N` | Left-align in N columns |
| `/right` `N` | Right-align in N columns |
| `/center` `N` | Center in N columns |
| `/pairs` | Convert a map to key-value pairs |

```text
$for(author)$/uppercase$author$endfor$
```

## Variable Sources

Variables come from three sources, in order of precedence (highest first):

1. **CLI variables** (`-V`/`--variable`, `--variable-json`)
2. **Metadata** (document metadata, `--metadata`, `--metadata-file`)
3. **Default values** (set by the writer)

### Metadata Variables

Common metadata fields available in templates:

| Variable | Description |
|----------|-------------|
| `title` | Document title |
| `author` | Author(s) |
| `date` | Date |
| `subtitle` | Subtitle |
| `abstract` | Abstract |
| `abstract-title` | Title for abstract section |
| `keywords` | Keywords |
| `subject` | Subject |
| `description` | Description |
| `category` | Category |

### Automatically Set Variables

These variables are set by pandoc automatically:

| Variable | Description |
|----------|-------------|
| `body` | The rendered document body |
| `date-meta` | ISO 8601 date from metadata |
| `header-includes` | Contents of `--include-in-header` |
| `include-before` | Contents of `--include-before-body` |
| `include-after` | Contents of `--include-after-body` |
| `meta-json` | JSON serialization of metadata |
| `numbersections` | True if `--number-sections` |
| `sourcefile` | Source file name |
| `outputfile` | Output file name |
| `pdf-engine` | PDF engine being used |
| `curdir` | Current working directory |
| `pandoc-version` | Pandoc version string |
| `toc` | True if TOC is included |
| `toc-title` | Title for the TOC |

## Per-Format Variables

### HTML Variables

| Variable | Description |
|----------|-------------|
| `document-css` | Use pandoc's default CSS |
| `mainfont` | Main font family |
| `fontsize` | Base font size |
| `fontcolor` | Text color |
| `linkcolor` | Link color |
| `monofont` | Monospace font |
| `monobackgroundcolor` | Code background color |
| `linestretch` | Line spacing |
| `maxwidth` | Max content width |
| `backgroundcolor` | Page background color |
| `margin-*` | Page margins (`margin-left`, `margin-right`, etc.) |

### HTML Math Variables

| Variable | Description |
|----------|-------------|
| `classoption` | Class options for math rendering |

### HTML Slides Variables

| Variable | Description |
|----------|-------------|
| `institute` | Institution |
| `revealjs-url` | URL to reveal.js |
| `s5-url` | URL to S5 |
| `slidy-url` | URL to Slidy |
| `slideous-url` | URL to Slideous |
| `title-slide-attributes` | Attributes for title slide |
| `highlightjs-theme` | highlight.js theme |

### Beamer Variables

| Variable | Description |
|----------|-------------|
| `aspectratio` | Slide aspect ratio (e.g., `169`) |
| `beameroption` | Beamer class option |
| `logo` | Logo image |
| `logooptions` | Logo options |
| `navigation` | Navigation controls |
| `section-titles` | Show section title pages |
| `theme` | Beamer theme |
| `colortheme` | Beamer color theme |
| `fonttheme` | Beamer font theme |
| `innertheme` | Beamer inner theme |
| `outertheme` | Beamer outer theme |
| `titlegraphic` | Title page graphic |
| `shorttitle` | Short title for footer |
| `shortauthor` | Short author for footer |

### LaTeX Variables

**Layout:**

| Variable | Description |
|----------|-------------|
| `block-headings` | Use block-style headings |
| `classoption` | Document class option |
| `documentclass` | Document class (default `article`) |
| `geometry` | Page geometry options |
| `hyperrefoptions` | hyperref options |
| `indent` | Use paragraph indentation |
| `linestretch` | Line spacing |
| `margin-*` | Page margins |
| `pagestyle` | Page style |
| `papersize` | Paper size (e.g., `a4`, `letter`) |
| `secnumdepth` | Section numbering depth |
| `csquotes` | Use csquotes package |

**Fonts:**

| Variable | Description |
|----------|-------------|
| `fontenc` | Font encoding |
| `fontfamily` | Font family |
| `fontfamilyoptions` | Font family options |
| `fontsize` | Font size |
| `mainfont` | Main font |
| `sansfont` | Sans-serif font |
| `monofont` | Monospace font |
| `mathfont` | Math font |
| `CJKmainfont` | CJK main font |
| `mainfontoptions` | Main font options |
| `mainfontfallback` | Main font fallbacks |
| `babelfonts` | Babel fonts |

**Links:**

| Variable | Description |
|----------|-------------|
| `colorlinks` | Color links |
| `boxlinks` | Box links |
| `linkcolor` | Internal link color |
| `filecolor` | File link color |
| `citecolor` | Citation link color |
| `urlcolor` | URL link color |
| `toccolor` | TOC link color |
| `links-as-notes` | Render links as footnotes |
| `urlstyle` | URL style |

**Front matter:**

| Variable | Description |
|----------|-------------|
| `lof` | List of figures |
| `lot` | List of tables |
| `thanks` | Thanks/acknowledgment |
| `toc` | Table of contents |
| `toc-depth` | TOC depth |

**BibLaTeX:**

| Variable | Description |
|----------|-------------|
| `biblatexoptions` | biblatex options |
| `biblio-style` | Bibliography style |
| `biblio-title` | Bibliography title |
| `bibliography` | Bibliography file |
| `natbiboptions` | natbib options |

**Other:**

| Variable | Description |
|----------|-------------|
| `pdf-trailer-id` | PDF trailer ID |
| `pdfstandard` | PDF standard (PDF/A, PDF/X, PDF/UA) |
| `beamerarticle` | Beamer article mode |
| `handout` | Beamer handout mode |

### ConTeXt Variables

| Variable | Description |
|----------|-------------|
| `fontsize` | Font size |
| `headertext` | Header text |
| `footertext` | Footer text |
| `indenting` | Indentation settings |
| `interlinespace` | Line spacing |
| `layout` | Page layout |
| `linkcolor` | Link color |
| `contrastcolor` | Contrast color |
| `linkstyle` | Link style |
| `mainfont` | Main font |
| `sansfont` | Sans-serif font |
| `monofont` | Monospace font |
| `mathfont` | Math font |
| `margin-*` | Page margins |
| `pagenumbering` | Page numbering style |
| `papersize` | Paper size |
| `pdfa` | PDF/A compliance |
| `pdfaiccprofile` | ICC profile |
| `pdfaintent` | PDF/A intent |
| `toc` | Table of contents |
| `urlstyle` | URL style |
| `whitespace` | Whitespace handling |
| `includesource` | Include source |

### Typst Variables

| Variable | Description |
|----------|-------------|
| `template` | Typst template |
| `margin` | Page margins |
| `papersize` | Paper size |
| `mainfont` | Main font |
| `fontsize` | Font size |
| `section-numbering` | Section numbering format |
| `page-numbering` | Page numbering format |
| `columns` | Number of columns |
| `thanks` | Thanks/acknowledgment |
| `mathfont` | Math font |
| `codefont` | Code font |
| `linestretch` | Line spacing |
| `linkcolor` | Link color |
| `filecolor` | File link color |
| `citecolor` | Citation link color |

### Man Page Variables

| Variable | Description |
|----------|-------------|
| `adjusting` | Text adjustment |
| `footer` | Footer text |
| `header` | Header text |
| `section` | Man section number |

### wkhtmltopdf Variables

| Variable | Description |
|----------|-------------|
| `footer-html` | Footer HTML file |
| `header-html` | Header HTML file |
| `margin-*` | Page margins |
| `papersize` | Paper size |

### Texinfo Variables

| Variable | Description |
|----------|-------------|
| `version` | Version string |
| `filename` | Info file name |

### ms Variables

| Variable | Description |
|----------|-------------|
| `fontfamily` | Font family |
| `indent` | Paragraph indent |
| `lineheight` | Line height |
| `pointsize` | Font point size |

## Context Construction

Internally, the template context is constructed by `metaToContext` / `metaToContext'` in `Text.Pandoc.Writers.Shared`. Metadata (`Meta`) is converted to a template `Context` (a key-value map), then CLI variables are overlaid on top. For details, see [architecture.md](architecture.md).

## Related References

- [cli-options.md](cli-options.md) — `--standalone`, `--template`, `-V`, `-D`
- [defaults.md](defaults.md) — Specifying variables in defaults files
- [architecture.md](architecture.md) — Internal template engine and context construction
