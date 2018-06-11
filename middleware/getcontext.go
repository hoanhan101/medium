package middleware

import (
	"context"
	"net/http"
)

// createContext returns a context value with key=fooID, value=bar
func createContext(ctx context.Context, r *http.Request) context.Context {
	foo := r.Header.Get("X-Foo-ID")
	if foo == "" {
		foo = "bar"
	}

	return context.WithValue(ctx, "foo", foo)
}

// ContextHandler persists the context request and passes to the next handler.
func ContextHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := createContext(r.Context(), r)
			next.ServeHTTP(w, r.WithContext(ctx))
		}()
	})
}
