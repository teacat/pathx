package pathx

import (
	"os"
	"path/filepath"
	"strings"
)

// separator 是本地系統路徑符號。
var separator = string(os.PathSeparator)

// doubleSeparator 是雙重本地系統路徑符號。
var doubleSeparator = separator + separator

// toLocalPath 會將傳入的路徑轉換成本地系統路徑符號。
func toLocalPath(path string) string {
	return filepath.FromSlash(path)
}

// Path 能夠移除單個路徑中多餘的符號，並且確保結尾沒有目錄符號。
func Path(path string) string {
	path = toLocalPath(path)
	path = filepath.Clean(path)
	path = strings.TrimSuffix(path, separator)
	return path
}

// PathDir 能夠移除單個路徑中多餘的符號，並且確保結尾有目錄符號。
func PathDir(path string) string {
	return Path(path) + separator
}

// Join 會將多個路徑彙整成一個。
func Join(paths ...string) string {
	var path string
	for _, v := range paths {
		path += v + "/"
	}
	return Path(path)
}

// JoinDir 會將多個路徑彙整乘一個並且在結尾增加目錄符號。
func JoinDir(paths ...string) string {
	return PathDir(Join(paths...))
}

// Name 能夠取得路徑中最後的檔案名稱。
func Name(path string) string {
	basename := filepath.Base(path)
	return strings.TrimSuffix(basename, filepath.Ext(basename))
}

// Ext 能夠取得路徑中最後的檔案副檔名（無起始 `.` 符號）。
func Ext(path string) string {
	return strings.TrimLeft(filepath.Ext(path), ".")
}

// Contains 表示路徑中是否有包含指定路徑。
func Contains(path string, subpath string) bool {
	return strings.Contains(toLocalPath(path), toLocalPath(subpath))
}

// Executable 能夠取得執行檔案的路徑位置。
func Executable() string {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return e
}

// ExecutableDir 能夠取得執行檔案父資料夾的路徑位置。
func ExecutableDir() string {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(e)
}

// Getwd 能夠取得目前的目錄路徑。
func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
