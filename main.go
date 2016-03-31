package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tucobenedicto/jsonWebToken"
)

const (
	myPrivKeyPath   = "/Users/kubiak/keys/app.rsa"
	myPublicKeyPath = "/Users/kubiak/keys/app.rsa.pub"
	// how the JWT will be saved in the header
	tokenHeaderName = "jwt"
)

var (
	privateKey *rsa.PrivateKey
)

func init() {
	privateKey = jsonWebToken.GetPrivateKeyFromPath(myPrivKeyPath)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles(r.URL.Path[1:])
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	tem.Execute(w, nil)
}

func jwtHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	claims := r.FormValue("claims")

	var claims_json map[string]interface{}
	if err := json.Unmarshal([]byte(claims), &claims_json); err != nil {
		// TODO 404
		log.Println("Couldn't unmarshal: ", err)
	}

	// TODO if claims is empty?

	t := jsonWebToken.NewToken()
	t.SetPrivateKey(privateKey)
	for k, v := range claims_json {
		t.AddData(k, v)
	}

	tokenString, err := t.GenerateToken()
	if err != nil {
		log.Println("Error during token gen: ", err)
	} else {
		w.Write([]byte(tokenString))
	}
}

func apiReportHandler(w http.ResponseWriter, r *http.Request) {
	// allow requests from localhost
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", tokenHeaderName)

	log.Println("Method: ", r.Method)
	if r.Method == "OPTIONS" {
		// If you make a request between localhosts,
		// OPTIONS checks a pre-flight response, allow it
		// this will allow a subsequent GET request to be made
		w.WriteHeader(http.StatusOK)
		return
	}
	jwt := r.Header.Get(tokenHeaderName)

	log.Println("Token: ", jwt)
	if _, ok := jsonWebToken.AuthorizeToken(jwt, &privateKey.PublicKey); ok {
		brands := jsonWebToken.GetFromToken(jwt, &privateKey.PublicKey, "Brn")
		w.Header().Set("Content-Type", "text/json")
		w.Write([]byte(fmt.Sprintf("{\"yourbrands\":%v}", brands)))
	} else {
		// TODO 404?
		log.Println("Couldn't authorize token.")
	}

}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/getjwt", jwtHandler)
	http.HandleFunc("/api/report", apiReportHandler)
	port := ":8081"
	log.Printf("Running on %v...\n", port)
	http.ListenAndServe(port, nil)
}
