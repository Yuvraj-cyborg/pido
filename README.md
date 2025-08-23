# pido

A powerful and fast command-line tool for compressing your files. `pido` handles images (JPEG, PNG), PDFs, and text files with ease.

## Overview

`pido` is designed to be a simple yet effective file compression utility. Whether you're a developer needing to optimize assets or just someone looking to save disk space, `pido` provides a straightforward way to reduce file sizes without a complicated setup.

## Features

- **Multi-Format Support**: Compress JPEGs, PNGs, PDFs, and TXT files.
- **Adjustable Quality**: Fine-tune the compression level for images and PDFs.
- **Directory Scanning**: Recursively find and compress files in a specified directory.
- **Flexible Output**: Choose a custom output directory for compressed files.
- **Cross-Platform**: Runs on macOS, Linux, and Windows.

## Installation

### macOS & Linux (via curl)

You can install `pido` with a single command. This will download the latest binary from the GitHub releases and move it to a common executable path, making it available system-wide.

```sh
curl -sL https://github.com/Yuvraj-cyborg/pido/releases/latest/download/pido-amd64-darwin -o /usr/local/bin/pido && chmod +x /usr/local/bin/pido
```
_Note: For Linux, replace `pido-amd64-darwin` with `pido-amd64-linux`._

You may need to run the command with `sudo` if you don't have write permissions for `/usr/local/bin`.

### Windows

1.  Download the latest `pido.exe` from the [Releases page](https://github.com/Yuvraj-cyborg/pido/releases).
2.  Place the `pido.exe` file in a directory of your choice (e.g., `C:\Program Files\pido`).
3.  Add that directory to your system's `PATH` environment variable to run `pido` from any terminal.

## Usage

The basic command structure is `pido compress [options] [file paths...]`.

### Examples

**1. Compress all supported files in a directory:**

This will scan the `my_assets` folder and compress all images, PDFs, and text files it finds, using a quality setting of 75.

```sh
pido compress --dir ./my_assets --quality 75
```

**2. Compress specific files:**

You can pass one or more file paths directly as arguments.

```sh
pido compress --quality 80 path/to/image.jpg path/to/document.pdf
```

**3. Specify an output directory:**

Compressed files will be saved with a `-pido` suffix in the specified folder. If the output folder doesn't exist, `pido` will create it.

```sh
pido compress --dir ./input_folder --quality 60 --out ./compressed_files
```

### Flags

-   `--dir string`: Directory to scan for files (optional).
-   `--quality int`: Compression quality percentage (0-100). This is a **required** flag.
-   `--out string`: Optional output folder for compressed files.

## Building from Source

If you prefer to build `pido` from the source, you'll need to have Go installed (version 1.21 or newer).

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Yuvraj-cyborg/pido.git
    cd pido
    ```

2.  **Build the binary:**
    ```sh
    go build -o pido ./cmd/main.go
    ```

3.  **Install the binary (optional):**
    You can move the built binary to a directory in your `PATH` to make it globally accessible.
    ```sh
    # For macOS and Linux
    sudo mv pido /usr/local/bin/
    ```

## License

This project is licensed under the MIT License.