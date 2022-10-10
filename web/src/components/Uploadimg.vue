<template>
  <div>
    <b-form-file
      v-model="img"
      browse-text="选择图像"
      placeholder="图像地址"
      accept="image/*"
    ></b-form-file>
    <div class="d-none">{{imgname}}</div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  props: ['imgname'],
  data () {
    return {
      img: null
    }
  },
  computed: {
    ...mapState('manageModule', {
      imgUrl: state => state.imgUrl
    })
  },
  methods: {
    ...mapActions('manageModule', { upimg: 'upimg' }),
    handleImgChange () {
      this.upimg({'img': this.img, 'imgname': this.imgname})
    }
  },
  watch: {
    img: function () {
      this.handleImgChange()
    }
  }
}
</script>

<style scoped lang="scss">
</style>
