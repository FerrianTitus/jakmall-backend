package main

import (
	"encoding/json"
	"github.com/ferriantitus/jakmall-backend/models"
	"io/ioutil"
	"log"
)

func main() {
	product, _ := ioutil.ReadFile("products.json")
	var dataProduct models.ListProduct
	err := json.Unmarshal(product, &dataProduct)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}

	review, _ := ioutil.ReadFile("reviews.json")
	var dataReview models.ListReview
	err = json.Unmarshal(review, &dataReview)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}
}
