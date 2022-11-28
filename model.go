package main

type Product struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Category    string `json:"category"`
}
