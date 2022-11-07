import Axios from 'axios'
import { API_PORT, IPADDRESS } from '../config/index'

const axios = Axios.create({
  baseURL: `http://${IPADDRESS}:${API_PORT}`
})

export default axios
