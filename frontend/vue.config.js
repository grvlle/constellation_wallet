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

const path = require("path");
// const PrerenderSPAPlugin = require("prerender-spa-plugin");

// module.exports = {
//
//   plugins: [
//     new PrerenderSPAPlugin({
//       // Required - The path to the webpack-outputted app to prerender.
//       staticDir: path.join(__dirname, "dist"),
//       // Required - Routes to render.
//       routes: ["/", "/login", "/dashboard"],
//     }),
//   ],
// };

module.exports = {
  transpileDependencies: [
    "secp256k1",
    "keccak"
  ],
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
    // config.plugin("html")
    //   .tap(args => {
    //     args[0].template = path.join(__dirname, "src", "index.html")
    //     return args
    //   })
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
    // resolve: {
    //   alias: {
    //     "keccak": path.resolve(__dirname, 'node_modules/@stardust-collective/dag-keystore/shim/keccak/'),
    //     "secp256k1": path.resolve(__dirname, 'node_modules/@stardust-collective/dag-keystore/shim/secp256k1/')
    //   }
    // }
  },
  devServer: {
    disableHostCheck: true,
    host: "localhost",
  },
};
