import request from '@/utils/request.js'
import qs from 'qs'

// 用户登录
const login = ({ name, password }) => (request.post('login', qs.stringify({ name, password })))
// // 获取用户信息
const info = () => (request.post('userinfo'))
// // 修改密码
// const modpasswd = ({ name, password }) => (request.post('register', qs.stringify({ name, password })))
export default {
  // modpasswd,
  login,
  info
}
