package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
    var err error
    dsn := "root:12345678@tcp(localhost:3306)/test"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }

    DB.SetMaxOpenConns(100)
    DB.SetMaxIdleConns(16)
    return nil
}

type User struct {
    // 内容可能为空的字段类型应该设置成NullString或NullInt64等

    Id int `db:"id"`
    Name sql.NullString `db:"string"`
    Age int `db:"age"`
}

func main() {
    err := initDb()
    if err != nil {
        fmt.Printf("init db failed, err:%v\n", err)
        return
    }
    // 查询一条
    //testQueryData()
    // 查询多条
    //testQueryMultiRow()
    // 插入一条
    //testInsertData()
    // 更新记录
    //testUpdateData()
    // 删除记录
    //testDeleteData()

    // prepare预处理
    //testPrepareData()
    testPrepareInsertData()
}

func testQueryData() {
    sqlstr := "select id, name, age from user where id=?"
    row := DB.QueryRow(sqlstr, 2)

    var user User
    err := row.Scan(&user.Id, &user.Name, &user.Age)
    if err != nil {
        fmt.Printf("scan failed, err:%v\n", err)
        return
    }

    fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
}

func testQueryMultiRow() {
    sqlstr := "select id, name, age from user where id>?"
    rows, err := DB.Query(sqlstr, 0)
    // 重点：rows对象一定要close掉
    defer func() {
        if rows != nil {
            rows.Close()
        }
    }()
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }

    for rows.Next() {
        var user User
        err := rows.Scan(&user.Id, &user.Name, &user.Age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("user:%#v\n", user)
    }
}

func testInsertData() {
    sqlstr := "insert into user(name, age) values (?, ?)"
    result, err := DB.Exec(sqlstr, "tom", 18)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Printf("get last insert id failed, err:%v\n", err)
        return
    }
    fmt.Printf("id is %d\n", id)
}

func testUpdateData() {
    sqlstr := "update user set name=? where id=?"
    result, err := DB.Exec(sqlstr, "jim", 3)
    if err != nil {
        fmt.Printf("update failed, err:%v\n", err)
        return
    }
    affectedRows, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("get affected rows failed, err:%v\n", err)
    }
    fmt.Printf("Affected rows is %d\n", affectedRows)
}

func testDeleteData() {
    sqlstr := "delete from user where id=?"
    result, err := DB.Exec(sqlstr,  3)
    if err != nil {
        fmt.Printf("delete failed, err:%v\n", err)
        return
    }
    affectedRows, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("get affected rows failed, err:%v\n", err)
    }
    fmt.Printf("delete db succ, Affected rows is %d\n", affectedRows)
}

func testPrepareData() {
    sqlstr := "select id,name, age from user where id > ?"
    stmt, err := DB.Prepare(sqlstr)
    if err != nil {
        fmt.Printf("prepare failed, err:%v\n", err)
        return
    }
    defer func() {
        if stmt != nil {
            stmt.Close()
        }
    }()

    rows, err := stmt.Query(0)
    // 重点：rows对象一定要close掉
    defer func() {
        if rows != nil {
            rows.Close()
        }
    }()
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }

    for rows.Next() {
        var user User
        err := rows.Scan(&user.Id, &user.Name, &user.Age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("user:%#v\n", user)
    }
}

func testPrepareInsertData() {
    sqlstr := "insert into user(name, age) values (?, ?)"
    stmt, err := DB.Prepare(sqlstr)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    // stmt关闭
    defer func() {
        if stmt != nil {
            stmt.Close()
        }
    }()
    result, err := stmt.Exec("jim", 100)

    id, err := result.LastInsertId()
    if err != nil {
        fmt.Printf("get last insert id failed, err:%v\n", err)
        return
    }
    fmt.Printf("id is %d\n", id)
}