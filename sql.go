package transfer

const (
	TransXlsx transfer = iota
	TransCsv
	TransXml
	TransJson
)

const (
	DbPGSql DbType = iota
)

const (
	SqlText sqlType = iota
	SqlFloat
	SqlGeometry
)

type (
	DbInfo struct {
		User     string
		Password string
		Host     string
		Port     string
		Schema   string
	}
	Opinions struct {
		PrimaryKey   string   // 指定主键列，
		BeforeCreate []string // 建表前
		AfterCreate  []string // 建表后
		BeforeInsert []string // 插入前
		AfterInsert  []string // 插入后
	}
)

type (
	sqlType  int
	transfer int
	DbType   int
	Field    struct {
		Arg     string
		Index   int
		Comment string
		Class   sqlType
	}
	Value struct {
		Arg   string
		Index int
		Value string
		Class sqlType
	}
	Information struct {
		Fields []Field
		Values [][]Value
	}
	Collection struct {
		Information *Information
		CreateSql   string
		InsertSql   []string
	}

	SqlGeneratorInterface interface {
		GetCollection() *Collection
		GetCreateSql() string
		GetInsertSql() []string
		Write(dsn string, opts ...Opinions) (int, error)
		WriteFile(filename string, opts ...Opinions) (int, error)
	}
)

func (i *Information) SqlGenerator(dbType DbType, table string) SqlGeneratorInterface {
	switch dbType {
	case DbPGSql:
		return pg{
			info:  i,
			table: table,
		}
	default:
		return pg{
			info:  i,
			table: table,
		}
	}
}

func (t sqlType) toString() string {
	switch t {
	case SqlText:
		return "text"
	default:
		return "text"
	}
}
