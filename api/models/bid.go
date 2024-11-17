package models

type Bid struct {
	Price        int    `json:"price"`
	DeliveryTime string `json:"delivery_time"`
	Comments     string `json:"comments"`
}
