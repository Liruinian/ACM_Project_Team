<script setup>
defineProps({
    url:{
        required: true,
        type: String,
    }
})
</script>
<template>
    <a :href="url">
    <div class="menuitem">
    <slot></slot>
    </div>
</a>
</template>

<style scoped>
.menuitem{
    width: fit-content;
    background-color: lightblue;
}
a{
    text-decoration: none;
    color:black
}
</style>