package main

import (
	"bufio"
	"flag"
	"github.com/thorduri/pushover"
	"log"
	"os"
	"runtime"
)

func scan() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := NewConfig()

	flag.StringVar(&config.User, "u", config.User, "user name")
	flag.StringVar(&config.Token, "t", config.Token, "token")
	flag.StringVar(&config.Message, "m", config.Message, "message")
	flag.Parse()

	po, err := pushover.NewPushover(config.Token, config.User)
	if err != nil {
		log.Fatal(err)
	}

	if config.Message == "" {
		config.Message = scan()
	}

	err = po.Message(config.Message)
	if err != nil {
		log.Fatal(err)
	}
}
