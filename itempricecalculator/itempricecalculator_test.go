package itempricecalculator_test

import (
	"testing"

	"github.com/sarahk189/tddplayground/itempricecalculator"
	"github.com/stretchr/testify/assert"
)

/*
Implement an item price calculator that computes the total price of a list of items based on their type and quantity.
TRUCK items costs 100.0 each and PARCEL items costs 25.0 each.

If TRUCK items weigh more than 100kg, they cost an additional 50.0 per piece.
If PARCEL items weigh more than 10kg, they cost an additional 25.0 per piece.

Items without weight information should be assumed to be outside the weight limit.

Non-existing item types should return an error.
*/

// 1. One arbitrary item should cost 100 - An item
// 2. Two arbitrary items should cost 200 - List of items
// 3. One Truck item, should cost 100 - One Truck Item
// 4. Three Truck items, should cost 300 - Three Truck Items
// 5. One Parcel item, should cost 25 - One Parcel Item
// 6. Quantity of items should be considered in the price calculation - One Parcel Item with quantity 2

func Test_TruckItemShouldCost100(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	item := []itempricecalculator.Item{
		{
			Type:     "TRUCK",
			Quantity: 1,
		},
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(item)

	//ASSERT
	assert.Equal(t, 100.0, price)
}

func Test_TwoTruckItemsShouldCost200(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			Type:     "TRUCK",
			Quantity: 1,
		},
		{
			Type:     "TRUCK",
			Quantity: 1,
		},
	}
	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 200.0, price)
}

func Test_OneTruckItemShouldCost100(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			Type:     "TRUCK",
			Quantity: 1,
		},
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 100.0, price)
}

func Test_ThreeTruckItemsShouldCost300(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			Type:     "TRUCK",
			Quantity: 1,
		},
		{
			Type:     "TRUCK",
			Quantity: 1,
		}, {
			Type:     "TRUCK",
			Quantity: 1,
		},
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 300.0, price)
}

func Test_OneParcelItemShouldCost25(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			Type:     "PARCEL",
			Quantity: 1,
		},
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 25.0, price)
}

func Test_OneParcelItemWithQuantity2ShouldCost50(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			Type:     "PARCEL",
			Quantity: 2,
		},
	}

	//ACT
	price := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 50.0, price)
}
