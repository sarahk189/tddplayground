package itempricecalculator_test

import (
	"math/rand"
	"testing"

	"github.com/sarahk189/tddplayground/itempricecalculator"
	"github.com/stretchr/testify/assert"
)

/*
Implement an item price calculator that computes the total price of a list of items based on their type and quantity.
TRUCK items costs 100.0 each and PARCEL items costs 25.0 each.

If a TRUCK item weighs more than 100kg, they cost an additional 50.0 per piece (quantity).
If a PARCEL item weighs more than 10kg, they cost an additional 25.0 per piece (quantity).

item := []itempricecalculator.Item{
		{
			Type:     "TRUCK",
			Quantity: 1,
			Weight:   105,
		},

		{
			Type:     "TRUCK",
			Quantity: 1,
			Weight:   50,
		},
	}

Items without weight information should be assumed to be outside the weight limit.

Non-existing item types should return an error.
*/

// 1. One arbitrary item should cost 100 - An item
// 2. Two arbitrary items should cost 200 - List of items
// 3. One Truck item, should cost 100 - One Truck Item
// 4. Three Truck items, should cost 300 - Three Truck Items
// 5. One Parcel item, should cost 25 - One Parcel Item
// 6. Quantity of items should be considered in the price calculation - One Parcel Item with quantity 2

func Test_CalculatePriceForItemsBasedOnNumberOfItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedPrice float64
	}{
		"One item": {
			items: []itempricecalculator.Item{
				{},
			},
			expectedPrice: 100.0,
		},

		"Two items": {
			items: []itempricecalculator.Item{
				{},
				{},
			},
			expectedPrice: 200.0,
		},
		"Three items": {
			items: []itempricecalculator.Item{
				{},
				{},
				{},
			},
			expectedPrice: 300.0,
		},
	}

	//ACT
	for testName, testCases := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			price := itemPriceCalculator.CalculatePrice(testCases.items)

			//ASSERT
			assert.Equal(t, testCases.expectedPrice, price)
		})

	}
}

func Test_CalculatePriceForRandomNumberOfItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{}

	numberOfItems := rand.Intn(20)
	for i := 0; i < numberOfItems; i++ {
		items = append(items, itempricecalculator.Item{})
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, float64(numberOfItems*100), price)
}
