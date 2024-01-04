<template>
  <h1>聊天室</h1>
  <p>当前用户id:{{ $route.params.targetId }}</p>
  <!--  显示当前targetId的聊天记录-->
  <ul>
    <li v-for="message in  messageStore.messageList[newMessage.message.targetId]" :key="message.id">
      <span>{{ message.ownerId }}</span>
      <span>信息:{{ message.content }}</span>
      <!--      <span>{{ message }}</span>-->
    </li>
  </ul>

  <input v-model="newMessage.message.content" placeholder="Type your message"/>
  <button @click="sendMessage">Send</button>
</template>

<script setup>

import {useUserStore} from "@/stores/modules/user";
import {useRoute} from 'vue-router';
import {useMessageStore} from "@/stores/modules/chatMessage";
import {reactive} from "vue";
import {useWebSocket} from "@/hooks/websocket";

const userStore = useUserStore()
const messageStore = useMessageStore()
// 新建一个消息 用于发送 格式为发送者id 接受者id 信息内容 信息类型 发送时间
const newMessage = reactive({
  message: {
    ownerId: userStore.userInfo.userId,
    targetId: useRoute().params.targetId,
    content: '',
    type: 0,
    time: new Date().getTime()
  }
})
const ws = useWebSocket(userStore.userInfo.userId, useRoute().params.targetId)

// 监听输入框的输入
// const inputChange = (event) => {
//   newMessage.value.content = event.target.value
// }

// 发送信息
const sendMessage = async () => {
  // 不能发送空消息
  if (!newMessage.message.content.trim()) {
    console.log("不能发送空消息")
    return
  }
  console.log("newMessage", newMessage.message)
  // console.log("2", messageStore.messageList[newMessage.message.targetId])
  // 更新时间
  newMessage.message.time = new Date().getTime()
  messageStore.addMessage(newMessage.message.targetId, JSON.parse(JSON.stringify(newMessage.message)))
  ws.ws.send(JSON.stringify(newMessage.message));
  // 发送消息 把自己的加入到message中
  // response.data
  // 清空输入框
  newMessage.message.content = ""
};

</script>
<style scoped>

</style>