package models

type Review struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Rating    int `json:"rating"`
}

type ListReview []Review

type OutputReview struct {
	TotalReviews	int		`json:"total_reviews"`
	AverageRatings float64	`json:"average_ratings"`
	Star5          int		`json:"5_star"`
	Star4          int		`json:"4_star"`
	Star3          int		`json:"3_star"`
	Star2          int		`json:"2_star"`
	Star1          int		`json:"1_star"`
}