function getCookie(cname) {
  var name = cname + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i].trim();
    if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
  }
  return "";
}

function logout() {
  document.cookie = "username=; expires=0; path=/";
  document.cookie = "login_token=; expires=0; path=/";
  document.cookie = "admin_token=; expires=0; path=/";
  window.location.href = "index.html";
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
  var keyWord = document.getElementById("search_inp").value;
  list = articles;
  var arr = [];
  for (var i = 0; i < list.length; i++) {
    if (list[i].title.indexOf(keyWord) >= 0) {
      arr.push(list[i]);
    }
  }

  for (var i = 0; i < list.length; i++) {
    if (list[i].content.indexOf(keyWord) >= 0) {
      if (arr.length != 0) {
        for (var j = 0; j < arr.length; j++) {
          if (arr[j].id == list[i].id) {
            break;
          }
          if (j == arr.length - 1) {
            arr.push(list[i]);
          }
        }
      } else {
        arr.push(list[i]);
      }
    }
  }

  search_load(arr);
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
        return;
      }
      articles = jsonart.Articles;
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
        let comment_cont = document.createElement("div");

        article_cont.classList.add("article_container");
        a_title.classList.add("a_title");
        a_category.classList.add("a_category");
        a_content.classList.add("a_content");
        a_information.classList.add("a_information");
        a_categ_cont.classList.add("a_categ_cont");
        comment_cont.classList.add("comment_container");

        comment_cont.innerHTML +=
          `        <div class="leave_comment">
        <textarea class="comment_inp" id="comment_inp_` +
          art.id +
          `" type="text" placeholder="留个评论吧~"></textarea>
        <div class="comment_sub" onclick="comment(` +
          art.id +
          `)" >提交</div>
      </div>`;
        article_cont.id = art.title;
        a_title.innerHTML = art.title;
        a_category.innerHTML = " <span class='b_text'>No. " + art.id + " </span>" + art.category;
        a_content.innerHTML = marked(art.content).replace("\n", "<br />");
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
        art_href.appendChild(a_categ_cont);
        art_href.appendChild(a_title);
        art_href.appendChild(a_content);
        art_href.appendChild(a_information);
        article_cont.appendChild(art_href);
        article_cont.appendChild(comment_cont);
        //Here
        article_cont.onmouseover = getComment(art.id, comment_cont);
        page_cont.appendChild(article_cont);
        a_categ_cont.innerHTML +=
          `
  <div class="a_time" id="a_time"><i class="fa fa-clock-o"></i>&nbsp` + art.time.replace("T", " ");
        +`</div>`;
      }
      search_load(articles);
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

        document.getElementById("loading-box").classList.add("loaded");

        setTimeout(function () {
          document.getElementById("logo-center").style.display = "none";
          document.getElementById("logo").classList.add("logo_animate");
          document.getElementById("menu").classList.add("menu_animate");
          document.getElementById("page_container").classList.add("page_animate");
        }, 200);
        iziToast.show({
          timeout: 2500,
          icon: false,
          title: "欢迎回家！" + username,
          message: "成功登入PRTS信息整合终端",
        });
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

function getComment(id, comment_cont) {
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/comments/" + id, true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var return_data = httpRequest.responseText;
      return_data = JSON.parse(return_data);
      if (return_data.msg == "success") {
        if (return_data.Comments != null) {
          for (com of return_data.Comments) {
            comment_cont.innerHTML +=
              `
          <div class="a_comment">
            <div class="avatar"><i class="fa fa-2x fa-user-circle"></i></div>
            <div class="comment_text">
              <div class="comment_title">
                <div class="c_user">` +
              com.user +
              `</div><div>cid:</div><div id="commentid">` +
              com.id +
              `</div>
                <div class="comment_like" onclick="likecomment(` +
              com.id +
              `,` +
              com.thumbUp +
              `)"><i class="fa fa-thumbs-up"> </i><span id="comment_like_` +
              com.id +
              `">` +
              com.thumbUp +
              `</span></div>
              </div>
              <div class="c_content">` +
              com.commentText +
              `</div><div class="a_comment_time">` +
              com.time.replace("T", " ") +
              `</div>
              <hr />
            </div>
          </div>`;
          }
        }
      } else {
        alert(return_data.msg);
        window.location = "index.html";
      }
    }
  };
}

function timestampToTime(times) {
  let time = times[1];
  let mdy = times[0];
  mdy = mdy.split("/");
  let month = parseInt(mdy[0]);
  let day = parseInt(mdy[1]);
  let year = parseInt(mdy[2]);
  return year + "-" + month + "-" + day + "T" + time;
}

function comment(id) {
  var inp = document.getElementById("comment_inp_" + id);
  var commentText = inp.value;
  let time = new Date();
  let nowTime = timestampToTime(time.toLocaleString("en-US", { hour12: false }).split(" "));

  var SendData = {
    commentText: commentText,
    time: nowTime,
  };
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/create-comment/" + id, true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send(JSON.stringify(SendData));
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var return_data = httpRequest.responseText;
      return_data = JSON.parse(return_data);
      console.log(return_data);
      if (return_data.msg == "success") {
        iziToast.show({
          timeout: 2500,
          icon: false,
          title: "评论创建成功！",
          message: commentText,
        });
        setTimeout(function () {
          window.location.reload();
        }, 2500);
      } else {
        alert(return_data.msg);
        window.location = "index.html";
      }
    }
  };
}

function likecomment(id, thumpUp) {
  var c_like = document.getElementById("comment_like_" + id);
  var httpRequest = new XMLHttpRequest();
  httpRequest.open("POST", "https://api.liruinian.top/article/like-comment/" + id, true);
  httpRequest.setRequestHeader("Content-type", "application/raw");
  httpRequest.setRequestHeader("username", getCookie("username"));
  httpRequest.setRequestHeader("authorization", getCookie("login_token"));
  httpRequest.setRequestHeader("adminauth", getCookie("admin_token"));
  httpRequest.send();
  httpRequest.onreadystatechange = function () {
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {
      var return_data = httpRequest.responseText;
      return_data = JSON.parse(return_data);
      console.log(return_data);
      if (return_data.msg == "success") {
        iziToast.show({
          timeout: 2500,
          icon: false,
          title: "点赞评论成功！",
          message: "当前评论赞数：" + String(parseInt(c_like.innerHTML) + 1),
        });

        c_like.innerHTML = String(parseInt(c_like.innerHTML) + 1);
      } else {
        alert(return_data.msg);
        window.location = "index.html";
      }
    }
  };
}
