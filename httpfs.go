package fschroot

import (
	"net/http"
	"path"
)

// FsChroot wrap a http.FileSystem that only allow access to given subdirectory only
type FsChroot struct {
	Root string

	child http.FileSystem
}

// New create a FsChroot. root must begins with / otherwise this function will panic
func New(root string, child http.FileSystem) FsChroot {
	if root[0] != '/' {
		panic("root must starts with /")
	}
	if root[len(root)-1] == '/' {
		root = root[:len(root)-1]
	}
	return FsChroot{
		Root:  root,
		child: child,
	}
}

// Open implements http.FileSystem.Open
func (fs FsChroot) Open(name string) (http.File, error) {
	// path.Clean does not sanitize ../ if it appears as prefix
	if name[0] != '/' {
		name = "/" + name
	}

	cleanedPath := path.Clean(name)
	return fs.child.Open(fs.Root + cleanedPath)
}
