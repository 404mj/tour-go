package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//四种连接数据库的形式
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/leetcode?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	// CRUD(db)
	mapSelect(db)

}

//将查询结果映射到map上！
//*** 具体怎么映射有几种方案，不同的方案方法不一样。
//这段代码看了很久才明白！！！草！
func mapSelect(db *sql.DB) {
	//查询数据，取所有字段
	rows2, _ := db.Query("select * from Employee")
	//返回所有列
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里
	//**************这是我忽视的地方，其实不需要这么麻烦写的！！
	//https://godoc.org/database/sql#Row.Scan
	//还有一个忽视的地方是二维数组和切片的使用赋值有点和普通语言不一样，类C。。见fiftyShadows.go
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	//******************

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	fmt.Println(result[2]["Id"])
}

//学习CRUD操作, 注意是*sql.DB类型，查文档得到！！！
func CRUD(db *sql.DB) {

	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()

	//简单查询，指定字段名
	rows, _ := db.Query("select Id, Name from Employee")
	//所有的列1
	// fmt.Println(rows.Columns())
	id := 0
	name := ""
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	//查询一行数据
	rows3 := db.QueryRow("select Id, Name from Employee where Id = ?", 1)
	rows3.Scan(&id, &name)
	fmt.Println(id, name)

	//插入一行数据
	ret, _ := db.Exec("insert into Employee2(Name,Salary,DepartmentId) values('go-',12,1)")
	//获取插入ID
	ins_id, _ := ret.LastInsertId()
	fmt.Println(ret.RowsAffected())
	fmt.Println(ins_id)
	//总是返回0！发现是mysql的问题！以为是我mysql配置的不对，发现其实是用的这个表没设置主键和auto_increment!!
	//https://stackoverflow.com/questions/13880267/mysql-last-insert-id-returns-0/25731698

	//更新数据
	ret2, _ := db.Exec("update Employee2 set Name = 'go-test-up' where Id = ?", ins_id)
	//获取影响行数
	aff_nums, _ := ret2.RowsAffected()
	fmt.Println("RowsAffected: ", aff_nums)

	//删除数据
	ret3, _ := db.Exec("delete from Employee2 where Id = ?", ins_id)
	//获取影响行数
	del_nums, _ := ret3.RowsAffected()
	fmt.Println(del_nums)

	//预处理语句
	stmt, _ := db.Prepare("select Id,Name from Employee2 where Id = ?")
	rows4, _ := stmt.Query(3)
	//注意这里需要Next()下，不然下面取不到值
	rows4.Next()
	rows4.Scan(&id, &name)
	fmt.Println(id, name)

	//事务处理
	tx, _ := db.Begin()

	ret4, _ := tx.Exec("update Employee2 set Salary = Salary + 10 where id = ?", 1)
	ret5, _ := tx.Exec("update Employee2 set Salary = Salary - 10 where id = ?", 2)
	upd_nums1, _ := ret4.RowsAffected()
	upd_nums2, _ := ret5.RowsAffected()

	if upd_nums1 > 0 && upd_nums2 > 0 {
		//只有两条更新同时成功，那么才提交
		tx.Commit()
	} else {
		//否则回滚
		tx.Rollback()
	}
}
