package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("----------")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var onConnect MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("--CONNECTED--")
}

var onConnectionLost MQTT.ConnectionLostHandler = func(client MQTT.Client, errmsg error) {
	fmt.Println("--DISCONNECTED--")
	fmt.Printf("error: %s\n", errmsg.Error())
}

var (
	serverAddrOpt = flag.String("s", "127.0.0.1", "MQTT server address")
	clientIDOpt   = flag.String("i", "sub", "MQTT client ID")
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
	opts.SetDefaultPublishHandler(f)
	opts.SetOnConnectHandler(onConnect)
	opts.SetConnectionLostHandler(onConnectionLost)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(*topicOpt, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(3600 * time.Second)

	if token := c.Unsubscribe(*topicOpt); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
