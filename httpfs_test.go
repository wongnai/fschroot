package fschroot_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wongnai/fschroot"
)

func assertFileContent(t *testing.T, fp io.ReadCloser, content string) {
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, content, string(data))
}

func TestFsChroot(t *testing.T) {
	assert := assert.New(t)
	_, filename, _, _ := runtime.Caller(0)
	realFs := http.Dir(filepath.Join(filepath.Dir(filename), "test_assets"))

	fp, err := realFs.Open("inaccessible")
	assert.Nil(err, "Sanity check")

	chroot := fschroot.New("/root/", realFs)

	for i := 0; i <= 10; i++ {
		_, err = chroot.Open(strings.Repeat("../", i) + "inaccessible")
		assert.NotNil(err, "File in outer directory must not be accessible")
	}

	for _, name := range []string{"outer.txt", "/outer.txt", "/something/../outer.txt", "./outer.txt"} {
		fp, err = chroot.Open(name)
		assert.Nil(err)
		assertFileContent(t, fp, "outer")
	}

	for _, name := range []string{"subfolder/file.txt", "/subfolder/file.txt", "./subfolder/file.txt", "/subfolder/./file.txt", "/subfolder/../subfolder/file.txt"} {
		fp, err = chroot.Open(name)
		assert.Nil(err)
		assertFileContent(t, fp, "passed")
	}
}
