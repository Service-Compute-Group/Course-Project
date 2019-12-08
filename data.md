```bash
get /api
```

```json
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

```bash
get /api/user
```

```json
{
    "data": [
        "un1",
        "un2",
        "un3",
        "un4"
    ]
}
```

```bash
get /api/user/{user}/article
```

```json
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

```bash
get /api/article
```

```json
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

```bash
get /api/article/{article}
```

```json
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

```bash
get /api/tag
```

```json
{
    "data": [
        "tagname1",
        "tagname2",
        "tagname3",
        "tagname4"
    ]
}
```

```bash
get /api/tag/{tag}
```

```json
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



