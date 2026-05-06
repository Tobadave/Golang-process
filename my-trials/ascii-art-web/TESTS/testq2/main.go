// package testq2
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	demoJson()
}

func demoJson() {
	type Employee struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	fmt.Println(">>>>>>>>>>>Struct to JSON...")
	emp := &Employee{Name: "John", Age: 32, Email: "test@gmail.com"}
	b, err := json.Marshal(emp)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var newEmp Employee
	str := `{"name":"John Doe", "age":"32", "email":"hello@gmail.com"}` //this is a json data
	json.Unmarshal([]byte(str), &newEmp)                                //so yes, this stores the "str" as bytes and stores in the adress of newEMployee.. thats like anew copy of the initial one

	fmt.Printf("Name: %s", newEmp.Name)
	fmt.Printf("Age: %d", newEmp.Age)
	fmt.Printf("Email: %s", newEmp.Email)

}
