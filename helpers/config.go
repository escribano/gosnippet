package helpers

import (
	"errors"
	"fmt"
	"github.com/laurent22/toml-go/toml"
	"os"
)

const configFile string = "gosnippet.toml"

const configLocation string = "/etc/gosnippet/"

var Config toml.Document

func ParseConfig() error {

	f := configFile
	if FileExists(f) {
		fmt.Println("Using config file at: " + f)
	} else {
		f := configLocation + configFile
		if FileExists(f) {
			fmt.Println("Using config file at: " + f)
		} else {
			return errors.New("Config file not found at: " + f)
		}
	}

	var parser toml.Parser
	Config = parser.ParseFile(f)

	for _, key := range []string{
		"database.database",
		"database.username",
		"database.password",
	} {
		if Config.GetString(key) == "" {
			return errors.New(
				fmt.Sprintf(
					"helpers/config.go: Required config value not found for %s in %s",
					key,
					f,
				),
			)
		}
	}

	return nil
}

func FileExists(f string) bool {
	if _, err := os.Stat(f); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
