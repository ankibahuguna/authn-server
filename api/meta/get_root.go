package meta

import (
	"bytes"
	"net/http"

	"github.com/keratin/authn-server/api"
	"github.com/keratin/authn-server/api/views"
)

func getRoot(app *api.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		views.Root(&buf)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	}
}
