package connections

import (
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func StablishConnection(address, username, password string) error {
	err := downloadProfile(address, username, password)

	if err != nil {
		return nil
	}

	return nil
}

func downloadProfile(address, username, password string) error {
	url := fmt.Sprintf("http://%s/rest/GetUserlogin", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println("[Download Profile] Error creating request")
		return err
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Set("Authentication", "Basic "+basicAuth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("[Download Profile] Error on profile download")
		return err
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println("[Download Profile] Error on request body reading")
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf(
			"[Download Profile] Request error. status: %d, body: %s",
			res.StatusCode,
			resBody,
		)
	}

	err = os.WriteFile("./profiles/profile", resBody, 0644)

	if err != nil {
		log.Println("[Download Profile] Error writing file")
		return err
	}

	return nil
}
