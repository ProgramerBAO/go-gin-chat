<template>
  <!--  <h2>登陆</h2>-->
  <div>
    <input v-model="user.name" placeholder="用户名"/>
    <input v-model="user.pwd" placeholder="密码"/>
    <br>
    <button :disabled="!canSubmit" @click="submitData">提交</button>
  </div>

</template>

<script setup>
// ref 响应式
import {computed, reactive} from "vue";
import axios from "axios";
import {useUserStore} from "@/stores/modules/user";
import router from "@/router/route";

const user = reactive({
  name: "",
  pwd: "",
})
// 应用
const userStore = useUserStore()
// 赋值
const submitData = async () => {
  // console.log(user)
  // post 第二个参数是body json格式
  console.log("user", user)
  const response = await axios.post("api/login", user)

  userStore.setUserInfo(user.name, response.data.ID)
  console.log("userStore=" + userStore.userInfo)
  userStore.setToken(response.data.token)
  console.log("user是" + userStore.userInfo)

  // 登陆之后路由跳转到home
  await router.push({
    path: 'home',
  })

  console.log(response)
}
const canSubmit = computed(() => {
  return Boolean(user.name && user.pwd)
})
</script>

<style scoped>
input {
  width: 100%;
  box-sizing: border-box; /* 防止宽度受到边框和填充的影响 */
}
</style>