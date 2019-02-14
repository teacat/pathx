package pathx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/d"
	assert.Equal(`a\b\c\d`, Path(path))
	path = "a/b/c/d/"
	assert.Equal(`a\b\c\d`, Path(path))
}

func TestPathDir(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/d"
	assert.Equal(`a\b\c\d\`, PathDir(path))
	path = "a/b/c/d/"
	assert.Equal(`a\b\c\d\`, PathDir(path))
}

func TestJoin(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/"
	path2 := "/d/e/f/"
	path3 := "g/h/i/"
	assert.Equal(`a\b\c\d\e\f\g\h\i`, Join(path, path2, path3))
	path3 = "g/h/i"
	assert.Equal(`a\b\c\d\e\f\g\h\i`, Join(path, path2, path3))
}

func TestJoinDir(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/"
	path2 := "/d/e/f/"
	path3 := "g/h/i"
	assert.Equal(`a\b\c\d\e\f\g\h\i\`, JoinDir(path, path2, path3))
	path3 = "g/h/i/"
	assert.Equal(`a\b\c\d\e\f\g\h\i\`, JoinDir(path, path2, path3))
}
func TestName(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c.exe"
	assert.Equal("c", Name(path))
	path = "a/b/c"
	assert.Equal("c", Name(path))
}
func TestExt(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c.exe"
	assert.Equal("exe", Ext(path))
	path = "a/b/c"
	assert.Equal("", Ext(path))
	path = "a/b/c.zip.gz"
	assert.Equal("gz", Ext(path))
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/d/e"
	assert.True(Contains(path, `c\d\e`))
}

func TestExecutable(t *testing.T) {
	assert := assert.New(t)
	e, _ := os.Executable()
	assert.Equal(e, Executable())
}

func TestGetwd(t *testing.T) {
	assert := assert.New(t)
	e, _ := os.Getwd()
	assert.Equal(e, Getwd())
}
