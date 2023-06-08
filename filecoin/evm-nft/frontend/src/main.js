import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // Make sure router is properly imported
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/saga-blue/theme.css';   
import 'primevue/resources/primevue.min.css';           
import 'primeicons/primeicons.css'; 

createApp(App)
  .use(router)  // Make sure you're using the router here
  .use(PrimeVue)
  .mount('#app')
