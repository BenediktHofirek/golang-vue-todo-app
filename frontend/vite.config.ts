import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import svgLoader from 'vite-svg-loader'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), mode === 'development' ? '' : 'process.env.')

  return {
    envDir: false,
    define: {
      "process.env.VITE_GCP_API_KEY": `"${env.VITE_GCP_API_KEY}"`,
      "process.env.VITE_GCP_AUTH_DOMAIN": `"${env.VITE_GCP_AUTH_DOMAIN}"`
    },
    plugins: [vue(), vueJsx(), vueDevTools(), svgLoader(), tailwindcss()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server:{
      host: '0.0.0.0',
      port: parseInt(env.APP_PORT)
    }
  }
})
