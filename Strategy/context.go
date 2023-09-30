package main

type Traveler interface {
    Travel()
}

type Context struct {
    traveler Traveler
}

func (c *Context) SetTraveler(t Traveler) {
    c.traveler = t
}

func (c *Context) Travel() {
    c.traveler.Travel()
}