// data-simulator/main.go
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type DataMessage struct {
    DeviceID  int     `json:"device_id"`
    FieldA    float64 `json:"field_a"`
    FieldB    float64 `json:"field_b"`
    RuleField int     `json:"rule_fld"`
}

func main() {
    var devices int
    var rate int
    var controllerURL string

    flag.IntVar(&devices, "devices", 5, "Number of devices to simulate")
    flag.IntVar(&rate, "rate", 1, "Number of messages per second per device")
    flag.StringVar(&controllerURL, "controller", "http://localhost:8081/data", "URL of the IoT Controller")
    flag.Parse()

    rand.Seed(time.Now().UnixNano())

    fmt.Printf("Simulating %d devices, each sending %d msgs/s to %s\n", devices, rate, controllerURL)

    for i := 0; i < devices; i++ {
        go simulateDevice(i, rate, controllerURL)
    }

    // Чтобы приложение не завершалось
    select {}
}

func simulateDevice(deviceID, rate int, controllerURL string) {
    ticker := time.NewTicker(time.Second / time.Duration(rate))
    defer ticker.Stop()

    for {
        <-ticker.C

        // Генерируем случайное сообщение
        msg := DataMessage{
            DeviceID:  deviceID,
            FieldA:    rand.Float64() * 10,
            FieldB:    rand.Float64() * 10,
            RuleField: rand.Intn(100),
        }

        // Преобразуем в JSON
        data, err := json.Marshal(msg)
        if err != nil {
            log.Printf("Error marshaling JSON: %v", err)
            continue
        }

        // Отправляем POST-запрос в IoT Controller
        resp, err := http.Post(controllerURL, "application/json", bytes.NewBuffer(data))
        if err != nil {
            log.Printf("Error sending request from device %d: %v", deviceID, err)
            continue
        }
        resp.Body.Close()
    }
}
