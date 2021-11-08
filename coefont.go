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
	"strconv"
	"time"
)

const URL = "https://api.coefont.cloud/v1/text2speech"

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

	reqBody.CoeFont = p.CoeFont
	reqBody.Text = p.Text
	reqBody.Speed = p.Speed
	reqBody.Pitch = p.Pitch
	reqBody.Kuten = p.Kuten
	reqBody.Toten = p.Toten
	reqBody.Volume = p.Volume
	reqBody.Intonation = p.Intonation

	m := map[string]interface{}{}
	m["coefont"] = p.CoeFont
	m["text"] = p.Text

	jsonned, err := json.Marshal(&m)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	b, err := sendRequest(jsonned, &now, p.ClientSecret, p.Accesskey)

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

	reqBody.CoeFont = p.CoeFont
	reqBody.Text = p.Text
	reqBody.Speed = p.Speed
	reqBody.Pitch = p.Pitch
	reqBody.Kuten = p.Kuten
	reqBody.Toten = p.Toten
	reqBody.Volume = p.Volume
	reqBody.Intonation = p.Intonation

	m := map[string]interface{}{}
	m["coefont"] = p.CoeFont
	m["text"] = p.Text

	jsonned, err := json.Marshal(&m)

	if err != nil {
		return err
	}

	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.In(time.UTC).Unix())

	b, err := sendRequest(jsonned, &now, p.ClientSecret, p.Accesskey)

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

func sendRequest(jsonBytes []byte, date *time.Time, secret string, accesskey string) ([]byte, error) {
	a := jsonBytes

	req, err := http.NewRequest("POST", URL, bytes.NewReader(a))

	if err != nil {
		return nil, err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(date.Unix(), 10) + string(a)))

	req.Header.Set("X-Coefont-Date", strconv.FormatInt(date.Unix(), 10))
	req.Header.Set("X-Coefont-Content", hex.EncodeToString(mac.Sum(nil)))
	req.Header.Set("Authorization", accesskey)
	req.Header.Set("Content-Type", "application/json")

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

	if err != nil {
		return nil, err
	}

	return resBody, nil
}
