<template>
  <div class="w-400 m-center m-t-30">
    <el-form ref="form" label-width="80px">
      <el-form-item label="用户名">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" show-password></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="goLogin">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { login } from '@/api/user'
export default {
  name: 'login',
  props: {
    msg: String
  },
  data() {
    return {
      form: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    async goLogin() {
      let postData = JSON.parse(JSON.stringify(this.form))
      login(postData).then((res) => {
        const token = res.data.token
        Lockr.set('x-token', token)
        this.$message({
          message: '登录成功',
          type: 'success'
        })
        setTimeout(() => {
          router.push('/console/workplatform')
        }, 1500)
      })
    }
  }
}
</script>

<style>

</style>
