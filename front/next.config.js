const { resolve } = require('path')

const nextConfig = {
  env: {
    baseUrl: process.env.BASE_URL ? process.env.BASE_URL : 'http://localhost:3000',
    apiUrl: process.env.API_URL ? process.env.API_URL : 'http://localhost:8080',
  }
}

module.exports = nextConfig
