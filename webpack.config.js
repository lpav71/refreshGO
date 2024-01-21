const path = require('path');
const { VueLoaderPlugin } = require('vue-loader');
const webpack = require('webpack');

module.exports = {
    mode: "development",
    entry: './app.js',
    output: {
        path: path.resolve(__dirname, './views/js'),
        filename: 'main.bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: 'vue-loader'
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            },
            {
                test: /\.(scss)$/,
                use: [
                    {
                        loader: 'style-loader'
                    },
                    {
                        loader: 'css-loader'
                    },
                    {
                        loader: 'postcss-loader',
                        options: {
                            postcssOptions: {
                                plugins: () => [
                                    require('autoprefixer')
                                ]
                            }
                        }
                    },
                    {
                        loader: 'sass-loader'
                    }
                ]
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.vue'],
        alias: {
            vue: 'vue/dist/vue.esm-bundler.js'
        }
    },
    plugins: [
        new VueLoaderPlugin(),
        new webpack.DefinePlugin({
            '__VUE_OPTIONS_API__': JSON.stringify(true),
            '__VUE_PROD_DEVTOOLS__': JSON.stringify(false),
        }),
    ],
};

/*
Описание webpack.config.js

"dev": "npx webpack --mode development", - сборка проекта в режиме разработчика (npm run dev)
"prod": "npx webpack --mode production" - сборка проекта в режиме Production (npm run prod)

1. Импортируются необходимые модули: path, VueLoaderPlugin и webpack.
2. Экспортируется объект конфигурации webpack.
3. Указывается точка входа для сборки - файл app.js в каталоге src.
4. Указывается опция output, которая указывает настройки выходного файла. Здесь используется опция clean: true, которая указывает на необходимость очистки выходной директории перед сборкой.
5. Описываются правила загрузки модулей. Здесь определены правила для загрузки файлов с расширением .vue и .css. Для загрузки файлов .vue используется vue-loader, а для загрузки файлов .css - style-loader и css-loader.
6. Описываются настройки разрешения импортов модулей. Здесь указываются расширения файлов .js и .vue, а также создается псевдоним для модуля vue, чтобы использовать бандлер с поддержкой опций API Vue 3.
7. Описываются плагины, используемые в сборке. Здесь используется VueLoaderPlugin для загрузки файлов .vue, а также DefinePlugin для определения глобальных переменных, таких как VUE_OPTIONS_API и VUE_PROD_DEVTOOLS.

Описание package.json

- "@vue/compiler-sfc": "^3.3.1": модуль для компиляции файлов .vue в JavaScript код.
- "babel-loader": "^9.1.2": загрузчик для использования Babel с webpack.
- "css-loader": "^6.7.3": загрузчик для обработки файлов CSS в JavaScript код.
- "style-loader": "^3.3.2": загрузчик для добавления стилей в DOM при использовании модулей CSS.
- "vue": "^3.2.36": основной модуль Vue.js версии 3.
- "vue-loader": "^17.0.1": загрузчик для компиляции файлов .vue в JavaScript код.
- "webpack": "^5.82.1": основной модуль webpack версии 5.
- "webpack-cli": "^5.1.1": интерфейс командной строки для webpack.
 */
