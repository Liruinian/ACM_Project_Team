<style scoped>
*{
    color:white;
}
.header{
    background-color:black;
    display: flex;
}
.header_logo{
    width: 150px;
}
.page_title{
    padding: 29px 0;
    font-size: 40px;
}
</style>
<template>
    <div class="header">
        <div class="logo">
            <img class="header_logo" src="./header_logo.png" alt="logo">
        </div>
        <div class="header_items">
            <div class="page_title">
                <slot name="ptitle"></slot>
            </div>
            <div class="hmenu" v-for="(item,index) in array">
            {{ item }} -- {{ index }}</div>
        </div>
    </div>
</template>

<script scoped>

</script>