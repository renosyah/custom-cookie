package cookie

import "github.com/gorilla/securecookie"
import "net/http"

var cookiehandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func set_cookie(cookie_name string, cookie_value string, res http.ResponseWriter) {
	value := map[string]string{
		"name": cookie_value,
	}
	if encoded, err := cookiehandler.Encode(cookie_name, value); err == nil {
		cookie_ku := &http.Cookie{
			Name:  cookie_name,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(res, cookie_ku)
	}
}
func get_cookie(cookie_name string, req *http.Request) (name_usernya string) {
	if cookie_ini, err := req.Cookie(cookie_name); err == nil {
		nilai_cookie := make(map[string]string)
		if err = cookiehandler.Decode(cookie_name, cookie_ini.Value, &nilai_cookie); err == nil {
			name_usernya = nilai_cookie["name"]
		}

	}
	return name_usernya
}

func clear_cookie(cookie_name string, res http.ResponseWriter) {
	bersihkan_cookie_ku := &http.Cookie{
		Name:   cookie_name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(res, bersihkan_cookie_ku)
}
