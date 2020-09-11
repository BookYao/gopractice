/**
 * @Author: BookYao
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/9/11 10:14
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/vl/audio_attendance/result", rollcallResMsgHandle)
	log.Fatal(http.ListenAndServe(":8003", nil))
}
