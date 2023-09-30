package main

import "fmt"

type Bike struct{}

func (b *Bike) Travel() {
    fmt.Println("Traveling by bike")
}