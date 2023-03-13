# NOLOGIN NEAUACMSYS TIMETABLE PROJ.
# A19220064
# 20230312

from flask import Flask, redirect, url_for, request, render_template, make_response
from flask_cors import CORS, cross_origin
import requests
import sqlite3
import json
import re

ACMURLlogin = "https://www.neauacm.cn/login/index.php"
ACMURLupload = "https://www.neauacm.cn/upload_file/"
NEAULogin = "http://zhjwxs.neau.edu.cn/j_spring_security_check"
NEAUTimeTable = "http://zhjwxs.neau.edu.cn/student/courseSelect/thisSemesterCurriculum/callback"

user_id = "A19220064"
ACMpassword = "b0b388ee173661219b7f39c4aa716eb5"

global gcookies
gcookies = {}

app = Flask(__name__)
CORS(app, resources=r'/*')


@app.route('/')
def index():
    f = open(r'cookies.txt', 'r')
    cookies = f.read()
    f.close()
    f = open(r'ddllist.txt', 'r')
    ddl = ""
    for line in f.read().split('\n'):
        dname, dtime, disdone = line.strip().split('|')
        if disdone == "x":
            ddl += "<li>"+dname+" "+dtime+"</li>"
        else:
            ddl += "<li><del>"+dname+" "+dtime+"</del></li>"
    f.close()

    return render_template("index.html", projtodo=checkUpload(), neaujson=loginNEAUServer(), cookies=cookies, acmloginuser=user_id, ddllist=ddl)


@app.route('/update-cookies', methods=["POST"])
def upd_cookies():

    f = open(r'cookies.txt', "w")
    f.write(request.data.decode())
    f.close
    return "success"


def loginNEAUServer():
    f = open(r'cookies.txt', 'r')
    cookies = {}
    for line in f.read().split(';'):
        name, value = line.strip().split('=', 1)
        cookies[name] = value
    resp = requests.post(NEAUTimeTable, cookies=cookies,
                         verify=False, timeout=10)
    serverres = resp.text
    return serverres


def getCenterString(s):
    rule = r'<td>(.*?)</td>'
    slotList = re.findall(rule, s)
    return slotList


def loginACMTeam():
    data = {
        'user_id': (None, user_id),
        'password': (None, ACMpassword),
    }
    resp = requests.post(ACMURLlogin, files=data, verify=False, timeout=10)
    cookies = resp.cookies
    global gcookies
    gcookies = cookies
    print("登录成功")
    if resp.text == 'ok':
        return cookies
    else:
        print("登录失败")
        return None


def checkUpload():
    global gcookies
    if gcookies != {}:
        cookies = gcookies
        print("use cached cookies")
    else:
        cookies = loginACMTeam()
    resp = requests.get(ACMURLupload, cookies=cookies)
    text = resp.text
    print(cookies)
    text = text.split("<tbody>")[1]
    text = text.split("</tbody>")[0]
    sped = text.split("<tr>")[1:]

    list1 = []
    for i in range(len(sped)):
        if i < 5:
            list1.append(getCenterString(sped[i]))
    print(list1)
    return list1


if __name__ == '__main__':
    app.debug = True
    app.run()
