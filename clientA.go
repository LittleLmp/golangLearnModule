package web

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

func GetWebAllInfo(p *string) {
	website := *p
	httpRes, err := http.Get(website)
	if err != nil {
		log.Println("fetch web resourse failed: ",err)
		return 
	}
	defer httpRes.Body.Close()
	if httpRes.StatusCode != http.StatusOK {
		log.Println("status code: ",httpRes.StatusCode)
	}
	httpBody, _ := ioutil.ReadAll(httpRes.Body)
		fmt.Println("body: ",string(httpBody))
	for k, v := range httpRes.Header {
        fmt.Println(k, ":", strings.Join(v, ""))
    }
}