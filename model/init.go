package model

import "github.com/boltdb/bolt"

var userDB, articleDB, articleTagDB, tagDB, reviewDB, articleReviewDB *bolt.DB

func init() {
	userDB, _ = bolt.Open("db/user.db", 0666, nil)
	articleDB, _ = bolt.Open("db/article.db", 0666, nil)
	tagDB, _ = bolt.Open("db/tag.db", 0666, nil)
	articleTagDB, _ = bolt.Open("db/articleTag.db", 0666, nil)
	reviewDB, _ = bolt.Open("db/review.db", 0666, nil)
	articleReviewDB, _ = bolt.Open("db/articleReview.db", 0666, nil)
}
