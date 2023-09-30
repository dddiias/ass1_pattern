package main

import "fmt"

type Walk struct{}

func (w *Walk) Travel() {
    fmt.Println("Traveling on foot")
}