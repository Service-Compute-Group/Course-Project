package main

import (
	"github.com/Service-Compute-Group/Course-Project/router"
	"github.com/boltdb/bolt"
)

func generateDBInfo() {

	// 生成user.db 和相关数据

	db, _ := bolt.Open("user.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("usernameBucket"))
		b.Put([]byte("uid1"), []byte("un1"))
		b.Put([]byte("uid2"), []byte("un2"))
		b.Put([]byte("uid3"), []byte("un3"))
		b.Put([]byte("uid4"), []byte("un4"))
		return nil
	})

	// 生成article.db和相关数据
	db, _ = bolt.Open("article.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("uidBucket"))
		b.Put([]byte("articleid1"), []byte("uid1"))
		b.Put([]byte("articleid2"), []byte("uid1"))
		b.Put([]byte("articleid3"), []byte("uid3"))
		b.Put([]byte("articleid4"), []byte("uid4"))
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("titleBucket"))
		b.Put([]byte("articleid1"), []byte("title1"))
		b.Put([]byte("articleid2"), []byte("title2"))
		b.Put([]byte("articleid3"), []byte("title3"))
		b.Put([]byte("articleid4"), []byte("title4"))
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("bodyBucket"))
		b.Put([]byte("article1"), []byte("cont1"))
		b.Put([]byte("article2"), []byte("cont2"))
		b.Put([]byte("article3"), []byte("cont3"))
		b.Put([]byte("article4"), []byte("cont4"))
		return nil
	})

	// 生成tag.db和相关数据
	db, _ = bolt.Open("tag.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("tagnameBucket"))
		b.Put([]byte("tagid1"), []byte("tagname1"))
		b.Put([]byte("tagid2"), []byte("tagname2"))
		b.Put([]byte("tagid3"), []byte("tagname3"))
		b.Put([]byte("tagid4"), []byte("tagname4"))
		return nil
	})

	// 生成review.db和相关数据
	db, _ = bolt.Open("review.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("messageBucket"))
		b.Put([]byte("reviewid1"), []byte("message1"))
		b.Put([]byte("reviewid2"), []byte("message2"))
		b.Put([]byte("reviewid3"), []byte("message3"))
		b.Put([]byte("reviewid4"), []byte("message4"))
		return nil
	})

	// 生成articleTag.db和相关数据
	db, _ = bolt.Open("articleTag.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("articleidBucket"))
		b.Put([]byte("atId1"), []byte("articleid1"))
		b.Put([]byte("atId2"), []byte("articleid2"))
		b.Put([]byte("atId3"), []byte("articleid3"))
		b.Put([]byte("atId4"), []byte("articleid4"))
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("tagidBucket"))
		b.Put([]byte("atId1"), []byte("tagid1"))
		b.Put([]byte("atId2"), []byte("tagid1"))
		b.Put([]byte("atId3"), []byte("tagid3"))
		b.Put([]byte("atId4"), []byte("tagid4"))
		return nil
	})

	// 生成articleReview.db和相关数据
	db, _ = bolt.Open("articleReview.db", 0666, nil)
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("articleidBucket"))
		b.Put([]byte("arId1"), []byte("articleid1"))
		b.Put([]byte("arId2"), []byte("articleid1"))
		b.Put([]byte("arId3"), []byte("articleid3"))
		b.Put([]byte("arId4"), []byte("articleid4"))
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("reviewidBucket"))
		b.Put([]byte("arId1"), []byte("reviewid1"))
		b.Put([]byte("arId2"), []byte("reviewid1"))
		b.Put([]byte("arId3"), []byte("reviewid3"))
		b.Put([]byte("arId4"), []byte("reviewid4"))
		return nil
	})
}

func main() {
	// 要生成数据把下面注释反过来即可
	// generateDBInfo()
	r := router.CreateServer()
	r.Run()
}
