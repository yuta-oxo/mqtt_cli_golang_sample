package main

import (
	"flag"
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var onConnect MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("--CONNECTED--")
}

var onConnectionLost MQTT.ConnectionLostHandler = func(client MQTT.Client, errmsg error) {
	fmt.Println("--DISCONNECTED--")
	fmt.Printf("error: %s\n", errmsg.Error())
}

var (
	serverAddrOpt = flag.String("s", "127.0.0.1", "MQTT server address")
	clientIDOpt   = flag.String("i", "pub", "MQTT client ID")
	topicOpt      = flag.String("t", "topic", "MQTT topic")
)

func main() {
	flag.Parse()

	fmt.Printf("SERVER: %s\n", *serverAddrOpt)
	fmt.Printf("Clinet ID: %s\n", *clientIDOpt)
	fmt.Printf("TOPIC: %s\n", *topicOpt)

	broker := "tcp://" + *serverAddrOpt + ":1883"
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(*clientIDOpt)
	opts.SetOnConnectHandler(onConnect)
	opts.SetConnectionLostHandler(onConnectionLost)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	messageFmt := "msg #%d from " + *clientIDOpt + " (%s)"
	for i := 0; i < 3600; i++ {
		t := time.Now()
		text := fmt.Sprintf(messageFmt, i, t)
		fmt.Println(text)
		token := c.Publish(*topicOpt, 1, false, text)
		token.Wait()
		time.Sleep(1000 * time.Millisecond)
	}
	c.Disconnect(250)
}
