# Release v1.0.2

- Added options to save certificates from connection and display content of
certificate files.

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

Please do not forget to verify Checksums before installing.

```bash
$ sha256sum -c ssl-tls-verify-v1.0.2-checksums.txt
ssl-tls-verify-v1.0.2-darwin-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.2-darwin-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.2-linux-amd64.d.tar.gz: OK
ssl-tls-verify-v1.0.2-linux-arm.d.tar.gz: OK
ssl-tls-verify-v1.0.2-linux-arm64.d.tar.gz: OK
ssl-tls-verify-v1.0.2-windows-386.d.zip: OK
ssl-tls-verify-v1.0.2-windows-amd64.d.zip: OK
ssl-tls-verify-v1.0.2-raspberry-pi.d.tar.gz: OK
ssl-tls-verify-v1.0.2-raspberry-pi-jessie.d.tar.gz: OK
```

Sample content of binary archives is shown below. After extracting, copy/rename the
binary somewhere in your PATH.

```bash
➤ unzip -l ssl-tls-verify-v1.0.2-windows-amd64.d.zip
Archive:  ssl-tls-verify-v1.0.2-windows-amd64.d.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1072  10-03-2025 19:42   ssl-tls-verify-v1.0.2-windows-amd64.d/LICENSE
    43936  10-03-2025 19:42   ssl-tls-verify-v1.0.2-windows-amd64.d/README.md
      903  10-03-2025 19:42   ssl-tls-verify-v1.0.2-windows-amd64.d/platforms.txt
  4931072  10-03-2025 19:42   ssl-tls-verify-v1.0.2-windows-amd64.d/ssl-tls-verify-v1.0.2-windows-amd64.exe
    46650  10-03-2025 19:42   ssl-tls-verify-v1.0.2-windows-amd64.d/ssl-tls-verify.1
---------                     -------
  5023633                     5 files
```

```bash
➤ tar -tvf ssl-tls-verify-v1.0.2-linux-amd64.d.tar.gz
-rw-r--r--  0 muquit staff    1072 Oct  3 19:42 ssl-tls-verify-v1.0.2-linux-amd64.d/LICENSE
-rw-r--r--  0 muquit staff   43936 Oct  3 19:42 ssl-tls-verify-v1.0.2-linux-amd64.d/README.md
-rw-r--r--  0 muquit staff     903 Oct  3 19:42 ssl-tls-verify-v1.0.2-linux-amd64.d/platforms.txt
-rwxr-xr-x  0 muquit staff 4763832 Oct  3 19:42 ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify-v1.0.2-linux-amd64
-rw-r--r--  0 muquit staff   46650 Oct  3 19:42 ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify.1
```
# Install
## Extract (unzip or untar)
Example:
```bash
➤ tar -xvf ssl-tls-verify-v1.0.2-linux-amd64.d.tar.gz
x ssl-tls-verify-v1.0.2-linux-amd64.d/LICENSE
x ssl-tls-verify-v1.0.2-linux-amd64.d/README.md
x ssl-tls-verify-v1.0.2-linux-amd64.d/platforms.txt
x ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify-v1.0.2-linux-amd64
x ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify.1
```
```bash
➤ sudo /bin/cp ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify-v1.0.2-linux-amd64 \
     /usr/local/bin/ssl-tls-verify

➤ sudo cp cp ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify.1 \
        /usr/local/share/man/man1/

➤ ssl-tls-verify -version
ssl-tls-verify v1.0.2
https://github.com/muquit/ssl-tls-verify

➤ man ssl-tls-verify

```





Cross-compiled and released with [go-xbuild-go](https://github.com/muquit/go-xbuild-go)
