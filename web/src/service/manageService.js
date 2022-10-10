import request from '@/utils/request.js'
import qs from 'qs'

const config = { 'Content-Type': 'multipart/form-data' }

const upimg = (img, imgname) => {
  let formData = new FormData()
  formData.append('img', img)
  formData.append('imgname', imgname)
  console.log(imgname)
  return request.post('http://localhost:8080/upimg', formData, config)
}
const addart = ({title, type, content}) => (request.post('addart', qs.stringify({title, type, content})))
const modart = ({artid, title, type, content}) => (request.post('modart', qs.stringify({artid, title, type, content})))
const pubart = ({artid}) => (request.post('pubart', qs.stringify({artid})))
const allart = ({pageid, hide}) => (request.post('allart', qs.stringify({pageid, hide})))
const modinfo = ({navid, key}) => (request.post('modinfo', qs.stringify({navid, key})))
const secretart = ({artid}) => (request.post('secretart', qs.stringify({artid})))

export default {
  upimg,
  addart,
  modart,
  pubart,
  allart,
  modinfo,
  secretart
}
