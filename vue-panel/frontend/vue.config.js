const path = require('path');
const webpack = require('webpack');

function resolveSrc(_path) {
  return path.join(__dirname, _path);
}
module.exports = {
  publicPath: process.env.NODE_ENV === 'production' ? './' : '/',
  filenameHashing: false,
  lintOnSave: true,
  productionSourceMap: false,
  configureWebpack: {
    resolve: {
      alias: {
        src: resolveSrc('src'),
        assets: resolveSrc('src/assets'),
      }
    },
    plugins: [
      new webpack.optimize.LimitChunkCountPlugin({
        maxChunks: 2
      })
    ],
  },
  css: {
    sourceMap: process.env.NODE_ENV !== 'production'
  }
};
