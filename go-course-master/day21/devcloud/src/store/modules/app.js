const state = {
    sidebar: {
      opened: true,
    },
    size: 'medium',
    system: 'dashboard'
  }
  
  const mutations = {
    SET_SIZE: (state, size) => {
      state.size = size
    },
    SET_SYSTEM: (state, system) => {
      state.system = system
    },
    TOGGLE_SIDEBAR: (state) => {
      state.sidebar.opened = !state.sidebar.opened
    }
  }
  
  const actions = {
    setSize({ commit }, size) {
      commit('SET_SIZE', size)
    },
    setSystem({ commit }, system) {
      commit('SET_SYSTEM', system)
    },
    toggleSideBar({ commit }) {
      commit('TOGGLE_SIDEBAR')
    }
  }
  
  export default {
    namespaced: true,
    state,
    mutations,
    actions
  }
  