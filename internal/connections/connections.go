package connections

import (
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func StablishConnection(address, username, password string) error {
	_, err := downloadProfile(address, username, password)

	log.Println(err)

	if err != nil {
		return err
	}

	return nil
}

func downloadProfile(address, username, password string) (string, error) {
	rawUrl, err := url.Parse(address)

	if err != nil {
		log.Println("[Download Profile] Error on Address Parsing")
		return "", err
	}

	url := fmt.Sprintf("https://%s/rest/GetUserlogin", rawUrl.Host)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println("[Download Profile] Error creating request")
		return "", err
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Set("Authorization", "Basic "+basicAuth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("[Download Profile] Error on download request")
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println("[Download Profile] Error on request body reading")
		return "", err
	}

	if res.StatusCode > 299 {
		return "", fmt.Errorf(
			"[Download Profile] Request error. status: %d, body: %s",
			res.StatusCode,
			resBody,
		)
	}

	fileName := time.Now().String() + "-profile.ovpn"
	filePath := "./profiles/" + fileName
	err = os.WriteFile(filePath, resBody, 0644)

	if err != nil {
		log.Println("[Download Profile] Error writing file")
		return "", err
	}

	return fileName, nil
}
