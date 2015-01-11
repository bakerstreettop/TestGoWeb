// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package BakerStreetMain

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

type Person struct {  
	Name    string  
	Age     int  
	Emails  []string  
	Company string  
	Role    string  
}  
	 
type OnlineUser struct {  
	User      []*Person  
	LoginTime string  
}  
	  
func Handler(w http.ResponseWriter, r *http.Request) {  
	SH := Person{  
	Name: "Sherlock",   
	Age: 27,   
	Emails: []string{"SherlockInChina@gmail.com", "Sherlock@bakerstreet.top"},  
	Company: "BakerStreetTop",  
	Role: "SE"}  
	  
	JW := Person{Name: "JohnWarson", Age: 27, Emails: []string{"kira@gmail.com", "kira@bakerstreet.top"}}  
	  
	onlineUser := OnlineUser{User: []*Person{&SH, &JW}}  
	  
	//t := template.New("Person template")  
	//t, err := t.Parse(templ)  
	t, err := template.ParseFiles("BakerStreetMain.html")  
	checkError(err)  
	 
	err = t.Execute(w, onlineUser)  
	checkError(err)  
}  


func init() {

	http.Handle("css/", http.FileServer(http.Dir("Static")))
	http.Handle("js/", http.FileServer(http.Dir("Static")))
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8888", nil)  
}

// func main() {  
// 	http.HandleFunc("/", Handler)  
//  http.ListenAndServe(":8888", nil)  	
// }  
	  
func checkError(err error) {  
	if err != nil {  
	fmt.Println("Fatal error ", err.Error())  
	os.Exit(1)  
	}  
}  




// func handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<html><meta charset='gb2312'/><title> ☎ BakerStreet </title><body>你好，旅行者。。!</body></html>")
// }
