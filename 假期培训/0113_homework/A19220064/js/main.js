function download_app(){
    window.location.href = "https://www.zhihu.com/app/";
}
function org_signup(){
    window.location.href = "https://www.zhihu.com/org/signup";
}



function check(){
    let account = document.getElementById("account");
    let password = document.getElementById("password");
    if(account.value == ''){
        account.attributes["placeholder"].value = "请正确输入手机号或邮箱";
        account.classList.add("error");
    } else{
        account.attributes["placeholder"].value = "手机号或邮箱";
        account.classList.remove("error");
    }
    if(password.value == '' && account.value == ''){
        password.attributes["placeholder"].value = "请正确输入密码";
        password.classList.add("error");
    }else{
        password.attributes["placeholder"].value = "密码";
        password.classList.remove("error");
    }
}

function login() {
    let account = document.getElementById("account").value;
    let password = document.getElementById("password").value;
    let p_test = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/;
    let em_test = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    if ( account == "" || password == "") {
        alert("手机号，邮箱或密码不能为空");
        return false;
    }

    if (p_test.test(account) || em_test.test(account) ) {
        // TODO: post 发送到server
        // return true;
        alert("登录信息：\n账号：" + account + "\n密码：" + password);
    }else{
        alert("请正确输入手机号或邮箱!");
        return false;
    }
}


const update_width = () => {
    if(window.innerWidth < 800){
        document.getElementById("left").style.display = "none";
        document.getElementsByClassName("container")[0].style.width = "auto";
        
    }else{
        document.getElementById("left").style.display = "flex";
        document.getElementsByClassName("container")[0].style.width = "750px";
    }
};

window.onload = () => {
    update_width();};
window.addEventListener('resize', update_width);