import Vue from 'vue';
import Vuetify from 'vuetify';
import DaySpanVuetify from 'dayspan-vuetify';
import fr from 'dayspan-vuetify/src/locales/fr';
// import da from 'dayspan-vuetify/src/locales/da';
import axios from 'axios';

import 'vuetify/dist/vuetify.min.css';
import 'material-design-icons-iconfont/dist/material-design-icons.css';
import 'dayspan-vuetify/dist/lib/dayspan-vuetify.min.css';

import { Session } from '@/plugins/cryptouser';

import App from './App';
import router from './router';

Vue.use(Vuetify);
Vue.use(DaySpanVuetify, {
});

Vue.$dayspan.addLocale('fr', fr);
// Vue.$dayspan.addLocale('da', da);

axios.defaults.baseURL = '/api/v0/';

Vue.prototype.$axios = axios;

const global = new Vue({
  data: {
    username: Session.userUsername,
    loggedIn: Session.userLoggedIn,
    data: null,
  },
  async created() {
    if (Session.userLoggedIn) {
      this.data = await Session.userGetData();
    }
  },
  methods: {
    async logIn(username, password) {
      await Session.userLogIn('', username, password);
      this.data = await Session.userGetData();
      this.refresh();
    },
    async logOut() {
      await Session.userLogOut();
      this.data = null;
      this.refresh();
    },
    async register(username, password) {
      await Session.userRegister('', username, password, {
        version: '1',
        events: [],
      });
      this.data = await Session.userGetData();
      this.refresh();
    },
    async refresh() {
      this.username = Session.userUsername;
      this.loggedIn = Session.userLoggedIn;
    },
    async exportData() {
      return Session.userExportData();
    },
    async importData(str) {
      return Session.userImportData(str);
    },
    async updatePassword(oldPassword, newPassword) {
      return Session.userUpdatePassword(oldPassword, newPassword);
    },
  },
});

Vue.prototype.$global = global;

Vue.prototype.$readData = async () =>
  Session.userGetData();

Vue.prototype.$updateData = async data =>
  Session.userSetData(data);

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App),
});
