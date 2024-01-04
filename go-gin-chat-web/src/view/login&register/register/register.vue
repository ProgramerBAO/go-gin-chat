<script setup>
import {reactive} from "vue";
import axios from "axios";

const user = reactive({
  name: "",
  pwd: "",
  phoneNum: "",
  email: "",
  rePwd: ""
})

const submitData = async () => {
  console.log(user)
  const response = await axios.post("api/register", user)
  console.log(response)
}
</script>

<template>
  <!--  <h2>注册</h2>-->
  <div>
    <input v-model="user.name" placeholder="用户名"/>
    <input v-model="user.pwd" placeholder="密码"/>
    <input v-model="user.rePwd" placeholder="重复密码"/>
    <input v-model="user.phoneNum" placeholder="手机号"/>
    <input v-model="user.email" placeholder="邮箱"/>
    <br>
    <button @click="submitData">提交</button>
  </div>

</template>

<style scoped>
input {
  width: 100%;
  box-sizing: border-box; /* 防止宽度受到边框和填充的影响 */
}
</style>