const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require('webpack');

module.exports = {
  entry: './client/index.js',
  output: {
    path: path.resolve(__dirname, './dist'),
    filename: 'bundle.js',
  },
  plugins: [
    new HtmlWebpackPlugin({ title: 'All the brands you know and love!' }),
  ],
  devtool: 'source-map',
  module: {
    loaders: [
      {
        test: /\.(js|jsx)$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.css$/,
        loaders: ['style-loader', 'css-loader'],
      },
      {
        test: /\.(svg|ttf|woff|woff2|eot)$/,
        loader: 'url-loader?limit=5000',
      },
      {
        test: /\.png$/,
        loader: 'url-loader',
        query: { mimetype: 'image/png' },
      },
    ],
  },
};

