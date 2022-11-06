package main

import "fmt"
// Import the math/rand package
import "math/rand"

func main() {
    
    for i := 1; i <= 5; i++ {
        fmt.Println(rand.Intn(10))
        
    }
}
