<template>
  <div class="p-a-10">
    <div class="">
      <div class="container-area">
        <el-row>
          <el-col :span="12">
            <div>公告日志</div>
            <div class="log-container c-wh ovf-y-auto" id="announceLogContainer">
              <div v-for="item in announceLog">
                <span>[{{ item.Time | timeSeconds }}]：</span>
                <span>{{ item.Content }}</span>
              </div>
            </div>
            <div class="tx-c m-t-10">
              <el-button type="primary" size="small" @click="clearLog('announceLog')">清空日志</el-button>
            </div>
          </el-col>
          <el-col :span="12">
            <div>新币公告日志</div>
            <div class="log-container c-wh ovf-y-auto" id="newListingLogContainer">
              <div v-for="item in newListingLog">
                <span>[{{ item.Time | timeSeconds }}]：</span>
                <span>{{ item.Content }}</span>
              </div>
            </div>
            <div class="tx-c m-t-10">
              <el-button type="primary" size="small" @click="clearLog('newListingLog')">清空日志</el-button>
            </div>
          </el-col>
        </el-row>
        <el-row class="m-t-20">
          <el-col :span="12">
            <div>接口日志</div>
            <div class="log-container c-wh ovf-y-auto" id="apiLogContainer">
              <div v-for="item in apiLog">
                <span>[{{ item.Time | timeSeconds }}]：</span>
                <span>{{ item.Content }}</span>
              </div>
            </div>
            <div class="tx-c m-t-10">
              <el-button type="primary" size="small" @click="clearLog('apiLog')">清空日志</el-button>
            </div>
          </el-col>
          <el-col :span="12">
            <div>错误日志</div>
            <div class="log-container c-wh ovf-y-auto" id="errorLogContainer">
              <div v-for="item in errorLog">
                <span>[{{ item.Time | timeSeconds }}]：</span>
                <span>{{ item.Content }}</span>
              </div>
            </div>
            <div class="tx-c m-t-10">
              <el-button type="primary" size="small" @click="clearLog('errorLog')">清空日志</el-button>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
var fn = null
import moment from 'moment'
import { getEngineLog } from '@/api/engine'
export default {
  data() {
    return {
      engineRunning: false,
      announceLog: [],            // 公告日志
      newListingLog: [],          // 新币公告日志
      apiLog: [],                 // 接口日志
      errorLog: []                // 错误日志
    }
  },
  methods: {
    getEngineLog() {
      getEngineLog().then((res) => {
        // console.log("getEngineLog res = ", res)
        const result = res.data
        const logData = result.logData
        this.engineRunning = result.engineRunning
        // 赋值日志
        this.announceLog = logData.announceLog
        this.newListingLog = logData.newListingLog
        this.apiLog = logData.apiLog
        this.errorLog = logData.errorLog
        this.scrollLogToBottom()
        if (!this.engineRunning) {
          clearInterval(fn)
        }
      })
    },
    scrollLogToBottom() {
      setTimeout(() => {
        var announceLogContainer = document.getElementById('announceLogContainer')
        announceLogContainer.scrollTop = announceLogContainer.scrollHeight

        var newListingLogContainer = document.getElementById('newListingLogContainer')
        newListingLogContainer.scrollTop = newListingLogContainer.scrollHeight

        var apiLogContainer = document.getElementById('apiLogContainer')
        apiLogContainer.scrollTop = apiLogContainer.scrollHeight

        var errorLogContainer = document.getElementById('errorLogContainer')
        errorLogContainer.scrollTop = errorLogContainer.scrollHeight

      }, 0)
    },
    clearLog(type) {
      if (type == 'announceLog') {
        this.announceLog = []
      } else if (type == 'newListingLog') {
        this.newListingLog = []
      } else if (type == 'apiLog') {
        this.apiLog = []
      } else if (type == 'errorLog') {
        this.errorLog = []
      }
    },
    init() {
      this.getEngineLog()
      setTimeout(() => {
        if (this.engineRunning) {
          fn = setInterval(this.getEngineLog, 1000 * 10)
        }
      }, 1000)
    }
  },
  created() {
    this.init()
  },
  filters: {
    timeSeconds: val => moment.unix(val).format('YYYY-MM-DD HH:mm:ss')
  }
}
</script>

<style>
.container-area {
  background: #fff;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 10px;
}
.log-container {
  width: 98%;
  height: 250px;
  margin-top: 5px;
  padding: 2px 5px;
  background: #000;
  border-radius: 5px;
  font-size: 12px
}
</style>
