<template>
  <div>
    <navbar />
    <div class="show-page02 mx-auto">
      <h2 class="text-center">全部文章列表</h2>
      <hr>
      <div class="info-list mx-4 my-3" v-for="(i,index) in artlist" :key="index">
        <div class="info-list-item">
          <a :href="'/#/article?id='+i.artid" target="_blank">
            <b-badge pill variant="light">{{i.artid}}</b-badge >
            {{i.title}}
          </a>
          <span class="fr">{{i.create}}</span>
          <hr>
        </div>
      </div>
      <div class="my-3">
        <b-button-group>
          <b-button variant="success" @click="beforeArtList">前一页</b-button>
          <b-button variant="light">第{{pageid}}页</b-button>
          <b-button variant="info" @click="nextArtList">后一页</b-button>
        </b-button-group>
      </div>
    </div>
    <footbar />
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

import Home from './Home'
import Navbar from '@/components/Navbar.vue'
import Footbar from '@/components/Footbar.vue'

export default {
  components: { Navbar, Footbar, Home },
  created () {
    this.$store.dispatch('infoModule/artlist', {'pageid': 1})
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
  },
  methods: {
    ...mapActions('infoModule', { getartlist: 'artlist' }),
    nextArtList () {
      this.pageid++
      this.getartlist({'pageid': this.pageid})
    },
    beforeArtList () {
      if (this.pageid > 1) {
        this.pageid--
      }
      this.getartlist({'pageid': this.pageid})
    }
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
