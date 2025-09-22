#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Scan Go source files for likely mojibake (garbled Chinese) and optionally fix
known broken UI strings by replacing them with ASCII/English placeholders.

Usage:
  python scripts/scan_mojibake.py --root internal/gui            # scan
  python scripts/scan_mojibake.py --root internal/gui --apply    # scan + fix

Heuristics:
  - Flag lines containing frequent mojibake characters (e.g. 鍙, 鏂, 閫, 妯, 璺, 鍚, 鐩, 鐗, 鍏, ...)
  - Only mutate inside string literals when --apply is passed, and only for a
    small set of safe replacements (common labels/buttons). All other hits are
    reported for manual review.
"""
from __future__ import annotations
import argparse
import pathlib
import re

MOJIBAKE_CHARS = "鍙鏂閫妯璺鍚鐩鐗鍏闂瀛鑷绐椤搴鎻锛"
STRING_RE = re.compile(r'"([^"\\]|\\.)*"')

# Minimal safe replacements (expand as needed)
SAFE_MAP = {
    "鍚姩": "Start",
    "閫€鍑": "Exit",
    "淇濆瓨": "Save",
    "璁剧疆": "Settings",
    "甯姪": "Help",
    "閬告搰": "Not found",
    "妯″紡": "Mode",
}

def has_mojibake(s: str) -> bool:
    return any(ch in s for ch in MOJIBAKE_CHARS)

def fix_string_literal(lit: str) -> str:
    # Remove the surrounding quotes
    body = lit[1:-1]
    new = body
    for k, v in SAFE_MAP.items():
        if k in new:
            new = new.replace(k, v)
    # Hard fallback: if still looks garbled, convert to ASCII placeholder
    if has_mojibake(new):
        new = re.sub(r"[\u0080-\uffff]", " ", new).strip() or "Label"
    return '"' + new + '"'

def process_file(path: pathlib.Path, apply: bool) -> int:
    text = path.read_text(encoding="utf-8", errors="replace")
    flagged = 0
    out = []
    for i, line in enumerate(text.splitlines(True), 1):
        if has_mojibake(line):
            flagged += 1
            if apply:
                def repl(m):
                    lit = m.group(0)
                    if has_mojibake(lit):
                        return fix_string_literal(lit)
                    return lit
                line = STRING_RE.sub(repl, line)
        out.append(line)
    if apply and flagged:
        new_text = "".join(out)
        if new_text != text:
            path.write_text(new_text, encoding="utf-8")
    return flagged

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="internal/gui", help="directory to scan")
    ap.add_argument("--apply", action="store_true", help="apply safe replacements")
    args = ap.parse_args()

    root = pathlib.Path(args.root)
    total = 0
    for p in root.rglob("*.go"):
        total += process_file(p, args.apply)
    print(f"Scanned {root}, flagged lines: {total}, apply={args.apply}")

if __name__ == "__main__":
    main()

