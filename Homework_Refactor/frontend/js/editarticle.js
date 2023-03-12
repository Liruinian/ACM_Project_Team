function getCookie(cname) {
  var name = cname + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i].trim();
    if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
  }
  return "";
}

var editor;
$(function () {
  editor = editormd("article-editor", {
    theme: "dark",
    previewTheme: "dark",
    editorTheme: "pastel-on-dark",
    width: "100%",
    height: 600,
    emoji: false,
    // markdown: "xxxx",     // dynamic set Markdown text
    path: "lib/", // Autoload modules mode, codemirror, marked... dependents libs path
  });
});
function newart() {
  var title = document.getElementById("a_title_inp").value;
  var category = document.getElementById("a_category_inp").value;
  var author = document.getElementById("a_author_inp").value;
  var time = document.getElementById("a_time_inp").value;
  var views = document.getElementById("a_views_inp").value;
  var href = document.getElementById("a_href_inp").value;
  var content = editor.getMarkdown();

  var artform = {
    edit: false,
    title: title,
    category: category,
    author: author,
    time: time,
    views: views,
    href: href,
    content: content,
  };
  console.log(JSON.stringify(artform));
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/create", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send(JSON.stringify(artform));
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      iziToast.show({
        timeout: 2500,
        icon: false,
        title: "服务器返回值：",
        message: json.msg,
      });
      setTimeout(function () {
        location.reload();
      }, 3000);
    }
  };
}
function delart() {
  var id = document.getElementById("edit_art_id").innerHTML;
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("DELETE", "https://api.liruinian.top/article/delete/" + id, true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      iziToast.show({
        timeout: 2500,
        icon: false,
        title: "服务器返回值：",
        message: json.msg,
      });
      setTimeout(function () {
        location.reload();
      }, 3000);
    }
  };
}
function submit() {
  var id = document.getElementById("edit_art_id").innerHTML;
  var title = document.getElementById("a_title_inp").value;
  var category = document.getElementById("a_category_inp").value;
  var author = document.getElementById("a_author_inp").value;
  var time = document.getElementById("a_time_inp").value;
  var views = document.getElementById("a_views_inp").value;
  var href = document.getElementById("a_href_inp").value;
  var content = editor.getMarkdown();
  var artform = {
    edit: true,
    id: parseInt(id),
    title: title,
    category: category,
    author: author,
    time: time,
    views: views,
    href: href,
    content: content,
  };

  console.log(JSON.stringify(artform));
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/edit", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send(JSON.stringify(artform));
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      iziToast.show({
        timeout: 2500,
        icon: false,
        title: "服务器返回值：",
        message: json.msg,
      });
      setTimeout(function () {
        location.reload();
      }, 3000);
    }
  };
}
function home() {
  window.location.href = "article.html";
}
function logout() {
  document.cookie = "username=; expires=0; path=/";
  document.cookie = "login_token=; expires=0; path=/";
  document.cookie = "admin_token=; expires=0; path=/";
  window.location.href = "index.html";
}
function load_articles() {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/list", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var jsonart = httpRequest.responseText;
      jsonart = JSON.parse(jsonart);
      if (jsonart.msg != "success") {
        alert(jsonart.msg);
        window.location.href = "index.html";
      }
      articles = jsonart.Articles;
      console.log(articles);
      articles.reverse();
      let searchol = document.getElementById("searchol");
      searchol.innerHTML = "";

      let getart = getQueryString("art");
      console.log(getart);
      if (getart == null) {
        document.getElementById("a_author_inp").value = username;
        var today = new Date();
        var DD = String(today.getDate()).padStart(2, "0");
        var MM = String(today.getMonth() + 1).padStart(2, "0");
        var yyyy = today.getFullYear();
        hh = String(today.getHours()).padStart(2, "0");
        mm = String(today.getMinutes()).padStart(2, "0");
        ss = String(today.getSeconds()).padStart(2, "0");
        today = yyyy + "-" + MM + "-" + DD + " " + hh + ":" + mm + ":" + ss;
        document.getElementById("a_time_inp").value = today;
      }
      for (art of articles) {
        if (getart == art.id) {
          console.log(art.content);
          document.getElementById("edit_art_id").innerHTML = art.id;
          document.getElementById("a_title_inp").value = art.title;
          document.getElementById("a_category_inp").value = art.category;
          document.getElementById("a_author_inp").value = art.author;
          document.getElementById("a_time_inp").value = art.time;
          document.getElementById("a_views_inp").value = art.views;
          document.getElementById("a_href_inp").value = art.href;
          editor = editormd("article-editor", {
            theme: "dark",
            previewTheme: "dark",
            editorTheme: "pastel-on-dark",
            width: "100%",
            height: 600,
            emoji: false,
            markdown: art.content,
            path: "lib/",
          });
        }
        let searchli = document.createElement("li");
        let searcha = document.createElement("a");
        searchli.classList.add("menuitem");
        searchli.classList.add("m_text");

        searcha.href = "editarticle.html?art=" + art.id;
        searchli.innerHTML = art.title;
        searcha.appendChild(searchli);
        searchol.appendChild(searcha);
      }
    }
  };
}
var username;
function user_classify() {
  // user admin 后端判断后返回到cookies

  let uInfo = "";
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/user/info", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      uInfo = JSON.parse(json);
      console.log(uInfo);
      if (uInfo.msg == "success") {
        username = uInfo.AForm[0].username;
        let usrgroupE = document.getElementById("usrgroup_s");
        let usernameE = document.getElementById("username_s");
        usrgroupE.innerHTML = uInfo.AForm[0].usertype;
        usernameE.innerHTML = username;
        if (usrgroupE.innerText != "admin") {
          let admin_comp = document.getElementsByClassName("admin");
          for (adm of admin_comp) {
            adm.style.display = "none";
          }
        } else {
          let admin_comp = document.getElementsByClassName("admin");
          for (adm of admin_comp) {
            adm.style.display = "block";
          }
        }
      }
    }
  };
}

function getQueryString(name) {
  var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
  var r = window.location.search.substr(1).match(reg);
  if (r != null) return unescape(r[2]);
  return null;
}

window.onload = () => {
  user_classify();
  load_articles();
};
