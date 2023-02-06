function logout() {
  window.location.href = "/logout";
}

function home() {
  window.location.href = "#";
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
    httpRequest.open("POST", "http://127.0.0.1:5000/article", true);
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

    searcha.href = "#" + art[1];
    searchli.innerHTML = art[1];
    searcha.appendChild(searchli);
    searchol.appendChild(searcha);
  }
}

var articles = "";

var httpRequest = new XMLHttpRequest();
httpRequest.open("GET", "http://127.0.0.1:5000/get-articles", true);
httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
httpRequest.send();
httpRequest.onreadystatechange = function () {
  if (httpRequest.readyState == 4 && httpRequest.status == 200) {
    articles = httpRequest.responseText;
    articles = JSON.parse(articles);
    console.log(articles);
  }
};

// articles.reverse();
function load_articles() {
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
    article_cont.id = art[1];
    a_title.innerHTML = art[1];
    a_category.innerHTML = " <span class='b_text'>No. " + art[0] + " </span>" + art[2];
    a_content.innerHTML = art[3];
    a_information.innerHTML =
      `
    <div class="a_i_item" id="a_author"><i class="fa fa-user"> 作者：` + art[4];
    +`</i></div>`;
    a_information.innerHTML +=
      `
    <div class="a_i_item" id="a_views"><i class="fa fa-eye"> 浏览量：` + art[6];
    +`</i></div>`;
    art_href.href = art[7];
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
  <div class="a_time" id="a_time"><i class="fa fa-clock-o"></i>&nbsp` + art[5];
    +`</div>`;
  }
  search_load(articles);
}

function add_art() {
  let a_title = document.getElementById("a_title_inp").value;
  let a_category = document.getElementById("a_category_inp").value;
  let a_content = document.getElementById("a_content_inp").value;
  let a_author = document.getElementById("a_author_inp").value;
  let a_time = document.getElementById("a_time_inp").value;
  let a_views = document.getElementById("a_views_inp").value;
  let a_href = document.getElementById("a_href_inp").value;
  let a_num = articles.length + 1;

  // let article = [a_num, a_title, a_category, a_content, a_author, a_time, a_views, a_href];
  // articles.reverse();
  // articles.push(article);
  // articles.reverse();
  // load_articles();

  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://127.0.0.1:5000/art-admin", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send(
    "change_type=" +
      "add" +
      "&id=" +
      a_num +
      "&title=" +
      a_title +
      "&category=" +
      a_category +
      "&content=" +
      a_content +
      "&author=" +
      a_author +
      "&time=" +
      a_time +
      "&views=" +
      a_views +
      "&href=" +
      a_href +
      ""
  );
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var return_data = httpRequest.responseText;
      if (return_data == "success") {
        location.reload();
      } else {
        alert(return_data);
      }
    }
  };
}

function rm_art() {
  let rm_art_num = prompt("请输入要删除的文章编号:");
  // let i = 0;
  // for (art of articles) {
  //   i++;
  //   if (art[0] == rm_art_num) {
  //     articles.splice(i - 1, 1);
  //   }
  // }
  // load_articles();

  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "http://127.0.0.1:5000/art-admin", true);
  httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  httpRequest.send("change_type=" + "del" + "&id=" + rm_art_num);
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var return_data = httpRequest.responseText;
      if (return_data == "success") {
        location.reload();
      } else {
        alert(return_data);
      }
    }
  };
}

function edit_art() {
  let edit_div = document.getElementById("editing");
  if (document.getElementById("a_edit_inp").value == "") {
    edit_div.style.display = "block";
    let edit_art_num = prompt("请输入要修改的文章编号:");
    let i = articles.length + 1;
    for (art of articles) {
      i--;
      if (art[0] == edit_art_num) {
        document.getElementById("a_edit_inp").value = art[0];
        document.getElementById("a_title_inp").value = art[1];
        document.getElementById("a_category_inp").value = art[2];
        document.getElementById("a_content_inp").value = art[3];
        document.getElementById("a_author_inp").value = art[4];
        document.getElementById("a_time_inp").value = art[5];
        document.getElementById("a_views_inp").value = art[6];
        document.getElementById("a_href_inp").value = art[7];
      }
    }
  } else {
    let a_edit = parseInt(document.getElementById("a_edit_inp").value);
    let a_title = document.getElementById("a_title_inp").value;
    let a_category = document.getElementById("a_category_inp").value;
    let a_content = document.getElementById("a_content_inp").value;
    let a_author = document.getElementById("a_author_inp").value;
    let a_time = document.getElementById("a_time_inp").value;
    let a_views = document.getElementById("a_views_inp").value;
    let a_href = document.getElementById("a_href_inp").value;
    edit_div.style.display = "none";
    // for (art of articles) {
    //   if (art[0] == a_edit) {
    //     art[1] = a_title;
    //     art[2] = a_category;
    //     art[3] = a_content;
    //     art[4] = a_author;
    //     art[5] = a_time;
    //     art[6] = a_views;
    //     art[7] = a_href;
    //   }
    // }

    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", "http://127.0.0.1:5000/art-admin", true);
    httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    httpRequest.send(
      "change_type=" +
        "edit" +
        "&id=" +
        a_edit +
        "&title=" +
        a_title +
        "&category=" +
        a_category +
        "&content=" +
        a_content +
        "&author=" +
        a_author +
        "&time=" +
        a_time +
        "&views=" +
        a_views +
        "&href=" +
        a_href +
        ""
    );
    httpRequest.onreadystatechange = function () {
      if (httpRequest.readyState == 4 && httpRequest.status == 200) {
        var return_data = httpRequest.responseText;
        if (return_data == "success") {
          location.reload();
        } else {
          alert(return_data);
        }
      }
    };

    document.getElementById("a_edit_inp").value = "";
    document.getElementById("a_title_inp").value = "";
    document.getElementById("a_category_inp").value = "";
    document.getElementById("a_content_inp").value = "";
    document.getElementById("a_author_inp").value = "";
    document.getElementById("a_time_inp").value = "";
    document.getElementById("a_views_inp").value = "";
    document.getElementById("a_href_inp").value = "";

    load_articles();
  }
}

function change() {
  let change = document.getElementById("change");
  if (change.classList.contains("change_open")) {
    change.classList.remove("change_open");
    change.classList.add("change_close");
  } else {
    change.classList.remove("change_close");
    change.classList.add("change_open");
  }
}

function user_classify() {
  // user admin 后端判断后返回
  if (document.getElementById("usrgroup_s").innerText != "admin") {
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
