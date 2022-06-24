<script setup lang="ts">
import VLink from '@/components/VLink.vue';

import IconGithub from '~icons/mdi/github';
</script>

<template>
    <div class="w-full h-[100vh] flex flex-col justify-center items-center">
        <div class="bg-guest-lighten rounded-xl p-8 flex flex-col items-center shadow-md">
            <slot />
        </div>
        <template v-if="$slots.or">
            <div class="mt-4 mb-2 text-guest-light text-xs select-none">OR</div>
            <slot name="or" />
        </template>
        <div class="fixed bottom-0 p-4">
            <v-link
                href="https://github.com/playmean/guest"
                class="text-xl text-zinc-500 hover:text-zinc-400 focus:text-zinc-400"
                target="_blank"
            >
                <icon-github />
            </v-link>
        </div>
    </div>
</template>
