package transfer

import (
	"github.com/xuri/excelize/v2"
	"io"
)

// todo
// excel的信息较为固定，但是我们依然希望用户能够指定某行作为字段信息
// 针对中文的字段，我们有采用拼音全拼的形式进行存储，并为其添加注释

type Xlsx struct {
	HeadMap map[int]string // 表字段映射关系
	Field   int            // 当 Head 没有写入时，改行作为字段头行，默认 0
	Value   int            // 参数所在行[起始行，末尾行],起始行默认 0 ，结尾行<=0时则不存在末尾行
	Sheet   string         // 被导入数据所在sheet页

	fields []string
	values [][]string
}

func OpenXlsxReader(xlsx Xlsx, reader io.Reader, opts ...excelize.Options) (*Information, error) {
	xl, err := excelize.OpenReader(reader, opts...)
	defer xl.Close()
	if err != nil {
		return nil, err
	}
	rows, err := xl.GetRows(xlsx.Sheet)
	xlsx.fields = rows[xlsx.Field]
	xlsx.values = rows[xlsx.Value:]
	return xlsx.Information(), nil
}

func OpenXlsxFile(xlsx Xlsx, filename string, opts ...excelize.Options) (*Information, error) {
	excelize.OpenFile(filename, opts...)
	return nil, nil
}

func (x Xlsx) WriteFile(filename string) error {
	return nil
}
func (x Xlsx) Information() *Information {
	fields := make([]Field, 0)
	values := make([][]Value, 0)
	if x.HeadMap != nil && len(x.HeadMap) > 0 {
		for i, head := range x.HeadMap {
			fields = append(fields, Field{
				Arg:     head,
				Index:   i,
				Comment: x.fields[i],
				Class:   SqlText,
			})
		}
	} else {
		for i, field := range x.fields {
			fields = append(fields, Field{
				Arg:     field,
				Index:   i,
				Comment: field,
				Class:   SqlText,
			})
		}
	}
	for _, value := range x.values {
		v := make([]Value, 0)
		for _, field := range fields {
			if field.Index < len(value) {
				v = append(v, Value{
					Arg:   field.Arg,
					Index: field.Index,
					Value: value[field.Index],
					Class: field.Class,
				})
			}
		}
		values = append(values, v)
	}
	return &Information{
		Fields: fields,
		Values: values,
	}
}
