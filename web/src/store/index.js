import Vue from 'vue'
import Vuex from 'vuex'
import userModule from './module/user'
import manageModule from './module/manage'
import infoModule from './module/info'

Vue.use(Vuex)

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  state: {
    showNav: true
  },
  mutations: {
    SET_SHOW_NAV (state, showNav) {
      state.showNav = showNav
    }
  },
  actions: {
  },
  modules: {
    userModule,
    manageModule,
    infoModule
  }
})
