package main

import (
       "os"
)


func main() {
     file, err := os.Open("filename_test")
     if err != nil {
     	panic(err)
     }
     defer file.Close()
}
