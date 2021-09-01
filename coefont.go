package coefontcloud

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const URL = "https://api.coefont.cloud/text2speech"

func NewParam() *coeFontParameter {
	ret := coeFontParameter{}
	ret.Pitch = 0
	ret.Kuten = 0.7
	ret.Speed = 1.0
	ret.Toten = 0.4
	ret.Volume = 1.0
	ret.Intonation = 1.0

	return &ret
}

func validate(p *coeFontParameter) bool {
	if p.CoeFont == "" {
		return false
	}

	if p.ClientSecret == "" {
		return false
	}

	if p.Text == "" {
		return false
	}

	if p.Accesskey == "" {
		return false
	}

	return true
}

func CallCoeFont(p *coeFontParameter) ([]byte, error) {
	if !validate(p) {
		return nil, errors.New("The requied paramter should not be empty.")
	}

	reqBody := CoeFontReqestBody{}

	reqBody.Accesskey = p.Accesskey
	reqBody.CoeFont = p.CoeFont
	reqBody.Text = p.Text

	mac := hmac.New(sha256.New, []byte(p.ClientSecret))
	mac.Write([]byte(p.Text))

	reqBody.Signature = hex.EncodeToString(mac.Sum(nil))
	reqBody.Speed = p.Speed
	reqBody.Pitch = p.Pitch
	reqBody.Kuten = p.Kuten
	reqBody.Toten = p.Toten
	reqBody.Volume = p.Volume
	reqBody.Intonation = p.Intonation

	jsonned, err := json.Marshal(&reqBody)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonned))

	b, err := sendRequest(jsonned)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func DownloadCoeFont(p *coeFontParameter, filename string) error {
	if !validate(p) {
		return errors.New("The requied paramter should not be empty.")
	}

	reqBody := CoeFontReqestBody{}

	reqBody.Accesskey = p.Accesskey
	reqBody.CoeFont = p.CoeFont
	reqBody.Text = p.Text

	mac := hmac.New(sha256.New, []byte(p.ClientSecret))
	mac.Write([]byte(p.Text))

	reqBody.Signature = hex.EncodeToString(mac.Sum(nil))
	reqBody.Speed = p.Speed
	reqBody.Pitch = p.Pitch
	reqBody.Kuten = p.Kuten
	reqBody.Toten = p.Toten
	reqBody.Volume = p.Volume
	reqBody.Intonation = p.Intonation

	jsonned, err := json.Marshal(&reqBody)

	if err != nil {
		return err
	}

	fmt.Println(string(jsonned))

	b, err := sendRequest(jsonned)

	if err != nil {
		return err
	}

	out, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer out.Close()

	io.Copy(out, bytes.NewReader(b))

	return nil
}

func sendRequest(jsonBytes []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", URL, bytes.NewReader(jsonBytes))

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	req2, err := http.NewRequest("GET", res.Header.Get("Location"), nil)

	if err != nil {
		return nil, err
	}

	res2, err := client.Do(req2)

	if err != nil {
		return nil, err
	}

	defer res2.Body.Close()

	resBody, err := ioutil.ReadAll(res2.Body)

	return resBody, nil
}
