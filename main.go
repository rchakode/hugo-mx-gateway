/*
Copyright 2020 Rodrigue Chakode and contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		"SendMail",
		"POST",
		"/sendmail",
		MuxSecAllowedDomainsHandler(
			MuxSecReCaptchaHandler(
				http.HandlerFunc(SendMail))),
	},
	Route{
		"Healthz",
		"GET",
		"/healthz",
		http.HandlerFunc(Healthz),
	},
}

func MuxLoggerHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := MuxLoggerHandler(route.Handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_TLS_CERT", "/etc/cert/cert.pem")
	viper.SetDefault("SERVER_TLS_PRIVATEKEY", "/etc/cert/privkey.pem")
	viper.SetDefault("SMTP_SERVER_ADDR", "127.0.0.1:465")
	viper.SetDefault("SMTP_CLIENT_USERNAME", "")
	viper.SetDefault("SMTP_CLIENT_PASSWORD", "")
	viper.SetDefault("CONTACT_REPLY_EMAIL", "noreply@company.com")
	viper.SetDefault("CONTACT_REPLY_CC_EMAIL", "contact@company.com")
	viper.SetDefault("EMAIL_SUBJECT", "Thanks to try our product")
	viper.SetDefault("DEMO_URL", "http://company.com/product-demo")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Infof("Defaulting to port %s", port)
	}

	router := NewRouter()

	log.Infof("Listening on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
