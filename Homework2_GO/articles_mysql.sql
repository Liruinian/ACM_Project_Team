USE articles;
CREATE TABLE IF NOT EXISTS articles(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    title char(100) NOT NULL,
    category char(50) NOT NULL,
    content TEXT NOT NULL,
    author char(30) NOT NULL,
    time char(30) NOT NULL,
    views char(10) NOT NULL,
    href char(100) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;