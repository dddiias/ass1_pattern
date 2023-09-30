package main

import "fmt"

type Car struct{}

func (c *Car) Travel() {
    fmt.Println("Traveling by car")
}