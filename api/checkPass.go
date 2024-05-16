package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/frit2000/go_final_project/env"
	"github.com/golang-jwt/jwt"
)

type AuthPass struct {
	Pass string `json:"password"`
}

type AuthPassError struct {
	MyTocken string `json:"token,omitempty"`
	Err      string `json:"error,omitempty"`
}

var AuthResult AuthPassError
var buf bytes.Buffer
var auth AuthPass

func CheckPass(w http.ResponseWriter, r *http.Request) {

	//получить данные от запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//переводим данные в стркутуру auth
	if err = json.Unmarshal(buf.Bytes(), &auth); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//проверить на совпадение TODO_PASSWORD и тела запроса
	if auth.Pass == os.Getenv("TODO_PASSWORD") {
		//сформировать jwt токен
		secret := []byte(auth.Pass)
		jwtToken := jwt.New(jwt.SigningMethodHS256)
		AuthResult.MyTocken, err = jwtToken.SignedString(secret)
		if err != nil {
			log.Printf("failed to sign jwt: %s\n", err)
		}
	} else {
		AuthResult.Err = "Неверный пароль"
	}

	//возвратить токен в поле tocken или ошибку
	resp, err := json.Marshal(&AuthResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Println("Не удалось записать данные в html:", err)
	}

}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwt string // JWT-токен из куки
		// смотрим наличие пароля
		pass := env.SetPass()
		if len(pass) > 0 {
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			if jwt != AuthResult.MyTocken {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
