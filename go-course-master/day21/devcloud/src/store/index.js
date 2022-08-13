import Vue from "vue";
import Vuex from "vuex";
import user from './modules/user'
import app from './modules/app'
import getters from './getters'
import VuexPersistence from 'vuex-persist'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})

Vue.use(Vuex);

const modules = {user, app}

export default new Vuex.Store({
  modules,
  getters,
  plugins: [vuexLocal.plugin]
});
