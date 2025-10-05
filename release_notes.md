# Release v1.0.3

Bug fixes and improvements

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

Please do not forget to verify Checksums before installing.

```bash
$ sha256sum -c ssl-tls-verify-v1.0.3-checksums.txt
ssl-tls-verify-v1.0.3-darwin-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.3-darwin-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.3-linux-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.3-linux-arm.d.tar.gz: OK
ssl-tls-verify-v1.0.3-linux-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.3-windows-386.d.zip: OK
ssl-tls-verify-v1.0.3-windows-amd64.d.zip: OK
ssl-tls-verify-v1.0.3-raspberry-pi.d.tar.gz: OK
ssl-tls-verify-v1.0.3-raspberry-pi-jessie.d.tar.gz: OK
```

**Note for Windows users:** Microsoft Edge may flag Go binaries as potentially harmful due to false positives in its virus detection software. This is a known issue with Go-compiled executables. Please use Chrome or Firefox to download, or build from source if you prefer.

Sample content of binary archives is shown below. The `tar` command is available in Windows 10 (1803) and later, or you can use the GUI (right-click → Extract All). After extracting, copy/rename the binary somewhere in your PATH.

```bash
➤ tar -tf ssl-tls-verify-v1.0.3-windows-amd64.d.zip
ssl-tls-verify-v1.0.3-windows-amd64.d/LICENSE
ssl-tls-verify-v1.0.3-windows-amd64.d/README.md
ssl-tls-verify-v1.0.3-windows-amd64.d/platforms.txt
ssl-tls-verify-v1.0.3-windows-amd64.d/ssl-tls-verify-v1.0.3-windows-amd64.exe
ssl-tls-verify-v1.0.3-windows-amd64.d/ssl-tls-verify.1
```

```bash
➤ tar -tvf ssl-tls-verify-v1.0.3-linux-amd64.d.tar.gz
-rw-r--r--  0 muquit staff    1072 Oct  5 14:30 ssl-tls-verify-v1.0.3-linux-amd64.d/LICENSE
-r--r--r--  0 muquit staff   39614 Oct  5 14:30 ssl-tls-verify-v1.0.3-linux-amd64.d/README.md
-rw-r--r--  0 muquit staff     903 Oct  5 14:30 ssl-tls-verify-v1.0.3-linux-amd64.d/platforms.txt
-rwxr-xr-x  0 muquit staff 4772024 Oct  5 14:30 ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify-v1.0.3-linux-amd64
-rw-r--r--  0 muquit staff   46778 Oct  5 14:30 ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify.1
```

# Install

## Extract (unzip or untar)

Example:

```bash
➤ tar -xvf ssl-tls-verify-v1.0.3-linux-amd64.d.tar.gz
x ssl-tls-verify-v1.0.3-linux-amd64.d/LICENSE
x ssl-tls-verify-v1.0.3-linux-amd64.d/README.md
x ssl-tls-verify-v1.0.3-linux-amd64.d/platforms.txt
x ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify-v1.0.3-linux-amd64
x ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify.1
```

```bash
➤ sudo /bin/cp ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify-v1.0.3-linux-amd64 \
     /usr/local/bin/ssl-tls-verify
➤ sudo cp ssl-tls-verify-v1.0.3-linux-amd64.d/ssl-tls-verify.1 \
        /usr/local/share/man/man1/
➤ ssl-tls-verify -version
ssl-tls-verify v1.0.3
https://github.com/muquit/ssl-tls-verify
➤ man ssl-tls-verify
```

Cross-compiled and released with [go-xbuild-go](https://github.com/muquit/go-xbuild-go)
