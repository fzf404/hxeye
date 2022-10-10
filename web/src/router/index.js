import Vue from 'vue'
import Router from 'vue-router'
import store from '@/store'
import manageRoutes from './module/manage'
import infoRoutes from './module/info'

Vue.use(Router)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Index.vue')
  },
  ...manageRoutes,
  ...infoRoutes,
  {
    path: '*',
    name: '404',
    component: () => import('../views/404.vue')
  }
]

const router = new Router({
  routes
})
router.beforeEach((to, from, next) => {
  if (to.meta.auth) {
    // 判断用户是否登录
    if (store.state.userModule.token) {
      next()
    } else {
      router.push({ name: 'Home' })
    }
  } else {
    next()
  }
})

export default router
