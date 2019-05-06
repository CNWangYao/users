package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io"
	"encoding/json"
	_"github.com/Go-SQL-Driver/MySQL"
	"database/sql"
)

type server struct {
	//数据库
	//db        *gorm.DB
	////所有配置
	//appConfig AppConfig
	////通过访问者IP获得地理位置
	//geoipdb   *geoip2.Reader
}





func main() {
	//defer_call()
	r:=mux.NewRouter()
    //s:=server{}
	r.HandleFunc("/api/register", resister)

	for {
		http.ListenAndServe(":8080",r)
}


    //r.HandleFunc("/api/login",s.login)
	fmt.Println("hello")

}


type result struct{
	Code int
	Msg string
	data   interface{}
}


func resister(w http.ResponseWriter,r * http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	username,found1:=r.Form["username"]
	password,found2:=r.Form["password"]

	if !(found1 && found2) {
		io.WriteString(w,"请勿非法访问")
		return
	}


	db,err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/users?charset=utf8")
	if err!=nil {
		io.WriteString(w,"数据库链接失败")
		return
	}

	defer  db.Close()

    sql:="insert into user(username, password) values(?,?)"
	_,err = db.Exec(sql,username[0],password[0])
	if err!=nil {
		arr:=&result{
			500,
			"注册失败",
			err,
		}
		b,json_err:=json.Marshal(arr)

		if json_err!=nil {
			fmt.Println("enocoding failed")
		}else{
			io.WriteString(w,string(b))
		}
	}else{
		arr:=&result{
			200,
			"注册成功",
			[]string{},
		}
		b,_json_err:=json.Marshal(arr)
		if _json_err!=nil {
			println(_json_err)
		}else{
			io.WriteString(w,string(b))
		}
	}



}




//func defer_call() {
//
//	defer func() {
//		fmt.Println("打印前")
//	}()
//
//	defer func() {
//		fmt.Println("打印中")
//	}()
//
//	defer func() {
//		fmt.Println("打印后")
//	}()
//
//	panic("触发异常")
//
//
//}
