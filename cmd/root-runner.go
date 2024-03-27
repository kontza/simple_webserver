package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func rootRunner(cmd *cobra.Command, args []string) {
	log.Info().Str("port", args[0]).Msg("Using")
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", args[0]), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Processing request...")
	log.Info().Msg("  URL:")
	log.Info().Str("path", r.URL.Path).Msg("    ")
	log.Info().Msg("  Headers:")
	for k, v := range r.Header {
		log.Info().Strs(k, v).Msg("    ")
	}
	log.Info().Msg("  Query parameters:")
	for k, v := range r.URL.Query() {
		log.Info().Strs(k, v).Msg("    ")
	}
	log.Info().Msg("  Body:")
	if bodyData, err := io.ReadAll(r.Body); err != nil {
		log.Error().Err(err).Msg("    Body read failed due to:")
	} else {
		jsonMap := make(map[string](interface{}))
		err := json.Unmarshal(bodyData, &jsonMap)
		if err != nil {
			log.Warn().Err(err).Msg("    Failed to unmarshal JSON due to")
			log.Info().Str("as string", string(bodyData)).Msg("    ")
		} else {
			log.Info().Interface("as map", jsonMap).Msg("    ")
		}
	}
	log.Info().Msg("Processing done")
}
