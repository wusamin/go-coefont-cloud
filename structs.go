package coefontcloud

type CoeFontParameter struct {
	CoeFont      string
	Text         string
	Accesskey    string
	ClientSecret string
	Speed        float32
	Pitch        int16
	Kuten        float32
	Toten        float32
	Volume       float32
	Intonation   float32
}

type CoeFontReqestBody struct {
	CoeFont    string  `json:"coefont"`
	Text       string  `json:"text"`
	Accesskey  string  `json:"accesskey"`
	Signature  string  `json:"signature"`
	Speed      float32 `json:"speed"`      // 0.1 to 10. Default is 1.0.
	Pitch      int16   `json:"pitch"`      // -3000 to 3000. Default is 0.
	Kuten      float32 `json:"kuten"`      // 0 to 5. Default is 0.7.
	Toten      float32 `json:"toten"`      // 0.2 to 2.0. Default is 0.4.
	Volume     float32 `json:"volume"`     // 0 to 5. Default is 1.0.
	Intonation float32 `json:"intonation"` // 0 to 2. Default is 1.0.
}
