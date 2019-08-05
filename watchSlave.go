package per

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func SqlGet(Dsn string) {
	//Dsn: "bak:CQQlSvrY4mtKiE5rNZhL@tcp(172.31.2.243:3306)/collection?charset=utf8",
	//Dsn: "root:HQzAp4w5vA@tcp(192.168.37.132:3310)/mysql?charset=utf8",
	db, err := sql.Open("mysql", Dsn)
	defer db.Close()
	query_sql := `show slave status`
	rows, err := db.Query(query_sql)
	//fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	// for rows.Next() {
	//  err := rows.Scan(&cloumns[0], &cloumns[1], &cloumns[2])
	//  if err != nil {
	//      log.Fatal(err)
	//  }
	//  fmt.Println(cloumns[0], cloumns[1], cloumns[2])
	// }
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var comment string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			var rePrint string
			if cloumns[i] == "Last_SQL_Errno" {
				if value != "0" {
					rePrint += value
				} else if cloumns[i] == "Last_SQL_Error" {
					if value != "" {
						rePrint = rePrint + ": " + value
					}
				}
			}
			if rePrint != "" {
				fmt.Println(rePrint)
				//per.Mail("[slave 同步]", rePrint, "xujiabin@9fbank.com.cn", "172.31.1.243", 465)
				Mail("[slave 同步]", rePrint, "xujiabin@9fbank.com.cn", "172.31.1.243", 465)
			}
			comment = comment + fmt.Sprintf("%v:%v\n", cloumns[i], value)
		}

	}
	fmt.Println(comment)
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
		return
	}
}
