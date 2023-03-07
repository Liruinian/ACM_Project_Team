-- SQLite articles
DROP TABLE articles;
CREATE TABLE articles(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title string NOT NULL,
    category string NOT NULL,
    content string NOT NULL,
    author string NOT NULL,
    time string NOT NULL,
    views string NOT NULL,
    href string NOT NULL
);
INSERT INTO articles (
        title,
        category,
        content,
        author,
        time,
        views,
        href
    )
VALUES (
        "只是一个测试",
        "Text Test",
        "测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage测试文本TestMessage",
        "Amazing",
        "2023-01-19 19:39:00",
        "111",
        "#"
    );
INSERT INTO articles (
        title,
        category,
        content,
        author,
        time,
        views,
        href
    )
VALUES (
        "文章概要中可以插入html标签",
        "HTML Test",
        "blockquote test<br><blockquote>这里是blockquote属性测试<br>插入html标签测试</blockquote>",
        "Amazing",
        "2023-01-19 19:39:00",
        "111",
        "#"
    );
INSERT INTO articles (
        title,
        category,
        content,
        author,
        time,
        views,
        href
    )
VALUES (
        "更多的html标签",
        "Advanced HTML",
        "甚至可以试一试<pre>pre标签</pre><i class='fa fa-smile-o'> i标签</i>",
        "Amazing",
        "2023-01-19 19:39:00",
        "111",
        "#"
    );
INSERT INTO articles (
        title,
        category,
        content,
        author,
        time,
        views,
        href
    )
VALUES (
        "文章的标题4",
        "Category",
        "测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本测试文本",
        "Amazing",
        "2023-01-19 19:39:00",
        "111",
        "#"
    );