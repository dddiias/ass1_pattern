package main

func main() {
    bike := &Bike{}
    car := &Car{}
    walk := &Walk{}

    context := &Context{}

    context.SetTraveler(bike)
    context.Travel()

    context.SetTraveler(car)
    context.Travel()

    context.SetTraveler(walk)
    context.Travel()
}