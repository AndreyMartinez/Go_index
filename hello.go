package main

import "fmt"

/**
* función para entrenar funcionamiento de las mismas en go
*/
func add(a , b int) int{
  return a + b 
}

/**
* función para entrenar funcionamiento de una función con retorno de dos valores
*/
func golandTwoFunciton(a , b string)(string,string){
	return a,b
}

/**
*función para entranar funcionamiento de la misma retornando valores con nombre
*/
func nameReturn(x,b string)(c string){
	  c = x + b 
	  return
}

/**
*función para entrenar manejo de variables 
*/
func varReturnValue()(string){
firstHello := "hola"
fmt.Println(firstHello)
return firstHello
}

/**
*función para manejo de for dentro de Go
*/
func forContent()(a int){
	a = 1
	for ; a < 100; {
		a += 1
	}
	return 
}

func main() {
	fmt.Println("hello, \n",add(1,2))
	a,b := golandTwoFunciton("hola","mundo")
	fmt.Println(a,b)
	fmt.Println(nameReturn("hola","mundo"))
	varReturnValue()
	fmt.Println(forContent())
}