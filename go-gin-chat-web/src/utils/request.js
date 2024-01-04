import axios from "axios";
import {useUserStore} from "@/stores/modules/user";
import router from "@/router/route";

const axiosService = axios.create({
    headers: {
        token: ""
    }
})

// 拦截器

// 请求拦截器
axiosService.interceptors.request.use(
    requestConfig => {
        const userStore = useUserStore()
        // 在这里实现了持久化
        requestConfig.headers = {
            'Content-Type': 'application/json',
            'Token': userStore.token,
            'X-Token': userStore.token,
            ...requestConfig.headers
        }
        return requestConfig
    },
// 响应拦截器(中间件)
    axiosService.interceptors.response.use((response) => {
            // 对响应数据做些什么
            return response;
        },
        (error) => {
            // 对响应错误做些什么
            if (error.response && error.response.status === 401) {
                // 如果状态码为 401，清空数据重定向到登录页面
                const userStore = useUserStore()
                // 清空数据
                userStore.ClearStorage()
                router.push('/');
            }
            return Promise.reject(error);
        }
    );

export default axiosService