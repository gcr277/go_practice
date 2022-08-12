package main

import (
    "fmt"
    "test/pac2"
    "test/pac3"
    "test/pac1"
    "test/info"
)
func init(){
    fmt.Printf("[%v]\n", info.CurrFuncName())
}
func main(){
    fmt.Printf("[%v]%v|%v|%v\n", info.CurrFuncName(),pac1.NUM1,pac2.NUM2,pac3.NUM3)
}