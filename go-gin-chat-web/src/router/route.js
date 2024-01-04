import {createRouter, createWebHashHistory} from 'vue-router'


// 定义路由并将路由和资源绑定
const routes = [
    {
        path: "/",
        name: "index",
        // 这样的话就能在index下面展示了
        children: [
            {
                path: "/login",
                name: "login",
                // 路由元信息 认证
                component: () => import("@/view/login&register/login/login.vue")
            },
            {
                path: "/register",
                name: "register",
                component: () => import("@/view/login&register/register/register.vue")
            },
        ],
        // 路由元信息 认证
        meta: {requiresAuth: true},
        component: () => import("@/view/login&register/index.vue")
    },
    {
        path: '/chat/:targetId',
        name: 'chat',
        component: ()=>import("@/view/chat/ChatRoom.vue"),
    },
    {
        // 404 页面 正则匹配 应该放在最后面 好理解 机制是全部遍历完没有匹配上就出现404
        path: "/:path(.*)",
        name: "NotFound",
        component: () => import("@/components/NotFound.vue")
    },
    {
        path: "/home",
        name: "home",
        children: [
            {
                path: "/nav",
                name: "navigation",
                component: () => import("@/view/home/Navigation.vue")
            },
            {
                path: "/SearchUserByName",
                name: "SearchUserByName",
                component: () => import("@/view/home/SearchUserByName.vue")
            },
        ],
        component: () => import("@/view/home/home.vue")
    }]
// 创建实例 就是说路由表有了,准备好了 还要创建个实例暴露出去
const router = createRouter({
    // 应该是打开了哪些网页的记录
    history: createWebHashHistory(),
    routes,
})
//导出 方便被挂载
export default router