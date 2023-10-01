package main
import paquete1 "truco/paquete1"


import "fmt"

func cambiarPorReferencia(numero *int)int{
   *numero = *numero + 2 
   return  *numero
}

func main(){

    var mazo paquete1.Deck 

    cualquier := 20
    cambiarPorReferencia(&cualquier)
    println(cualquier)




    

    fmt.Println(mazo)

}
