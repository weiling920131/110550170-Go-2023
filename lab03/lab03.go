package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	res1 := strings.Split(r.URL.Path, "/")
	// fmt.Fprintf(w, "Error!")
	// fmt.Println(res1)
	if (len(res1) != 4){
		fmt.Fprintf(w, "Error!")
	} else{
		a,err:=strconv.Atoi(res1[2])
		if err != nil {
			// Add code here to handle the error!
			fmt.Fprintf(w, "Error!")
			return
		 }
		b,err:=strconv.Atoi(res1[3])

		if err != nil {
			// Add code here to handle the error!
			fmt.Fprintf(w, "Error!")
			return
		 }
		if res1[1] == "add"{
			fmt.Fprintf(w, "%s + %s = %s", res1[2], res1[3], strconv.Itoa(a+b))
			return 
		}else if res1[1] == "sub" {
			fmt.Fprintf(w, "%s - %s = %s", res1[2], res1[3], strconv.Itoa(a-b))
			return
		}else if res1[1] == "mul"{
			fmt.Fprintf(w, "%s * %s = %s", res1[2], res1[3], strconv.Itoa(a*b))
			return
		}else if res1[1] == "div"{
			if b==0{
				fmt.Fprintf(w, "Error!")
			}else{
				fmt.Fprintf(w, "%s / %s = %s, reminder = %s", res1[2], res1[3], strconv.Itoa(a/b),strconv.Itoa(a%b))
				return
			}
			
		}else{
			fmt.Fprintf(w, "Error!")
			return
		}

	}

}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
