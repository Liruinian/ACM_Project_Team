CREATE TABLE IF NOT EXISTS login(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    phone VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    usertype VARCHAR(30) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO login VALUES(1,15765505517,'2941330150@qq.com','tim_lrn2016','@a20040207','admin');
INSERT INTO login VALUES(2,18846089512,'15204635097@163.com','usertest','ABCabc123456','user');