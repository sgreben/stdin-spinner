# ${APP}

`${APP}` is a terminal spinner indicating if data is coming in on stdin.

For when you need to show the user that something is happening, but don't want to spam them with raw logs.

![no-spin](doc/nospin.gif) ![spin](doc/spin.gif)

- [Get it](#get-it)
- [Usage](#usage)

## Get it

Using go get:

```bash
go get -u github.com/sgreben/${APP}
```

Or [download the binary](https://github.com/sgreben/${APP}/releases/latest) from the releases page.

```bash
# Linux
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_linux_x86_64.zip
unzip ${APP}_${VERSION}_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_osx_x86_64.zip
unzip ${APP}_${VERSION}_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_windows_x86_64.zip
unzip ${APP}_${VERSION}_windows_x86_64.zip
```

## Usage

```text
$USAGE
```
