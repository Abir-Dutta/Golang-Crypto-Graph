"# Golang-Crypto-Graph" 


Date: 6/13/2018

Data Types
1. boolean
2. int
3. float32
4. float64
5. string
https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet

```go
package main

import ("fmt"
)

// var a int = 4
// var b float64 = 4.5
type Emp struct {
	id int
	name string
	dept string
}
func (e Emp) display(){ //Value Type
	e.name ="Abir"
	fmt.Println("inside display valuetype ", &e.name)
	//return e
}
func (e *Emp) display1(){ //Pointer Type
	e.name ="Abir1"
	fmt.Println("inside display ptrtype ", &e.name)
}

// func display (a int, b float64) (int,float64){
// 	return a,b
// }

func main() {
	// fmt.Printf("datatype %T",a)
	// fmt.Printf("datatype %T",b)

	// fmt.Printf("value %v",a)
	// fmt.Printf("value %v",b)
	e := Emp{id:456,dept:"IT"}
	fmt.Println(&e.name)
	e.display()
	e.display1()
	//fmt.Println(&e)
	fmt.Println(e)
}
```
