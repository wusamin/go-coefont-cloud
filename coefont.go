package coefontcloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const URL = "https://api.coefont.cloud/text2speech"

func DownloadCoeFont(p *CoeFontParameter, filename string) error {
	reqBody := CoeFontReqestBody{}

	mac := hmac.New(sha256.New, []byte(p.ClientSecret))
	mac.Write([]byte(p.Text))

	reqBody.Signature = hex.EncodeToString(mac.Sum(nil))
	reqBody.Accesskey = p.Accesskey
	reqBody.CoeFont = p.CoeFont
	reqBody.Text = p.Text
	fmt.Println(p.Speed)
	reqBody.Speed = p.Speed
	fmt.Println(reqBody.Speed)
	reqBody.Pitch = p.Pitch
	reqBody.Kuten = p.Kuten
	reqBody.Toten = p.Toten
	reqBody.Volume = p.Volume
	reqBody.Intonation = p.Intonation

	jsoned, err := json.Marshal(&reqBody)

	if err != nil {
		return err
	}

	fmt.Println(string(jsoned))

	return nil
}
