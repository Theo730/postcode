package postcode

import (
    "fmt"
    "net/http"
    "net/url"
    ylog	"github.com/postcode/lib/ylog"
    "github.com/julienschmidt/httprouter"
)

// Read - конвертирование строки URL параметров запроса в структуру данных.
func Read(str url.Values) (map[string]string, error) {
	val := make(map[string]string)

	for k, v := range str {
		val[k] = v[0]
	}
	return val, nil
}

// BasicAuth - аутентификация basic.
func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			h(w, r, ps)
		} else {
			ylog.YLog(3, Ident, "Error: user or/and password incorrect "+fmt.Sprintf("%v", r.RemoteAddr)+" User:'"+user+"' Password:'"+password+"'")
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// GetVal - получение параметра name из url.
func GetVal(w http.ResponseWriter, ps httprouter.Params, name string) string {
	id := string(ps.ByName(name))
	return id
}
