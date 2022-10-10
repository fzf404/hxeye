import request from '@/utils/request.js'
import qs from 'qs'

const active = () => (request.get('active'))
const getart = ({artid}) => (request.post('getart', qs.stringify({artid})))
const artlist = ({pageid}) => (request.post('artlist', qs.stringify({pageid})))
const search = ({key}) => (request.post('search', qs.stringify({key})))

export default {
  active,
  getart,
  artlist,
  search
}
