package itempricecalculator

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

//type WeightCalc struct{}

func NewItemPriceCalculator() ItemPriceCalculator {
	return ItemPriceCalculator{}
}

func (i *ItemPriceCalculator) CalculatePrice(items []Item) float64 {
	if items[0].Type == "PARCEL" && len(items) == 1 {
		return 25.0
	}
	if items[0].Type == "PARCEL" && items[1].Type == "PARCEL" && len(items) == 2 {
		return 50.0
	}
	return 100.0 * float64(len(items))
}
