package config

import (
	"flag"
	"os"
	"strings"
)

func getSetting(key string, defValue string, description string) *string {
	pref := os.Getenv(strings.ToUpper(key))
	flg := flag.String(key, defValue, description)
	if pref == "" {
		return flg
	}
	return &pref
}

func getBool(key string, defValue bool, description string) *bool {
	pref := os.Getenv(strings.ToUpper(key))
	flg := flag.Bool(key, defValue, description)
	if pref != "true" && pref != "false" {
		return flg
	}
	val := pref == "true"
	return &val
}
