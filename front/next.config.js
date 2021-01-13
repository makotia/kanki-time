const { resolve } = require('path')
const isProd = process.env.NODE_ENV === "production"

const nextConfig = {
  env: {
    baseUrl: isProd ? '' : 'http://localhost:3000',
    apiUrl: isProd ? '' : 'http://localhost:8080'
  }
}

module.exports = nextConfig
