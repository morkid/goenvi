package goenv

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// New Environment
func New() *Environment {
	return &Environment{}
}

var e *Environment = New()

// Initialize environment
func Initialize() {
	e.Initialize()
}

// Register custom viper
func Register(v *viper.Viper, optional bool) {
	e.Register(v, optional)
}

// AddFlagSetProvider add flag set provider
func AddFlagSetProvider(f FlagSetProvider) {
	e.AddFlagSetProvider(f)
}

// Add from local file
func Add(envType string, fileName string) {
	e.Add(envType, fileName)
}

// FlagSetProvider interface
type FlagSetProvider interface {
	VisitAll(fn func(*pflag.FlagSet))
}

// Environment struct
type Environment struct {
	optionalConfig   []*viper.Viper
	mainConfig       []*viper.Viper
	flagSetProviders []FlagSetProvider
}

// Initialize environment
func (env *Environment) Initialize() {
	if nil != env.optionalConfig {
		for i := range env.optionalConfig {
			optionalViperKeys := env.optionalConfig[i].AllKeys()
			for j := range optionalViperKeys {
				viper.Set(optionalViperKeys[j], env.optionalConfig[i].Get(optionalViperKeys[j]))
			}
		}
	}

	env.Add("dotenv", ".env")

	if nil != env.mainConfig {
		for i := range env.mainConfig {
			mainViperKeys := env.mainConfig[i].AllKeys()

			for j := range mainViperKeys {
				viper.Set(mainViperKeys[j], env.mainConfig[i].Get(mainViperKeys[j]))
			}
		}
	}

	if nil != env.flagSetProviders && len(env.flagSetProviders) > 0 {
		for i := range env.flagSetProviders {
			env.loadParameters(env.flagSetProviders[i])
		}
	}

	keys := viper.AllKeys()

	for i := range keys {
		keyName := keys[i]
		keyNameUpper := env.fixKeyName(keyName)
		strValue := viper.GetString(keyName)

		if strValue != "" {
			os.Setenv(keyNameUpper, strValue)
		} else {
			originalValue := viper.Get(keyName)

			if nil != originalValue {
				jsonValue, err := json.Marshal(originalValue)

				if nil == err {
					os.Setenv(keyNameUpper, string(jsonValue))
				}
			}
		}
	}
}

// Register custom viper
func (env *Environment) Register(v *viper.Viper, optional bool) {
	if optional {

		if nil == env.optionalConfig {
			env.optionalConfig = []*viper.Viper{}
		}

		if nil != v {
			env.optionalConfig = append(env.optionalConfig, v)
		}

	} else {

		if nil == env.mainConfig {
			env.mainConfig = []*viper.Viper{}
		}

		if nil != v {
			env.mainConfig = append(env.mainConfig, v)
		}

	}
}

// AddFlagSetProvider add flag set provider
func (env *Environment) AddFlagSetProvider(f FlagSetProvider) {
	if nil == env.flagSetProviders {
		env.flagSetProviders = []FlagSetProvider{}
	}

	env.flagSetProviders = append(env.flagSetProviders, f)
}

// Add from local file
func (env *Environment) Add(envType string, fileName string) {
	v := viper.New()
	v.SetConfigType(envType)
	v.SetConfigFile(fileName)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); nil == err {
		keys := v.AllKeys()

		for i := range keys {
			keyName := env.fixKeyName(keys[i])

			if strings.ToUpper(keys[i]) != keyName {
				viper.Set(keys[i], v.Get(keys[i]))
			}

			viper.Set(keyName, v.Get(keys[i]))
		}
	}
}

func (env *Environment) loadParameters(f FlagSetProvider) {
	if nil == f {
		return
	}

	v := viper.New()

	f.VisitAll(func(fSet *pflag.FlagSet) {
		if err := v.BindPFlags(fSet); nil == err {
			keys := v.AllKeys()

			for i := range keys {
				if value := v.Get(keys[i]); nil != value {
					keyName := env.fixKeyName(keys[i])
					viper.Set(keyName, value)
				}
			}
		}
	})
}

func (Environment) fixKeyName(keyName string) string {
	re := regexp.MustCompile(`[^A-z0-9\_]+`)
	keyByte := re.ReplaceAll([]byte(keyName), []byte("_"))

	return strings.ToUpper(string(keyByte))
}
