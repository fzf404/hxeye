<template>
  <div class="footer bg-light py-5">
    <div class="container">
      <div class="row text-secondary">

        <div class="col-md-2">
          <ul class="nav flex-column">
            <li class="nav-item mx-3 my-2">
              <h5>
                  {{activeItem.footbar01[0]}}
              </h5>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=8" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar01[1]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=9" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar01[2]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=10" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar01[3]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=11" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar01[4]}}
              </a>
            </li>
          </ul>
        </div>

        <div class="col-md-2">
          <ul class="nav flex-column">
            <li class="nav-item mx-3 my-2">
              <h5>
                {{activeItem.footbar02[0]}}
              </h5>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=12" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar02[1]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=13" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar02[2]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=14" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar02[3]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=15" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar02[4]}}
              </a>
            </li>
          </ul>
        </div>
        <div class="col-md-2">
          <ul class="nav flex-column">
            <li class="nav-item mx-3 my-2">
              <h5>
                {{activeItem.footbar03[0]}}
              </h5>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=16" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar03[1]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=17" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar03[2]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=18" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar03[3]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=19" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar03[4]}}
              </a>
            </li>
          </ul>
        </div>

        <div class="col-md-2">
          <ul class="nav flex-column">
            <li class="nav-item mx-3 my-2">
              <h5>
                {{activeItem.footbar04[0]}}
              </h5>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=20" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar04[1]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=21" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar04[2]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=22" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar04[3]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=23" class="nav-link text-secondary" target="_blank">
                {{activeItem.footbar04[4]}}
              </a>
            </li>
          </ul>
        </div>

        <div class="col-md-1"></div>

        <div class="col-md-3">
          <ul class="nav flex-column">

            <li class="nav-item mx-3 my-2">
              <h5>
                &ensp;{{activeItem.footbar05[0]}}
              </h5>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=24" class="nav-link text-secondary" target="_blank">
                <b-icon icon="envelope-open-fill" />
                &ensp;{{activeItem.footbar05[1]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=25" class="nav-link text-secondary" target="_blank">
                <b-icon icon="phone" />
                &ensp;{{activeItem.footbar05[2]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=26" class="nav-link text-secondary" target="_blank">
                <b-icon icon="cursor-fill" />
                &ensp;{{activeItem.footbar05[3]}}
              </a>
            </li>
            <li class="nav-item">
              <a href="/#/article?id=27" class="nav-link text-secondary" target="_blank">
                <b-icon icon="chat-fill" />
                &ensp;{{activeItem.footbar05[4]}}
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState('infoModule', {
      activeItem: state => state.active
    })
  }
}
</script>

<style scoped>
h5{
  color:#454545;
}
</style>
