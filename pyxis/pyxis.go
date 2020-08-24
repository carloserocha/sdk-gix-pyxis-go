package pyxis

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const DEFAULT_URL = "shx-integrador/srv/"

type Client struct {
	Username string
	Password string
	URL      string
}

var context = &Client{
	Username: "",
	Password: "",
	URL:      "",
}

// new client user
func NewClient(customURL string) {
	context.URL = customURL
}

// send request
func NewRequest(op, url string, data []byte) []byte {
	timeout := time.Duration(20 * time.Second)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Timeout:   timeout,
		Transport: tr, // ignore self-certification
	}

	reqURI := context.URL + DEFAULT_URL + url

	authorization := handleAuthorization(context.Username, context.Password)

	request, _ := http.NewRequest(http.MethodPost, reqURI, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func handleAuthorization(username string, password string) string {
	encryptedLogin := fmt.Sprintf("%s:%s", username, password)
	str := "Basic " + base64.StdEncoding.EncodeToString([]byte(encryptedLogin)) // arrumar essa porcaria.

	return str
}
