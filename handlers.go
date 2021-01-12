package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var AppUname string
var AppPwd string
var AppExtEndpoint string

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CLUSTER_REGION:", os.Getenv("CLUSTER_REGION"))
	fmt.Fprint(w, "Welcome!\n This app is running on:"+os.Getenv("CLUSTER_REGION"))
}

func defHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"status\":\"UP\"}")
}

func custHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"status\":\"UP\"}")
}

func callecho(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Calling endpoint. ..." + AppExtEndpoint)
	req, err := http.NewRequest("GET", AppExtEndpoint, nil)
	//req.SetBasicAuth(AppUname, AppPwd)

	if err != nil {
		//log.Fatal("Error reading request. ", err)
		fmt.Printf("Error reading request. ...")
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error reading response. ...")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading reponse body. ...")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(string(body)); err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", body)

}

func printenv(w http.ResponseWriter, r *http.Request) {

	envpath, _ := os.LookupEnv("ENVVARS_PATH")
	clustername, _ := os.LookupEnv("CLUSTER_REGION")
	appversion, _ := os.LookupEnv("APP_VERSION")
	dbusername, _ := os.LookupEnv("DB_UNAME")
	dbpwd, _ := os.LookupEnv("DB_PWD")

	fmt.Fprint(w, "CLUSTER_REGION:", clustername, "\n")
	fmt.Fprint(w, "ENVVARS_PATH:", envpath, "\n")
	fmt.Fprint(w, "APP_VERSION:", appversion, "\n")
	fmt.Fprint(w, "DB_UNAME:", dbusername, "\n")
	fmt.Fprint(w, "DB_PWD:", dbpwd, "\n")
	fmt.Fprint(w, "APP_EXT_ENDPOINT:", AppExtEndpoint, "\n")
	fmt.Fprint(w, "APP_UNAME:", AppUname, "\n")
	fmt.Fprint(w, "APP_PWD:", AppPwd, "\n")
}
