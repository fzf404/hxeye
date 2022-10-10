<template>
  <div class="navbar-container" id="navbar-container">
    <div class="navbar-logo" v-if="navheaderimg">
      <img src="/images/indexLogo01.jpg" class="col-md-4" style="max-height: 80px">
      <img src="/images/indexLogo02.jpg" class="col-md-4" style="max-height: 80px">
      <div class="col-md-4 ">
        <b-form-input v-model="text" placeholder="请输入搜索内容" class="input-search" @keyup.enter="$router.push('/search?key='+text)"></b-form-input>
      </div>
    </div>
    <div class="navbar-base" id="navbar-logo">
      <b-navbar toggleable="lg" type="light" variant="faded" :sticky="true">
        <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav >
          <b-nav-item class="navbar-item-all navbar-item" href="/#/" target="_blank">{{activeItem.navbar[0]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=1" target="_blank">{{activeItem.navbar[1]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=2" target="_blank">{{activeItem.navbar[2]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=3" target="_blank">{{activeItem.navbar[3]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=4" target="_blank">{{activeItem.navbar[4]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=5" target="_blank">{{activeItem.navbar[5]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/list" target="_blank">{{activeItem.navbar[6]}}</b-nav-item>
          <b-nav-item class="navbar-item-all navbar-item" href="/#/article?id=7" target="_blank">{{activeItem.navbar[7]}}</b-nav-item>
        </b-navbar-nav>
        </b-collapse>
      </b-navbar>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  created () {
    this.$store.dispatch('infoModule/active')
  },
  mounted () {
    window.addEventListener('scroll', this.handleScroll)
  },
  data () {
    return {
      text: '',
      navheaderimg: true
    }
  },
  computed: {
    ...mapState('infoModule', {
      activeItem: state => state.active
    })
  },
  methods: {
    handleScroll () {
      let scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
      if (scrollTop < 130) {
        document.getElementById('navbar-logo').style.background = '#fff'
        this.navheaderimg = true
      } else if (scrollTop > 130) {
        document.getElementById('navbar-logo').style.background = '#1890ff'
        this.navheaderimg = false
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.navbar-logo{
  max-width: 90%;
  margin: auto;
  display: flex;
  justify-content: space-around;
  margin: 40px 50px 30px;
  img {
    padding: 0 60px;
  }
  .input-search{
    width: 70%;
    margin-left:80px;
    margin-top: 4px;
    border-radius: 18px;
  }
}
.navbar-container{
  width: 100%;
  z-index: 10000;
  position: fixed;
  background: #fff;
  top: 0;
  right: 0;
  .navbar-base {
    display: flex;
    justify-content: center;
    font-size: 18px;
    letter-spacing:3px;

    .navbar-item-all{
      padding: 0 34px;
    }
    .navbar-item {
      border-right: 2px solid #eee;
      border-left-width: 100px;
    }
  }
}
</style>
