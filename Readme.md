API的输入输出使用JSON，应用REST风格
操作成功时候默认返回HTTP200状态

所有的API调用通过访问如下网址实现：
www.example.com/api

api 地址末尾没有斜杠，api/ 不是合法的地址。

www.example.com/api    使用GET获取可用API列表
www.example.com/api/user   使用GET获取用户列表，返回为一个JSON数组，每个元素包含一个name字段，代表该用户的用户名
www.example.com/api/user/用户名/article   使用GET获取该用户发表的所有文章列表，返回为一个JSON数组，每个元素包含一个id字段（代表文章唯一标识符）、一个title字段（文章标题）
www.example.com/api/article   使用GET获取全站所有文章列表，返回的JSON包含一个一个id字段（文章唯一标识符）、title字段（文章标题）
www.example.com/api/article/文章ID   使用GET获取该文章的具体内容，返回的JSON包含一个user字段（作者的用户名）、一个title字段（文章标题）、一个body字段（文章正文）、一个tag字段（使用数组表示，数组元素是标签名）、一个review字段（数组表示，数组元素是评论内容）
www.example.com/api/tag   使用GET获取所有tag列表，返回是一个JSON数组，数组元素是标签名
www.example.com/api/tag/标签名   使用GET获取所有拥有此tag的文章列表，返回是一个JSON数组，每个元素包含一个id字段（代表文章唯一标识符）、一个title字段（文章标题）

数据恐怕要自行导入，我们没有写内容修改API的时间了

数据库部分：uid, username 唯一
任务要求使用boltDB，只有键值对没有那些花里胡哨的，所以不得不更改结构，但思路还是按照邓设计的表

user.db 
	usernameBucket
		k: uid v: username

article.db
	uidBucket
		k: articleid v: uid   
	titleBucket
		k: articleid v: title
	bodyBucket
		k: articleid v: body

tag.db
	tagnameBucket
		k: tagid v: tagname

review.db
	messageBucket
		k: reviewid v: message

articleTag.db
	articleidBucket
		k: atId v: articleid
	tagidBucket
		k: atId v: tagid

articleReview.db
	articleidBucket
		k: arId v: articleid
	reviewidBucket
		k: arId v: reviewid

(另，data.md是初步测试的结果)