package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func initDb() error {
    var err error
    dsn := "root:12345678@tcp(localhost:3306)/test"
    DB, err = sqlx.Open("mysql", dsn)
    if err != nil {
        return err
    }

    DB.SetMaxOpenConns(100)
    DB.SetMaxIdleConns(16)
    return nil
}

type User struct {
    // 内容可能为空的字段类型应该设置成NullString或NullInt64等

    Id int64 `db:"id"`
    Name sql.NullString `db:"name"`
    Age int `db:"age"`
}

func main() {
    err := initDb()
    if err != nil {
        fmt.Printf("init db failed, err:%v\n", err)
        return
    }

    // 查询一条
    //testQuery()

    // 查询多条
    //testQueryMulti()

    // 更新
    testUpdateData()
}

func testQuery() {
    sqlstr := "select id, name, age from user where id=?"
    var user User

    err := DB.Get(&user, sqlstr, 2)
    if err != nil {
        fmt.Printf("get failed, err:%v\n", err)
        return
    }
    fmt.Printf("user:%#v\n", user)
}

func testQueryMulti() {
    sqlstr := "select id, name, age from user where id > ?"
    var user []User

    err := DB.Select(&user, sqlstr, 0)
    if err != nil {
        fmt.Printf("get failed, err: %v\n", err)
        return
    }
    fmt.Printf("user:%#v\n", user)
}

func testUpdateData() {
    sqlstr := "update user set name=? where id=?"
    result, err := DB.Exec(sqlstr, "abc", 1)
    if err != nil {
        fmt.Printf("update failed, err:%v\n", err)
        return
    }

    count, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("affected rows failed, err:%v\n", err)
        return
    }
    fmt.Printf("affected rows: %d\n", count)
}