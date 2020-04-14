package main

import "fmt"

func main() {
 var linea string
 fmt.Scanln(&linea)
 
//  bs := []byte(linea)
 fmt.Print(string(linea[0]))
 for i:= int(0); i < len(linea); i++ {
     if string(linea[i]) == "-" && i+1 != len(linea) {
         fmt.Print(string(linea[i+1]))
     }
 }
 
}
