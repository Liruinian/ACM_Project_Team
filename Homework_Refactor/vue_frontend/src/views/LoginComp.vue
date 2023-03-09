<template>
          <div id="login_container" class="login_cont fade_left">
        <div class="m_password_login underlined">密码登录 <span class="notetext">Password Login</span></div>
        <div class="right">
          <div class="d_password_login">
            <input type="text" id="account" class="input" name="account" onchange="check()" placeholder="手机号或邮箱" /><br />
            <input type="password" id="password" class="input" name="password" onchange="check()" placeholder="密码" />
            <div class="pass_tray">
              <span class="signup_url" id="signup"><span style="cursor: pointer" onclick="signup()">注册 PRTS</span></span>
              <span class="forgot text" id="forget_pass"><span style="cursor: pointer" onclick="signup()">忘记密码</span></span>
            </div>
            <input type="submit" value="submit" id="submit" style="display: none" />
            <div class="login" id="login" onclick="login()">登录</div>
          </div>
          <div class="d_other_methods">
            <div class="divider">
              <div class="line"></div>
              <div class="notetext">其他方式登录</div>
              <div class="line"></div>
            </div>
            <div class="method_login_cont">
              <a href="#wechat">
                <div class="circle wechat" id="wechat_login">
                  <i class="fa fa-wechat"></i>
                </div> </a
              ><a href="#qq">
                <div class="circle qq" id="qq_login">
                  <i class="fa fa-qq"></i>
                </div> </a
              ><a href="#weibo">
                <div class="circle weibo" id="weibo_login">
                  <i class="fa fa-weibo"></i>
                </div>
              </a>
            </div>
          </div>
          <div class="notetext">登录即代表同意《PRTS协议》《隐私保护指引》</div>
        </div>
    </div>
</template>

<style scoped>
.login_cont{
  margin: 0 auto;
  overflow: hidden;
  background-color: rgba(23, 73, 94, 0.2);
  backdrop-filter: blur(10px);
}
.right {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.m_password_login {
  margin: 0;
  padding: 5px 20px;
  text-align: left;
  font-weight: bold;
  font-size: 25px;
  width: 100%;
}
.login {
  width: 320px;
  height: 30px;
  margin: 20px 0;
  background-color: rgb(34, 187, 255);
  color: white;
  padding-top: 7px;
}
.login:hover {
  background-color: rgb(0, 162, 232);
  cursor: pointer;
}
.divider {
  display: flex;
  align-items: center;
}

.input {
  color: #a5a5a5;
  width: 320px;
  height: 35px;
  font-size: 1em;
  border: none;
  background: none;
  border-radius: 0;
  border-bottom: 2px solid rgba(102, 102, 102, 0.22);
  margin-bottom: 10px;
  font-weight: bold;
  font-family: "BenderLight";
}
input:focus {
  outline: none;
}
.underlined {
  border-bottom: 4px solid rgb(255, 238, 34);
  margin-bottom: 20px;
}
.login_method,

.line {
  flex: 1;
  width: 124px;
  border-top: 1px solid #ebebeb;
}

.signup_url {
  font-size: 14px;
  color: rgb(26, 161, 219);
}
.signup,
.forgot {
  flex: 1;
}
.forgot {
  text-align: right;
}
.pass_tray {
  display: flex;
  text-align: justify;
  margin: 0;
}
.circle {
  padding: 10px;
  margin: 10px;
  width: 20px;
  height: 20px;
  background-color: rgba(255, 255, 255, 0.2);
}
.method_login_cont {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 200px;
  margin-left: 70px;
}
.wechat {
  color: #4abf5d;
}
.qq {
  color: #70c6fc;
}
.weibo {
  color: #ec6e2d;
}
</style>