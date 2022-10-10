<template>
  <div class="manage-contianer">
    <p class="d-none">{{user}}</p>
    <div class="row m-5">

      <b-list-group class="col-md-2 amange-nav">
        <b-list-group-item >
          <h4 style="line-height:1.8em">
            <b-badge :class="{'badge-success':user.Super,'badge-info':!user.Super}" >{{user.Super?'管理员':'拟稿人'}}</b-badge>
            <br/>
            你好，{{user.name}}
            <br/>
            <b-button variant="outline-secondary" @click="logout">退出登录</b-button>
            <hr>
            <b-button variant="info" href="/#/secret/edit" target="_blank">新建文章</b-button>

          </h4>
        </b-list-group-item>
        <b-list-group-item :class="{'active':(choice==1)}" @click="handleChoice(1)">首页图片修改</b-list-group-item>
        <b-list-group-item :class="{'active':(choice==2)}" @click="handleChoice(2)">基本页修改</b-list-group-item>
        <b-list-group-item :class="{'active':(choice==3)}" @click="handleChoice(3)">全部文章列表</b-list-group-item>
        <b-list-group-item :class="{'active':(choice==4)}" @click="handleChoice(4)">待发布文章列表</b-list-group-item>
      </b-list-group>

      <div class="col-10" style="line-height:2.2em" v-if="choice==1">
        <b-card no-body>
          <b-tabs pills card>

            <b-tab title="首页Logo" active>
              <b-card-text>
                <h4>电话图片修改</h4>
                <uploadimg :imgname="'indexLogo01.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>Logo图片修改</h4>
                <uploadimg :imgname="'indexLogo02.jpg'"/>
              </b-card-text>
            </b-tab>

            <b-tab title="导航栏" >

              <div class="row">
                <div class="col-md-4 p-3">
                  <h2>首页导航栏</h2>
                  <b-card-text v-for="(i,index) in active.navbar" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index,newname[index])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>
                <div class="col-md-4 p-3">
                  <h2>首页新闻栏</h2>
                  <b-card-text v-for="(i,index) in active.newbar" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+8]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index+8,newname[index+8])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>
              </div>
            </b-tab>

            <b-tab title="首页轮播图01">
              <b-alert show>最佳展示尺寸：1600*400</b-alert>
              <b-card-text>
                <h4>轮播图01</h4>
                <uploadimg :imgname="'indexAD0101.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图02</h4>
                <uploadimg :imgname="'indexAD0102.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图03</h4>
                <uploadimg :imgname="'indexAD0103.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图04</h4>
                <uploadimg :imgname="'indexAD0104.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图05</h4>
                <uploadimg :imgname="'indexAD0105.jpg'"/>
              </b-card-text>
            </b-tab>

            <b-tab title="首页新闻图">
              <b-alert show>最佳展示尺寸：400*240</b-alert>
              <b-card-text>
                <h4>新闻图01</h4>
                <uploadimg :imgname="'indexNEW01.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>新闻图02</h4>
                <uploadimg :imgname="'indexNEW02.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>新闻图03</h4>
                <uploadimg :imgname="'indexNEW03.jpg'"/>
              </b-card-text>
            </b-tab>

            <b-tab title="首页轮播图02">
              <b-alert show>最佳展示尺寸：1600*400</b-alert>
              <b-card-text>
                <h4>轮播图01</h4>
                <uploadimg :imgname="'indexAD0201.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图02</h4>
                <uploadimg :imgname="'indexAD0202.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图03</h4>
                <uploadimg :imgname="'indexAD0203.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图04</h4>
                <uploadimg :imgname="'indexAD0204.jpg'"/>
              </b-card-text>
              <b-card-text>
                <h4>轮播图05</h4>
                <uploadimg :imgname="'indexAD0205.jpg'"/>
              </b-card-text>
            </b-tab>

            <b-tab title="底部导航栏">
              <div class="row">

                <div class="col-md-4 p-3">
                  <h2>底部导航栏01</h2>
                  <b-card-text v-for="(i,index) in active.footbar01" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+11]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index+11,newname[index+11])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>

                <div class="col-md-4 p-3">
                  <h2>底部导航栏02</h2>
                  <b-card-text v-for="(i,index) in active.footbar02" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+16]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index+16,newname[index+16])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>

                <div class="col-md-4 p-3">
                  <h2>底部导航栏03</h2>
                  <b-card-text v-for="(i,index) in active.footbar03" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+21]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index+21,newname[index+21])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>

                <div class="col-md-4 p-3">
                  <h2>底部导航栏04</h2>
                  <b-card-text v-for="(i,index) in active.footbar04" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+26]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(+26,newname[index+26])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>

                <div class="col-md-4 p-3">
                  <h2>底部导航栏05</h2>
                  <b-card-text v-for="(i,index) in active.footbar05" :key="index">
                    <b-input-group :prepend="i" class="mt-3">
                      <b-form-input placeholder="新名字" v-model="newname[index+31]"></b-form-input>
                      <b-input-group-append>
                        <b-button variant="outline-success" @click="handleModInfo(index+31,newname[index+31])">保存</b-button>
                      </b-input-group-append>
                    </b-input-group>
                  </b-card-text>
                </div>

              </div>
            </b-tab>

          </b-tabs>
        </b-card>

      </div>

      <div class="col-10" style="line-height:2.2em" v-if="choice==2">
        <b-card no-body>
          <b-tabs pills card>

            <b-tab title="导航栏文章" >

              <div class="row">
                <div class="col-md-4 p-3">
                  <h3 class="my-3">首页导航栏</h3>
                  <div class="my-2" v-for="(i,index) in active.navbar" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index))">文章修改</b-button>
                    </div>
                  </div>
                </div>
              </div>
            </b-tab>

            <b-tab title="底部导航栏">
              <div class="row">

                <div class="col-md-4 p-3">
                  <h3 class="my-3">底部导航栏01</h3>
                  <div class="my-2" v-for="(i,index) in active.footbar01" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index+7))">文章修改</b-button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4 p-3">
                  <h3 class="my-3">底部导航栏02</h3>
                  <div class="my-2" v-for="(i,index) in active.footbar02" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index+11))">文章修改</b-button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4 p-3">
                  <h3 class="my-3">底部导航栏03</h3>
                  <div class="my-2" v-for="(i,index) in active.footbar03" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index+15))">文章修改</b-button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4 p-3">
                  <h3 class="my-3">底部导航栏04</h3>
                  <div class="my-2" v-for="(i,index) in active.footbar04" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index+19))">文章修改</b-button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4 p-3">
                  <h3 class="my-3">底部导航栏05</h3>
                  <div class="my-2" v-for="(i,index) in active.footbar05" :key="index">
                    <div v-if="index!=0">
                      <b>{{i}}：</b>
                      <b-button variant="outline-success" @click="$router.push('/secret/edit?id='+(index+23))">文章修改</b-button>
                    </div>
                  </div>
                </div>

              </div>
            </b-tab>
          </b-tabs>
        </b-card>

      </div>

      <div class="col-10" style="line-height:2.2em" v-if="choice==3">
        <b-card no-body>
          <b-tabs pills card>

            <b-tab title="全部文章列表" active @click="handleChoice(3)">
              <div class="info-list mx-4 my-3" v-for="(i,index) in artlist" :key="index">
                <div class="info-list-item">
                  <b-badge pill variant="light">{{i.artid}}</b-badge >
                  <b-badge :class="{'badge-success':i.public,'badge-info':!i.public}" >{{i.public?'已发布':'未发布'}}</b-badge>
                  <a :href="'/#/article?id='+i.artid" target="_blank">{{i.title}}</a>
                  <div class="float-right">
                    <b-button variant="info" :href="'/#/secret/edit?id='+i.artid" target="_blank">编辑</b-button>
                    <b-button variant="success" @click="handlePubArt(i.artid)">发布</b-button>
                  </div>
                  <hr>
                </div>
              </div>
              <div class="m-3">
                <b-button-group>
                  <b-button variant="success" @click="beforeArtList">前一页</b-button>
                  <b-button variant="light">第{{pageid}}页</b-button>
                  <b-button variant="info" @click="nextArtList">后一页</b-button>
                </b-button-group>
              </div>
            </b-tab>

            <b-tab title="未发布文章列表" @click="handleChoice(5)">
              <div class="info-list mx-4 my-3" v-for="(i,index) in artlist" :key="index">
                <div class="info-list-item">
                  <b-badge :class="{'badge-success':i.public,'badge-info':!i.public}" >{{i.public?'已发布':'未发布'}}</b-badge>
                  <a :href="'/#/article?id='+i.artid" target="_blank">{{i.title}}</a>
                  <span>{{i.create}}</span>
                  <div class="float-right">
                    <b-button variant="info" :href="'/#/secret/edit?id='+i.artid" target="_blank">编辑</b-button>
                    <b-button variant="success" @click="handlePubArt(i.artid)">发布</b-button>
                  </div>
                  <hr>
                </div>
              </div>
              <div class="m-3">
                <b-button-group>
                  <b-button variant="success" @click="beforeArtList">前一页</b-button>
                  <b-button variant="light">第{{pageid}}页</b-button>
                  <b-button variant="info" @click="nextArtList">后一页</b-button>
                </b-button-group>
              </div>
            </b-tab>

          </b-tabs>
        </b-card>

      </div>

    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Uploadimg from '@/components/Uploadimg.vue'

export default {
  components: { Uploadimg },
  created () {
    this.$store.dispatch('infoModule/active')
  },
  data () {
    return {
      choice: 1,
      pageid: 1,
      newname: [],
      hide: 0
    }
  },
  computed: {
    ...mapState('userModule', {
      user: state => state.userInfo
    }),
    ...mapState('infoModule', {
      active: state => state.active
    }),
    ...mapState('manageModule', {
      artlist: state => state.artlist
    })
  },
  methods: {
    ...mapActions('userModule', { logout: 'logout' }),
    ...mapActions('manageModule', { modinfo: 'modinfo', allart: 'allart', pubart: 'pubart' }),
    handleChoice (choice) {
      this.choice = choice
      if (choice === 3) {
        this.hide = 0
        this.allart({'pageid': this.pageid, 'hide': this.hide})
      }
      if (choice === 5) {
        this.choice = 3
        this.hide = 1
        this.allart({'pageid': this.pageid, 'hide': this.hide})
      }
    },
    handleModInfo (navid, key) {
      this.modinfo({'navid': navid + 1, 'key': key})
    },
    handlePubArt (artid) {
      this.pubart({'artid': artid})
    },
    nextArtList () {
      this.pageid++
      this.allart({'pageid': this.pageid, 'hide': this.hide})
    },
    beforeArtList () {
      if (this.pageid > 1) {
        this.pageid--
      }
      this.allart({'pageid': this.pageid, 'hide': this.hide})
    }
  }
}
</script>

<style scoped lang="scss">
body{
  background: #f5f5f5;
}
h4{
  margin: 12px 0;
}
.manage-nav{
  border-radius: 8px;
}
</style>
