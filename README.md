# Image to WebP Converter

A fast, fully parallelized Command Line Interface (CLI) tool designed to recursively discover and convert all `.jpg`, `.jpeg`, and `.png` files within a directory (and its subdirectories) into the highly efficient WebP format.

> **Note:** This tool is primarily designed and optimized for **Windows** (including seamless native integration with Windows Explorer).

Built with Go and beautifully styled using [Lipgloss](https://github.com/charmbracelet/lipgloss).

## ✨ Features

- **Recursive Search:** Automatically processes all supported images in the executable's directory and all of its subdirectories.
- **Multi-Core Performance:** Leverages all available CPU cores via Go routines for lightning-fast, parallelized generation.
- **Smart Skipping:** Saves time by skipping files that have already been converted during subsequent runs.
- **Preserved Extensions:** The original file extension is kept in the format `filename.png.webp` or `filename.jpg.webp` to quickly identify the original source file type.
- **Double-Click Support:** When launched by double-clicking in Windows Explorer, the program automatically pauses at the end so you can review the results before the window closes.
- **Beautiful UI:** Features an elegant, color-coded, and clean command-line presentation.

## 🚀 Installation

### Pre-compiled Binary (Windows)
Simply download `img2webp.exe` from the Releases tab and place it into the folder containing your images. It's completely standalone—no extra dependencies or libraries (like `libwebp`) are required!

### Build from Source
Because this project utilizes the CGO-based `github.com/chai2010/webp` encoder, you will need **Go** and a **C-Compiler** (e.g., GCC via MinGW-w64 on Windows).

```bash
git clone <your-repo-link>
cd img2webp
go mod tidy
make
# or manually: go build -o bin/img2webp.exe ./cmd/img2webp
```

## 🛠️ Usage

The tool is designed for plug-and-play simplicity.

### Method 1: The Easy Way (Double-Click)
Copy `img2webp.exe` into the topmost directory where your original images are located (or any parent directory of those images) and **double-click** it. The tool securely scans for all images, converts them using the default quality setting (80%), and awaits a keystroke before exiting so you can read the output.

### Method 2: Command Line (PowerShell / CMD)
For more granular control, open your terminal in the target directory and run the executable with flags.

```powershell
# Default conversion (80% quality)
.\img2webp.exe

# Convert with 50% quality for smaller file sizes
.\img2webp.exe -q 50

# Force overwrite of existing conversions
.\img2webp.exe -o
```

#### Available Flags
| Flag                 | Description                                                                 |
| -------------------- | --------------------------------------------------------------------------- |
| `-h`, `--help`       | Displays the help message.                                                 |
| `-q`, `--quality`    | Set WebP export quality (float between `0` and `100`). Default is `80.0`. |
| `-o`, `--overwrite`  | Overwrite existing WebP copies instead of skipping them.                    |

## 📦 Built With

- [chai2010/webp](https://github.com/chai2010/webp) - Reliable CGO-binding for the official libwebp C-Library.
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - For styling the brilliant terminal interface.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
