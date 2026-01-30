package itempricecalculator

type Item struct {
	ID       string
	Type     string
	Quantity int
}

//type WeightService struct {
//}
//
//func (w *WeightService) GetWeight(itemID string) float64 {
//	return 0.0
//}

//func NewWeightService() *WeightService {
//	return &WeightService{}
//}

type ItemWeightProvider interface {
	GetWeight(itemID string) float64
}

type ItemPriceCalculator struct {
	service ItemWeightProvider
}

func NewItemPriceCalculator(service ItemWeightProvider) *ItemPriceCalculator {
	return &ItemPriceCalculator{
		service: service,
	}
}

func (i *ItemPriceCalculator) CalculatePrice(items []Item) (float64, error) {
	totalPrice := 0.0

	for _, item := range items {
		weight := i.service.GetWeight(item.ID)
		//if item.Type == "TRUCK" {
		totalPrice += 100.0
		if weight > 100.0 {
			totalPrice += 50.0
		}
		//} else if item.Type == "PARCEL" {
		//	totalPrice += 25.0
		//}
	}

	return totalPrice, nil
}
