package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3305)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type Result struct {
		Name    string
		Type    string
		Comment string
	}

	type Table struct {
		Name    string
		Comment string
	}

	var result []Result
	var tables []Table

	db_name := "test"

	table_sql := "SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '" + db_name + "' "
	db.Raw(table_sql).Scan(&tables)
	for _, table := range tables {

		table_name := table.Name

		fmt.Println("当前表为；" + table_name + " 注释为：" + table.Comment)
		db.Raw("SELECT COLUMN_NAME AS Name, DATA_TYPE AS Type, COLUMN_COMMENT AS Comment   FROM INFORMATION_SCHEMA.COLUMNS   WHERE TABLE_SCHEMA = '" + db_name + "'  AND TABLE_NAME = '" + table_name + "';  ").Scan(&result)
		for _, k := range result {
			fmt.Println(k.Name + "|" + k.Type + "|" + k.Comment)
		}
	}

}
