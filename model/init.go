package model

import "github.com/boltdb/bolt"

var userDB, articleDB, articleTagDB, tagDB, reviewDB, articleReviewDB *bolt.DB

func init() {
	userDB, _ = bolt.Open("user.db", 0666, nil)
	articleDB, _ = bolt.Open("article.db", 0666, nil)
	tagDB, _ = bolt.Open("tag.db", 0666, nil)
	articleTagDB, _ = bolt.Open("articleTag.db", 0666, nil)
	reviewDB, _ = bolt.Open("review.db", 0666, nil)
	articleReviewDB, _ = bolt.Open("articleReview.db", 0666, nil)
}
