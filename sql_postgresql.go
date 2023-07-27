package transfer

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"strings"
)

const (
	server_name   = "{SERVER_NAME}"         // 数据库连接地址
	server_type   = "{SERVER_TYPE}"         // 数据库类型
	host          = "{HOST}"                // 数据库连接地址
	database      = "{DATABASE}"            // 数据库名称
	schema        = "{SCHEMA}"              // 模式名称
	date          = "{02/01/2006 15:04:05}" // SQL生成时间
	table         = "{TABLE}"               // 数据表名称
	create_sql    = "{CREATE_SQL}"          // 建表语句
	insert_sql    = "{INSERT_SQL}"          // 插入表语句
	self_add      = "{SELF_ADD}"            // 自增id函数名称
	primary_key   = "{PRIMARY_KEY}"         // 主键名称
	before_create = "{BEFORE_CREATE}"       // 建表前SQL
	after_create  = "{AFTER_CREATE}"        // 建表后SQL
	before_insert = "{BEFORE_INSERT}"       // 插入前SQL
	after_insert  = "{AFTER_INSERT}"        // 插入后SQL
)

var pgModel = `/*
 ONLINE-MAP-MAKING Premium Data Transfer

 Source Server         : {SERVER_NAME}
 Source Server Type    : {SERVER_TYPE}
 Source Host           : {HOST}
 Source Catalog        : {DATABASE}
 Source Table          : {TABLE}
 Source Schema         : {SCHEMA}

 Date: {02/01/2006 15:04:05}
*/
-- ----------------------------
-- Before create for {TABLE}
-- ----------------------------
{BEFORE_CREATE}

-- ----------------------------
-- Table structure for {TABLE}
-- ----------------------------
{CREATE_SQL}

-- ----------------------------
-- After create for {TABLE}
-- ----------------------------
{AFTER_CREATE}

-- ----------------------------
-- Before insert for {TABLE}
-- ----------------------------
{BEFORE_INSERT}

-- ----------------------------
-- Records of {TABLE}
-- ----------------------------
{INSERT_SQL}

-- ----------------------------
-- After insert for {TABLE}
-- ----------------------------
{AFTER_INSERT}
`

type pg struct {
	info  *Information
	table string
}

func (p pg) GetCreateSql() string {
	argStr := ""
	for _, field := range p.info.Fields {
		argStr = fmt.Sprintf("%s, %s %s", argStr, field.Arg, field.Class.toString())
	}
	argStr = strings.TrimPrefix(argStr, ",")
	return fmt.Sprintf("CREATE TABLE %s (%s);", p.table, argStr)
}
func (p pg) GetInsertSql() []string {
	// INSERT INTO users (name, age, gender) VALUES ('张三', 25, 'male')
	insertSql := make([]string, 0)
	for _, value := range p.info.Values {
		argStr := ""
		valueStr := ""
		for _, v := range value {
			argStr += fmt.Sprintf("%s,'%s'", argStr, v.Arg)
			valueStr = fmt.Sprintf("%s,'%s'", valueStr, v.Value)
		}
		argStr = strings.TrimPrefix(argStr, ",")
		valueStr = strings.TrimPrefix(valueStr, ",")
		insertSql = append(insertSql, fmt.Sprintf("INSERT INTO %s (%s) VALUE (%s);", p.table, argStr, valueStr))
	}
	return insertSql
}

func (p pg) GetCollection() *Collection {
	return &Collection{
		Information: p.info,
		CreateSql:   p.GetCreateSql(),
		InsertSql:   p.GetInsertSql(),
	}
}

func (p pg) Write(dsn string, opts ...Opinions) (int, error) {
	client, err := pgClient(dsn)
	defer client.Close()
	if err != nil {
		return 0, err
	}
	tx, _ := client.Begin()
	c := p.GetCollection()
	for _, opt := range opts {
		if len(opt.BeforeCreate) > 0 {

		}
		if len(opt.AfterCreate) > 0 {

		}
		if len(opt.BeforeInsert) > 0 {

		}
		if len(opt.AfterInsert) > 0 {

		}
	}
	tx.Exec(c.CreateSql)
	for _, insertSql := range c.InsertSql {
		tx.Exec(insertSql)
	}
	tx.Commit()
	return 0, nil
}
func (p pg) WriteFile(filename string, opts ...Opinions) (int, error) {
	coll := p.GetCollection()
	l, parse := pgParse(coll, opts...)
	_, err := createFile(filename)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(filename, []byte(parse), 0644)
	return l, err
}

func PgQuery(dsn string, sql string, p *pg) {

}

func pgClient(dsn string) (*sql.DB, error) {
	schema := "public"
	user := "user"
	password := ""
	host := ""
	port := ""
	dbname := ""
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s schema=%s sslmode=%s",
		host, port, user, password, dbname, schema, "disable"))
	err = db.Ping()
	return db, err
}

func pgParse(coll *Collection, opts ...Opinions) (int, string) {
	model := pgModel
	model = strings.ReplaceAll(model, create_sql, coll.CreateSql)
	model = strings.ReplaceAll(model, insert_sql, strings.Join(coll.InsertSql, "\n"))
	for _, opt := range opts {
		if len(opt.BeforeCreate) > 0 {
			model = strings.ReplaceAll(model, before_create, strings.Join(opt.BeforeCreate, "\n"))
		}
		if len(opt.AfterCreate) > 0 {
			model = strings.ReplaceAll(model, after_create, strings.Join(opt.AfterCreate, "\n"))
		}
		if len(opt.BeforeInsert) > 0 {
			model = strings.ReplaceAll(model, before_insert, strings.Join(opt.BeforeInsert, "\n"))
		}
		if len(opt.AfterInsert) > 0 {
			model = strings.ReplaceAll(model, after_insert, strings.Join(opt.AfterInsert, "\n"))
		}
	}
	return len(coll.InsertSql), model
}
