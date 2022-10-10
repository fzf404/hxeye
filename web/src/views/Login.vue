<template>
  <div class="login row">
    <div class="col-md-4"></div>
    <b-form @submit="onSubmit" class="col-md-4">
      <h2 class="mb-4">后台管理</h2>
      <b-form-group id="input-group-1" label="用户名:" label-for="input-1">
        <b-form-input
          id="input-2"
          v-model="user.name"
          required
          placeholder="请输入用户名"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-1" label="密码:" label-for="input-1">
        <b-form-input
          id="input-2"
          v-model="user.password"
          required
          placeholder="请输入密码"
        ></b-form-input>
      </b-form-group>
      <br />
      <b-button type="submit" variant="primary" block>登录</b-button>
    </b-form>
    <div class="col-md-4"></div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  data () {
    return {
      user: {
        name: '',
        password: ''
      }
    }
  },
  methods: {
    ...mapActions('userModule', { userlogin: 'login' }),
    onSubmit (evt) {
      evt.preventDefault()
      this.userlogin(this.user).then(() => {
        this.$router.push({ name: 'Manage' })
      }).catch((err) => {
        console.log(err)
      })
    }
  }
}
</script>

<style scoped >
body{
  background: #f5f5f5;
}
.login {
  margin-top: 80px;
}
</style>
