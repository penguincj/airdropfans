package market

import (
	"fmt"
	"time"

	"airdrop/market/models"
	"airdrop/market/services"
)

const EPSINON = 0.0000001

var tokenList = map[int64]string{
	0: "btcusdt",
	1: "ethusdt",
	2: "eosusdt",
	3: "htusdt",
	4: "ltcusdt",
	5: "xrpusdt",
	6: "dashusdt",
	7: "etcusdt",
}

var tokenPriceMap map[string]models.TokenPrice

func getTokenPrice(name string) models.TokenPrice {

	fmt.Println("name", name)

	htKLine := services.GetKLine(name, "1day", 1)
	if htKLine.Status != "ok" {
		fmt.Printf("get %s failed", name)
		return models.TokenPrice{}
	}

	kLineData := htKLine.Data[0]

	token := models.TokenPrice{
		Name:  name,
		Open:  kLineData.Open,
		Close: kLineData.Close,
	}

	if (token.Open >= -EPSINON) && (token.Open <= EPSINON) {
		fmt.Println("open price is zero")
		return models.TokenPrice{}
	}

	floatPercent := ((token.Close - token.Open) / token.Open) * 100
	floatPercentS := fmt.Sprintf("%0.2f", floatPercent)
	token.FloatPercent = floatPercentS
	fmt.Printf("%s float %0.2f \n", name, floatPercent)

	return token
}

func RunMarket() {

	tokenPriceMap = make(map[string]models.TokenPrice, 8)

	t := time.NewTimer(1 * time.Second)


	for {
		select {
		case <-t.C:
			t.Reset(1 * time.Minute)
			fmt.Println("111111")
			for _, tokenName := range tokenList {
				token := getTokenPrice(tokenName)
				tokenPriceMap[tokenName] = token
			}

			fmt.Println("cj")
			for name, tokenPrice := range tokenPriceMap {
				fmt.Printf("%s float %s \n", name, tokenPrice.FloatPercent)
			}
		}
	}
}
