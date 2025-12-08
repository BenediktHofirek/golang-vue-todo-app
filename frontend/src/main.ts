import './assets/style.css'
import '@fontsource-variable/nunito-sans'
import PrimeVue from 'primevue/config'

import { createApp } from 'vue'
import { VueQueryPlugin } from '@tanstack/vue-query'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)
app.use(VueQueryPlugin)
app.use(PrimeVue, { unstyled: true })

app.mount('#app')
