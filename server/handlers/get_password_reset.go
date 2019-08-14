package handlers

import (
	"net/http"

	"github.com/keratin/authn-server/app"
	"github.com/keratin/authn-server/app/services"
)

func GetPasswordReset(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account, err := app.AccountStore.FindByUsername(r.FormValue("username"))
		if err != nil {
			panic(err)
		}

		// run in the background so that a timing attack can't enumerate usernames
		go func() {
			err := services.PasswordResetSender(app.Config, account)
			if err != nil {
				app.Reporter.ReportRequestError(err, r)
			}
		}()

		// w.WriteHeader(http.StatusOK)
		WriteData(w, http.StatusOK, map[string]string{
			"message": "ok",
		})
	}
}
