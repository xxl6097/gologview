const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath:'.',
  outputDir: '../logview',
  productionSourceMap:false//不生成map文件
})
