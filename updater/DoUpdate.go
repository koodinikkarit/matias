package updater

import (
	"net/http"

	update "github.com/inconshreveable/go-update"
)

func DoUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err2 := update.Apply(resp.Body, update.Options{})
	if err2 != nil {
		// error handling
	}
	return err
}
