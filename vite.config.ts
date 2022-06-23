import { join } from 'path';
import IconsResolver from 'unplugin-icons/resolver';
import Icons from 'unplugin-icons/vite';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';

import vue from '@vitejs/plugin-vue';

export default defineConfig({
    mode: process.env.MODE,
    resolve: {
        alias: [
            {
                find: '@/',
                replacement: join(__dirname, 'src') + '/',
            },
        ],
    },
    plugins: [
        vue(),
        Components({
            resolvers: [IconsResolver()],
        }),
        Icons(),
    ],
    build: {
        sourcemap: true,
    },
});
