package models

type Bid struct {
	Price        int    `json:"price"`
	DeliveryTime int `json:"delivery_time"`
	Comments     string `json:"comments"`
}
