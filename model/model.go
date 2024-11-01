// Package model has domain models used throughout the application.
package model

type Language string

func (l Language) String() string {
	return string(l)
}

const (
	LanguageDanish Language = "dansk"
	LanguageIldsk  Language = "ildsk"
)
