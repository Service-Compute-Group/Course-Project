Vue.component('reviewitem', {
    props: ['review'],
    template: '<div class=reviewitem>{{review.msg}}</div>'
})

Vue.component('tagitem', {
    props: ['tag'],
    template: '<div class=tagitem onclick="seeTag(event)">{{tag.name}}</div>'
})

var app = new Vue({
    el: '#app',
    data: {
        title: "",
        author: "",
        tags: [],
        content: "",
        reviews: []
    },
    methods: {
        seeTag: function(event) {
            console.log(event)
        }
    }
})

function get_article_id(_url) {
    var u = new URL(_url)
    var parts = u.pathname.split('/')
    return parts[parts.length - 1]
}

function load(id) {
    var xhr = new XMLHttpRequest()
    xhr.addEventListener("load", function(evt) {
        try {
            json = JSON.parse(xhr.responseText)
            render_article(json)
        } catch (error) {
            alert("发生内部错误")
        }
    })
    //xhr.open("GET", "/api/article/" + articleid)
    xhr.open("GET", "http://rest.apizza.net/mock/51c1b2b3f6568f3bda100dc161396207/api/article/" + id)
    xhr.send()
} 

function render_article(detail) {
    // console.log(detail.data)
    app.title = detail.data.Title 
    app.author = detail.data.User
    app.content = detail.data.Body

    atags = []
    for (var i = 0; i < detail.data.Tag.length; i ++) {
        atags.push({id: i, name: detail.data.Tag[i]})
    }
    app.tags = atags

    areviews = []
    for (var i = 0; i < detail.data.Review.length; i ++) {
        areviews.push({id: i, msg: detail.data.Review[i]})
    }
    app.reviews = areviews
}

var articleid = get_article_id(window.location.href)

window.onload = function() {
    // console.log(articleid)
    // this.load(articleid)
    this.load("1")
}

function seeAuthor() {
    window.location.href = "/user/" + app.author
}

function seeTag(event) {
    window.location.href = "/tag/" + event.target.innerText
}