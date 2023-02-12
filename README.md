# NOTICE

Coefont API is now only avalable in the Enterprise Plan.
the Enterprise Plan users will probably use own code, so this repository has been archived.

# go-coefont-cloud

go-coefont-cloud is golang library CoeFont CLOUD.
> https://coefont.cloud/

## Install
```
go get -u github.com/wusamin/go-coefont-cloud
```

## Usage
```
// Get request body with this method.
p := coefontcloud.NewParam()

p.CoeFont = "Millial"
p.Accesskey = "access key"
p.ClientSecret = "client secret"
p.Text = "Hello world!"
p.Speed = 0.9

// Download file.
if err := coefontcloud.DownloadCoeFont(p, "C:/temp/voice.wav"); err != nil {
	fmt.Println(err)
} else {
	fmt.Println("successed!!!")
}
```
