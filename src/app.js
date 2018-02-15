import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import VueRouter from 'vue-router'

Vue.use(VueRouter)
Vue.use(Buefy)

// components
import Index from './components/pages/index.vue'
import Ranking from './components/pages/ranking.vue'
import NotFound from './components/pages/notfound.vue'

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: "/", component: Index },
    { path: "/ranking", component: Ranking },
    { path: "*", component: NotFound },
  ]
})

const app = new Vue({
  router
}).$mount("#app")
