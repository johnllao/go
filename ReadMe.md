# React Environment Setup

* Install NodeJS
  NodeJS is ncessary to have the npm to get all the packages necessary
* Install Webpack
  `npm install webpack --save`
* Install Babel
  `npm install babel-core babel-loader babel-preset-es2015 babel-preset-react --save`
* Install React
  `npm install react react-dom --save`
* Install Axios
  `npm install axios --save`
* Run `npm init` from the root directory
* Confgure `package.json`
  ```
  {
    "name": "playreact",
    "version": "1.0.0",
    "description": "React play ground",
    "main": "index.js",
    "scripts": {
      "test": "echo \"Error: no test specified\" && exit 1"
    },
    "author": "John L. Lao",
    "license": "ISC",
    "dependencies": {
      "babel-core": "^6.25.0",
      "babel-loader": "^7.0.0",
      "babel-preset-es2015": "^6.24.1",
      "babel-preset-react": "^6.24.1",
      "react": "^15.5.4",
      "react-dom": "^15.5.4",
      "webpack": "^2.6.1"
    },
    "babel": {
        "presets": [ "es2015", "react" ]
      ]
    }
  }
  ```
* Create `webpack.config.js` to the root dir
* Configure `webpack.config.js`
  ```
  var webpack = require('webpack');
  var path    = require('path');

  var BUILD_DIR = path.resolve(__dirname, '<output path of the packed javascript>');
  var APP_DIR   = path.resolve(__dirname, '<path of the jsx files>');

  var config = {
    entry : APP_DIR + '/index.jsx',
    output: {
      path    : BUILD_DIR,
      filename: 'index.js'
    },
    module: {
      loaders: [{
        include: APP_DIR,
        loader: "babel-loader" 
      }]
    }
  };

  module.exports = config;
  ```
* Create `index.html`
  ```
  <!DOCTYPE html>
  <html>
 
  <head>
    <title>React! React! React!</title>
  </head>
 
  <body>
    <div id="container"></div>
 
    <script src="index.js"></script>
  </body>
 
  </html>
  ```