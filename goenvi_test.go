package goenv

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentJson(t *testing.T) {
	assert := assert.New(t)
	env := New()
	env.Add("json", "config_test.json")
	env.Initialize()

	assert.Equal("hello world", viper.GetString("MESSAGE"), "Invalid message")
	assert.Equal("hello world", os.Getenv("MESSAGE"), "Invalid message")
}

func TestEnvironmentDotenv(t *testing.T) {
	assert := assert.New(t)
	env := New()
	env.Add("dotenv", ".env.test")
	env.Initialize()

	assert.Equal("hello world", viper.GetString("MESSAGE"), "Invalid message")
	assert.Equal("hello world", os.Getenv("MESSAGE"), "Invalid message")
}

func TestEnvironmentYaml(t *testing.T) {
	assert := assert.New(t)
	env := New()
	env.Add("yaml", "config_test.yml")
	env.Initialize()

	assert.Equal("3.8", viper.GetString("VERSION"), "Invalid version")
	assert.Equal("3.8", os.Getenv("VERSION"), "Invalid version")
	assert.Equal("golang:1.15-buster", viper.GetString("services.go.image"), "Invalid image")
	assert.Equal("golang:1.15-buster", os.Getenv("SERVICES_GO_IMAGE"), "Invalid image")
	assert.Equal("", os.Getenv("services.go.image"), "Invalid image")
	assert.Equal("", os.Getenv("SERVICES.GO.IMAGE"), "Invalid image")
}
