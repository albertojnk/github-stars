import Vue from "vue";
import VModal from "vue-js-modal";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import axios from "axios";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

Vue.config.productionTip = false;
Vue.prototype.$http = axios;
Vue.use(VModal);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
