package main

/*
  한줄로는 import 가 여러개 동시에 안되는?
*/
import (
	"fmt"
	"os"
)

/*
  this function printName
  welcome to the GoLang world
*/
func printName() {
	fmt.Println("Hellow dongok")
}

//single line commnet out
func main() {
	fmt.Println("Heelo world")
	printName()

	//exit 는 echo $? 에서 결과로 사용됨
	os.Exit(1)
}
