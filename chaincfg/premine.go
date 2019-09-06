package chaincfg

type Premine struct {
	AddressString string
	Coins         float64
}

func Sum(premine []Premine) float64 {
	result := float64(0)
	for _, el := range premine {
		result = result + el.Coins
	}
	return result
}
