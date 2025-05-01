let api = {
  getURL: (i) => {
    let a = {

 
      /**
       * 登录, 登出接口
       */
      // 登录
      'S_LOGIN': 'auth-service/system/login',
      // 登出
      'S_LOGOUT': 'auth-service/system/logout',
      // 用户信息dev ? 'static/jsons/menu/user-info.json' :
      'S_USERINFO': 'usermgmt-service/menu/user-info',
 
      'C_GETINFO': 'getInfo',
      'C_STARTSCAN': 'startScan',
    
    }

    return window.config.API_MAPPING(a[i])
  }
}

export default api
