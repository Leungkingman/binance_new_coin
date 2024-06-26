<template>
  <div class="container-area w-950">
    <el-row>
      <el-col :span="4">
        <el-button type="primary" @click="initSystem" :loading="initing" :disabled="!canInit">初始化</el-button>
      </el-col>
      <el-col :span="4">
        <el-button type="success" @click="startEngine">运行</el-button>
      </el-col>
      <el-col :span="4">
        <el-button type="warning" @click="stopEngine">暂停</el-button>
      </el-col>
      <!-- <el-col :span="6" class="m-t-10">
        <span>运行状态：</span>
        <span v-if="engineRunning == true" class="c-green fw-bold">正在运行中</span>
        <span v-else class="c-red fw-bold">已停止</span>
      </el-col> -->
    </el-row>
    <el-row class="m-t-20">
      <el-col :span="12">
        <div>运行日志：</div>
        <div class="log-container ovf-y-auto" :id="runningLogContainer">
          <div v-for="item in running_log">
            <span>[{{ item.Time | mTimeSeconds }}]：</span>
            <span>{{ item.Content }}</span>
          </div>
        </div>
        <div class="tx-c m-t-10">
          <el-button type="primary" size="small" @click="clearLog('running')">清空日志</el-button>
        </div>
      </el-col>
      <el-col :span="12">
        <div>定时任务日志：</div>
        <div class="log-container ovf-y-auto" :id="exchangeLogContainer">
          <div v-for="item in exchange_log">
            <span>[{{ item.Time | mTimeSeconds }}]：</span>
            <span>{{ item.Content }}</span>
          </div>
        </div>
        <div class="tx-c m-t-10">
          <el-button type="primary" size="small" @click="clearLog('exchange')">清空日志</el-button>
        </div>
      </el-col>
    </el-row>
    <el-row class="m-t-20">
      <el-col :span="12">
        <div>执行日志：</div>
        <div class="log-container ovf-y-auto" :id="operateLogContainer">
          <div v-for="item in operate_log">
            <span>[{{ item.Time | mTimeSeconds }}]：</span>
            <span>{{ item.Content }}</span>
          </div>
        </div>
        <div class="tx-c m-t-10">
          <el-button type="primary" size="small" @click="clearLog('operate')">清空日志</el-button>
        </div>
      </el-col>
      <el-col :span="12">
        <div>错误日志：</div>
        <div class="log-container ovf-y-auto" :id="errorLogContainer">
          <div v-for="item in error_log">
            <span>[{{ item.Time | mTimeSeconds }}]：</span>
            <span>{{ item.Content }}</span>
          </div>
        </div>
        <div class="tx-c m-t-10">
          <el-button type="primary" size="small" @click="clearLog('error')">清空日志</el-button>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import moment from 'moment'
import axios from 'axios'
import Lockr from 'lockr'
import { Message, Loading } from 'element-ui'
export default {
  name: 'ServerPlatform',
  props: {
    serverHost: {
      type: String,
      default: ''
    },
    sIndex: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      runningLogContainer: `runningLogContainer${this.sIndex}`,
      exchangeLogContainer: `exchangeLogContainer${this.sIndex}`,
      operateLogContainer: `operateLogContainer${this.sIndex}`,
      errorLogContainer: `errorLogContainer${this.sIndex}`,
      apiService: null,
      canInit: true,
      initing: false,
      engineRunning: false,
      running_log: [],        // 运行日志
      error_log: [],          // 错误日志
      operate_log: [],        // 执行日志
      exchange_log: [],       // 初始化及定时任务日志
    }
  },
  methods: {
    initHttp() {
      const host = `http://${this.serverHost}:8880`
      const service = axios.create({
        baseURL: host,
        timeout: 99999,
        headers: {
          "Content-Type": "application/json;charset=utf-8",
          "x-token": Lockr.get("x-token") || ""
        }
      })
      service.interceptors.request.use(
        config => {
          const token = Lockr.get("x-token") || ""
          config.headers = {
            'Content-Type': 'application/json',
            'x-token': token
          }
          return config
        },
        error => {
          Message({
            showClose: true,
            message: error,
            type: 'error'
          })
          return Promise.reject(error);
        }
      );
      service.interceptors.response.use(
        response => {
          if (response.data.data && response.data.data.reload) {
            router.push('/')
          }
          if (response.data.code == 0) {
            return response.data
          } else {
            Message({
              showClose: true,
              message: response.data.msg,
              type: 'error',
            })
            return Promise.reject(response.data.msg)
          }
        },
        error => {
          Message({
            showClose: true,
            message: error,
            type: 'error'
          })
          return Promise.reject(error)
        }
      )
      this.apiService = service
    },
    async getInitialInfo() {
      const config = { url: "/exchange/getInitialInfo", method: 'get' }
      const res = await this.apiService(config)
      if (res.code == 0) {
        const result = res.data
        this.exchange_log = result.exchange_log
        this.scrollLogToBottom()
        if (!this.exchange_log || (this.exchange_log[this.exchange_log.length - 1]['Content'] != 'gate.io初始化数据完成')) {
          setTimeout(() => {
            this.getInitialInfo()
          }, 1000 * 5)
        }
      }
    },
    async initSystem() {
      this.canInit = false
      this.initing = true
      const config = { url: "/exchange/initialExchangeData", method: 'get' }
      const res = await this.apiService(config)
      if (res.code == 0) {
        this.initing = false
        this.$message({ message: res.msg, type: 'success' })
        // this.getInitialInfo()
      }
    },
    async startEngine() {
      const config = { url: '/engine/startEngine', method: 'get' }
      const res = await this.apiService(config)
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
        this.getEngineRunningStatus()
      }
    },
    async stopEngine() {
      const config = { url: "/engine/stopEngine", method: 'get' }
      const res = await this.apiService(config)
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
        this.getEngineRunningStatus()
      }
    },
    async getEngineRunningStatus() {
      // console.log('getEngineRunningStatus')
      try {
        const config = { url: "/engine/getEngineRunningStatus", method: 'get' }
        const res = await this.apiService(config)
        if (res.code == 0) {
          const result = res.data
          this.engineRunning = result.engineRunning
          this.running_log = result.running_log
          this.error_log = result.error_log
          this.operate_log = result.operate_log
          this.exchange_log = result.exchange_log
          this.scrollLogToBottom()
          if (this.engineRunning) {
            setTimeout(async() => {
              await this.getEngineRunningStatus()
            }, 1000 * 5)
          }
        }
      } catch (err) {
        console.log('err = ', err)
      }
    },
    async getNewEngineRunningStatus() {
      try {
        const config = { url: "/engine/getEngineRunningStatus", method: 'get' }
        const res = await this.apiService(config)
        if (res.code == 0) {
          const result = res.data
          this.running_log = result.running_log
          this.error_log = result.error_log
          this.operate_log = result.operate_log
          this.exchange_log = result.exchange_log
          this.scrollLogToBottom()
          setTimeout(async() => {
            await this.getNewEngineRunningStatus()
          }, 1000 * 5)
        }
      } catch (err) {
        console.log('err = ', err)
      }
    },
    scrollLogToBottom() {
      setTimeout(() => {
        var runningLogContainer = document.getElementById(this.runningLogContainer)
        runningLogContainer.scrollTop = runningLogContainer.scrollHeight
        var exchangeLogContainer = document.getElementById(this.exchangeLogContainer)
        exchangeLogContainer.scrollTop = exchangeLogContainer.scrollHeight
        var operateLogContainer = document.getElementById(this.operateLogContainer)
        operateLogContainer.scrollTop = operateLogContainer.scrollHeight
        var errorLogContainer = document.getElementById(this.errorLogContainer)
        errorLogContainer.scrollTop = errorLogContainer.scrollHeight
      }, 0)
    },
    clearLog() {

    },
    checkOrder() {

    },
    cancelOrder() {

    },
    testLog() {
      console.log('serverHost = ', this.serverHost)
    }
  },
  async created() {
    this.initHttp()
  },
  filters: {
    timeSeconds: val => moment.unix(val).format('YYYY-MM-DD HH:mm:ss'),
    mTimeSeconds: val => moment.unix(val / 1000).format('YYYY-MM-DD HH:mm:ss')
  }
}
</script>

<style>
.running-circle {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #33CC99;
}
.stop-circle {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #ff0000;
}
</style>