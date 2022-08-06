package main

endpoint := fmt.Sprintf("%s/%s@kline_%s", getWsEndpoint(), strings.ToLower(symbol), interval)
cfg := newWsConfig(endpoint)
c, _, err := websocket.DefaultDialer.Dial(cfg.Endpoint, nil)
if err != nil {
	return nil, nil, err
}
c.SetReadLimit(655350)
doneC = make(chan struct{}) /// may be it come from struct{}
stopC = make(chan struct{})
go func() {
	// This function will exit either on error from
	// websocket.Conn.ReadMessage or when the stopC channel is
	// closed by the client.
	defer close(doneC)
	if WebsocketKeepalive {
		keepAlive(c, WebsocketTimeout)
	}
	// Wait for the stopC channel to be closed.  We do that in a
	// separate goroutine because ReadMessage is a blocking
	// operation.
	silent := false
	go func() {
		select {
		case <-stopC:
			silent = true
		case <-doneC:
		}
		c.Close()
	}()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if !silent {
				func(err){
					fmt.Println(err)
				}
			}
			return
		}
		func(message){
			fmt.Println(event)
		}
	}
}()
