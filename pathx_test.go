package pathx

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/d"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd", separator, separator, separator), Path(path))
	path = "a/b/c/d/"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%s", separator, separator, separator, separator), Path(path))
}

func TestPathDir(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/d"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%s", separator, separator, separator, separator), PathDir(path))
	path = "a/b/c/d/"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%s", separator, separator, separator, separator), PathDir(path))
}

func TestJoin(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/"
	path2 := "/d/e/f/"
	path3 := "g/h/i/"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%se%sf%sg%sh%si", separator, separator, separator, separator, separator, separator, separator, separator), Join(path, path2, path3))
	path3 = "g/h/i"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%se%sf%sg%sh%si", separator, separator, separator, separator, separator, separator, separator, separator), Join(path, path2, path3))
}

func TestJoinDir(t *testing.T) {
	assert := assert.New(t)
	path := "a/b/c/"
	path2 := "/d/e/f/"
	path3 := "g/h/i"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%se%sf%sg%sh%si%s", separator, separator, separator, separator, separator, separator, separator, separator, separator), JoinDir(path, path2, path3))
	path3 = "g/h/i/"
	assert.Equal(fmt.Sprintf("a%sb%sc%sd%se%sf%sg%sh%si%s", separator, separator, separator, separator, separator, separator, separator, separator, separator), JoinDir(path, path2, path3))
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
	assert.True(Contains(path, fmt.Sprintf("c%sd%se", separator, separator)))
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
