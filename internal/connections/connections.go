package connections

import (
	"crypto/tls"
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
)

func StablishConnection(address, username, password string) error {
	_, err := downloadProfile(address, username, password)

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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(req)

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

	fileName := fmt.Sprintf("%d-profile.ovpn", time.Now().UnixMilli())
	filePath := "~/.openvpn-profiles/" + fileName

	err = os.Mkdir("~/.openvpn-profiles", 0766)

	if err != nil {
		log.Println("[Download Profile] Error creating profiles folder")
		return "", err
	}

	err = os.WriteFile(filePath, resBody, 0766)

	if err != nil {
		log.Println("[Download Profile] Error writing file")
		return "", err
	}

	return fileName, nil
}

func startOpenvpn(fileName string) {
	_ = exec.Command("openvpn3", "session-start", "-c", "~/.openvpn-profiles/"+fileName)
}
