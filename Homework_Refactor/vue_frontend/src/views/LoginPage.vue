<script>
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
</script>

<template>
    <div class="loginpage">
      <div class="header">
        <img alt="logo" src="../assets/imgs/prts_reversed.png" height="150" />
      </div>
      <div class="placeholder"></div>
<slot></slot>
      </div>

      <div class="loading" id="loading"><img src="../assets/imgs/loading@3x.gif" width="30px" /><br />loading...</div>

      <div class="footer">
        <p>Deploy version: Beta 1.0</p>
        <!-- <p>{{ config.deployVer }}</p> -->
        <p>东北农业大学 电气2204 李睿年 2023.2.14</p>
        <a href="https://beian.miit.gov.cn/" target="_blank">黑ICP备2023001389号</a>
      </div>

</template>

<style scoped>
.loginpage {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.header {
  position: absolute;
  top: 0;
  z-index: 999;
}
.placeholder {
  height: 150px;
}
.loading {
  display: none;
  padding: 5px;
  position: absolute;
  background-color: rgba(0, 0, 0, 0.3);
  border-radius: 5px;
}
.footer {
  align-items: center;
  color: white;
  text-shadow: 0 1px 2px #999;
  font-size: 12px;
}
</style>