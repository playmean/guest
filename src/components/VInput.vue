<script setup lang="ts">
import { useVModel } from '@vueuse/core';

import ShortcutHint from './Shortcut/ShortcutHint.vue';
import { Shortcuts } from './Shortcut/types';

const props = defineProps<{
    modelValue: string;
    label?: string;
    placeholder?: string;
    shortcut?: Shortcuts;
}>();

const localValue = useVModel(props, 'modelValue');
</script>

<template>
    <div class="relative flex items-center">
        <label v-if="label" class="flex flex-row">
            <div
                class="flex items-center px-2 border border-r-0 rounded-l-sm border-guest-light text-sm text-zinc-400 select-none"
            >
                {{ label }}
            </div>
            <input
                v-model="localValue"
                type="text"
                class="input rounded-l-none border-t border-b"
                :placeholder="placeholder"
            />
        </label>
        <input
            v-else
            v-model="localValue"
            type="text"
            class="input border"
            :placeholder="placeholder"
        />
        <div
            v-if="shortcut && localValue.length <= 10"
            class="absolute right-3 flex items-center text-xs text-zinc-500 pointer-events-none"
        >
            <shortcut-hint :shortcut="shortcut" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.input {
    @apply bg-guest-light px-3 py-2 rounded-sm outline-none border-transparent focus:border-guest-accent placeholder:text-zinc-600 placeholder:text-sm;
}
</style>
