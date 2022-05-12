// My goal here was a singleton that works like pydantic's `Settings` class,
// but it does not work thus far.
// In another package:
//
//     config.GetSettings().environment undefined
//     (type *config.Settings has no field or method environment)
//     compilerMissingFieldOrMethod
//

package config

import (
	"sync"

	"github.com/codingconcepts/env"
)

type Settings struct {
	environment string `env:"ENVIRONMENT"`
}

var settings *Settings
var once sync.Once

func initSettings() *Settings {
	s := Settings{}
	if err := env.Set(&s); err != nil {
		panic(err)
	}
	return &s
}

func GetSettings() *Settings {
	once.Do(func() {
		settings = initSettings()
	})
	return settings
}
