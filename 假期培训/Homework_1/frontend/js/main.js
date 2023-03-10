function check() {
  let account = document.getElementById("account");
  let password = document.getElementById("password");
  if (account.value == "") {
    account.attributes["placeholder"].value = "请正确输入手机号或邮箱";
    account.classList.add("error");
  } else {
    account.attributes["placeholder"].value = "手机号或邮箱";
    account.classList.remove("error");
  }
  if (password.value == "" && account.value == "") {
    password.attributes["placeholder"].value = "请正确输入密码";
    password.classList.add("error");
  } else {
    password.attributes["placeholder"].value = "密码";
    password.classList.remove("error");
  }
}

function login() {
  let login_cont = document.getElementById("login_container");
  let loader = document.getElementById("loading");

  loader.style.display = "block";

  let account = document.getElementById("account").value;
  let password = document.getElementById("password").value;
  let p_test = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/;
  let em_test = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
  if (account == "" || password == "") {
    alert("手机号，邮箱或密码不能为空");
    loader.style.display = "none";
    return false;
  }

  if ((p_test.test(account) || em_test.test(account)) && password.length >= 6) {
    // TODO: post 发送到server
    // return true;
    alert("登录信息：\n账号：" + account + "\n密码：" + password);
    login_cont.classList.remove("fade_left");
    login_cont.classList.add("fade_right");
    loader.style.display = "none";
    login_cont.addEventListener("animationend", function () {
      window.location.href = "./article.html?username=" + account;
    });

    return true;
  } else {
    alert("请正确输入手机号或邮箱!");
    loader.style.display = "none";
    return false;
  }
}

function signup() {
  let login_cont = document.getElementById("login_container");
  login_cont.classList.remove("fade_left");
  login_cont.classList.add("fade_right");
  login_cont.addEventListener("animationend", function () {
    window.location.href = "./signup.html";
  });
}

const update_width = () => {
  let login_cont = document.getElementById("login_container");
  let loader = document.getElementById("loading");

  if (window.innerWidth < 800) {
    login_cont.style.width = "auto";
  } else {
    login_cont.style.width = "650px";
  }

  loader.style.width = login_cont.style.width;
  loader.style.height = login_cont.style.height;
};

window.onload = () => {
  update_width();
};
window.addEventListener("resize", update_width);
