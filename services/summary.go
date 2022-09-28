package services

import (
	"github.com/ferriantitus/jakmall-backend/models"
	"math"
)

func TotalSummaryReview(dataReview models.ListReview) int{
	totalCountReviews := len(dataReview)

	return totalCountReviews
}

func SummaryAverageReview(dataReview models.ListReview) float64 {
	sumRating := 0

	for _, value := range dataReview {
		sumRating += value.Rating
	}

	average := float64(sumRating) / float64(len(dataReview))

	oneDecimal := math.Round(average*10)/10

	return oneDecimal
}

func SumRatingReview(dataReview models.ListReview) (int, int, int, int, int) {
	sumRating1 := 0
	sumRating2 := 0
	sumRating3 := 0
	sumRating4 := 0
	sumRating5 := 0

	for _, value := range dataReview {
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
	return sumRating5, sumRating4, sumRating3, sumRating2, sumRating1
}