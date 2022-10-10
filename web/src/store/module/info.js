
import infoService from '@/service/infoService'
import router from '@/router'

const infoModule = {
  namespaced: true,
  state: {
    active: {},
    article: {},
    artlist: {}
  },
  mutations: {
    SET_ACTIVE (state, data) {
      state.active = data
    },
    SET_LIST (state, data) {
      state.artlist = data
    },
    GET_ARTICLE (state, data) {
      state.article = data
    }
  },
  actions: {
    active (context) {
      return new Promise((resolve, reject) => {
        infoService.active().then((res) => {
          if (res.data.code !== 200) {
            router.push({ name: '404' })
            return
          }
          context.commit('SET_ACTIVE', res.data.data)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    getart (context, {artid}) {
      return new Promise((resolve, reject) => {
        infoService.getart({artid}).then((res) => {
          if (res.data.code !== 200) {
            router.push({ name: '404' })
            return
          }
          context.commit('GET_ARTICLE', res.data.data.article)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    artlist (context, {pageid}) {
      return new Promise((resolve, reject) => {
        infoService.artlist({pageid}).then((res) => {
          if (res.data.code !== 200) {
            alert('已经到达末尾了')
            return
          }
          context.commit('SET_LIST', res.data.data.arts)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    search (context, {key}) {
      return new Promise((resolve, reject) => {
        infoService.search({key}).then((res) => {
          if (res.data.code !== 200) {
            alert('什么都没找到')
            router.push({ name: '404' })
          }
          context.commit('SET_LIST', res.data.data.arts)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    }

  }
}

export default infoModule
