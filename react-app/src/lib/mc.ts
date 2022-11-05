import * as minio from 'minio'

const mc = new minio.Client({
  endPoint: "localhost",
  port: 9000,
  useSSL: false,
  accessKey: "admin",
  secretKey: "administrator"
})

export default mc
