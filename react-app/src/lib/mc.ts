import * as minio from 'minio'
import { IPADDRESS, MINIO_PORT, MINIO_ROOT_USER, MINIO_ROOT_PASS } from '../config/index'

const mc = new minio.Client({
  endPoint: IPADDRESS,
  port: MINIO_PORT,
  useSSL: false,
  accessKey: MINIO_ROOT_USER,
  secretKey: MINIO_ROOT_PASS
})

export default mc
