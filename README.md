# Calendar TUI

A beautiful, interactive **terminal calendar generator** written in Go.

Generates a full-year calendar in the terminal with support for **60+ languages/locales**, customizable first day of the week (Monday/Sunday), colored weekends, and output to console or file.

## Features

- Modern TUI interface (built with Bubble Tea + bubbles/table)
- **60+ locales** — Russian, English (US/UK), German, French, Hindi, Hebrew, Armenian, Kazakh, Ukrainian, Polish, Spanish, Italian, and many more
- Automatic locale detection based on system language
- Choose **Monday** or **Sunday** as first day of week (default per country)
- Year selection: 1900–2100
- Output options:
  - Colorful console (weekends in red)
  - Plain text file `calendar.txt`
- Clean, centered layout (when terminal size allows)

## Screenshots
1. Main menu
<img width="629" height="219" alt="Screenshot 2026-02-10 100047" src="https://github.com/user-attachments/assets/dc0cb288-d218-4423-ac32-15c8a9684917" />

<img width="631" height="214" alt="Screenshot 2026-02-10 100244" src="https://github.com/user-attachments/assets/5af755fb-6452-41bb-9aad-40cefa5f6618" />

<img width="627" height="223" alt="Screenshot 2026-02-10 100331" src="https://github.com/user-attachments/assets/26fe23d3-9cf7-4a11-89b8-bc565442eb86" />


2. Generated calendar in console (Poland example)
<img width="630" height="722" alt="Screenshot 2026-02-10 100502" src="https://github.com/user-attachments/assets/d18d461e-0b72-481c-8f9f-94b668fe8cca" />


4. Saved `calendar.txt` preview (Sweden example)
<img width="590" height="685" alt="Screenshot 2026-02-10 100857" src="https://github.com/user-attachments/assets/6eb9cedc-05e0-48c3-9d72-7f53d8baa34b" />

## Installation

### From source (recommended)

```bash
git clone https://github.com/yourusername/calendar-tui.git
cd calendar-tui
go mod tidy
go build -o calendar-tui
```

## Run

```
./calendar-tui          # Linux / macOS
calendar-tui.exe        # Windows
```

## Via go install (if you have Go ≥ 1.18)

```
go install github.com/yourusername/calendar-tui@latest
calendar-tui
```

## Usage

1. Run the program → interactive menu appears
2. Use ← / → arrows to change values
3. Select year (default: current year)
4. Choose first day of week (auto-suggested by country)
5. Select output: Console or File
6. Choose country/language
7. Press Enter on "Go >>>" line
8. Calendar is shown or saved as calendar.txt

Quit: q or Ctrl+C

## Supported Locales (partial list)
### Europe
Russian, English (US/UK/Ireland), German (Germany/Austria/Switzerland), French (France/Switzerland/Monaco), Italian, Spanish, Polish, Ukrainian, Belarusian, Hungarian, Greek, Czech, Danish, Norwegian, Swedish, Finnish, Dutch, Portuguese, Romanian, Serbian, Croatian, Bulgarian, Albanian, Slovenian, Slovak, Icelandic, Maltese, etc.
### Asia & others
Hebrew (Israel), Armenian, Kazakh, Kyrgyz, Tajik, Uzbek, Turkmen, Vietnamese, Turkish, Malay, etc.

### Unsupported locales
Chinese, Japanese, Persian (Iran), Hindi (India), Georgian,

All translations stored in locales/*.json + two built-in (RU + EN).

## CLI Flags & Usage

You can run the program in **non-interactive mode** (direct calendar generation) by passing flags.  
If **no flags** are provided → the interactive TUI menu will launch.

```
bash
calendar-tui [flags]
```
#### Available flags:
| Flag | Description | Default | Example|
| :--: | :--- | :--- | :---: |
| `--year N` | Calendar year (must be between 1900 and 2100) | Current year | `--year 2026` |
| `--file` | Save output to file instead of printing to console | False | `--file` |
| `--output FILE` | Custom output filename (automatically enables file mode) | calendar.txt | `--output mycal-2026.txt` |

## Contributing
Pull requests are very welcome!
Good first contributions:

* New locales (copy any existing JSON and translate)
* Better error messages / input validation

### License
MIT License

### Made with :heart: and the Charm ecosystem
Bubble Tea · Lip Gloss · Bubbles
