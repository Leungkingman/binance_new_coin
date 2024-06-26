<template>
  <div class="p-a-10 w-1400">
    <el-tabs v-model="activeTab" @tab-click="handleClick">
      <el-tab-pane label="控制台" name="console">
        <div class="fl w-400">
          <!-- 币种更新区域 -->
          <div class="container-area">
            <el-form ref="form" :inline="true" label-width="90px" label-position="left">
              <!-- <el-form-item label="已记录币种（英文逗号,隔开）" class="w-300">
                <el-input v-model="coins"></el-input>
              </el-form-item> -->
              <el-form-item label="已记录币种">
                <el-input v-model="coins"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="updateCoin" size="small">更新</el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="container-area m-t-10">
            <el-form ref="form" label-width="110px">
              <el-form-item label="U使用量" class="w-300">
                <el-input v-model="userInfo.buy_usdt"></el-input>
              </el-form-item>
              <el-form-item label="U最小使用量" class="w-300">
                <el-input v-model="userInfo.min_buy_usdt"></el-input>
              </el-form-item>
              <el-form-item label="滑点(%)" class="w-300">
                <el-input v-model="userInfo.slippage"></el-input>
              </el-form-item>
              <el-form-item label="线程数" class="w-300">
                <el-input v-model="userInfo.worker_count"></el-input>
              </el-form-item>
              <el-form-item label="队列间隔" class="w-300">
                <el-input v-model="userInfo.queue_time_gap"></el-input>
              </el-form-item>
              <el-form-item label="轮询间隔" class="w-300">
                <el-input v-model="userInfo.server_queue_time_gap"></el-input>
              </el-form-item>
              <el-form-item label="请求间隔" class="w-300">
                <el-input v-model="userInfo.request_time_gap"></el-input>
              </el-form-item>
              <el-form-item label="新接口请求间隔" class="w-300">
                <el-input v-model="userInfo.new_request_time_gap"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="updateUserInfo">更新</el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="container-area m-t-10">
            <el-form ref="form">
              <!-- <el-form-item>
                <el-button type="primary" @click="allInit">集体初始化</el-button>
                <el-button type="success" @click="queueStart">队列启动</el-button>
                <el-button type="danger" @click="allStop">全部暂停</el-button>
              </el-form-item> -->
              <el-form-item>
                <el-button type="primary" @click="allInit">集体初始化</el-button>
                <el-button type="success" @click="startMainServer">启动主服务器</el-button>
                <el-button type="danger" @click="stopMainServer">停止</el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="container-area m-t-10">
            <el-form ref="form">
              <el-form-item>
                <el-button type="primary" @click="testBinanceApi">测试请求</el-button>
                <el-button type="success" @click="getEngineLog">获取日志</el-button>
                <el-button @click="clearEngineLog">清除日志</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
        <div class="fl m-l-20" style="width: 972px;">
          <el-tabs v-model="activeServer" type="card" @tab-click="switchServer">
            <el-tab-pane v-for="(s, index) in server" :label="s.ip" :name="s.ip" class="bg-wh">
              <server-platform :ref="serverRefs[index]" :serverHost="s.ip" :sIndex="index"></server-platform>
            </el-tab-pane>
          </el-tabs>
        </div>
        <div style="clear: both"></div>
      </el-tab-pane>
      <el-tab-pane label="订单" name="order">
        <div class="m-t-10">
          <el-table
            :data="orderData"
            border
            style="width: 100%">
            <el-table-column
              prop="coin"
              label="币种"
              width="90">
            </el-table-column>
            <el-table-column
              prop="find_price"
              label="发现价"
              width="100">
            </el-table-column>
            <el-table-column
              prop="price"
              label="成交均价"
              width="100">
            </el-table-column>
            <el-table-column
              label="10%利润"
              width="110">
              <template slot-scope="scope">
                {{ scope.row.price | tenPrice }}
              </template>
            </el-table-column>
            <el-table-column
              label="20%利润"
              width="110">
              <template slot-scope="scope">
                {{ scope.row.price | twentyPrice }}
              </template>
            </el-table-column>
            <el-table-column
              label="30%利润"
              width="110">
              <template slot-scope="scope">
                {{ scope.row.price | thirdtyPrice }}
              </template>
            </el-table-column>
            <el-table-column
              label="40%利润"
              width="110">
              <template slot-scope="scope">
                {{ scope.row.price | fourtyPrice }}
              </template>
            </el-table-column>
            <el-table-column
              prop="amount"
              label="成交量"
              width="100">
            </el-table-column>
            <el-table-column
              label="成交金额"
              width="100">
              <template slot-scope="scope">
                {{ scope.row.total | saveTwo }}
              </template>
            </el-table-column>
            <el-table-column
              label="时间"
              width="180">
              <template slot-scope="scope">
                {{ parseInt(scope.row.create_time) | timeSeconds }}
              </template>
            </el-table-column>
            <el-table-column
              label="操作">
              <template slot-scope="scope">
                <el-button size="mini" @click="updateOrder(scope.row)">更新订单</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="域名管理" name="domain">
        <div>
          <el-button type="primary" @click="openDomainDialog">添加域名</el-button>
        </div>
        <div class="m-t-10">
          <el-table
            :data="domainData"
            border
            style="width: 100%">
            <el-table-column
              type="index"
              width="50">
            </el-table-column>
            <el-table-column
              prop="id"
              label="ID"
              width="90">
            </el-table-column>
            <el-table-column
              prop="domain"
              label="域名">
            </el-table-column>
            <el-table-column
              label="操作">
              <template slot-scope="scope">
                <el-button size="mini" @click="updateDomain(scope.row)">更新域名</el-button>
                <el-button size="mini" type="danger" @click="deleteDomain(scope.row.id)">删除域名</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-tab-pane>
    </el-tabs>
    <!-- 添加域名表单 -->
    <el-dialog title="域名管理" :visible.sync="domainDialogShow">
      <el-form>
        <el-form-item label="域名" label-width="100px">
          <el-input v-model="currentDomain.domain" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDomainDialog">取 消</el-button>
        <el-button type="primary" @click="addDomain">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import serverPlatform from '@/components/serverPlatform'
import service from '@/utils/request'
import moment from 'moment'
export default {
  data() {
    return {
      server: [
        { ip: '8.218.81.24' },
        { ip: '8.210.77.136' },
        { ip: '8.210.127.104' },
        { ip: '8.210.3.245' },
        { ip: '8.210.92.126' },
        { ip: '47.242.77.62' },
        { ip: '8.210.217.245' },
        { ip: '47.242.207.148' },
        { ip: '47.243.59.19' },
        { ip: '8.210.174.216' },
      ],
      serverRefs: [],
      currentStart: 0,
      currentStop: 0,
      currentInit: 0,
      canInit: true,
      initing: false,
      engineRunning: false,
      pwdDialogShow: false,
      domainDialogShow: false,
      coins: "",
      userInfo: {
        buy_usdt: "",
        min_buy_usdt: "",
        slippage: "",
        worker_count: "",
        queue_time_gap: "",
        server_queue_time_gap: "",
        request_time_gap: "",
        new_request_time_gap: "",
        long_profit: "",
        short_profit: "",
      },
      orderData: [],
      activeTab: 'console',
      activeServer: '',
      currentDomain: {
        id: null,
        domain: ''
      },
      domainData: []
    }
  },
  methods: {
    async getUserConfig() {
      const config = { url: "/user/getConfig", method: 'get' }
      const res = await service(config)
      if (res.code == 0) {
        this.userInfo.buy_usdt = res.data.buy_usdt
        this.userInfo.min_buy_usdt = res.data.min_buy_usdt
        this.userInfo.slippage = res.data.slippage
        this.userInfo.worker_count = res.data.worker_count
        this.userInfo.queue_time_gap = res.data.queue_time_gap
        this.userInfo.server_queue_time_gap = res.data.server_queue_time_gap
        this.userInfo.request_time_gap = res.data.request_time_gap
        this.userInfo.new_request_time_gap = res.data.new_request_time_gap
        this.userInfo.long_profit = res.data.long_profit
        this.userInfo.short_profit = res.data.short_profit
      }
    },
    async getMainQueueTaskStatus() {
      const config = { url: "/engine/getQueueTaskRunningStatus", method: 'get' }
      const res = await service(config)
      // console.log('res = ', res)
      if (res.code == 0) {
        this.engineRunning = res.data
        if (this.engineRunning) {
          await this.allChildComponentGetEngineRunningStatus()
        }
      }
      setTimeout(() => {
        this.getMainQueueTaskStatus()
      }, 1000 * 10)
    },
    async getCoinData() {
      const config = { url: "/coin/getCoin", method: 'get' }
      const res = await service(config)
      if (res.code == 0) {
        this.coins = res.data["coin"]
      }
    },
    async getOrders() {
      const config = { url: "/order/getGateOrders", method: 'get' }
      const res = await service(config)
      if (res.code == 0) {
        this.orderData = res.data
      }
    },
    async getDomains() {
      const config = { url: "/domain/getDomain", method: 'get' }
      const res = await service(config)
      // console.log('getDomains res = ', res)
      if (res.code == 0) {
        this.domainData = res.data
      }
    },
    async updateCoin() {
      const postData = { coins: this.coins }
      const res = await service({
        url: "/coin/updateCoin",
        method: 'post',
        data: postData
      })
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
      }
    },
    async updateUserInfo() {
      const postData = {
        buy_usdt: this.userInfo.buy_usdt,
        min_buy_usdt: this.userInfo.min_buy_usdt,
        slippage: this.userInfo.slippage,
        worker_count: this.userInfo.worker_count,
        queue_time_gap: this.userInfo.queue_time_gap,
        server_queue_time_gap: this.userInfo.server_queue_time_gap,
        request_time_gap: this.userInfo.request_time_gap,
        new_request_time_gap: this.userInfo.new_request_time_gap,
        long_profit: this.userInfo.long_profit,
        short_profit: this.userInfo.short_profit,
      }
      const res = await service({
        url: "/user/updateConfig",
        method: 'post',
        data: postData
      })
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
      }
    },
    async getEngineLog() {
      const config = { url: "/engine/getEngineLog", method: 'get' }
      const res = await service(config)
      // console.log('getEngineLog res = ', res)
    },
    async clearEngineLog() {
      const config = { url: "/engine/clearEngineLog", method: 'get' }
      const res = await service(config)
      // console.log('clearEngineLog res = ', res)
    },
    async startMainServer() {
      const ips = []
      this.server.forEach((item) => {
        ips.push(item.ip)
      })
      const postData = { ips: ips }
      const res = await service({
        url: "/engine/startMainQueueTask",
        method: 'post',
        data: postData
      })
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
      }
    },
    async allChildComponentGetEngineRunningStatus() {
      for (let i = 0; i < this.serverRefs.length; i++) {
        let ref = this.serverRefs[i]
        let chindComponent = this.$refs[ref][0]
        // await chindComponent.getEngineRunningStatus()
        await chindComponent.getNewEngineRunningStatus()
      }
    },
    async stopMainServer() {
      const config = { url: "/engine/stopQueueTaskRunningStatus", method: 'get' }
      const res = await service(config)
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
      }
    },
    queueStart() {
      const ref = this.serverRefs[this.currentStart]
      const chindComponent = this.$refs[ref][0]
      chindComponent.startEngine()
      this.currentStart += 1
      if (this.currentStart < this.serverRefs.length) {
        setTimeout(() => {
          this.queueStart()
        }, parseInt(this.userInfo.queue_time_gap))
      } else {
        this.currentStart = 0
      }
    },
    allStart() {
      for (let i = 0; i < this.serverRefs.length; i++) {
        let ref = this.serverRefs[i]
        let chindComponent = this.$refs[ref][0]
        chindComponent.startEngine()
      }
    },
    allStop() {
      const ref = this.serverRefs[this.currentStop]
      const chindComponent = this.$refs[ref][0]
      chindComponent.stopEngine()
      this.currentStop += 1
      if (this.currentStop < this.serverRefs.length) {
        this.allStop()
      } else {
        this.currentStop = 0
      }
    },
    allInit() {
      const ref = this.serverRefs[this.currentInit]
      const chindComponent = this.$refs[ref][0]
      chindComponent.initSystem()
      this.currentInit += 1
      if (this.currentInit < this.serverRefs.length) {
        this.allInit()
      }
    },
    async testBinanceApi() {
      const config = { url: "/test/testGetApi", method: 'get' }
      const res = await service(config)
      // console.log('testBinanceApi res = ', res)
    },
    setServerRefs() {
      for (let i = 0; i < this.server.length; i++) {
        const refName = "serverPlatform" + i
        this.serverRefs.push(refName)
      }
    },
    async init() {
      this.activeServer = this.server[0]['ip']
      this.setServerRefs()
      await this.getUserConfig()
      await this.getCoinData()
      await this.getOrders()
      await this.getDomains()
      await this.allChildComponentGetEngineRunningStatus()
    },
    handleClick(val) {
      // console.log('handleClick val = ', val)
    },
    switchServer(val) {
      console.log('switchServer val = ', val)
      this.activeServer = val.label
    },
    async updateOrder(val) {
      // console.log('val = ', val)
      const data = {
        order_id: val.order_id,
        currency_pair: val.coin,
      }
      const config = { url: "/order/updateGateOrder", method: 'post', data }
      const res = await service(config)
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
      }
    },
    openDomainDialog() {
      this.domainDialogShow = true
    },
    closeDomainDialog() {
      this.currentDomain = { id: null, domain: '' }
      this.domainDialogShow = false
    },
    async addDomain() {
      if (this.currentDomain.id) {        // 更新操作
        const data = {
          id: this.currentDomain.id,
          domain: this.currentDomain.domain
        }
        const config = { url: "/domain/updateDomain", method: 'post', data }
        const res = await service(config)
        console.log('res = ', res)
        if (res.code == 0) {
          this.closeDomainDialog()
          this.$message({ message: res.msg, type: 'success' })
          await this.getDomains()
        }
      } else {                            // 新增操作
        const data = {
          domain: this.currentDomain.domain
        }
        const config = { url: "/domain/addDomain", method: 'post', data }
        const res = await service(config)
        // console.log('res = ', res)
        if (res.code == 0) {
          this.closeDomainDialog()
          this.$message({ message: res.msg, type: 'success' })
          await this.getDomains()
        }
      }
    },
    async updateDomain(val) {
      this.currentDomain = JSON.parse(JSON.stringify(val))
      this.openDomainDialog()
    },
    async deleteDomain(val) {
      const data = {
        id: val
      }
      const config = { url: "/domain/deleteDomain", method: 'post', data }
      const res = await service(config)
      console.log('res = ', res)
      if (res.code == 0) {
        this.$message({ message: res.msg, type: 'success' })
        await this.getDomains()
      }
    }
  },
  async created() {
    await this.init()
  },
  components: {
    serverPlatform
  },
  filters: {
    timeSeconds: val => moment.unix(val).format('YYYY-MM-DD HH:mm:ss'),
    mTimeSeconds: val => moment.unix(val / 1000).format('YYYY-MM-DD HH:mm:ss'),
    orderState: val => {
      let state = ""
      if (val == "NEW") {
        state = "未成交"
      } else if (val == "FILLED") {
        state = "已成交"
      } else if (val == "PARTIALLY_FILLED") {
        state = "部分成交"
      } else if (val == "CANCELED") {
        state = "已撤单"
      } else if (val == "PARTIALLY_CANCELED") {
        state = "部分撤单"
      }
      return state
    },
    tenPrice: val => {
      const perce = val.length - val.indexOf('.')
      return (Number(val) * 1.1).toFixed(perce)
    },
    twentyPrice: val => {
      const perce = val.length - val.indexOf('.')
      return (Number(val) * 1.2).toFixed(perce)
    },
    thirdtyPrice: val => {
      const perce = val.length - val.indexOf('.')
      return (Number(val) * 1.3).toFixed(perce)
    },
    fourtyPrice: val => {
      const perce = val.length - val.indexOf('.')
      return (Number(val) * 1.4).toFixed(perce)
    },
    fiftyPrice: val => {
      const perce = val.length - val.indexOf('.')
      return (Number(val) * 1.5).toFixed(perce)
    },
    nextPrice: val => {
      const perce = val.length - val.indexOf('.') - 1
      const start_price = (Number(val) * 1.1).toFixed(perce )
      const end_price = (Number(val) * 1.2).toFixed(perce)
      return `${start_price} - ${end_price}`
    },
    saveTwo: val => {
      return Number(val).toFixed(2)
    }
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
  width: 96%;
  height: 300px;
  margin-top: 5px;
  padding: 2px 5px;
  background: #000;
  border-radius: 5px;
  font-size: 12px;
  color: #fff;
}
</style>
