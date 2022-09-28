package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ferriantitus/jakmall-backend/models"
	"github.com/ferriantitus/jakmall-backend/services"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "stop" { // break command
		fmt.Println("Choose your option : ")
		fmt.Println("(1) Get All Review Summary")
		fmt.Println("(2) Get Review Summary by Product ID")
		fmt.Println("(stop) End the program")
		fmt.Print("Enter your choice number: ")
		scanner.Scan()
		text = scanner.Text()
		if text == "1"{
			fmt.Println("--> All Summary review:")

			s5, s4, s3, s2, s1 := services.SumRatingReview(dataReview)

			j, _ := json.MarshalIndent(models.OutputReview{
				TotalReviews:   services.TotalSummaryReview(dataReview),
				AverageRatings: services.SummaryAverageReview(dataReview),
				Star5:          s5,
				Star4:          s4,
				Star3:          s3,
				Star2:          s2,
				Star1:          s1,
			}, "","  ")
			fmt.Println(string(j))

		} else if text == "2"{
			fmt.Print("Enter Product ID: ")
			scanner.Scan()
			var ProdID string
			ProdID = scanner.Text()
			ProdIDConv, _ := strconv.Atoi(ProdID)

			fmt.Println("--> Product Review: product", ProdIDConv)

			s5, s4, s3, s2, s1 := services.ProductRatingReview(dataReview, ProdIDConv)

			j, _ := json.MarshalIndent(models.OutputReview{
				TotalReviews:   services.ProductTotalReview(dataReview, ProdIDConv),
				AverageRatings: services.ProductAverageReview(dataReview, ProdIDConv),
				Star5:          s5,
				Star4:          s4,
				Star3:          s3,
				Star2:          s2,
				Star1:          s1,
			}, "","  ")
			fmt.Println(string(j))
		} else {
			fmt.Println("Invalid Input")
		}
		fmt.Println("\n----------------------------------------------")
	}
	fmt.Println("Program stopped")
}
