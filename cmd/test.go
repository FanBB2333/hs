package main

import (
    "fmt"
    "os"
)

func main() {
    // get all the environment variables
    envVars := os.Environ()
    for _, env := range envVars {
        fmt.Println(env)
    }
}
