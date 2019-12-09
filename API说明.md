## API列表
+ get/api:获取所有可以使用的API列表
data后面的内容是API列表以及各个API的说明
```
  {
    "data": {
        "/api": "使用GET获取可用API列表",
        "/api/article": "使用GET获取全站所有文章列表",
        "/api/article/{articalID}}": "使用GET获取该文章的具体内容",
        "/api/tag": "使用GET获取所有tag列表",
        "/api/tag/{tagname}": "使用GET获取所有拥有此tag的文章列表",
        "/api/user/{user}/article": "使用GET获取该用户发表的所有文章列表",
        "api/user": "使用GET获取用户列表"
    }
}
```
+  get /api/user: 获取所有用户列表
  所返回的结果是一个用户的列表，列表中每一个元素都是一个用户
```
{
    "data": [
        "un1",
        "un2",
        "un3",
        "un4"
    ]
}
```

+ get /api/user/{user}/article: 获取指定用户发表的所有文章
返回结果的列表中每一项中有两个元素，分别是文章的ID与文章标题，每一项代表的文章都是参数中指定的user发出来的
```
{
    "data": [
        {
            "articleID": "articleid1",
            "title": "title1"
        },
        {
            "articleID": "articleid2",
            "title": "title1"
        }
    ]
}
```
+  get /api/article:获取该站全部文章
与上一个API一样，返回的也是一个列表，每一项中也有两个元素，分别是文章ID与文章标题
```
{
    "data": [
        {
            "articleID": "articleid1",
            "title": "title1"
        },
        {
            "articleID": "articleid2",
            "title": "title1"
        },
        {
            "articleID": "articleid3",
            "title": "title3"
        },
        {
            "articleID": "articleid4",
            "title": "title4"
        }
    ]
}
```
+ get /api/article/{article}:获取某篇文章的具体内容
返回的是一个列表，列表中项包含了参数指定的文章的信息，有文章的作者，题目，内容，标签等
```
{
    "data": {
        "user": "un1",
        "title": "title1",
        "body": "",
        "tag": [
            "tagname1"
        ],
        "review": [
            "message1",
            "message1"
        ]
    }
}
```
+ get /api/tag: 获取所有标签列表
返回一个列表，列表中是各个标签
```
{
    "data": [
        "tagname1",
        "tagname2",
        "tagname3",
        "tagname4"
    ]
}
```
+ get /api/tag/{tag}:获得拥有特定标签的文章列表
返回一个列表，每一项中也有两个元素，分别是文章ID与文章标题，这里面每一项表示的文章都有上面参数中指定的tag标签
```
{
    "data": [
        {
            "articleID": "articleid1",
            "title": "title1"
        },
        {
            "articleID": "articleid2",
            "title": "title2"
        }
    ]
}
```
