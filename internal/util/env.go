package util

import (
	"github.com/rs/zerolog/log"
	"github.com/sourcehaven/mypass-api/internal/glob"
	"strings"
)

func ToEnv(env string) string {
	switch strings.ToLower(env) {
	case "0", glob.Local:
		return glob.Local
	case "1", "dev", "devel", glob.Development:
		return glob.Development
	case "2", "examples", glob.Testing:
		return glob.Testing
	case "3", "stage", glob.Staging:
		return glob.Staging
	case "4", "prod", glob.Production:
		return glob.Production
	default:
		log.Panic().Msg("Invalid env value: " + env)
		return "error"
	}
}
