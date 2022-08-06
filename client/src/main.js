import Vue from "vue";
import App from "./App.vue";
import router from './router'
import VueRouter from "vue-router";
import VueSidebarMenu from 'vue-sidebar-menu'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'

import '@fortawesome/fontawesome-free/css/all.css';
import '@fortawesome/fontawesome-free/js/all.js';



import {
  FormInputPlugin,
  NavbarPlugin,
  LayoutPlugin,
  IconsPlugin,
  BCard,
  BInputGroup,
  BButton,
} from "bootstrap-vue";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false;


Vue.use(VueRouter);
Vue.use(VueSidebarMenu)
Vue.use(FormInputPlugin);
Vue.use(NavbarPlugin);
Vue.use(LayoutPlugin);
Vue.component("b-card", BCard);
Vue.component("b-input-group", BInputGroup);
Vue.component("b-button", BButton);
Vue.use(IconsPlugin);



// Vue.prototype.$clientname="Welcome"
// Vue.prototype.$displayPicture = ""
Vue.prototype.$displayPicture = Vue.observable({
  url: './assests/bg.jpg'
});
Vue.prototype.$clientname = Vue.observable({
  name: 'Welcome',
  set: 0
});
Vue.prototype.$groups = Vue.observable({
  group1: 0,
  group2: 0
});
Vue.prototype.$notifications = Vue.observable({
  group1: 0,
  group2: 0,
  inbox: 0
});

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");

