package cmd

import (
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
	if bytedata, err := io.ReadAll(r.Body); err != nil {
		log.Error().Err(err).Msg("    Body read failed due to:")
	} else {
		reqBodyString := string(bytedata)
		log.Info().Str("data", reqBodyString).Msg("    ")
	}
	log.Info().Msg("Processing done")
}
