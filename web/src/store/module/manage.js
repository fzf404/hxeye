
import manageService from '@/service/manageService'
import router from '@/router'

const manageModule = {
  namespaced: true,
  state: {
    imgUrl: '',
    artid: '',
    artlist: [],
    article: null
  },
  mutations: {
    SET_IMGURL (state, imgUrl) {
      state.imgUrl = imgUrl
    },
    SET_ARTID (state, artid) {
      state.artid = artid
    },
    SET_ARTLIST (state, artlist) {
      state.artlist = artlist
    },
    SET_ARTICLE (state, article) {
      state.article = article
    }
  },
  actions: {
    upimg (context, { img, imgname }) {
      return new Promise((resolve, reject) => {
        manageService.upimg(img, imgname).then((res) => {
          alert(res.data.msg)
          if (res.data.code !== 200) {
            router.push({ name: 'Login' })
          }
          context.commit('SET_IMGURL', res.data.data[0])
          resolve(res)
        }).catch((err) => {
          reject(err)
        })
      })
    },
    addart (context, { title, type, content }) {
      return new Promise((resolve, reject) => {
        manageService.addart({title, type, content}).then((res) => {
          alert(res.data.msg)
          if (res.data.code !== 200) {
            router.push({ name: 'Login' })
            return
          }
          context.commit('SET_ARTID', res.data.data.artid)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    allart (context, { pageid, hide }) {
      return new Promise((resolve, reject) => {
        manageService.allart({pageid, hide}).then((res) => {
          if (res.data.code !== 200) {
            alert('已经到达末尾了')
            return false
          }
          context.commit('SET_ARTLIST', res.data.data.arts)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    modart (context, {artid, title, type, content}) {
      return new Promise((resolve, reject) => {
        manageService.modart({artid, title, type, content}).then((res) => {
          alert(res.data.msg)
          if (res.data.code !== 200) {
            router.push({ name: 'Login' })
            return
          }
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    pubart (context, {artid}) {
      return new Promise((resolve, reject) => {
        manageService.pubart({artid}).then((res) => {
          alert(res.data.msg)
          if (res.data.code !== 200) {
            return
          }
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    modinfo (context, {navid, key}) {
      return new Promise((resolve, reject) => {
        manageService.modinfo({navid, key}).then((res) => {
          alert(res.data.msg)
          if (res.data.code !== 200) {
            return
          }
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    },
    secretart (context, {artid}) {
      return new Promise((resolve, reject) => {
        manageService.secretart({artid}).then((res) => {
          if (res.data.code !== 200) {
            router.push({ name: '404' })
            return
          }
          context.commit('SET_ARTICLE', res.data.data.article)
          resolve(res)
        }).catch((err) => {
          alert('后端服务器出现故障，请联系管理员。')
          reject(err)
        })
      })
    }
  }
}

export default manageModule
