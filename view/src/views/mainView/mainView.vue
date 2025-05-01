<template>
  <el-container style="width: 100%;height: 100%;">
    <!-- <div style="height: 300px;width: 100%;box-shadow: 0 0 0 1px rgb(43 25 25 / 15%);">
    <div style="height: 239px;width: 100%;">
      <el-descriptions title="用户信息">
    <el-descriptions-item label="用户名">{{ workStatus }}</el-descriptions-item>
    <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="备注">
      <el-tag size="small">学校</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="联系地址">江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item>
</el-descriptions>
    </div>
    <div style="height: 20px; width: 100%;" @click="btn">CTL</div>
  </div> -->


    <el-container>
      <el-header style="height: 270px;">
        <div>
          <el-input placeholder="请输入内容" v-model="rootDir" class="input-with-select">
            <template slot="prepend">ScanRoot:</template>
            <el-button slot="append" icon="el-icon-search" @click="startScan"></el-button>
          </el-input>
        </div>
        <div><label>root：{{ Info.Root }}</label></div>
        <div><label>status：{{ workStatus }}</label></div>
        <div><label>fileCount:{{ Info.FileCount }}</label></div>
        <div><label>curDir:{{ Info.curDir }}</label></div>
        <div><label>startTIme:{{ Info.StartTime }}</label></div>
        <div><label>Msg:{{ Info.Msg }}</label></div>

        <div> <el-button type="primary" @click="getInfo">查询</el-button></div>
      </el-header>
      <el-container>

        <el-main>
          <div v-bind:key="group.Key" v-for="group in Info.Groups">
            <el-divider content-position="left">{{ group.Key }}({{ group.FileList.length }})</el-divider>
            <el-table :data="group.FileList" style="width: 100%">
              <el-table-column prop="fileName" label="文件名" width="220">
              </el-table-column>
              <el-table-column prop="filePath" label="位置">
              </el-table-column>
              <el-table-column prop="fileSize" label="大小"  width="180">
              </el-table-column>
             

            </el-table>
          </div>

          <div>
            <el-divider content-position="left">文件列表{{ Info.Files.length }}</el-divider>
            <el-table :data=" Info.Files" style="width: 100%">
              <el-table-column prop="fileName" label="文件名" width="220">
              </el-table-column>
              <el-table-column prop="filePath" label="位置">
              </el-table-column>
              <el-table-column prop="fileSize" label="大小"  width="180">
              </el-table-column>
            </el-table>
          </div>


        </el-main>
      </el-container>
    </el-container>



  </el-container>
</template>
<script>
export default {
  components: {},
  data() {
    return {
      rootDir: "D:\\Downloads",
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }],
      Info: {
        Root: "",
        Status: 0,
        Msg: "",
        FileCount: 0,
        StartTime: 0,
        CurDir: "",
        Files: [],
        Groups: []
      }
    }
  },
  computed: {
    workStatus: function () {
      if (this.Info.Status == 0) {
        return "空闲";
      }

      if (this.Info.Status == 1) {
        return "扫描中……";
      }

      if (this.Info.Status == 2) {
        return "已完成";
      }

      if (this.Info.Status == 0) {
        return "错误";
      }
    },
    tabName: function (item) {

    }
  },
  methods: {
    getInfo() {
      var _this = this;
      this.$http.get({
        api: 'C_GETINFO',
        params: {},
        callback: function (args) {
          _this.Info.Root = args.Root
          _this.Info.Status = args.Status
          _this.Info.Msg = args.Msg
          _this.Info.FileCount = args.FileCount
          _this.Info.StartTime = args.StartTime
          _this.Info.CurDir = args.CurDir

          _this.Info.Groups = []
          _this.Info.Files=[]
          if (args.Files) {
            for (var i = 0; i < args.Files.length; i++) {
              var group = args.Files[i]
              if (group.FileList.length > 1) {
                _this.Info.Groups.push(group)
              } else {
                _this.Info.Files.push(group.FileList[0])
              }
            }
          }
          //_this.Info.Files=args.Files

        }
      }

      )
    },
    startScan() {
      this.$http.get({
        api: 'C_STARTSCAN',
        params: {
          dir: this.rootDir
        },
        callback: function (args) {

        }
      }

      )
    },
  },
  updated() { },
  created() {
    var _this = this
  },
  mounted() {

  },
  watch: {
  }
}

</script>