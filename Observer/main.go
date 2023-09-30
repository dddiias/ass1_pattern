package main

import "fmt"

func main() {
	// Create a new order system.
	orderSystem := NewOrderSystem()

	// Create some customers.
	customer1 := &Customer{id: "1"}
	customer2 := &Customer{id: "2"}

	// Register the customers as observers.
	orderSystem.RegisterObserver(customer1)
	orderSystem.RegisterObserver(customer2)

	// Create some items.
	item1 := &Item{name: "item1", inStock: true}
	item2 := &Item{name: "item2", inStock: true}

	// Place some orders.
	orderSystem.PlaceOrder(customer1, item1)
	orderSystem.PlaceOrder(customer2, item2)

	// Remove one of the customers as an observer.
	orderSystem.RemoveObserver(customer2)

	// Place another order.
	orderSystem.PlaceOrder(customer1, item2)
}

// OrderSystem represents the order system that customers can place orders on.
type OrderSystem struct {
	orders    []*Order
	observers []Observer
}

// NewOrderSystem creates a new OrderSystem.
func NewOrderSystem() *OrderSystem {
	return &OrderSystem{}
}

// RegisterObserver registers an observer to the order system.
func (os *OrderSystem) RegisterObserver(observer Observer) {
	os.observers = append(os.observers, observer)
}

// RemoveObserver removes an observer from the order system.
func (os *OrderSystem) RemoveObserver(observer Observer) {
	for i, o := range os.observers {
		if o == observer {
			os.observers = append(os.observers[:i], os.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies all observers of the order system.
func (os *OrderSystem) NotifyObservers() {
	for _, observer := range os.observers {
		observer.update(os)
	}
}

// PlaceOrder places an order with the given customer and item.
func (os *OrderSystem) PlaceOrder(customer *Customer, item *Item) {
	order := &Order{
		Customer: customer,
		Item:     item,
	}
	os.orders = append(os.orders, order)
	os.NotifyObservers()
}

// Order represents an order with a customer and item.
type Order struct {
	Customer *Customer
	Item     *Item
}

// Update sends an email to the customer for the given order.
func (c *Customer) Update(subject Subject) {
	if os, ok := subject.(*OrderSystem); ok {
		for _, order := range os.orders {
			if order.Customer == c {
				fmt.Printf("Sending email to customer %s for item %s\n", c.getID(), order.Item.name)
			}
		}
	}
}
