package handlers

import (
	"log"
	"net/http"

	"github.com/keratin/authn-server/app"
	"github.com/keratin/authn-server/app/services"
	"github.com/keratin/authn-server/server/sessions"
)

func DeleteSession(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := services.SessionEnder(app.RefreshTokenStore, sessions.GetRefreshToken(r))
		if err != nil {
			app.Reporter.ReportRequestError(err, r)
		}

		sessions.Set(app.Config, w, "")

		w.WriteHeader(http.StatusOK)

		// w.Header().Set("Access-Control-Allow-Origin", "*")

		log.Println("Set cookie at DELETE /session")
		log.Println(r.Cookies())
		log.Println(w.Header())
		log.Println(w.Header().Get("Set-Cookie"))
	}
}
