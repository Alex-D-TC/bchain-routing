package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Contracts struct {
	Swiss string
	Relay string
}

func ReadContractsConfig(rawUrl string) (Contracts, error) {

	file, err := os.Open(rawUrl)
	if err != nil {
		return Contracts{}, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Contracts{}, err
	}

	var contracts Contracts
	err = json.Unmarshal(bytes, &contracts)
	return contracts, err
}
