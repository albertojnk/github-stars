import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    id: null,
    repositories: null,
    loading: null,
    loaded: null
  },
  mutations: {
    setNewUser(state, payload) {
      state.id = payload.id;
      state.repositories = payload.repositories;
    },
    setLoading(state, payload) {
      state.loading = payload.loading;
      state.loaded = payload.loaded;
    },
    setLoaded(state, payload) {
      state.loading = payload.loading;
      state.loaded = payload.loaded;
    }
  },
  actions: {}
});
