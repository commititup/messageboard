import Vue from 'vue'
import App from './App.vue'
import Notifications from 'vue-notification';
Vue.use(Notifications);
// vue-good-table setup ,import the styles and register it globally
import VueGoodTablePlugin from 'vue-good-table';
import 'vue-good-table/dist/vue-good-table.css'
Vue.use(VueGoodTablePlugin);
import 'vue-loading-overlay/dist/vue-loading.css';

new Vue({
  el: '#app',
  render: h => h(App)
})
