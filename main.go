package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Param struct {
	Type   string
	Length int
}
type ambar struct {
	mas map[int]interface{}
}

func (ambar *ambar) handler(w http.ResponseWriter, r *http.Request) {
	var body Param
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println("err")
	}
	r.Body.Close()

	an := map[string]interface{}{}
	an["id"] = len(ambar.mas)
	b, e := json.Marshal(an)
	if e != nil {
		fmt.Println("ошибка json")
	}

	defer r.Body.Close()

	w.Write([]byte(string(b)))

	if body.Type == "int" {
		i, _ := strconv.ParseInt(RandIntRunes(body.Length), 10, 64)
		ambar.mas[len(ambar.mas)] = i
	} else {
		ambar.mas[len(ambar.mas)] = RandStringRunes(body.Length)
	}

}

func (ambar *ambar) handler2(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	am := map[string]interface{}{}
	am["value"] = ambar.mas[int(i)]

	b, e := json.Marshal(am)
	if e != nil {
		fmt.Println("ошибка json")
	}
	w.Write([]byte(string(b)))

}

func main() {

	ambar := ambar{map[int]interface{}{}}
	r := chi.NewRouter()

	srv := http.Server{Addr: ":3000", Handler: r}

	defer srv.Shutdown(context.TODO())

	r.Post("/api/generate/", ambar.handler)
	r.Get("/api/retrieve/", ambar.handler2)
	srv.ListenAndServe()
}

func RandStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandIntRunes(n int) string {
	letterRunes := []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
