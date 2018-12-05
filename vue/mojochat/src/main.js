import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueNativeSock from 'vue-native-websocket'

Vue.use(BootstrapVue)
Vue.use(VueNativeSock, 'ws://localhost:8080', {
  connectManually: true,
  reconnection: true
})

const vm = new Vue({
  el: '#app',
  render: h => h(App)
})
