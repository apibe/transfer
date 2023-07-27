package outpackage

import (
	"fmt"
	"github.com/masoner/transfer"
	"os"
	"testing"
)

func TestOpenXlsx(t *testing.T) {
	open, err := os.Open("../testfile/测试.xlsx")
	if err != nil {
		panic(err.Error())
	}
	xlsx := transfer.Xlsx{
		HeadMap: map[int]string{0: "YOU", 1: "I"},
		Field:   0,
		Value:   1,
		Sheet:   "Sheet1",
	}
	info, err := transfer.OpenXlsxReader(xlsx, open)
	i, err := info.SqlGenerator(transfer.DbPGSql, "t_test").WriteFile("./xx.sql")
	//marshal, err := json.Marshal(generator.GetCollection())
	//fmt.Println(string(marshal))
	fmt.Println(i)
	if err != nil {
		fmt.Println(err.Error())
	}
}
