import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// 动态获取本机ip地址
const localIp = require('ip')
console.log("本机地址是",localIp.address())

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  server: {
    host: '0.0.0.0',
    proxy: {
      // 实现跨域
      '/api': {
        target: 'http://192.168.3.57:8080', // 后端服务的地址
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})

export {
  localIp,
}
