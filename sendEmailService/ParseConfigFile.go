package sendEmailService

import (
	"io/ioutil"
	"os"
	"strings"
)

func ParseConfigFile() (map[string]string, error)  {
	emailAndPath := make(map[string]string)
	configFile, err := ioutil.ReadFile("./config")
	if err != nil {
		os.Create("./config")
		return nil, err
	}

	configLines := strings.Split(string(configFile), "\n")
	for i := 0; i < len(configLines); i++ {
		if (configLines[i] != "") {
			configLine := strings.Split(string(configLines[i]), ";")
			emailAndPath[configLine[0]] = configLine[1]
		}
	}
	return emailAndPath, nil
}
