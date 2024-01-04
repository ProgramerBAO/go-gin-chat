import {defineStore} from "pinia";
import {ref} from "vue";

export const useMessageStore = defineStore('messageList', () => {
        // 消息列表 存的是map 键是targetId 值是massage
        const messageList = ref({})
        //新增消息
        const addMessage = (targetId, val) => {
            console.log("聊天记录", val)
            // 判断targetId是否存在,不存在新建后添加message,存在直接新增消息
            if (!messageList.value[targetId]) {
                messageList.value[targetId] = []
            }
            messageList.value[targetId].push(val)
        }
        // 根据targetId获取消息
        const getMessage = (targetId) => {
            return messageList.value[targetId]
        }
        return {
            messageList,
            addMessage,
            getMessage
        }
    },
    {
        // 这里是持久化的关键
        persist: true,
    },)