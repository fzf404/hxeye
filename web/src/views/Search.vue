<template>
  <div>
    <navbar />
    <div class="show-page02 mx-auto">
      <h2 class="text-center">搜索结果</h2>
      <hr>
      <div class="info-list mx-4 my-3" v-for="(i,index) in artlist" :key="index">
        <div class="info-list-item">
          <a :href="'/article?id='+i.artid" target="_blank">
            <b-badge pill variant="light">{{i.artid}}</b-badge >
            {{i.title}}
          </a>
          <span class="fr">{{i.create}}</span>
          <hr>
        </div>
      </div>
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
    this.$store.dispatch('infoModule/search', {'key': this.$route.query.key})
  },
  data () {
    return {
      pageid: 1
    }
  },
  computed: {
    ...mapState('infoModule', {
      artlist: state => state.artlist
    })
  }
}

</script>

<style scoped lang="scss">
.fr{
  float: right;
}
.show-page02{
  margin-top: 190px;
  max-width: 60%;
}
</style>
