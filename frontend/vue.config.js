let loaderOptions = {
  sass: {
    prependData: `
    @import "@/assets/sass/paper/_variables.scss";
    @import "@/assets/sass/paper/_themes.scss";
    `,
  },
};

let cssConfig = {
  loaderOptions,
};

if (process.env.NODE_ENV == "production") {
  cssConfig = {
    extract: {
      filename: "[name].css",
      chunkFilename: "[name].css",
    },
    loaderOptions,
  };
}

module.exports = {
  chainWebpack: (config) => {
    let limit = 9999999999999999;
    config.module
      .rule("images")
      .test(/\.(png|gif|jpg|jpeg)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .tap((options) => Object.assign(options, { limit: limit }));
    config.module
      .rule("svg")
      .test(/\.svg$/)
      .use("vue-svg-loader")
      .loader("vue-svg-loader");
    config.module
      .rule("fonts")
      .test(/\.(woff2?|eot|ttf|otf)(\?.*)?$/i)
      .use("url-loader")
      .loader("url-loader")
      .options({
        limit: limit,
      });
  },
  css: cssConfig,
  configureWebpack: {
    target: 'web',
    output: {
      filename: "[name].js",
    },
    optimization: {
      splitChunks: false,
    },
  },
  devServer: {
    disableHostCheck: true,
    host: "localhost",
  },
};
