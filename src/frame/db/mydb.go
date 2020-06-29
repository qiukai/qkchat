package mydb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@(localhost:3306)/qk?charset=utf8&parseTime=True&loc=Local")
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(10)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("创建数据库连接成功")
}


func Tx(fun func(sqldb *sql.DB) error) {
	tx, err := db.Begin()
	defer db.Close()
	if err != nil {
		log.Fatal("开启事物失败", err)
	}
	err = fun(db)
	if err != nil {
		log.Fatal("执行失败", err)
		err = tx.Rollback()
		if err != nil {
			log.Fatal("事物回滚失败", err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			log.Fatal("事物提交失败", err)
		}
	}

}

func Query(sqlStr string) {

	rows, err := db.Query(sqlStr)
	defer rows.Close()
	defer db.Close()

	if err != nil {
		log.Fatal("查询出错", err)
	}

	//获取列名
	columns, _ := rows.Columns()

	for i := range columns {
		fmt.Print(columns[i], "\t")
	}
	fmt.Println()

	//定义一个切片,长度是字段的个数,切片里面的元素类型是sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	//定义一个切片,元素类型是interface{} 接口
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		//把sql.RawBytes类型的地址存进去了
		scanArgs[i] = &values[i]
	}
	//获取字段值
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i := range scanArgs {
			fmt.Print(string(values[i]), "\t")
		}
		fmt.Println()
	}
}
