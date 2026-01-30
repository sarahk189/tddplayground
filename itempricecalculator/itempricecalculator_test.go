package itempricecalculator_test

import (
	"fmt"
	"math/rand"
	"slices"
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

//func TestItemPriceCalculator_OneTruckItemShouldCost100(t *testing.T) {
//	t.Parallel()
//
//	calculator := itempricecalculator.NewItemPriceCalculator()
//
//	items := []itempricecalculator.Item{
//		{
//			Type: "TRUCK",
//		},
//	}
//
//	res, _ := calculator.CalculatePrice(items)
//
//	assert.Equal(t, 100.00, res)
//}
//
//func TestItemPriceCalculator_TwoTruckItemsShouldCost200(t *testing.T) {
//	t.Parallel()
//
//	calculator := itempricecalculator.NewItemPriceCalculator()
//
//	items := []itempricecalculator.Item{
//		{
//			Type: "TRUCK",
//		}, {
//			Type: "TRUCK",
//		},
//	}
//
//	res, _ := calculator.CalculatePrice(items)
//
//	assert.Equal(t, 200.00, res)
//}

type MockWeightService struct {
	wasCalled           bool
	numberOfInvocations int
	itemsIDs            []string
	weightToReturn      float64

	itemWeights map[string]float64
}

func (m *MockWeightService) GetWeight(itemID string) float64 {
	m.wasCalled = true
	m.numberOfInvocations++

	m.itemsIDs = append(m.itemsIDs, itemID)

	return m.itemWeights[itemID]
}

func TestName_ShouldAdd50ToTruckItemWeighingMoreThan100KG(t *testing.T) {
	t.Parallel()

	// Arrange
	service := &MockWeightService{
		itemWeights: map[string]float64{
			"item1": 175.0,
		},
	}

	calculator := itempricecalculator.NewItemPriceCalculator(service)

	items := []itempricecalculator.Item{
		{
			ID:   "item1",
			Type: "TRUCK",
		},
	}

	// Act
	res, _ := calculator.CalculatePrice(items)

	// Assert
	assert.Equal(t, 150.00, res)
}

func TestItemPriceCalculator_ShouldCallWeightServiceAllProvidedItems(t *testing.T) {
	t.Parallel()

	// Arrange
	service := &MockWeightService{}
	calculator := itempricecalculator.NewItemPriceCalculator(service)

	items := make([]itempricecalculator.Item, 0)
	for i := 1; i <= 2; i++ {
		items = append(items, itempricecalculator.Item{
			ID:   fmt.Sprintf("item%d", rand.Intn(10)),
			Type: "TRUCK",
		})
	}

	// Act
	_, _ = calculator.CalculatePrice(items)

	// Assert
	assert.True(t, slices.Contains(service.itemsIDs, items[0].ID))
	assert.True(t, slices.Contains(service.itemsIDs, items[1].ID))
}

func TestItemPriceCalculator_ShouldCallWeightServiceWhenCalculatingPrice(t *testing.T) {
	t.Parallel()

	// Arrange
	service := &MockWeightService{}
	calculator := itempricecalculator.NewItemPriceCalculator(service)

	items := []itempricecalculator.Item{
		{
			Type: "TRUCK",
		},
	}

	// Act
	_, _ = calculator.CalculatePrice(items)

	// Assert
	assert.True(t, service.wasCalled)
}

func TestItemPriceCalculator_ShouldReturnCorrectTruckArticlePrice(t *testing.T) {
	t.Parallel()

	// Arrange
	service := &MockWeightService{}
	calculator := itempricecalculator.NewItemPriceCalculator(service)

	items := make([]itempricecalculator.Item, 0)
	numberOfItems := rand.Intn(10)

	for i := 0; i < numberOfItems; i++ {
		items = append(items, itempricecalculator.Item{
			Type: "TRUCK",
		})
	}

	// Act
	res, _ := calculator.CalculatePrice(items)

	// Assert
	assert.Equal(t, float64(numberOfItems)*100.00, res)
}

func TestNameFooBob2(t *testing.T) {
	t.Parallel()

	// Arrange
	testCases := map[string]struct {
		items         []itempricecalculator.Item
		expectedPrice float64
	}{
		"One TRUCK item should cost 100.00": {
			items: []itempricecalculator.Item{
				{Type: "TRUCK"},
			},
			expectedPrice: 100.00,
		},
		"Two TRUCK items should cost 200.00": {
			items: []itempricecalculator.Item{
				{Type: "TRUCK"},
				{Type: "TRUCK"},
			},
			expectedPrice: 200.00,
		},
		"Three TRUCK items should cost 200.00": {
			items: []itempricecalculator.Item{
				{Type: "TRUCK"},
				{Type: "TRUCK"},
				{Type: "TRUCK"},
			},
			expectedPrice: 300.00,
		},
	}

	service := &MockWeightService{}
	calculator := itempricecalculator.NewItemPriceCalculator(service)

	// Act
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			res, _ := calculator.CalculatePrice(testCase.items)

			// Assert
			assert.Equal(t, testCase.expectedPrice, res)
		})
	}
}
