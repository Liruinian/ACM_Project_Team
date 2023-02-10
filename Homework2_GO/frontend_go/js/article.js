function logout() {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://101.43.177.2:8880/logout", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var json = httpRequest.responseText;
      if (json != '"success"') {
        alert(json);
      } else {
        window.location.href = "index.html";
      }
    }
  };
}
function change() {
  window.location.href = "editarticle.html";
}
function home() {
  window.location.href = "article.html";
}
function search() {
  let search = document.getElementById("search");
  if ((search.style.height = "0px")) {
    search.style.height = "";
  }
  if (search.classList.contains("search_open")) {
    search.classList.remove("search_open");
    search.classList.add("search_close");
  } else {
    search.classList.remove("search_close");
    search.classList.add("search_open");
  }
}
function search_inp() {
  si = document.getElementById("search_inp");
  if (si.value != "") {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", "http://101.43.177.2:8880/article", true);
    httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    httpRequest.send("search=" + si.value);
    httpRequest.onreadystatechange = function () {
      if (httpRequest.readyState == 4 && httpRequest.status == 200) {
        var return_data = httpRequest.responseText;
        return_data = JSON.parse(return_data);
        search_load(return_data);
      }
    };
  } else {
    search_load(articles);
  }
}

function search_load(articles) {
  let searchol = document.getElementById("searchol");
  searchol.innerHTML = "";

  for (art of articles) {
    let searchli = document.createElement("li");
    let searcha = document.createElement("a");
    searchli.classList.add("menuitem");
    searchli.classList.add("m_text");

    searcha.href = "#" + art.title;
    searchli.innerHTML = art.title;
    searcha.appendChild(searchli);
    searchol.appendChild(searcha);
  }
}
var articles = "";

function load_articles() {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://101.43.177.2:8880/get-articles", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var jsonart = httpRequest.responseText;
      if (jsonart == '"Please Login"') {
        window.location.href = "index.html";
      }
      articles = JSON.parse(JSON.parse(jsonart));
      console.log(articles);
      articles.reverse();
      let i = articles.length + 1;
      let page_cont = document.getElementById("page_container");

      page_cont.innerHTML = "";

      for (art of articles) {
        i--;
        let article_cont = document.createElement("div");
        let art_href = document.createElement("a");
        let a_title = document.createElement("div");
        let a_category = document.createElement("span");
        let a_content = document.createElement("div");
        let a_information = document.createElement("div");
        let a_categ_cont = document.createElement("div");

        article_cont.classList.add("article_container");
        a_title.classList.add("a_title");
        a_category.classList.add("a_category");
        a_content.classList.add("a_content");
        a_information.classList.add("a_information");
        a_categ_cont.classList.add("a_categ_cont");
        article_cont.id = art.title;
        a_title.innerHTML = art.title;
        a_category.innerHTML = " <span class='b_text'>No. " + art.id + " </span>" + art.category;
        a_content.innerHTML = marked(art.content).replace(/\n/g, "<br />");
        a_information.innerHTML =
          `
    <div class="a_i_item" id="a_author"><i class="fa fa-user"> 作者：` + art.author;
        +`</i></div>`;
        a_information.innerHTML +=
          `
    <div class="a_i_item" id="a_views"><i class="fa fa-eye"> 浏览量：` + art.views;
        +`</i></div>`;
        art_href.href = art.href;
        art_href.target = "_blank";

        a_categ_cont.appendChild(a_category);
        article_cont.appendChild(a_categ_cont);
        article_cont.appendChild(a_title);
        article_cont.appendChild(a_content);
        article_cont.appendChild(a_information);
        art_href.appendChild(article_cont);
        page_cont.appendChild(art_href);
        a_categ_cont.innerHTML +=
          `
  <div class="a_time" id="a_time"><i class="fa fa-clock-o"></i>&nbsp` + art.time;
        +`</div>`;
      }
      search_load(articles);
    }
  };
}

function user_classify() {
  // user admin 后端判断后返回到cookies

  let uInfo = "";
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://101.43.177.2:8880/userinfo", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var jsonart = httpRequest.responseText;
      uInfo = JSON.parse(JSON.parse(jsonart));
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

const update_width = () => {
  let page_cont = document.getElementById("page_container");
  let menu = document.getElementById("menu");
  let logo = document.getElementsByClassName("logo")[0];

  page_cont.style.height = window.innerHeight - 20 + "px";
  menu.style.height = window.innerHeight - 140 + "px";
  if (window.innerWidth < 1020) {
    page_cont.style.width = "auto";
    if (window.innerWidth < 600) {
      menu.style.display = "none";
      logo.style.display = "none";
      page_cont.style.marginLeft = "0px";
    } else {
      menu.style.display = "block";
      logo.style.display = "block";
      page_cont.style.marginLeft = "220px";
    }
  } else {
    page_cont.style.width = "700px";
    page_cont.style.marginLeft = window.innerWidth / 2 - 300 + "px";
  }
};

window.onload = () => {
  update_width();
  load_articles();
  user_classify();
};
window.addEventListener("resize", update_width);
