import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

// components
import Index from './components/pages/index.vue'

Vue.component("index", Index)

new Vue({
  el: "#app"
})
