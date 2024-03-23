// /go-workspace/src/myapp/main.go

package main

import "pkg" // Importing our custom package

func main() {
    pkg.SayHello() // Using function from the custom package
    println(pkg.SomeVariable) // Using variable from the custom package
}
