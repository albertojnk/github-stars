import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    id: null,
    repositories: null
  },
  mutations: {
    setNewUser(state, payload) {
      state.id = payload.id;
      state.repositories = payload.repositories;
    }
  },
  actions: {}
});
