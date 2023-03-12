# NOLOGIN NEAUACMSYS TIMETABLE PROJ.
# A19220064
# 20230312

from flask import Flask, redirect, url_for, request, render_template, make_response
from flask_cors import CORS, cross_origin
import requests
import sqlite3
import json
import re

URLlogin = "https://www.neauacm.cn/login/index.php"
URLupload = "https://www.neauacm.cn/upload_file/"

user_id = "A19220064"
password = "b0b388ee173661219b7f39c4aa716eb5"

app = Flask(__name__)
CORS(app, resources=r'/*')


@app.route('/')
def index():
    return render_template("index.html")


def getCenterString(s):
    rule = r'<td>(.*?)</td>'
    slotList = re.findall(rule, s)
    return slotList

def loginFunc():
    data = {
    'user_id':(None, user_id),
    'password': (None, password),
    }
    resp = requests.post(URLlogin, files=data,verify=False, timeout=10)
    cookies = resp.cookies
    print(cookies)
    print("登录成功")

    return resp.text,cookies

def checkUpload():
    returnText,cookies = loginFunc()
    if returnText == 'ok':
        resp = requests.get(URLupload,cookies=cookies)
        text = resp.text
        text = text.split("<tbody>")[1]
        text = text.split("</tbody>")[0]
        sped = text.split("<tr>")[1:]

        list1 = []
        for i in range(len(sped)):
            list1.append(getCenterString(sped[i]))
        print(list1)


if __name__ == '__main__':
    app.debug = True
    app.run()