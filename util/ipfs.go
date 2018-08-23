package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func IPFSReadFile(addr string) ([]byte, error) {

	resp, err := http.Get(fmt.Sprintf("http://localhost:5001/api/v0/cat?arg=", addr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	return rawJSON, err
}

func IPFSAddFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}

	io.Copy(part, file)
	writer.Close()

	resp, err := http.Post("http://localhost:5001/api/v0/add", writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(result), nil
}
