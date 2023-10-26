package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	// "strings"
)

// TODO: Create a struct to hold the data sent to the template
type container struct{
	Expression	string
	Result		string
}
func gcd(n1 int, n2 int) int{
	if(n2%n1 == 0){
		return n1
	}
	return gcd(n2%n1, n1)
}

func lcm(n1 int , n2 int) int{
	return (n1/gcd(n1, n2))*(n2/gcd(n1, n2))*gcd(n1, n2)

}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function
	
	op := r.URL.Query().Get("op")
	num1 := r.URL.Query().Get("num1")
	num2 := r.URL.Query().Get("num2")
	// fmt.Fprintf(w, op)
	// fmt.Fprintf(w, num1)
	// fmt.Fprintf(w, num2)
	if(op == "" || num1=="" || num2==""){
		http.ServeFile(w, r, "error.html")
		return
	}
	n1, err := strconv.Atoi(num1)
	n2, err := strconv.Atoi(num2)
	
	if err != nil {
		// Add code here to handle the error!
		http.ServeFile(w, r, "error.html")
		return
	 }

	if op == "add"{
		var test = container{
			Expression: num1+" + "+num2,
			Result: strconv.Itoa(n1+n2),
		}
		err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
		return 
	}else if op == "sub" {
		var test = container{
			Expression: num1+" - "+num2,
			Result: strconv.Itoa(n1-n2),
		}
		err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
		return
	}else if op == "mul"{
		var test = container{
			Expression: num1+" * "+num2,
			Result: strconv.Itoa(n1*n2),
		}
		err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
		return
	}else if op == "div"{
		if n2==0{
			http.ServeFile(w, r, "error.html")
		}else{
			var test = container{
				Expression: num1+" / "+num2,
				Result: strconv.Itoa(n1 / n2),
			}
			err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
			return
		}
		
	}else if op == "gcd"{
		var test = container{
			Expression: "GCD("+num1+", "+num2+")",
			Result: strconv.Itoa(gcd(n1,n2)),
		}
		err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
	}else if op == "lcm"{
		var test = container{
			Expression: "LCM("+num1+", "+num2+")",
			Result: strconv.Itoa(lcm(n1, n2)),
		}
		err = template.Must(template.ParseFiles("index.html")).Execute(w, test)
	}else{
		http.ServeFile(w, r, "error.html")
		return
	}

	// }
	
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
