package main

type Navigator struct{}

func (n *Navigator) Navigate(t Traveler) {
    context := &Context{}
    context.SetTraveler(t)
    context.Travel()
}