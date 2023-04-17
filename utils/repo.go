package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Res(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	resData, err := json.Marshal(h)
	if err != nil {
		fmt.Println("err", err)
	}
	w.Write(resData)
}

func ResList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	resData, err := json.Marshal(h)
	if err != nil {
		fmt.Println("err", err)
	}
	w.Write(resData)

}

func ResFail(w http.ResponseWriter, msg string) {
	Res(w, -1, nil, msg)

}

func ResOk(w http.ResponseWriter, data interface{}, msg string) {
	Res(w, 0, data, msg)

}

func ResOkList(w http.ResponseWriter, data interface{}, total interface{}) {
	ResList(w, 0, data, total)

}
