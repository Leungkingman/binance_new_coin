const path = require('path')
const resolve = (dir) => {
  return path.join(__dirname, dir)
}

module.exports = {
	lintOnSave: false,
	chainWebpack: (config) => {
    config.resolve.alias
      .set('@', resolve('src'))
      .set('assets', resolve('src/assets'))
  }
}