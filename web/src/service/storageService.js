// 本地缓存服务
const PREFIX = 'fzf404.com01_'

const USER_PREFIX = `${PREFIX}user_`
const USER_TOKEN = `${USER_PREFIX}token`
const USER_INFO = `${USER_PREFIX}info`

const set = (key, data) => {
  localStorage.setItem(key, data)
}

const get = (key) => localStorage.getItem(key)

export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO
}
