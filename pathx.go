package pathx

import (
	"os"
)

// Path 能夠移除單個路徑中多餘的符號，並且確保結尾沒有目錄符號。
func Path(path string) string {

}

// PathDir 能夠移除單個路徑中多餘的符號，並且確保結尾有目錄符號。
func PathDir(path string) string {

}

// Join 會將多個路徑彙整成一個。
func Join(paths ...string) string {

}

// JoinDir 會將多個路徑彙整乘一個並且在結尾增加目錄符號。
func JoinDir(paths ...string) string {

}

// Name 能夠取得路徑中最後的檔案名稱。
func Name(path string) {

}

// Ext 能夠取得路徑中最後的檔案副檔名（無起始 `.` 符號）。
func Ext(path string) string {

}

// Merge 會將路徑中重複的路徑截斷並且重新組合。
func Merge(path string) string {

}

// Contains 表示路徑中是否有包含指定路徑。
func Contains() bool {

}

// Executable 能夠取得執行檔案的路徑位置。
func Executable() string {

}

// Getwd 能夠取得目前的目錄路徑。
func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
