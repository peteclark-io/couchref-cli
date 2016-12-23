module.exports = {
    entry: ['whatwg-fetch', './app/main.js'],
    output: {
        path: __dirname + '/../public/js',
        publicPath: '/js/',
        filename: "bundle.js"
    },
    module: {
        loaders: [{
            test: /\.json$/,
            loader: "json-loader"
        }, {
            test: /\.js$/,
            exclude: /(node_modules|bower_components)/,
            loader: 'babel',
            query: {
                presets: ['es2015']
            }
        }, {
            test: /\.css$/,
            loader: 'style-loader'
        }, {
            test: /\.css$/,
            loader: 'css-loader',
            query: {
                sourceMap: false,
                modules: true,
                localIdentName: '[local]-[hash:base64:4]',
                minimize: true,
            }
        }, {
            test: /\.css$/,
            loader: 'postcss-loader'
        }, {
            test: /\.scss$/,
            loaders: ["style", "css", "sass"]
        }, {
            test: /\.(png|jpe?g|gif|svg|woff|woff2|ttf|eot|ico)(\?.*)?$/,
            loader: 'url-loader?limit=50000'
        }]
    }
};
