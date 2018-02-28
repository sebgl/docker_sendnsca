package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sebgl/nscacheck"
)

var (
	host    = flag.String("host", "", "Hostname for the check payload")
	service = flag.String("service", "", "Service name for the check payload")
	state   = flag.Int("state", 0, "State for the check payload (eg. 0, 1, 2)")
	msg     = flag.String("msg", "", "Message for the check payload")
)

func nscaSettingsFromEnv() nscacheck.NSCASettings {
	server := os.Getenv("NSCA_SERVER")
	key := os.Getenv("NSCA_KEY")
	port, err := strconv.Atoi(os.Getenv("NSCA_PORT"))
	if err != nil {
		log.Fatal("Invalid $NSCA_PORT")
	}
	encryption, err := strconv.Atoi(os.Getenv("NSCA_ENCRYPTION"))
	if err != nil {
		log.Fatal("Invalid $NSCA_ENCRYPTION")
	}
	return nscacheck.NSCASettings{
		Server:     server,
		Port:       port,
		Key:        key,
		Encryption: encryption,
	}
}

func main() {
	flag.Parse()
	nscaSettings := nscaSettingsFromEnv()
	err := nscacheck.Send(nscaSettings, *host, *service, *state, *msg)
	if err != nil {
		fmt.Println(err)
	}
}
