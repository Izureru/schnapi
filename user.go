package main

import "time"

type User struct {
	Name     string    `json:"name"`
	Wardrobe []Clothes `json:"wardrobe"`
}

type Users []User

type Clothes struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Image    string    `json:"image"`
	Colour   string    `json:"colour"`
	Lastworn time.Time `json:"lastworn"`
	Wearing  bool      `json:"wearing"`
}
