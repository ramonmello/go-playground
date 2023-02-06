package main

import "fmt"

type Address struct {
	Street string
	Number int
	city   string
	state  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

// Only methods
type StatusHandler interface {
	Disable()
}

func (c *Client) Disable() {
	c.Active = false
}

func HandleDisable(userStatus StatusHandler) {
	userStatus.Disable()
}

func main() {
	user := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	// user.Disable()
	HandleDisable(&user)

	fmt.Printf("Name: %s, Age: %d, Active: %t", user.Name, user.Age, user.Active)
}
