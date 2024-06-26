package testing

import (
	"binanceNewCoin/server/tool"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetRandom(t *testing.T) {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(4)
	fmt.Printf("result = " + tool.IntToString(result))
}
