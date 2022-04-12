package coefontcloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetDictionary(category string, credential *ConefontCredential) ([]*Dict, error) {
	result, err := sendDictRequest("https://api.coefont.cloud/v1/dict", category, credential.ClientSecret, credential.Accesskey)
	if err != nil {
		return nil, err
	}
	return result, err
}

func sendDictRequest(url, category, secret, accesskey string) ([]*Dict, error) {
	req, err := http.NewRequest("GET", url, nil)

	now := time.Now()

	if err != nil {
		return nil, err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(now.Unix(), 10)))

	req.Header.Set("X-Coefont-Date", strconv.FormatInt(now.Unix(), 10))
	req.Header.Set("X-Coefont-Content", hex.EncodeToString(mac.Sum(nil)))
	req.Header.Set("Authorization", accesskey)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("category", category)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret []*Dict
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
