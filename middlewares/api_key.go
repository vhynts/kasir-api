package middlewares

import "net/http"

func APIKey(validApiKey string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			// API key wajib ada
			if apiKey == "" {
				http.Error(w, "API Key is required", http.StatusUnauthorized)
				return
			}

			// API key harus sesuai
			if apiKey != validApiKey {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			// lanjut ke handler berikutnya
			next(w, r)
		}
	}
}
