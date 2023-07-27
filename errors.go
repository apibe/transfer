package transfer

import (
	"errors"
	"strings"
)

func (i Information) verify() error {
	errs := make([]string, 0)
	if len(i.Values) < 1 {
		errs = append(errs, "无可插入数据")
	}
	if len(i.Fields) < 1 {
		errs = append(errs, "无可建表字段")
	}
	if len(i.Fields) != kidMaxLen(i.Values) {
		errs = append(errs, "表字段长度不一致！")
	}
	if len(errs) > 0 {
		return errors.New("invalid database information : " + strings.Join(errs, " "))
	}
	return nil
}

func (x Xlsx) verify() error {
	return nil
}

func kidMaxLen(kids [][]Value) int {
	max := 0
	for _, kid := range kids {
		if len(kid) > max {
			max = len(kid)
		}
	}
	return max
}
