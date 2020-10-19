package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main(){
	fmt.Println("Program Begins 🤓")
	resp,_ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	byteData,_ :=ioutil.ReadAll(resp.Body)
	stringData := string(byteData)
	fmt.Println(stringData)
	fmt.Println("End of program  😈"  )
}