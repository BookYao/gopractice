/**
 * @Author: BookYao
 * @Description:
 * @File:  rollbackRes
 * @Version: 1.0.0
 * @Date: 2020/9/11 10:13
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type rollcallRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success string `json:"success"`
	Data    string `json:"data"`
}

func rollcallResMsgBuild() ([]byte, error) {
	var resMsg = rollcallRes{
		Code:    "200",
		Message: "ok",
		Success: "true",
		Data:    "111",
	}
	jsonResMsg, err := json.Marshal(resMsg)
	if err != nil {
		log.Println("rollcall json marshal failed!")
		return nil, err
	}
	return jsonResMsg, nil
}

func rollcallResMsgHandle(w http.ResponseWriter, r *http.Request) {
	msg, err := rollcallResMsgBuild()
	if err != nil {
		log.Println("rollcall res msg build failed.")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "")
		return
	}

	fmt.Fprintf(w, "%s\n", msg)
}
