package model

import "github.com/boltdb/bolt"

func GetUserNameList() []string {
	var opt []string
	userDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("usernameBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			opt = append(opt, string(v))
		}
		return nil
	})
	return opt

}

type ArticleListResp struct {
	ArticleID string `json:"articleID"`
	Title     string `json:"title"`
}

func GetArticleList(username string) []ArticleListResp {

	uid := getUidWithUsername(username)

	var opt []ArticleListResp
	articleDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("uidBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// 找到属于该user的article id
			if string(v) == uid {
				tmp := ArticleListResp{
					ArticleID: string(k),
				}
				// 根据articleID 获取相应article
				articleDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("titleBucket"))
					tmp.Title = string(b.Get(k))
					return nil
				})
				opt = append(opt, tmp)
			}
		}
		return nil
	})
	return opt
}

func GetAllArticles() []ArticleListResp {
	var opt []ArticleListResp
	articleDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("titleBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tmp := ArticleListResp{
				ArticleID: string(k),
				Title:     string(v),
			}
			opt = append(opt, tmp)
		}
		return nil
	})
	return opt
}

type ConcreteInfo struct {
	User   string   `json:"user"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Tag    []string `json:"tag"`
	Review []string `json:"review"`
}

func GetConcreteInfo(articleID string) ConcreteInfo {
	opt := ConcreteInfo{}
	articleDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("titleBucket"))
		opt.Title = string(b.Get([]byte(articleID)))
		b = tx.Bucket([]byte("bodyBucket"))
		opt.Body = string(b.Get([]byte(articleID)))
		b = tx.Bucket([]byte("uidBucket"))
		uid := b.Get([]byte(articleID))
		userDB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("usernameBucket"))
			opt.User = string(b.Get(uid))
			return nil
		})
		return nil
	})
	opt.Tag = getTagsWithArticleID(articleID)
	opt.Review = getMessageWithArticleID(articleID)
	return opt
}

func getTagsWithArticleID(articleID string) []string {
	var opt []string
	articleTagDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("articleidBucket"))
		b.ForEach(func(k, v []byte) error {
			if string(v) == articleID {
				articleTagDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("tagidBucket"))
					tagid := b.Get(k)
					tagDB.View(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("tagnameBucket"))
						tmp := b.Get(tagid)
						opt = append(opt, string(tmp))
						return nil
					})
					return nil
				})
			}
			return nil
		})
		return nil
	})
	return opt
}

func getMessageWithArticleID(articleID string) []string {
	var opt []string
	articleReviewDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("articleidBucket"))
		b.ForEach(func(k, v []byte) error {
			if string(v) == articleID {
				articleReviewDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("reviewidBucket"))
					reviewid := b.Get(k)
					reviewDB.View(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("messageBucket"))
						tmp := b.Get(reviewid)
						opt = append(opt, string(tmp))
						return nil
					})
					return nil
				})
			}
			return nil
		})
		return nil
	})
	return opt
}

func GetAllTags() []string {
	var opt []string
	tagDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tagnameBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			opt = append(opt, string(v))
		}
		return nil
	})
	return opt
}

func GetArticleListWithTag(tagname string) []ArticleListResp {
	tagID := getTagIDWithTagname(tagname)

	var opt []ArticleListResp
	articleTagDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tagidBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// 找到属于该tag的article id
			if string(v) == tagID {
				var tmp ArticleListResp
				// 根据atId 获取article id
				articleTagDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("articleidBucket"))
					tmp.ArticleID = string(b.Get(k))
					return nil
				})
				// 根据articleID 获取title
				articleDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("titleBucket"))
					tmp.Title = string(b.Get([]byte(tmp.ArticleID)))
					return nil
				})
				opt = append(opt, tmp)
			}
		}
		return nil
	})
	return opt
}

func getUidWithUsername(username string) string {
	var uid string
	userDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("usernameBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) == username {
				uid = string(k)
				break
			}
		}
		return nil
	})
	return uid
}

func getTagIDWithTagname(tagname string) string {
	var tagID string
	tagDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tagnameBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) == tagname {
				tagID = string(k)
				break
			}
		}
		return nil
	})
	return tagID
}
