// My goal here was a singleton that works like pydantic's `Settings` class,
// but it does not work thus far.
// In another package:
//
//     config.GetSettings().environment undefined
//     (type *config.Settings has no field or method environment)
//     compilerMissingFieldOrMethod
//
// My workaround is `NewSettings`, which is not a singleton
// and may have some creation overhead.

package config

import (
	"sync"

	"github.com/codingconcepts/env"
)

type settings struct {
	ENVIRONMENT  string `env:"ENVIRONMENT"`
	POSTGRES_DSN string `env:"POSTGRES_DSN"`
}

func NewSettings() *settings {
	s := settings{}
	if err := env.Set(&s); err != nil {
		panic(err)
	}
	return &s
}

var Settings *settings
var once sync.Once

func OnceSettings() {
	once.Do(func() {
		s := settings{}
		if err := env.Set(&s); err != nil {
			panic(err)
		}
		Settings = &s
	})
}
