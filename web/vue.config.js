module.exports = {
  publicPath: './',
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://localhost:8082',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
