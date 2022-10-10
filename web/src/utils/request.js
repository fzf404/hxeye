import axios from 'axios'
import storageService from '../service/storageService'

const service = axios.create({
  baseURL: 'http://localhost:8080/',
  timeout: 1000 * 3
})
// Add a request interceptor
service.interceptors.request.use((config) => {
  // Do something before request is sent
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` })
  return config
}, (error) => (Promise.reject(error)))

export default service
