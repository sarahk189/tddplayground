package itempricecalculator_test

import (
	"errors"
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

func Test_CalculatePriceForTruckItemsBasedOnNumberOfItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedPrice float64
	}{
		"One item": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "TRUCK",
					Quantity: 1,
				},
			},
			expectedPrice: 100.0,
		},

		"Two items": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "TRUCK",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "TRUCK",
					Quantity: 1,
				},
			},
			expectedPrice: 200.0,
		},
		"Three items": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "TRUCK",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "TRUCK",
					Quantity: 1,
				},
				{
					ID:       "ART9012",
					Type:     "TRUCK",
					Quantity: 1,
				},
			},
			expectedPrice: 300.0,
		},
	}

	//ACT
	for testName, testCases := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			price, _ := itemPriceCalculator.CalculatePrice(testCases.items)

			//ASSERT
			assert.Equal(t, testCases.expectedPrice, price)
		})

	}
}

func Test_CalculatePriceForRandomNumberOfTruckItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{}

	numberOfItems := rand.Intn(20)
	for i := 0; i < numberOfItems; i++ {
		items = append(items, itempricecalculator.Item{
			Type: "TRUCK",
		})
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, float64(numberOfItems*100), price)
}

func Test_CalculatePriceForOneParcelItem(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	item := []itempricecalculator.Item{
		{
			ID:       "ART1234",
			Type:     "PARCEL",
			Quantity: 1,
		},
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(item)

	//ASSERT
	assert.Equal(t, 25.0, price)
}

func Test_CalculatePriceForTwoParcelItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	item := []itempricecalculator.Item{
		{
			ID:       "ART1234",
			Type:     "PARCEL",
			Quantity: 1,
		},
		{
			ID:       "ART5678",
			Type:     "PARCEL",
			Quantity: 1,
		},
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(item)

	//ASSERT
	assert.Equal(t, 50.0, price)
}

func Test_CalculatePriceForRandomNumberOfParcels(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{}

	numberOfItems := rand.Intn(20)
	for i := 0; i < numberOfItems; i++ {
		items = append(items, itempricecalculator.Item{
			Type: "PARCEL",
		})
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, float64(numberOfItems*25), price)

}

func Test_CalculatePriceForRandomAmountOfParcelAndTruckItems(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{}

	numberOfTruckItems := rand.Intn(10)
	for i := 0; i < numberOfTruckItems; i++ {
		items = append(items, itempricecalculator.Item{
			Type: "TRUCK",
		})
	}

	numberOfParcelItems := rand.Intn(10)
	for i := 0; i < numberOfParcelItems; i++ {
		items = append(items, itempricecalculator.Item{
			Type: "PARCEL",
		})
	}

	expectedPrice := float64(numberOfTruckItems*100 + numberOfParcelItems*25)

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, expectedPrice, price)
}

func Test_CalculatePriceForTypoInItemType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedPrice float64
	}{
		"Test mixed casing Truck": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "trUCK",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "TRuCK",
					Quantity: 1,
				},
				{
					ID:       "ART9012",
					Type:     "Truck",
					Quantity: 1,
				},
			},
			expectedPrice: 300.0,
		},

		"Test mixed casing Parcel": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "paRcel",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "Parcel",
					Quantity: 1,
				},
				{
					ID:       "ART9012",
					Type:     "PARcel",
					Quantity: 1,
				},
			},
			expectedPrice: 75.0,
		},
	}

	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			price, _ := itemPriceCalculator.CalculatePrice(testCase.items)

			assert.Equal(t, testCase.expectedPrice, price)
		})
	}
}

func Test_ShouldReturnErrorForInvalidItemTypes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedError error
	}{
		"Test first item is invalid item type": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "DUCK",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "TRUCK",
					Quantity: 1,
				},
				{
					ID:       "ART9012",
					Type:     "PARCEL",
					Quantity: 1,
				},
			},
			expectedError: errors.New("invalid item type: DUCK"),
		},

		"Test second item is invalid item type": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "PARCEL",
					Quantity: 1,
				},
				{
					ID:       "ART5678",
					Type:     "GOOSE",
					Quantity: 1,
				},
				{
					ID:       "",
					Type:     "TRUCK",
					Quantity: 1,
				},
			},
			expectedError: errors.New("invalid item type: GOOSE"),
		},
	}

	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			_, err := itemPriceCalculator.CalculatePrice(testCase.items)

			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func Test_ItemIsAParcelOrTruckItem(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedPrice float64
		expectedError error
	}{
		"Test item is a truck item": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "TRUCK",
					Quantity: 1,
				},
			},
			expectedPrice: 100.0,
			expectedError: nil,
		},

		"Test item is a parcel item": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "PARCEL",
					Quantity: 1,
				},
			},
			expectedPrice: 25.0,
			expectedError: nil,
		},

		"Test item is an invalid item type one": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1232",
					Type:     "BICYCLE",
					Quantity: 1,
				},
			},
			expectedPrice: 0.0,
			expectedError: errors.New("invalid item type: BICYCLE"),
		},

		"Test item is an invalid item type two": {
			items: []itempricecalculator.Item{
				{
					ID:       "ART1234",
					Type:     "RandomItemType",
					Quantity: 1,
				},
			},
			expectedPrice: 0.0,
			expectedError: errors.New("invalid item type: RandomItemType"),
		},
	}

	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			price, err := itemPriceCalculator.CalculatePrice(testCase.items)

			assert.Equal(t, testCase.expectedPrice, price)
			assert.Equal(t, testCase.expectedError, err)

		})
	}
}

func Test_CalculatePriceForItemsWithQuantityInformationForTruck(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			ID:       "ART1234",
			Type:     "TRUCK",
			Quantity: 5,
		},
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 500.0, price)
}

func Test_CalculatePriceForItemsWithQuantityInformationForParcel(t *testing.T) {
	t.Parallel()

	//ARRANGE
	itemPriceCalculator := itempricecalculator.NewItemPriceCalculator()
	items := []itempricecalculator.Item{
		{
			ID:       "ART1234",
			Type:     "PARCEL",
			Quantity: 5,
		},
	}

	//ACT
	price, _ := itemPriceCalculator.CalculatePrice(items)

	//ASSERT
	assert.Equal(t, 125.0, price)
}
