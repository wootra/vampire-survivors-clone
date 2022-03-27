package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("server is running")
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("./public")));
	if(err != nil){
		fmt.Println("failed:", err)
		return
	}
	fmt.Println("server is finished")
}