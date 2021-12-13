package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type grocery struct {
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
}

func display(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minQuantity, _ := strconv.Atoi(vars["quantity"])
	flag := false
	fruits, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/fruits")
	response, _ := ioutil.ReadAll(fruits.Body)
	var responseItems []grocery
	var fruit []grocery
	json.Unmarshal(response, &fruit)
	responseItems = append(responseItems, fruit...)
	vegetables, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/vegetables")
	response, _ = ioutil.ReadAll(vegetables.Body)
	var vegetable []grocery
	json.Unmarshal(response, &vegetable)
	responseItems = append(responseItems, vegetable...)
	grains, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/grains")
	response, _ = ioutil.ReadAll(grains.Body)
	var grain []grocery
	json.Unmarshal(response, &grain)
	responseItems = append(responseItems, grain...)
	sort.SliceStable(responseItems, func(i, j int) bool {
		return responseItems[i].Name < responseItems[j].Name
	})
	for _, item := range responseItems {
		if item.Quantity <= minQuantity {
			flag = true
			fmt.Fprintln(w, item)
		}
	}
	if flag == false {
		fmt.Fprintln(w, "NotFound")
	}
}
func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/quest/{quantity}", display).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequests()
}
