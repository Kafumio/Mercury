package account

type OfficialAccount struct {
	AppID          string `json:"app_id"`           // app_id
	AppSecret      string `json:"app_secret"`       // app_secret
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // encoding_aes_key
}
