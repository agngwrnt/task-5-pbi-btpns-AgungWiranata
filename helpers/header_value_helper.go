package helpers

import "net/http"

func GetHeaderValue(r *http.Request, headerName string) string {
	return r.Header.Get(headerName)
}
