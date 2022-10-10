<template>
  <div>
    <navbar />
    <div class="show-page02 mx-auto">
      <h1 class="text-center">{{article.title}}</h1>
      <hr>
      <p class="float-left ml-5">发布时间：{{article.create}}</p>
      <p class="float-right mr-5">更新时间：{{article.update}}</p>
      <div class="" v-html="article.content"></div>
    </div>
    <footbar />
  </div>
</template>

<script>
import { mapState } from 'vuex'

import Home from './Home'
import Navbar from '@/components/Navbar.vue'
import Footbar from '@/components/Footbar.vue'

export default {
  components: { Navbar, Footbar, Home },
  created () {
    this.$store.dispatch('infoModule/getart', {'artid': this.$route.query.id})
  },
  data () {
    return {
    }
  },
  computed: {
    ...mapState('infoModule', {
      article: state => state.article
    })
  }
}

</script>

<style scoped lang="scss">
body{
  background: #f5f5f5;
}
.show-page02{
  margin-top: 190px;
  max-width: 60%;
}
</style>
