package services

import (
	"github.com/ferriantitus/jakmall-backend/models"
	"math"
)

func ProductTotalReview(dataReview models.ListReview, productID int) int{
	var totalReview = 0

	for _, value := range dataReview {
		if value.ProductID == productID {
			totalReview += 1
		}
	}
	return totalReview
}

func ProductAverageReview(dataReview models.ListReview, productID int) float64 {
	totalRating := 0.0
	totalReview := 0

	for _, value := range dataReview {
		if value.ProductID == productID {
			totalReview += 1
		}
	}

	for _, value := range dataReview {
		if value.ProductID == productID {
			totalRating += float64(value.Rating)
		}
	}

	averageRatingProduct := totalRating / float64(totalReview)

	oneDecimal := math.Round(averageRatingProduct*10)/10

	return oneDecimal
}

func ProductRatingReview(dataReview models.ListReview, productID int) (int,int,int,int,int) {
	sumRating1 := 0
	sumRating2 := 0
	sumRating3 := 0
	sumRating4 := 0
	sumRating5 := 0

	for _, value := range dataReview {
		if value.ProductID == productID {
			if value.Rating == 1 {
				sumRating1 += 1
			}
			if value.Rating == 2 {
				sumRating2 += 1
			}
			if value.Rating == 3 {
				sumRating3 += 1
			}
			if value.Rating == 4 {
				sumRating4 += 1
			}
			if value.Rating == 5 {
				sumRating5 += 1
			}
		}
	}
	return sumRating5, sumRating4, sumRating3, sumRating2, sumRating1
}