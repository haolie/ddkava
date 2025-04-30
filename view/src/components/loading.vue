<template>
  <div class="alertLoading" id="loading" v-show="loadDom" style="display: block;position: absolute;top: 0;right: 0;overflow: hidden;width:100%;height: 100%">
    <div class="layerbg"></div>
    <div class="layerbox" style="width: 520px;margin: 288px auto">
      <img src="static/assets/images/loading.gif" style="margin:0 auto" >
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    loadDom () {
      return this.$store.state.loadDom
    }
  }
}
</script>
