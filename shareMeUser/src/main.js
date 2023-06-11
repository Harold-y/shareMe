import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import naive from 'naive-ui'

const app = createApp(App)
app.config.globalProperties.BASE_URL = "http://localhost:8569/"
app.use(router)
app.use(naive)
app.mount('#app')
