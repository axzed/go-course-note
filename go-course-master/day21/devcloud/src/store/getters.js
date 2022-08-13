const getters = {
    accessToken: state => state.user.accessToken,
    namespace: state => state.user.namespace,
    account: state => state.user.account,
    username: state => state.user.name,
    userType: state => state.user.type,
    userAvatar: state => state.user.avatar,
    size: state => state.app.size,
    system: state => state.app.system,
    sidebar: state => state.app.sidebar
  }
  
  export default getters