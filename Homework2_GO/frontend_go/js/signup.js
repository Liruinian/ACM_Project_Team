function check() {
  let phone = document.getElementById("phonenum");
  let email = document.getElementById("email");

  if (password.value == "" && phone.value == "" && email.value == "") {
    password.attributes["placeholder"].value = "请输入密码";
    password.classList.add("error");
  } else {
    password.attributes["placeholder"].value = "密码";
    password.classList.remove("error");
  }
  if (phone.value == "") {
    phone.attributes["placeholder"].value = "请输入手机号";
    phone.classList.add("error");
  } else {
    phone.attributes["placeholder"].value = "手机号";
    phone.classList.remove("error");
  }
  if (phone.value == "" && email.value == "") {
    email.attributes["placeholder"].value = "请输入邮箱";
    email.classList.add("error");
  } else {
    email.attributes["placeholder"].value = "邮箱";
    email.classList.remove("error");
  }
}

function signup() {
  let signup_cont = document.getElementById("signup_container");
  let loader = document.getElementById("loading");
  loader.style.display = "block";
  let phone = document.getElementById("phonenum").value;
  let email = document.getElementById("email").value;
  let name = document.getElementById("name").value;
  let password = document.getElementById("password").value;
  let p_reminder = document.getElementById("pass_reminder");
  let p_test = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/;
  let em_test = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;

  if (phone == "" || email == "" || password == "") {
    alert("手机号，邮箱或密码不能为空");
    loader.style.display = "none";
    return false;
  }
  if (!p_test.test(phone)) {
    document.getElementById("phonenum").value = "";
    document.getElementById("phonenum").placeholder = "手机号格式不正确";
  }
  if (!em_test.test(email)) {
    document.getElementById("email").value = "";
    document.getElementById("email").placeholder = "邮箱格式不正确";
  }
  if (p_test.test(phone) && em_test.test(email) && !p_reminder.classList.contains("error")) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.open("POST", "http://127.0.0.1:8880/signup", true);
    httpRequest.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    httpRequest.send("phone=" + phone + "&email=" + email + "&username=" + name + "&password=" + password + "");
    httpRequest.onreadystatechange = function () {
      if (httpRequest.readyState == 4 && httpRequest.status == 200) {
        var return_data = httpRequest.responseText;
        if (return_data == '"success"') {
          signup_cont.classList.remove("fade_left");
          signup_cont.classList.add("fade_right");
          loader.style.display = "none";
          signup_cont.addEventListener("animationend", function () {
            window.location.href = "index.html";
          });
          return true;
        } else {
          alert(return_data);
          loader.style.display = "none";
          return false;
        }
      }
    };
  } else {
    alert("请检查注册信息是否有误!");
    loader.style.display = "none";
    return false;
  }
}

function back() {
  let signup_cont = document.getElementById("signup_container");
  signup_cont.classList.remove("fade_left");
  signup_cont.classList.add("fade_left_back");
  signup_cont.addEventListener("animationend", function () {
    signup_cont.style.opacity = 0;
    window.location.href = "index.html";
  });
}

const update_width = () => {
  let signup_cont = document.getElementById("signup_container");
  let loader = document.getElementById("loading");
  let wel_eng = document.getElementById("wel_eng");
  if (window.innerWidth < 800) {
    signup_cont.style.width = "auto";
    wel_eng.style.fontSize = "45px";
  } else {
    signup_cont.style.width = "650px";
    wel_eng.style.fontSize = "70px";
  }

  loader.style.width = signup_cont.style.width;
  loader.style.height = signup_cont.style.height;
};

window.onload = () => {
  update_width();

  document.getElementById("password").onkeyup = function () {
    let password = document.getElementById("password");
    let p_reminder = document.getElementById("pass_reminder");

    let strongRegex = new RegExp("^(?=.{8,})(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*\\W).*$", "g");
    let mediumRegex = new RegExp("^(?=.{7,})(((?=.*[A-Z])(?=.*[a-z]))|((?=.*[A-Z])(?=.*[0-9]))|((?=.*[a-z])(?=.*[0-9]))).*$", "g");
    let enoughRegex = new RegExp("(?=.{6,}).*", "g");

    if (password.value.match(strongRegex)) {
      p_reminder.innerHTML = "密码强度：强";
      p_reminder.classList.remove("error");
      p_reminder.style.color = "rgb(75, 191, 94)";
    } else if (password.value.match(mediumRegex)) {
      p_reminder.innerHTML = "密码强度：中";
      p_reminder.classList.remove("error");
      p_reminder.style.color = "rgb(112, 198, 252)";
    } else if (password.value.match(enoughRegex)) {
      p_reminder.innerHTML = "密码强度：弱";
      p_reminder.classList.remove("error");
      p_reminder.style.color = "rgb(236, 110, 45)";
    } else {
      p_reminder.innerHTML = "推荐密码为八位及以上，并且包括字母数字特殊字符三项";
      p_reminder.style.color = "red";
      p_reminder.classList.add("error");
    }
  };
};
window.addEventListener("resize", update_width);
