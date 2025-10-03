# Release v1.0.2

- Added toptions to save certificates from connection and display content of
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

Sample content of binary archives is shown below. After extracting, copy the
binary somewhere in your PATH.

```bash
➤ unzip -l ssl-tls-verify-v1.0.2-windows-amd64.d.zip
Archive:  bin/ssl-tls-verify-v1.0.2-windows-amd64.d.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1072  10-03-2025 16:45   ssl-tls-verify-v1.0.2-windows-amd64.d/LICENSE
    44012  10-03-2025 16:45   ssl-tls-verify-v1.0.2-windows-amd64.d/README.md
      903  10-03-2025 16:45   ssl-tls-verify-v1.0.2-windows-amd64.d/platforms.txt
  4931072  10-03-2025 16:45   ssl-tls-verify-v1.0.2-windows-amd64.d/ssl-tls-verify-v1.0.2-windows-amd64.exe
---------                     -------
  4977059                     4 files
```

```bash
➤ tar -tvf bin/ssl-tls-verify-v1.0.2-linux-amd64.d.tar.gz
-rw-r--r--  0 muquit staff    1072 Oct  3 16:45 ssl-tls-verify-v1.0.2-linux-amd64.d/LICENSE
-r--r--r--  0 muquit staff   44012 Oct  3 16:45 ssl-tls-verify-v1.0.2-linux-amd64.d/README.md
-rw-r--r--  0 muquit staff     903 Oct  3 16:45 ssl-tls-verify-v1.0.2-linux-amd64.d/platforms.txt
-rwxr-xr-x  0 muquit staff 4763832 Oct  3 16:45 ssl-tls-verify-v1.0.2-linux-amd64.d/ssl-tls-verify-v1.0.2-linux-amd64
```



Cross-compiled and released with [go-xbuild-go](https://github.com/muquit/go-xbuild-go)
