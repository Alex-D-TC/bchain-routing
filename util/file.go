package util

import "io/ioutil"

func WriteToTemp(raw []byte) (string, error) {
	file, err := ioutil.TempFile("", "swissRelay")
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(raw)
	return file.Name(), nil
}
