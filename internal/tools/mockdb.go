package tools

import (
    "time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
    "rajat": {
        AuthToken: "123ABC",
        Username: "rajat",
    },
    "himanshu": {
        AuthToken: "456DEF",
        Username: "himanshu",
    },
}

var mockCoinDetails = map[string]CoinDetails{
    "rajat": {
        Coins: 100,
        Username: "rajat",
    },
    "himanshu": {
        Coins: 200,
        Username: "himanshu",
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
