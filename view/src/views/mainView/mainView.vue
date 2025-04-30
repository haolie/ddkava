<template>
  <el-container style="width: 100%;height: 100%;">
    <button @click="btn">button</button>
   vvvvv
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