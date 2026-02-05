package itempricecalculator

import (
	"fmt"
	"strings"
)

// Implement an item price calculator that computes the total price of a list of items based on their type and quantity.
// TRUCK items costs 100.0 each and PARCEL items costs 25.0 each.
//
// If TRUCK items weigh more than 100kg, they cost an additional 50.0 per piece.
// If PARCEL items weigh more than 10kg, they cost an additional 25.0 per piece.
//
// Items without weight information should be assumed to be outside the weight limit.
//
// Non-existing item types should return an error.
type Item struct {
	ID       string
	Type     string
	Quantity int
}

type ItemPriceCalculator struct {
}

func NewItemPriceCalculator() ItemPriceCalculator {
	return ItemPriceCalculator{}
}

func (i *ItemPriceCalculator) CalculatePrice(items []Item) (float64, error) {
	price := 0.0

	for _, item := range items {
		itemType := strings.ToUpper(item.Type)

		if itemType == "PARCEL" {
			price += 25.0
		} else if itemType == "TRUCK" {
			price += 100.0
		} else {
			return 0.0, fmt.Errorf("invalid item type: %s", item.Type)
		}
	}

	return price, nil
}
