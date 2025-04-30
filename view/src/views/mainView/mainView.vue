<template>
  <el-container style="width: 100%;height: 100%;"> 
   <div style="height: 300px;width: 100%;box-shadow: 0 0 0 1px rgb(43 25 25 / 15%);">
    <div style="height: 239px;width: 100%;">

    </div>
    <div style="height: 20px; width: 100%;">CTL</div>
  </div>
</el-container>
</template>
<script>  
export default {
  components: {},
  data () {
    return {
      LoginList:[]
    }
  },
  computed: { 
    tabName:function(item){
       
    }
  },
  methods: {  
    btn(){  
       this.$http.get({
        api:'C_GETINFO',
        params:{},
        callback:function(args){
           console.log(args)
        }
       }

       )
    }, 
  },
  updated () {}, 
  created () {
    var _this= this
  },
  mounted () { 
    
  },
  watch:{ 
  }
}

</script>