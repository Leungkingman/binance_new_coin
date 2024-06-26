const apiFunc = {
  props: {
    serverHost: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      apiService: null
    }
  },
  methods: {
    initHttp() {
      const host = `http://${this.serverHost}:8880`
      this.apiService = service       // in workerApi.js
    }
  }
}

export default apiFunc