package tools 

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"sebastian": {
		AuthToken: "vettel",
		Username: "sebastian",
	},
	"lewis": {
		AuthToken: "hamilton",
		Username: "lewis",
	},
	"daniel": {
		AuthToken: "ricciardo",
		Username: "daniel",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"sebastian": {
		Coins: 170,
		Username: "sebastian",
	},
	"lewis": {
		Coins: 210,
		Username: "lewis",
	},
	"daniel": {
		Coins: 150,
		Username: "daniel",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}