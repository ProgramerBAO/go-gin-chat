// 这里存放认证
import router from "@/router/route";
import {useUserStore} from "@/stores/modules/user";
// 白名单不需要被验证的路由
const whiteList = ['login']
// 路由跳转前的拦截
router.beforeEach(async (to, from, next) => {
    // 路由跳转前的拦截
    // 1.判断是否登录 不在白名单内 需要验证token
    if (whiteList.indexOf(to.name) > -1) {
        // 在白名单内 直接放行
        next()
    }
    const userStore = useUserStore()
    console.log("token=", userStore.token)
    if (!userStore.token) {
        console.log("未登录,请先登录!")
        await userStore.ClearStorage()
        next('login')
    } else {
        next()
    }

})
