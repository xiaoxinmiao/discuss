package main

import (
	"fmt"
	"runtime"

	"github.com/pangpanglabs/goutils/echomiddleware"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
}

type Config struct {
	Logger struct {
		Kafka echomiddleware.KafkaConfig
	}
}
