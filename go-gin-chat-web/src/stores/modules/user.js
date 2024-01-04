import {defineStore, acceptHMRUpdate} from "pinia";
import {ref, watch} from "vue";
import {createPersistedState} from "pinia-plugin-persistedstate";

// 以use开头 以Store结尾
export const useUserStore = defineStore('user', () => {
        // 用户信息
        const userInfo = ref({
            userName: "",
            userId: "",
        })
        const getUserInfo = () => {
            return userInfo.value
        }
        const setUserInfo = (userName, userId) => {
            userInfo.value.userName = userName
            userInfo.value.userId = userId
        }

        //好友列表
        const friendList = ref(JSON.parse(window.localStorage.getItem('friendList') || '[]'))
        const getFriendList = () => {
            return friendList.value
        }

        const setFriendList = (val) => {
            friendList.value = val
        }
        // 监听friendList变化
        watch(() => friendList.value, () => {
            window.localStorage.setItem('friendList', JSON.stringify(friendList.value))
        })
        // token
        const token = ref(window.localStorage.getItem('x-token') || '')
        // status
        const setToken = (val) => {
            token.value = val
        }

        watch(() => token.value, () => {
            window.localStorage.setItem('x-token', token.value)
            // window.localStorage.setItem('token', token.value)
        })        // 监听token变化


        /* 清理数据 */
        const ClearStorage = async () => {
            token.value = null
            userInfo.value = null
            friendList.value = null
            window.localStorage.clear()
            window.sessionStorage.clear()
        }

        return {
            userInfo,
            setUserInfo,
            getUserInfo,
            token,
            setToken,
            ClearStorage,
            friendList,
            setFriendList,
            getFriendList,

        }
    },
    {
        // 这里是持久化的关键
        persist: true,
    },)

// if (import.meta.hot) {
//     // 热更新
//     import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
// }







