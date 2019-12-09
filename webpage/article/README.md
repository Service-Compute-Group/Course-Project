这是文章内容的显示界面。

用户访问任何一个 ```{{host}}/article/{{articleID}}``` 网址，路由都应当返回这个页面，即```articleview.html```。页面在加载的时候会从url当中提取文章ID并向API发送数据请求。

需要事后做的处理：处理好资源文件（CSS和JS）的路由导航。它们是```{{host}}/static/articleview.css```和```{{host}}/static/articleview.js```。