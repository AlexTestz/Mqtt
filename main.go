package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := MQTT.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")
    opts.SetClientID("go_mqtt_client")
    opts.SetCleanSession(true)

    client := MQTT.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    fmt.Println("Conectado al broker MQTT")

    token := client.Publish("test/topic", 0, false, "Hola Mundo")
    token.Wait()

    fmt.Println("Mensaje publicado: Hola Mundo")

    time.Sleep(1 * time.Second)

    client.Disconnect(250)
    fmt.Println("Desconectado del broker MQTT")
}
