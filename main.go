package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:8395")
	viper.SetDefault("SERVER_TLS_CERT", "/etc/cert/cert.pem")
	viper.SetDefault("SERVER_TLS_PRIVATEKEY", "/etc/cert/privkey.pem")
	viper.SetDefault("SMTP_SERVER_ADDR", "127.0.0.1:465")
	viper.SetDefault("SMTP_CLIENT_USERNAME", "")
	viper.SetDefault("SMTP_CLIENT_PASSWORD", "")
	viper.SetDefault("CONTACT_REPLY_EMAIL", "noreply@company.com")
	viper.SetDefault("CONTACT_REPLY_CC_EMAIL", "contact@company.com")
	viper.SetDefault("EMAIL_SUBJECT", "Thanks to try our product")
	viper.SetDefault("DEMO_URL", "http://company.com/product-demo")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(viper.GetString("SERVER_ADDR"), router))
}
