<script setup>
import {computed, onMounted, reactive} from "vue";
import axiosService from "@/utils/request";
import {useUserStore} from "@/stores/modules/user";
import router from "@/router/route";

onMounted(() => {
  getFriendList()
})

const user = reactive({
  name: "",
})

const userStore = useUserStore()

const result = reactive({
  name: "",
  targetId: "",
  ownerId: userStore.userInfo.userId,
})
const submitData = async () => {
  // post 第二个参数是body json格式  json格式
  const response = await axiosService({
    method: "get",
    url: "/api/user/getUsers",
    params: user
  })
  console.log(response)
  // 获取参数是response.data
  result.targetId = response.data.user.ID
  result.name = response.data.user.Name
  result.ownerId = userStore.userInfo.userId
}

const addFriend = async () => {
  // post 第二个参数是body json格式  json格式
  const response = await axiosService({
    method: "get",
    url: "/api/user/addFriend",
    params: result
  })
  console.log(response)
  // 获取参数是response.data
  // 刷新好友列表
  if (response.data.friendList) {
    userStore.setFriendList(response.data.friendList)
  }
}
const deleteFriend = async () => {
  // post 第二个参数是body json格式  json格式
  const response = await axiosService({
    method: "get",
    url: "/api/user/deleteFriend",
    params: result
  })
  console.log(response)
  // 获取参数是response.data
  if (response.data.friendList) {
    userStore.setFriendList(response.data.friendList)
  }
}
const getFriendList = async () => {
  // post 第二个参数是body json格式  json格式
  const response = await axiosService({
    method: "get",
    url: "/api/user/getFriendList",
    params: result
  })
  console.log(response)
  // 获取参数是response.data
  if (response.data.friendList) {
    userStore.setFriendList(response.data.friendList)
  }
}

const canSubmit = computed(() => {
  return Boolean(user.name)
})
const goToChat = (targetId) => {
  // 使用路由导航到聊天页面，例如 /chat/:userId
  router.push(`/chat/${targetId}`);
};

</script>

<template>
  <div>
    <input v-model="user.name" placeholder="用户名"/>
    <button :disabled="!canSubmit" @click="submitData">提交</button>
  </div>
  <div>
    <h1>{{ result.name }}</h1>
    <h1>{{ result.targetId }}</h1>
    <button @click="addFriend">添加好友</button>
    <button @click="deleteFriend">删除好友</button>
    <button @click="getFriendList">显示好友</button>
    <!--    好友列表-->
    <div>
      <h1>好友列表</h1>
      <h1 v-if="! userStore.friendList">暂无好友</h1>
      <ul>
        <li v-for="(item, index) in userStore.friendList" :key="index" @click="goToChat(item.targetId)">{{
            item.desc
          }}
        </li>
      </ul>

    </div>
  </div>
</template>


<style scoped>

</style>