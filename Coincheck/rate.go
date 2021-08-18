package main

type Rate struct {
	client *CoinCheck
}

// Get the Rate of Buy Coin
func (r Rate)buy(pair string) string {
	return r.client.Request("GET", "api/rate/"+pair, "")

}
