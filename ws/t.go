package main
import (
	"fmt"
	"github.com/leek-box/sheep/util"
	"encoding/json"
)


// Endpoint 行情的Websocket入口
var Endpoint = "wss://real.okex.com:10441/websocket"


func main() {
	fmt.Println("connecting")
	fmt.Println(Endpoint)
	ws, err := util.NewSafeWebSocket(Endpoint)
	if err != nil {
		return 
	}		
	data := "{'event':'addChannel','channel':'ok_sub_spot_bch_btc_ticker' }"
	b, err := json.Marshal(data)
	ws.Send(b)
	// result := ws.Listen(r []byte)
    // fmt.Printf(result)
}
