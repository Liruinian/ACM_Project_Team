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
  httpRequest.open("PUT", "http://8.130.53.145:8880/articles", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.send(JSON.stringify(artform));
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      alert(json.status);
      location.reload();
    }
  };
}
function delart() {
  var id = document.getElementById("edit_art_id").innerHTML;
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("DELETE", "http://8.130.53.145:8880/articles/" + id, true);
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      alert(json.status);
      location.reload();
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
    id: id,
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
  httpRequest.open("PUT", "http://8.130.53.145:8880/articles", true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.send(JSON.stringify(artform));
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      alert(json.status);
      location.reload();
    }
  };
}
function home() {
  window.location.href = "article.html";
}
function logout() {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://8.130.53.145:8880/logout", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      if (json.status != "success") {
        alert(json.status);
      } else {
        window.location.href = "index.html";
      }
    }
  };
}
function load_articles() {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("GET", "http://8.130.53.145:8880/articles", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var jsonart = httpRequest.responseText;
      jsonart = JSON.parse(jsonart);
      if (jsonart.status != undefined) {
        alert(jsonart.status);
        window.location.href = "index.html";
      }
      articles = JSON.parse(jsonart);
      console.log(articles);
      articles.reverse();
      let searchol = document.getElementById("searchol");
      searchol.innerHTML = "";

      for (art of articles) {
        let getart = getQueryString("art");
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

function user_classify() {
  // user admin 后端判断后返回到cookies

  let uInfo = "";
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("GET", "http://8.130.53.145:8880/userinfo", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      json = JSON.parse(json);
      if (json.status != undefined) {
        return;
      }
      uInfo = JSON.parse(json);
    }

    let usrgroupE = document.getElementById("usrgroup_s");
    let usernameE = document.getElementById("username_s");
    usrgroupE.innerHTML = uInfo.usertype;
    usernameE.innerHTML = uInfo.username;
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