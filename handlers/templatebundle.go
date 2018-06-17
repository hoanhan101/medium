package handlers

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/hoanhan101/medium/common"
)

// TemplateBundleHandler creates a template bundle and sends to client.
func TemplateBundleHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var templateContentItemsBuffer bytes.Buffer
		enc := gob.NewEncoder(&templateContentItemsBuffer)

		// Obtain a map of template where the key is template's filename
		// without extension and the value is its string content.
		m := env.TemplateSet.Bundle().Items()
		err := enc.Encode(&m)
		if err != nil {
			log.Println("Encountered error while encoding gob:", err)
		}

		// application/octet-stream indicates that the response will be a
		// binary formated.
		w.Header().Set("Content-Type", "application/octet-stream")

		// Write back a slice of bytes.
		w.Write(templateContentItemsBuffer.Bytes())
	})
}
