package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InitApiServer() (*http.ServeMux) {
	server := http.NewServeMux()
	server.HandleFunc("/api/member/check", handleMemCheck)

	return server
}

type respMsg struct {
	Ok bool
	Member Mem
	Msg string
}

func handleMemCheck(w http.ResponseWriter, r *http.Request) {
	var param struct{
		Name string `json:"name"`
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)

		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(buf, &param)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// 在api接口调用这做判断是否返回错误
	ret, err := retrieveStudent(param.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 可以将包装的错误返回给上层，或者直接返回一个空的查询对象
		}
	}
	fmt.Println(ret)
	// 无措返回需要查询的人员信息
	return
}

func writeResult(w http.ResponseWriter, ok bool, msg string, mem Mem) {

}
