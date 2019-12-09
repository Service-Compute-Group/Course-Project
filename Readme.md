# API 规范
API的输入输出使用JSON格式，应用REST风格。

操作成功时候默认返回HTTP200状态。

所有的API调用通过访问如下网址实现：```{{host}}/api```

注意，地址末尾没有斜杠，```api/```不是合法的地址。

以下是支持的API列表。

| 网址 | 方法 | 用途 |
| --- | ---- | --- |
| ```{{host}}/api``` | ```GET``` | 获取可用API列表 |
| ```{{host}}/api/user``` | ```GET``` | 获取全部用户列表 |
| ```{{host}}/api/user/{{username}}/article``` | ```GET``` | 获取该用户发表的所有文章列表 |
| ```{{host}}/api/article``` | ```GET``` | 获取全站所有文章列表 |
| ```{{host}}/api/article/{{articleID}}``` | ```GET``` | 获取该文章的具体内容 |
| ```{{host}}/api/tag``` | ```GET```| 获取所有标签列表 |
| ```{{host}}/api/tag/{{tagname}}``` | ```GET``` | 获取全站所有拥有此标签的文章列表 |

注释：
1. 获取用户列表，返回为一个JSON数组，每个元素包含一个name字段，代表该用户的用户名
2. 该用户发表的所有文章列表，返回为一个JSON数组，每个元素包含一个id字段（代表文章唯一标识符）、一个title字段（文章标题）
3. 全站所有文章列表，返回的JSON包含一个一个id字段（文章唯一标识符）、title字段（文章标题）
4. 文章的具体内容，返回的JSON包含一个user字段（作者的用户名）、一个title字段（文章标题）、一个body字段（文章正文）、一个tag字段（使用数组表示，数组元素是标签名）、一个review字段（数组表示，数组元素是评论内容）
5. 所有tag列表，返回是一个JSON数组，数组元素是标签名
6. 所有拥有此tag的文章列表，返回是一个JSON数组，每个元素包含一个id字段（代表文章唯一标识符）、一个title字段（文章标题）

# 数据库设计
数据库部分：uid, username 唯一。

任务要求使用boltDB，只有键值对没有那些花里胡哨的，所以不得不更改结构，但思路还是按照邓设计的表。

+ user.db 
  + usernameBucket
    + k: uid v: username

+ article.db
  + uidBucket
    + k: articleid v: uid   
  + titleBucket
    + k: articleid v: title
  + bodyBucket
    + k: articleid v: body

+ tag.db
  + tagnameBucket
    + k: tagid v: tagname

+ review.db
  + messageBucket
    + k: reviewid v: message

+ articleTag.db
  + articleidBucket
    + k: atId v : articleid
  + tagidBucket
    + k: atId v: tagid

+ articleReview.db
  + articleidBucket
    + k: arId v: articleid
  + reviewidBucket
    + k: arId v: reviewid

(另，data.md是初步测试的结果)

# 网站设计

网站设计了三类页面。

* 访问```{{host}}/user/{{username}}```可以查看该用户的详细信息。
* 访问```{{host}}/article/{{articleID}}```可以查看该文章的详细内容。
* 访问```{{host}}/tag/{{tagname}}```可以查看拥有该标签的所有文章列表。
