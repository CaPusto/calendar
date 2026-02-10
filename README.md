# Calendar TUI

A beautiful, interactive **terminal calendar generator** written in Go.

Generates a full-year calendar in the terminal with support for **60+ languages/locales**, customizable first day of the week (Monday/Sunday), colored weekends, and output to console or file.

## Features

- Modern TUI interface (built with Bubble Tea + bubbles/table)
- **60+ locales** — Russian, English (US/UK), German, French, Chinese, Japanese, Persian, Hindi, Hebrew, Armenian, Georgian, Kazakh, Ukrainian, Polish, Spanish, Italian, and many more
- Automatic locale detection based on system language
- Choose **Monday** or **Sunday** as first day of week (default per country)
- Year selection: 1900–2100
- Output options:
  - Colorful console (weekends in red)
  - Plain text file `calendar.txt`
- Clean, centered layout (when terminal size allows)

## Screenshots

(добавьте 2–4 скриншота после первого коммита)

1. Main menu (Russian locale)  
2. Generated calendar in console (Chinese example)  
3. Saved `calendar.txt` preview

## Installation

### From source (recommended)

```bash
git clone https://github.com/yourusername/calendar-tui.git
cd calendar-tui
go mod tidy
go build -o calendar-tui
