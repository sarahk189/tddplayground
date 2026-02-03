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

func Test_TruckItemShouldCost100(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	//ACT
	price := itemPriceCalculator.CalculatePrice()

	//ASSERT
	assert.Equal(t, 100.0, price)
}
