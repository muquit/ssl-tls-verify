# Release v1.0.4

Bug fixes and improvements

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

Please do not forget to verify Checksums before installing.

```bash
$ sha256sum -c ssl-tls-verify-v1.0.4-checksums.txt
ssl-tls-verify-v1.0.4-darwin-amd64: OK
ssl-tls-verify-v1.0.4-darwin-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.4-darwin-arm64: OK
ssl-tls-verify-v1.0.4-darwin-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.4-linux-amd64: OK
ssl-tls-verify-v1.0.4-linux-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.4-linux-arm: OK
ssl-tls-verify-v1.0.4-linux-arm.d.tar.gz: OK
ssl-tls-verify-v1.0.4-linux-arm64: OK
ssl-tls-verify-v1.0.4-linux-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.4-windows-386.exe: OK
ssl-tls-verify-v1.0.4-windows-386.d.zip: OK
ssl-tls-verify-v1.0.4-windows-amd64.exe: OK
ssl-tls-verify-v1.0.4-windows-amd64.d.zip: OK
ssl-tls-verify-v1.0.4-raspberry-pi: OK
ssl-tls-verify-v1.0.4-raspberry-pi.d.tar.gz: OK
ssl-tls-verify-v1.0.4-raspberry-pi-jessie: OK
ssl-tls-verify-v1.0.4-raspberry-pi-jessie.d.tar.gz: OK
```

**Note for Windows users:** Microsoft Edge may flag Go binaries as potentially harmful due to false positives in its virus detection software. This is a known issue with Go-compiled executables. Please use Chrome or Firefox to download, or build from source if you prefer.

Sample content of binary archives is shown below. The `tar` command is available in Windows 10 (1803) and later, or you can use the GUI (right-click → Extract All). After extracting, copy/rename the binary somewhere in your PATH.

```bash
➤ tar -tf ssl-tls-verify-v1.0.4-windows-amd64.d.zip
ssl-tls-verify-v1.0.4-windows-amd64.d/
ssl-tls-verify-v1.0.4-windows-amd64.d/LICENSE
ssl-tls-verify-v1.0.4-windows-amd64.d/README.md
ssl-tls-verify-v1.0.4-windows-amd64.d/platforms.txt
ssl-tls-verify-v1.0.4-windows-amd64.d/ssl-tls-verify-v1.0.4-windows-amd64.exe
ssl-tls-verify-v1.0.4-windows-amd64.d/ssl-tls-verify.1
```

```bash
➤ tar -tvf ssl-tls-verify-v1.0.4-linux-amd64.d.tar.gz
drwxr-xr-x  0 muquit staff       0 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/
-rw-r--r--  0 muquit staff    1072 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/LICENSE
-r--r--r--  0 muquit staff   60232 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/README.md
-rw-r--r--  0 muquit staff     903 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/platforms.txt
-rwxr-xr-x  0 muquit staff 5050528 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify-v1.0.4-linux-amd64
-rw-r--r--  0 muquit staff   64783 Oct 31 19:24 ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify.1
```

# Install

## Extract (unzip or untar)

Example:

```bash
➤ tar -xvf ssl-tls-verify-v1.0.4-linux-amd64.d.tar.gz
x ssl-tls-verify-v1.0.4-linux-amd64.d/LICENSE
x ssl-tls-verify-v1.0.4-linux-amd64.d/README.md
x ssl-tls-verify-v1.0.4-linux-amd64.d/platforms.txt
x ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify-v1.0.4-linux-amd64
x ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify.1
```

```bash
➤ sudo /bin/cp ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify-v1.0.4-linux-amd64 \
     /usr/local/bin/ssl-tls-verify
➤ sudo cp ssl-tls-verify-v1.0.4-linux-amd64.d/ssl-tls-verify.1 \
        /usr/local/share/man/man1/
➤ ssl-tls-verify -version
ssl-tls-verify v1.0.4
https://github.com/muquit/ssl-tls-verify
➤ man ssl-tls-verify
```

Cross-compiled and released with [go-xbuild-go](https://github.com/muquit/go-xbuild-go)
