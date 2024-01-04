<script setup>

</script>

<template>
  <router-link to="/login">登录</router-link>
  <router-link to="/register">注册</router-link>
  <router-view/>
</template>

<style scoped>
input {
  width: 100%;
  box-sizing: border-box; /* 防止宽度受到边框和填充的影响 */
}
</style>