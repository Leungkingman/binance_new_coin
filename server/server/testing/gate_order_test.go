package testing

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/service"
	"fmt"
	"testing"
)

func TestAddOrder(t *testing.T) {
	service.InitDatabase()
	o := &model.DB_Gate_Order {
		Order_id: "76939752188",
		Coin: "GALA",
		Price: "0.02494",
		Amount: "120271.81",
		Total: "2925.0104192",
		Left: "0",
		Create_time: "1631498451",
	}
	err := global.DB_ENGINE.Create(o).Error
	fmt.Printf("err = %v \n", err)
	//fmt.Println("Before all tests")
}