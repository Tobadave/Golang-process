package main

import (
	"fmt"
	// "io/ioutil"
	"os"
	"log"
)

func main() {
	// os.WriteFile()

	// // write data into a file
	// fmt.Println("writing data into a file")
	// writeFile("welcome to go")
	// readFile()
	// // read data from a file
	// fmt.Println("reading data from a file")
	// readFile()
	calculate()
	fmt.Println("====================================")
	formatLogging()
	log.Println("testing this with details..")


}

func formatLogging() {
	fmt.Println("Testing right now")
	var warning *log.Logger //log logger acn be used to give more detials to a person .. 
	// similr to string builder? yes.


	//but from below you can redefine warning again.. into a new one and being able to add it as a clear output.
	// os.Stdout, waring o<- which is the text
	// /log.Logger

	warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile) //details to add, they ofteen use Log.Ltime, log.Ldate, log.Lshortfile?
		//log.Println, by default does another already..
	
	warning.Println("This is a warning message 1")
	warning.Println("This is a warning message 2")
}
/*
package main
import (
"fmt"
"sync"
"math/rand"
"time"
)
type Task struct {
value int
executedBy string
}
var total int
var task Task
var lock sync.Mutex
func main(){
fmt.Printf("synchronizing goroutines demo\n")
total = 0
task.value = 0
task.executedBy = ""
display()
// run background
go calculate()
go perform()
// press ENTER to exit
var input string
fmt.Scanln(&input)
fmt.Println("done")
}
func calculate(){for total < 10 {
lock.Lock()
task.value = rand.Intn(100)
task.executedBy = "from calculate()"
display()
total++
lock.Unlock()
time.Sleep(500 * time.Millisecond)
}
}
func perform(){
for total < 10 {
lock.Lock()
task.value = rand.Intn(100)
task.executedBy = "from perform()"
display()
total++
lock.Unlock()
time.Sleep(500 * time.Millisecond)
}
}
func display(){
fmt.Println("--------------------------")
fmt.Println(task.value)
fmt.Println(task.executedBy)
fmt.Println("--------------------------")
}
*/






/*
ate a file, maing.go, and write this code.
package main
import (
"fmt"
"time"
)
func main(){
fmt.Printf("goroutines demo\n")
// run func in background
go calculate()
index := 0
// run in background
go func() {
for index < 6 {
fmt.Printf("go func() index= %d \n", index)
var result float64 = 2.5 * float64(index)
fmt.Printf("go func() result= %.2f \n", result)
time.Sleep(500 * time.Millisecond)
index++
}
}()
// run in background
go fmt.Printf("go printed in the background\n")
// press ENTER to exit
var input string
fmt.Scanln(&input)
fmt.Println("done")
}
func calculate() {
i := 12
for i < 15 {
fmt.Printf("calculate()= %d \n", i)
var result float64 = 8.2 * float64(i)
fmt.Printf("calculate() result= %.2f \n", result)time.Sleep(500 * time.Millisecond)
i++
}
}
*/



/*
func main() {
fileLogging()
}
func fileLogging(){
fmt.Println("-----------file logging-----------")
file, err := os.OpenFile("d:/temp/myapp.log",
os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
fmt.Println("Failed to open log file")
return
}
var logFile *log.Logger
logFile = log.New(
file,
"APP: ",
log.Ldate|log.Ltime|log.Lshortfile)
logFile.Println("This is error message 1")
logFile.Println("This is error message 2")
logFile.Println("This is error message 3")
fmt.Println("Done")
}
*/

func calculate() {
	fmt.Println(">>>>>>>>>>> Trying out this holder")

	ans := 0
	defer func() {

		if error := recover(); error != nil {
			fmt.Println("Recovering...")
			fmt.Println(error)
		}
	}()
//a way to use defer as a guard function.. if 
// if (err := recover(); err != nil {})

	a := 10
	b := 0

	ans = a/b
	fmt.Printf("result = %d", ans)
	fmt.Println("Done")
}

// func writeFile(message string) {
// 	// bytes := []byte(message)
// 	// ioutil.WriteFile("./testgo.txt", bytes, 0644)
// 	// fmt.Println("created a file")

// 	bytes := []byte(message) //stores the message passed as bytes indexe (oriinal way for raw files)
// 	ioutil.WriteFile("./testgo.txt", bytes, 0644)
// 	fmt.Println("created a file")
// }
// func readFile() {
// 	data, _ := ioutil.ReadFile("./testgo.txt")
// 	fmt.Println("file content:")
// 	fmt.Println(string(data))
// }
