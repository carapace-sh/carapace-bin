# PDF Generation

Generating PDFs with pandoc: engines, LaTeX setup, alternative engines, reproducible builds, and accessible PDF standards.

> **Source of truth**: <https://pandoc.org/MANUAL.html#creating-a-pdf>. For PDF-specific CLI flags, see [cli-options.md](cli-options.md).

## Overview

Pandoc does not render PDFs directly. Instead, it produces an intermediate format (LaTeX, ConTeXt, HTML, Typst, or roff ms) and invokes an external PDF engine to render the final PDF.

```bash
pandoc input.md -o output.pdf                # Defaults to LaTeX
pandoc input.md -o output.pdf --pdf-engine=xelatex
pandoc input.md -o output.pdf --pdf-engine=typst
pandoc input.md -o output.pdf --pdf-engine=weasyprint
```

## PDF Engines

| Engine | Intermediate Format | Flag |
|--------|---------------------|------|
| pdflatex | LaTeX | `--pdf-engine=pdflatex` |
| lualatex | LaTeX | `--pdf-engine=lualatex` |
| xelatex | LaTeX | `--pdf-engine=xelatex` |
| latexmk | LaTeX | `--pdf-engine=latexmk` |
| tectonic | LaTeX | `--pdf-engine=tectonic` |
| context | ConTeXt | `--pdf-engine=context` |
| weasyprint | HTML | `--pdf-engine=weasyprint` |
| pagedjs-cli | HTML | `--pdf-engine=pagedjs-cli` |
| prince | HTML | `--pdf-engine=prince` |
| wkhtmltopdf | HTML | `--pdf-engine=wkhtmltopdf` |
| groff | roff ms | `--pdf-engine=groff` |
| pdfroff | roff ms | `--pdf-engine=pdfroff` |
| typst | Typst | `--pdf-engine=typst` |

If `--pdf-engine` is not specified, pandoc defaults to LaTeX and tries `pdflatex`, `lualatex`, or `xelatex` in order.

### Passing Options to PDF Engines

Use `--pdf-engine-opt` to pass command-line arguments to the PDF engine:

```bash
pandoc input.md -o output.pdf --pdf-engine=xelatex --pdf-engine-opt=-interaction=nonstopmode
```

Multiple `--pdf-engine-opt` flags can be used:

```bash
pandoc input.md -o output.pdf \
  --pdf-engine=typst \
  --pdf-engine-opt=--pdf-standard=a-2b
```

## LaTeX PDF Generation

When using a LaTeX engine, pandoc generates a LaTeX document and compiles it. The following LaTeX packages are commonly required:

`amsfonts`, `amsmath`, `lm`, `unicode-math`, `iftex`, `listings`, `fancyvrb`, `longtable`, `booktabs`, `graphicx`, `bookmark`, `xcolor`, `soul`, `geometry`, `setspace`, `babel`, `hyperref`, `polyglossia` (for XeLaTeX/LuaLaTeX with non-Latin scripts).

### Choosing a LaTeX Engine

| Engine | Best for |
|--------|----------|
| `pdflatex` | Standard Latin-script documents, fastest |
| `xelatex` | System fonts, CJK, full Unicode |
| `lualatex` | System fonts, CJK, Lua scripting, OTF features |
| `latexmk` | Automated multi-pass compilation (handles references) |
| `tectonic` | Automatic dependency fetching, no TeX Live needed |

### LaTeX Variables

PDF output is heavily controlled by template variables. See [templates.md](templates.md) for the full LaTeX variable reference. Key variables:

```yaml
---
variables:
  documentclass: article
  papersize: a4
  geometry:
    - margin=1in
  mainfont: "Source Sans Pro"
  fontsize: 12pt
  colorlinks: true
  linkcolor: "blue"
---
```

## HTML-to-PDF Generation

HTML-based engines (weasyprint, wkhtmltopdf, prince, pagedjs-cli) render HTML/CSS to PDF. This is useful when:

- The document uses complex CSS styling
- You need CSS paged media features
- You want HTML-quality typography without LaTeX

```bash
pandoc input.md -o output.pdf --pdf-engine=weasyprint --css=style.css
```

### wkhtmltopdf Variables

| Variable | Description |
|----------|-------------|
| `footer-html` | Footer HTML file |
| `header-html` | Header HTML file |
| `margin-*` | Page margins |
| `papersize` | Paper size |

## Typst PDF Generation

Typst is a modern typesetting system with a built-in PDF engine:

```bash
pandoc input.md -o output.pdf --pdf-engine=typst
```

Typst variables (see [templates.md](templates.md) for full list):

```yaml
---
variables:
  papersize: a4
  margin: [2cm, 3cm]
  mainfont: "Linux Libertine"
  fontsize: 12pt
  section-numbering: "1.1"
  page-numbering: "1"
  columns: 2
---
```

## ConTeXt PDF Generation

ConTeXt is an alternative TeX macro package:

```bash
pandoc input.md -o output.pdf --pdf-engine=context
```

## Reproducible Builds

To produce reproducible PDF output, set the `SOURCE_DATE_EPOCH` environment variable to a Unix timestamp. Pandoc uses this value instead of the current time for any time-stamped output.

```bash
SOURCE_DATE_EPOCH=1700000000 pandoc input.md -o output.pdf
```

For LaTeX output, the `pdf-trailer-id` metadata field sets the PDF trailer ID for reproducibility. For EPUB, set the `identifier` metadata field.

## Accessible PDF and PDF Standards

### PDF/A, PDF/X, PDF/UA (LaTeX)

With LuaLaTeX (TeX Live 2025+), use the `pdfstandard` variable:

```yaml
---
variables:
  pdfstandard: PDF/A-2b
---
```

Supported standards: `PDF/A-1b`, `PDF/A-2b`, `PDF/A-3b`, `PDF/X-4`, `PDF/UA-1`.

### Tagged PDF (ConTeXt)

Enable the `tagging` extension:

```bash
pandoc input.md -o output.pdf --pdf-engine=context -f markdown+tagging
```

### Accessible PDF (WeasyPrint)

```bash
pandoc input.md -o output.pdf --pdf-engine=weasyprint \
  --pdf-engine-opt=--pdf-variant=pdf/ua-1
```

### Tagged PDF (Prince)

```bash
pandoc input.md -o output.pdf --pdf-engine=prince \
  --pdf-engine-opt=--tagged-pdf
```

### PDF Standard (Typst)

```bash
pandoc input.md -o output.pdf --pdf-engine=typst \
  --pdf-engine-opt=--pdf-standard=a-2b
```

## Related References

- [cli-options.md](cli-options.md) — `--pdf-engine`, `--pdf-engine-opt` flags
- [templates.md](templates.md) — LaTeX, ConTeXt, Typst, wkhtmltopdf variables
- [defaults.md](defaults.md) — Specifying PDF options in defaults files
