<template>
  <div class="m-5 row">
    <div class="col-md-3">
      <b-list-group-item style="line-height:1.8em">

        <h4 style="line-height:1.8em">
          <b-badge :class="{'badge-success':user.Super,'badge-info':!user.Super}" >{{user.Super?'管理员':'拟稿人'}}</b-badge>
          <br/>
          你好，{{user.name}}
          <br/>
          <b-button variant="primary" href="/#/secret/manage" target="_blank">管理页面</b-button>
          <b-button variant="outline-secondary" @click="logout">退出登录</b-button>
        </h4>
        <hr/>
        <b-alert show>离开此界面时请先保存</b-alert>
        <b-button variant="info" @click="saveArt">保存草稿</b-button>
        <hr/>
        <b>说明</b>
        <br>
        只有管理员有权限发布文章
        <br>
        普通用户只可编辑
        <br>
        发布请回到<a href="/#/secret/mange" target="_blank">管理界面</a>
        <br>

      </b-list-group-item>
    </div>
    <div class="col-md-9">
    <h2 class="display-4">欢迎使用编辑器</h2>
    <b-input-group prepend="文章标题" class="mt-3">
      <b-form-input v-model="title"></b-form-input>
      <template v-slot:append>
        <b-form-select text="文章分类" variant="success" v-model="artype" :options="options">
        </b-form-select>
      </template>
    </b-input-group>
    <br>
    <div id="edit"></div>
    </div>
  </div>
</template>

<script>
import E from 'wangeditor'
import storageService from '../service/storageService'
import { mapState, mapActions } from 'vuex'

export default {
  created () {
    console.log(this.$route.query.id)
    this.$store.commit('manageModule/SET_ARTID', this.$route.query.id)
  },
  mounted () {
    const editor = new E('#edit')
    editor.config.height = 640
    editor.config.uploadImgServer = '/api/upimg'
    editor.config.uploadImgMaxSize = 5 * 1024 * 1024
    editor.config.uploadImgAccept = ['jpg', 'jpeg', 'png', 'gif']
    editor.config.uploadImgMaxLength = 1
    editor.config.zIndex = 100
    editor.config.uploadFileName = 'img'
    editor.config.uploadImgHeaders = {
      Authorization: 'Bearer ' + storageService.get(storageService.USER_TOKEN)
    }
    editor.config.uploadImgHooks = {
      // 图片上传并返回了结果，图片插入已成功
      success: function (xhr) {
        console.log('success', xhr)
      },
      // 图片上传并返回了结果，但图片插入时出错了
      fail: function (xhr, editor, resData) {
        console.log('fail', resData)
      },
      // 上传图片出错，一般为 http 请求的错误
      error: function (xhr, editor, resData) {
        console.log('error', xhr, resData)
      },
      // 上传图片超时
      timeout: function (xhr) {
        console.log('timeout')
      },
      // 把图片插入到编辑器中
      customInsert: function (insertImgFn, result) {
        insertImgFn(result.data[0])
      }
    }
    editor.create()
    this.editor = editor
    if (this.artid) {
      this.secretart({'artid': this.artid}).then((res) => {
        this.title = res.data.data.article.title
        this.artype = res.data.data.article.arttype

        this.editor.txt.html(res.data.data.article.content)
      })
    }
  },
  data () {
    return {
      editor: null,
      title: '',
      artype: 1,
      options: [
        { value: 1, text: '分类1' },
        { value: 2, text: '分类2' },
        { value: 3, text: '分类3' }
      ]
    }
  },
  computed: {
    ...mapState('userModule', {
      user: state => state.userInfo
    }),
    ...mapState('manageModule', {
      artid: state => state.artid,
      article: state => state.article
    })
  },
  methods: {
    ...mapActions('userModule', { logout: 'logout' }),
    ...mapActions('manageModule', { addart: 'addart', modart: 'modart', secretart: 'secretart' }),
    saveArt () {
      if (!this.artid) {
        this.addart({'title': this.title, 'type': this.artype, 'content': this.editor.txt.html()}).catch((err) => {
          console.log(err)
        })
        return
      }
      this.modart({'artid': this.artid, 'title': this.title, 'type': this.artype, 'content': this.editor.txt.html()}).catch((err) => {
        console.log(err)
      })
    }
  }
}

</script>

<style scoped lang="scss">
body{
  background: #f5f5f5;
}
h2{
  font-size: 36px !important;
}
</style>
