package main

import (
    "encoding/json"
    "os"
    "fmt"
	"flag"

	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	WebhookId string
	WebhookToken string
	IgnoreUsers []string
}

func main() {
	configFile := flag.String("config", "", "Location of the config file.")
	pamUser := os.Getenv("PAM_USER")
	pamType := os.Getenv("PAM_TYPE")
	pamRHost := os.Getenv("PAM_RHOST")
	hostname, _ := os.Hostname()
	payload := ""

	flag.Parse()

	if (len(pamUser) == 0 || len(pamType) == 0 || len(pamRHost) == 0) {
		fmt.Println("Required environmental variables are not set.")
		os.Exit(10)
	}

	if (len(*configFile) == 0) {
		fmt.Println("Config option is required")
		os.Exit(12)
	}

	if (pamType == "open_session") {
		payload = fmt.Sprintf("%s: %s logged in (remote host: %s).", hostname, pamUser, pamRHost)
	} else if (pamType == "close_session") {
		payload = fmt.Sprintf("%s: %s logged out (remote host: %s).", hostname, pamUser, pamRHost)
	} else {
		os.Exit(0)
	}

	// https://stackoverflow.com/a/16466189
	file, _ := os.Open(*configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	if (len(configuration.WebhookId) == 0 || len(configuration.WebhookToken) == 0) {
		fmt.Println("ssh-notify must be configured first.")
		os.Exit(11)
	}

	// https://stackoverflow.com/a/15323988
	for _, v := range configuration.IgnoreUsers {
		if v == pamUser {
			os.Exit(0)
		}
	}

	fmt.Println(payload)

	client := webhook.New(snowflake.MustParse(configuration.WebhookId), configuration.WebhookToken)
	_, err = client.CreateContent(payload)

	if err != nil {
		fmt.Println("error:", err)
	}
}