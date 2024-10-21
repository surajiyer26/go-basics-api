package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"suraj": {
		AuthToken: "123suraj",
		Username:  "suraj",
	},
	"pooraj": {
		AuthToken: "123pooraj",
		Username:  "pooraj",
	},
	"coolraj": {
		AuthToken: "123coolraj",
		Username:  "coolraj",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"suraj": {
		Coins:    100,
		Username: "suraj",
	},
	"pooraj": {
		Coins:    200,
		Username: "pooraj",
	},
	"coolraj": {
		Coins:    300,
		Username: "coolraj",
	},
}

func (d *mockDB) GetUserLoginDatabase(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoinDatabase(username string) *CoinDetails {
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
