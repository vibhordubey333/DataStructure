#!/usr/bin/env python3
"""
Convert all Blind 75 Markdown files to standalone, colorful HTML with syntax-highlighted code.

Install deps first:
  pip install -r scripts/requirements-md2html.txt
  # or: pip install markdown pygments
"""

import sys
from pathlib import Path

try:
    import markdown
except ImportError:
    print("Missing 'markdown'. Install with: pip install markdown pygments", file=sys.stderr)
    sys.exit(1)
try:
    from pygments.formatters import HtmlFormatter
    PYGMENTS_AVAILABLE = True
except ImportError:
    PYGMENTS_AVAILABLE = False

BLIND75_DIR = Path(__file__).resolve().parent.parent / "blind75"
# Skip README and empty/legacy files
SKIP_NAMES = {"README.md"}

# Pygments style: colorful, readable (options: default, tango, emacs, friendly, monokai, etc.)
PYGMENTS_STYLE = "tango"

HTML_HEAD = """<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{title}</title>
  <style>
    :root {
      --bg: #fafafa;
      --text: #1a1a1a;
      --border: #e0e0e0;
      --accent: #1565c0;
      --code-bg: #f5f5f5;
    }
    * { box-sizing: border-box; }
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
      line-height: 1.6;
      color: var(--text);
      background: var(--bg);
      margin: 0;
      padding: 1rem;
    }
    .container { max-width: 900px; margin: 0 auto; }
    h1 { font-size: 1.75rem; color: var(--accent); border-bottom: 2px solid var(--border); padding-bottom: 0.5rem; }
    h2 { font-size: 1.35rem; margin-top: 1.5rem; color: #333; }
    h3 { font-size: 1.1rem; margin-top: 1rem; color: #555; }
    p { margin: 0.5rem 0; }
    ul, ol { margin: 0.5rem 0; padding-left: 1.5rem; }
    table { border-collapse: collapse; width: 100%; margin: 1rem 0; font-size: 0.95rem; }
    th, td { border: 1px solid var(--border); padding: 0.5rem 0.75rem; text-align: left; }
    th { background: var(--code-bg); font-weight: 600; }
    hr { border: none; border-top: 1px solid var(--border); margin: 1.5rem 0; }
    pre { margin: 1rem 0; overflow-x: auto; border-radius: 6px; }
    code { font-family: "SF Mono", "Consolas", "Monaco", monospace; font-size: 0.9rem; }
    pre code { padding: 0; background: none; }
    .codehilite { margin: 1rem 0; border-radius: 6px; overflow-x: auto; }
    .codehilite pre { margin: 0; padding: 1rem; }
  </style>
  <style>
  {pygments_css}
  </style>
</head>
<body>
<div class="container">
"""

HTML_TAIL = """
</div>
</body>
</html>
"""


def get_pygments_css() -> str:
    if not PYGMENTS_AVAILABLE:
        return "/* Pygments not installed; code blocks use .code-block below */"
    formatter = HtmlFormatter(style=PYGMENTS_STYLE, cssclass="codehilite")
    return formatter.get_style_defs(".codehilite")


def md_to_html(md_path: Path, title: str, pygments_css: str) -> str:
    raw = md_path.read_text(encoding="utf-8")
    extensions = ["extra", "tables", "toc"]
    config = {}
    if PYGMENTS_AVAILABLE:
        extensions.append("codehilite")
        config["codehilite"] = {
            "css_class": "codehilite",
            "linenums": False,
            "use_pygments": True,
        }
    md = markdown.Markdown(extensions=extensions, extension_configs=config or None)
    body = md.convert(raw)
    full = (
        HTML_HEAD.replace("{title}", title).replace("{pygments_css}", pygments_css)
        + body
        + HTML_TAIL
    )
    return full


def main():
    blind75 = BLIND75_DIR
    if not blind75.is_dir():
        print(f"Directory not found: {blind75}")
        return 1
    pygments_css = get_pygments_css()
    count = 0
    for md_path in sorted(blind75.glob("*.md")):
        if md_path.name in SKIP_NAMES:
            continue
        if md_path.stat().st_size == 0:
            continue
        # e.g. TwoSum-LC-1-E -> "TwoSum LC 1 (E)"
        title = md_path.stem.replace("-", " ")
        for d in (" E", " M", " H"):
            if title.endswith(d):
                title = title[: -len(d)] + f" ({d.strip()})"
                break
        html_path = md_path.with_suffix(".html")
        try:
            html = md_to_html(md_path, title, pygments_css)
            html_path.write_text(html, encoding="utf-8")
            count += 1
            print(f"  {html_path.name}")
        except Exception as e:
            print(f"  Skip {md_path.name}: {e}")
    print(f"\nDone. Generated {count} HTML files in {blind75}")
    return 0


if __name__ == "__main__":
    exit(main())
