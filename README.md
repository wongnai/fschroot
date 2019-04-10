# FSChroot

[![GoDoc](https://godoc.org/github.com/wongnai/fschroot?status.svg)](https://godoc.org/github.com/wongnai/fschroot) [![CircleCI](https://circleci.com/gh/wongnai/fschroot.svg?style=svg)](https://circleci.com/gh/wongnai/fschroot)

Filter [http.FileSystem](https://golang.org/pkg/net/http/#FileSystem) to subpath only. Useful if you use a VFS such as [shurcooL/vfsgen](https://github.com/shurcooL/vfsgen).

## Usage

```go
root := http.Dir("/somewhere") // not really useful if you use http.Dir

chroot := fschroot.New("/child/path", root)
handler := http.FileServer(chroot)
```

## Security

If you believe you've found a security issue in this package, please don't open a GitHub issue. Instead, contact us at <security@wongnai.com> per our [responsible disclosure policy](https://www.wongnai.com/security?locale=en).

## License

[MIT License](LICENSE)
