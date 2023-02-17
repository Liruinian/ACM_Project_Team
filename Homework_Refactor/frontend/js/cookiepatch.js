function LoginCookieHandler(json) {
  json = JSON.parse(json);
  if (json.code == 2000) {
    var d = new Date();
    d.setTime(d.getTime() + 3 * 24 * 60 * 60 * 1000);
    document.cookie = "username=" + json.data.username + "; expires=" + d.toGMTString() + "; path=/";
    document.cookie = "login_token=" + json.data.login_token + "; expires=" + d.toGMTString() + "; path=/";
    document.cookie = "admintoken=" + json.data.admin_token + "; expires=" + d.toGMTString() + "; path=/";
  }
}
