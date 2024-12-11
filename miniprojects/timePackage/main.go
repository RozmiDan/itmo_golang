// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main(){
// 	var dataInput string
// 	fmt.Scanln(&dataInput)
// 	parsedData, err := time.Parse(time.RFC3339, dataInput)
// 	if err != nil{
// 		panic(err)
// 	}
// 	fmt.Print(parsedData.Format(time.UnixDate))
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	var inputData string 
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		if(len(scanner.Text()) == 0){
			break
		}
		inputData += scanner.Text()
	}
	resData, err := time.Parse("2006-01-02 15:04:05", inputData)
	if err != nil {
		panic(err)
	}

	noon := time.Date(
		resData.Year(),
		resData.Month(),
		resData.Day(),
		13, 0, 0, 0,
		resData.Location(),
	)
	if resData.After(noon){
		resData = resData.AddDate(0,0,1)
	}
	fmt.Print(resData.Format("2006-01-02 15:04:05"))
}