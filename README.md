# stdin-spinner

`stdin-spinner` is a terminal spinner indicating if data is coming in on stdin.

For when you need to show the user that something is happening, but don't want to spam them with raw logs.

![no-spin](doc/nospin.gif) ![spin](doc/spin.gif)

## Contents

- [Contents](#contents)
- [Get it](#get-it)
- [Usage](#usage)

## Get it

Using go get:

```bash
go get -u github.com/sgreben/stdin-spinner
```

Or [download the binary](https://github.com/sgreben/stdin-spinner/releases/latest) from the releases page.

```bash
# Linux
curl -LO https://github.com/sgreben/stdin-spinner/releases/download/1.0.0/stdin-spinner_1.0.0_linux_x86_64.zip
unzip stdin-spinner_1.0.0_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/stdin-spinner/releases/download/1.0.0/stdin-spinner_1.0.0_osx_x86_64.zip
unzip stdin-spinner_1.0.0_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/stdin-spinner/releases/download/1.0.0/stdin-spinner_1.0.0_windows_x86_64.zip
unzip stdin-spinner_1.0.0_windows_x86_64.zip
```

## Usage

```text
Usage of stdin-spinner:
  -a	(alias for -tee-append)
  -frames string
    	spinner animation frames (default "░▒▓█▓▒░")
  -r duration
    	(alias for -refresh) (default 100ms)
  -refresh duration
    	 (default 100ms)
  -tee string
    	copy stdin to this file
  -tee-append
    	append to the tee file (if specified)
  -w int
    	(alias for -width) (default 12)
  -width int
    	spinner animation width (default 12)
```
