package config

type EmailConfig struct {
	AppPassword     string
	EmailEnabled    bool
	EmailServerAddr string
	EmailServerHost string
	SenderEmail     string
}

func getEmailConfig() EmailConfig {
	return EmailConfig{
		AppPassword:     fatalGetString("EMAIL_APP_PASSWORD"),
		EmailEnabled:    fatalGetBool("EMAIL_ENABLED"),
		EmailServerAddr: fatalGetString("EMAIL_SERVER_ADDR"),
		EmailServerHost: fatalGetString("EMAIL_SERVER_HOST"),
		SenderEmail:     fatalGetString("EMAIL_SENDER"),
	}
}
