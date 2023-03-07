-- SQLite login
DROP TABLE login;
CREATE TABLE login(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone string NOT NULL,
    email string NOT NULL,
    username string NOT NULL,
    password string NOT NULL,
    usertype string NOT NULL
);
INSERT INTO login (phone, email, username, password, usertype)
VALUES (
        '15765505517',
        '2941330150@qq.com',
        'tim_lrn2016',
        '@a20040207',
        'admin'
    );