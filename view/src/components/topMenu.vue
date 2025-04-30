<template>
  <!--<header class="sys-header">-->
    <!--<div class="sys-logo"><img src="static/assets/images/logo.png" title=""></div>-->

    <!--<div class="sys-topbar">-->
      <!--<ul>-->
        <!--<li><a href="javacript:void(0);"><img src="static/assets/images/_userphoto.png" title=""><b>{{userName}}</b></a>-->
        <!--</li>-->
        <!--<li><a href="index.html"><i class="icon-logout"></i></a></li>-->
      <!--</ul>-->
    <!--</div>-->
    <!--<div class="sys-menu">-->
      <!--<ul>-->
        <!--&lt;!&ndash;<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>场所管理</span></a>-->
          <!--<ul class="dropdown-menu">-->
            <!--<li><a href="index.html">场所概览</a></li>-->
            <!--<li><a href="search.html">场所查询</a></li>-->
            <!--<li><a href="maintain.html">场所信息维护</a></li>-->
          <!--</ul>-->
        <!--</li>&ndash;&gt;-->
        <!--<li v-for="(menu,index) in menuItems" :key="index" class="dropdown">-->
          <!--<a class="dropdown-toggle" :class="{active:menu.url===currentUrl}" :href="'#/'+menu.url"><span>{{menu.name}}</span></a>-->
        <!--</li>-->

        <!--&lt;!&ndash;<li class="dropdown">&ndash;&gt;-->
        <!--&lt;!&ndash;<a class="dropdown-toggle active" href="place.html"><span>场所概览</span></a>&ndash;&gt;-->
        <!--&lt;!&ndash;</li>&ndash;&gt;-->
        <!--&lt;!&ndash;<li class="dropdown">&ndash;&gt;-->
        <!--&lt;!&ndash;<a class="dropdown-toggle" href="search.html"><span>场所查询</span></a>&ndash;&gt;-->
        <!--&lt;!&ndash;</li>&ndash;&gt;-->
        <!--&lt;!&ndash;<li class="dropdown">&ndash;&gt;-->
        <!--&lt;!&ndash;<a class="dropdown-toggle" href="maintain.html"><span>场所信息维护</span></a>&ndash;&gt;-->
        <!--&lt;!&ndash;</li>&ndash;&gt;-->
        <!--&lt;!&ndash;<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>数据质量分析</span></a>-->
        <!--</li>-->
        <!--<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>分析应用</span></a>-->
        <!--</li>-->
        <!--<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>任务管理</span></a>-->
        <!--</li>-->
        <!--<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>系统管理</span></a>-->
        <!--</li>-->
        <!--<li class="dropdown">-->
          <!--<a class="dropdown-toggle" href="#"><span>数据服务</span></a>-->
        <!--</li>&ndash;&gt;-->
      <!--</ul>-->
    <!--</div>-->
  <!--</header>-->
  <header class="sys-header">
    <div class="sys-logo"><img src="static/assets/images/logo.png" title=""></div>

    <div class="sys-topbar">
      <ul>
        <li><a href="javacript:void(0);"><img src="static/assets/images/_userphoto.png" title=""><b>{{userName}}</b></a></li>
        <li><a href="javascript:void(0)" @click="logout"><i class="icon-logout"></i></a></li>
      </ul>
    </div>
    <div class="sys-menu">
      <ul>
        <el-menu class="el-menu-demo" mode="horizontal" >
          <el-submenu v-for="(menu,index) in menuItems" :key="index" :index="index.toString()" >
            <li slot="title" ><a class="dropdown-toggle" :class="{active:currentUrl==menu.url}" :href="menu.child.length?'#/'+menu.child[0].url:'#/'+menu.url"><span>{{menu.name}}</span></a></li>
            <el-menu-item  v-for="(m,mi) in menu.child" :class="{active:$route.fullPath=='/'+m.url}" :key="m.id" :index="index+'_'+mi" ><a :href="'#/'+m.url">{{m.name}}</a></el-menu-item>
          </el-submenu>
        </el-menu>
      </ul>
    </div>
  </header>
</template>

<script>
let _this
export default {
  components: {},
  data () {
    return {
      menuItems: [],
      currentUrl: ''
    }
  },
  computed: {
    userName () {
      return window.sessionStorage.getItem('userName')
    }
  },
  methods: {
    setUrl () {
      var temp = this.$route.fullPath.split('/')
      this.currentUrl = temp.length > 1 ? temp[1] : ''
    },
    logout () {
      let self = this
      this.$http.get({
        api: 'S_LOGOUT',
        params: {},
        callback: (res) => {
          self.$store.commit('deleteToken')
          window.location.href = window.config.LOGIN_PAGE
        }
      })
    }
  },
  updated () {
  },
  watch: {
    $route () {
      this.setUrl()
    }
  },
  created () {
  },
  mounted () {
    _this = this
    _this.$http.get({
      api: 'S_MENU',
      params: {systemId: 100},
      callback: function (data) {
        _this.menuItems = data
        _this.setUrl()
      }
    })

    _this.$http.get({
      api: 'S_USERINFO',
      callback: function (data) {
        _this.userName = data.name
      }
    })

    $('.dropdown').click(function (e) {

    })
  }
}
</script>
