import '@/assets/main.css'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import {createApp} from 'vue'
// 引入pinia
import {createPinia} from 'pinia'
import App from '@/App.vue'
import route from "@/router/route";
// 引入权限
import '@/permission'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

createApp(App)
    .use(route)
    // 创建仓库并且把这个仓库配发给这个app
    // use(createPersistedState()) 持久化
    .use(pinia)
    // .use(axiosService)
    .mount('#app')
