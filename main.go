package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/Antonious-Stewart/idea-decider/configuration"
	"github.com/go-chi/chi/v5"
)

var port string
var numberOfScans int

func init() {
	numberOfScans = configuration.GetValues().NumberOfScans
	port = configuration.GetValues().Port
}

func main() {
	router := chi.NewRouter()

	router.Post("/api", func(writer http.ResponseWriter, request *http.Request) {
		type parameters struct {
			Ideas []string `json:"ideas"`
		}

		p := json.NewDecoder(request.Body)

		var params parameters

		err := p.Decode(&params)

		if err != nil {
			writer.WriteHeader(400)
			writer.Write([]byte("Bad Request"))
			return
		}

		possibilities := len(*&params.Ideas)

		if possibilities < 2 {
			log.Fatal("Please enter more than one idea.")
		}

		ideas := make(map[string]int)

		for _, idea := range params.Ideas {
			ideas[idea] = 0
		}

		for i := 0; i < numberOfScans; i++ {
			index := rand.Intn(possibilities)
			selection := params.Ideas[index]
			ideas[selection]++
		}

		var result string
		maxValue := 0

		for key, value := range ideas {
			if value > maxValue {
				maxValue = value
				result = key
			}
		}

		type response struct {
			Result string `json:"result"`
		}

		resp := response{
			Result: result,
		}

		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")

		json.NewEncoder(writer).Encode(resp)
	})

	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}
