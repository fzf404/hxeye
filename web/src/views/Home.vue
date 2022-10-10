<template>
  <div>
    <div class="index-show">
      <b-carousel
        id="carousel-1"
        v-model="slide"
        :interval="4000"
        controls
        indicators
        background="#ababab"
        img-width="1024"
        img-height="640"
        style="text-shadow: 1px 1px 2px #333;"
        @sliding-start="onSlideStart"
        @sliding-end="onSlideEnd"
      >

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0101.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0102.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0103.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0104.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0105.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

      </b-carousel>
    </div>

    <div class="info-container" v-if="activeItem.newbar">
      <div class="body-container row">

        <div class="col-md-4 news-contianer">
          <b-card
            img-src="/images/indexNEW01.jpg"
            img-alt="Image"
            img-left
            tag="article"
            class="mx-4"
          >
            <b-card-title style="text-align:center; line-height:1.4em">
            <b>{{activeItem.newbar[0].slice(0,2)}}</b>
            <br/>
            {{activeItem.newbar[0].slice(2)}}
            </b-card-title>
          </b-card>
          <div class="info-list mx-4 my-3" v-for="(i,index) in activeItem.art01" :key="index">
            <div class="info-list-item">
              <a :href="'/#/article?id='+i.artid" target="_blank">{{i.title.slice(0,16)}}...</a>
              <span>{{i.create}}</span>
            </div>
          </div>
        </div>

        <div class="col-md-4 news-contianer">
          <b-card
            img-src="/images/indexNEW02.jpg"
            img-alt="Image"
            img-left
            tag="article"
            class="mx-4"
          >
          <b-card-title style="text-align:center; line-height:1.4em">
            <b>{{activeItem.newbar[1].slice(0,2)}}</b>
            <br/>
            {{activeItem.newbar[1].slice(2)}}
            </b-card-title>
          </b-card>
          <div class="info-list mx-4 my-3" v-for="(i,index) in activeItem.art02" :key="index">
            <div class="info-list-item">
              <a :href="'/#/article?id='+i.artid" target="_blank">{{i.title.slice(0,16)}}...</a>
              <span>{{i.create}}</span>
            </div>
          </div>
        </div>

        <div class="col-md-4 news-contianer">
          <b-card
            img-src="/images/indexNEW03.jpg"
            img-alt="Image"
            img-left
            tag="article"
            class="mx-3"
          >
          <b-card-title style="text-align:center; line-height:1.4em">
            <b>{{activeItem.newbar[2].slice(0,2)}}</b>
            <br/>
            {{activeItem.newbar[2].slice(2)}}
            </b-card-title>
          </b-card>
          <div class="info-list mx-4 my-3" v-for="(i,index) in activeItem.art03" :key="index">
            <div class="info-list-item">
              <a :href="'/#/article?id='+i.artid" target="_blank">{{i.title.slice(0,16)}}...</a>
              <span>{{i.create}}</span>
            </div>
          </div>
        </div>

      </div>
    </div>
    <div class="index-show">
      <b-carousel
        id="carousel-1"
        v-model="slide"
        :interval="4000"
        controls
        indicators
        background="#ababab"
        img-width="1024"
        img-height="640"
        style="text-shadow: 1px 1px 2px #333;"
        @sliding-start="onSlideStart"
        @sliding-end="onSlideEnd"
      >

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0201.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0202.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0203.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0204.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

        <b-carousel-slide>
          <template v-slot:img>
            <img
              class="d-block img-fluid w-100"
              src="/images/indexAD0205.jpg"
              alt="image slot"
            >
          </template>
        </b-carousel-slide>

      </b-carousel>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {

  data () {
    return {
      slide: 0,
      sliding: null
    }
  },
  computed: {
    ...mapState('infoModule', {
      activeItem: state => state.active
    })
  },
  methods: {
    onSlideStart (slide) {
      this.sliding = true
    },
    onSlideEnd (slide) {
      this.sliding = false
    }
  }
}
</script>

<style scoped lang="scss">
.index-show{
  img{
    height: 400px;
  }
}
span{
  float: right;
}
.card{
  border-radius: 24px !important;
}
.body-container{
  max-width: 84%;
  margin: auto;
  border-radius: 10px;
  .news-contianer {
    padding-top:30px;
    // border-right: 2px solid #eee;
    img{
      border-top-left-radius: 18px;
      border-bottom-left-radius: 18px;
      width: 200px;
      height: 120px;
      margin: auto;
    }
    .info-list-item{
      border-bottom: 1px dashed #454545;
      line-height: 1.6em;
    }
  }
}
.info-container{
  background: rgba($color: #bddeff, $alpha: 0.4);
}
</style>
