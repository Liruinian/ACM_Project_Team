from flask import Flask, redirect, url_for, request, render_template, make_response
import sqlite3

app = Flask(__name__)


@app.route('/')
def index():
    return render_template("index.html")


@app.route('/login', methods=['POST'])
def login():
    user = request.form['username']
    password = request.form['password']
    login_type = request.form['login_type']
    print(user)
    print(password)
    print(login_type)
    data = lookdata(login_type, user)
    ip = request.remote_addr
    if data:
        if data[4] == password:
            res = make_response("success")
            res.set_cookie("login_flag", "1", max_age=3600)
            res.set_cookie("username", data[3], max_age=3600)
            res.set_cookie("usrgroup", data[5], max_age=3600)
            print("用户登录：" + ip+"\n用户名：" + data[3] + "\n用户组：" + data[5])
            # TODO: 加密
            return res
        else:
            return "wrong password"
    else:
        return "invalid username"


@app.route('/logout')
def logout():
    res = redirect(url_for("index"))
    res.delete_cookie("login_flag")
    res.delete_cookie("username")
    res.delete_cookie("usrgroup")
    ip = request.remote_addr
    print("用户登出：" + ip)
    return res


@app.route('/art-admin', methods=['POST'])
def art_admin():
    if (request.cookies.get("usrgroup") == "admin"):
        change = request.form['change_type']
        print(change)
        if (change == "add"):
            title = request.form['title']
            category = request.form['category']
            content = request.form['content']
            author = request.form['author']
            time = request.form['time']
            views = request.form['views']
            href = request.form['href']
            conn = sqlite3.connect("articles.db")
            c = conn.cursor()
            c.execute("INSERT INTO articles (title,category,content,author,time,views,href) VALUES ('%s','%s','%s','%s','%s','%s','%s')" % (
                title, category, content, author, time, views, href))
            conn.commit()
            conn.close()
            ip = request.remote_addr
            print("用户增加文章：" + ip+"\n文章标题：" + title + "\n文章分类：" + category + "\n文章内容" + content +
                  "\n文章作者：" + author + "\n文章时间：" + time + "\n文章浏览量：" + views + "\n文章链接：" + href)

            return "success"
        else:
            if (change == "del"):
                id = request.form['id']
                conn = sqlite3.connect("articles.db")
                c = conn.cursor()
                c.execute("DELETE FROM articles WHERE id = %s" % id)
                conn.commit()
                conn.close()
                ip = request.remote_addr
                print("用户删除文章：" + ip+"\n文章id："+id)

                return "success"
            else:
                if (change == "edit"):
                    id = request.form['id']
                    title = request.form['title']
                    category = request.form['category'].replace(" ", "+")
                    content = request.form['content']
                    author = request.form['author']
                    time = request.form['time']
                    title = request.form['title']
                    category = request.form['category']
                    content = request.form['content']
                    author = request.form['author']
                    time = request.form['time']
                    views = request.form['views']
                    href = request.form['href']
                    conn = sqlite3.connect("articles.db")

                    c = conn.cursor()
                    c.execute("UPDATE articles SET title='%s',category='%s',content='%s',author='%s',time='%s',views='%s',href='%s' WHERE id='%s'" % (
                        title, category, content, author, time, views, href, id))
                    conn.commit()
                    conn.close()
                    ip = request.remote_addr
                    print("用户修改文章：" + ip+"\n文章标题：" + title + "\n文章分类：" + category + "\n文章内容" + content +
                          "\n文章作者：" + author + "\n文章时间：" + time + "\n文章浏览量：" + views + "\n文章链接：" + href)
                    return "success"
                else:
                    return "Unknown Method"
    else:
        return "forbidden"


@app.route('/article', methods=['POST', 'GET'])
def article():
    if (request.cookies.get("login_flag") == "1"):
        search = request.form.get("search")
        if (search == None):
            user = request.cookies.get("username")
            group = request.cookies.get("usrgroup")
            conn = sqlite3.connect("articles.db")
            c = conn.cursor()
            c.execute('SELECT * FROM articles')
            data = c.fetchall()
            conn.close()
            print(data)
            return render_template("article.html", username=user, usrgroup=group, articles_f=data)
        else:
            conn = sqlite3.connect("articles.db")
            c = conn.cursor()
            c.execute(
                'SELECT * FROM articles where title like "%' + search + '%"')
            data = c.fetchall()
            conn.close()
            print(data)
            return (data)
    else:
        return redirect(url_for("index"))


@app.route('/signup', methods=['POST', 'GET'])
def signup():
    if (request.method == "POST"):
        phone = request.form['phone']
        email = request.form['email']
        username = request.form['username']
        password = request.form['password']

        return adduser(phone, email, username, password)
    else:
        return render_template("signup.html")


def lookdata(type, value):
    conn = sqlite3.connect('login.db')
    c = conn.cursor()
    if (type == "1"):
        print("use phone")
        c.execute("SELECT * FROM login WHERE phone='%s'" % value)
    if (type == "2"):
        print("use email")
        c.execute("SELECT * FROM login WHERE email='%s'" % value)
    data = c.fetchone()
    conn.close()
    print(data)
    return data


def exist(type, value):
    conn = sqlite3.connect('login.db')
    c = conn.cursor()
    sql = "SELECT * FROM login WHERE '%s'='%s'" % (type, value)
    c.execute(sql)
    data = c.fetchall()
    conn.close()

    if (len(data) == 0):
        return False
    else:
        return True


def adduser(phone, email, username, password):
    conn = sqlite3.connect('login.db')
    c = conn.cursor()
    print(c.execute("SELECT * FROM login WHERE phone='%s'" % phone))
    if (exist("phone", phone) == False):
        if (exist("email", email) == False):
            if (exist("username", username) == False):
                sql = "INSERT INTO login(phone, email, username, password, usertype) VALUES('%s', '%s', '%s','%s','%s')" % (
                    phone, email, username, password, 'user')
                res = c.execute(sql)
                conn.commit()
                conn.close()
                ip = request.remote_addr
                print("用户注册"+ip+"\n用户名：" + username + "\n密码：" +
                      password + "\n邮箱："+email+"\n手机号："+phone+"\n用户类型：user"+"\nres:"+str(res))

                return "success"
            else:
                return "username already exists"
        else:
            return "email already exists"
    else:
        return "phone already exists"


if __name__ == '__main__':
    app.run()
