package main

import (
	"log"

	"github.com/spf13/viper"
)

var (
	logLevel              string
	crowdsecBouncerAPIKey string
	crowdsecBouncerURL    string
	mikrotikHost          string
	username              string
	password              string
	async                 bool
	useTLS                bool
)

func initConfig() {
	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "1")
	viper.BindEnv("crowdsec_bouncer_api_key")
	viper.BindEnv("crowdsec_url")
	viper.SetDefault("crowdsec_url", "http://crowdsec:8080/")
	viper.BindEnv("mikrotik_host")
	viper.BindEnv("mikrotik_user")
	viper.BindEnv("mikrotik_pass")
	viper.BindEnv("mikrotik_tls")
	viper.SetDefault("mikrotik_tls", "true")

	logLevel = viper.GetString("log_level")

	crowdsecBouncerAPIKey = viper.GetString("crowdsec_bouncer_api_key")
	if crowdsecBouncerAPIKey == "" {
		log.Fatal("Crowdsec API key is not set")
	}
	crowdsecBouncerURL = viper.GetString("crowdsec_url")
	if crowdsecBouncerURL == "" {
		log.Fatal("Crowdsec URL is not set")
	}
	mikrotikHost = viper.GetString("mikrotik_host")

	username = viper.GetString("mikrotik_user")
	if username == "" {
		log.Fatal("Mikrotik username is not set")
	}

	password = viper.GetString("mikrotik_pass")
	if password == "" {
		log.Fatal("Mikrotik password is not set")
	}

	useTLS = viper.GetBool("mikrotik_tls")
	log.Printf("[INFO] Using config: %+v", viper.AllSettings())
}