<script setup lang="ts">
import type { Component } from 'vue';

import { useVModel } from '@vueuse/core';

const props = defineProps<{
    entries: [string, string][];
    modelValue: string;
    label?: Component | string;
    placeholder?: string;
}>();

const localValue = useVModel(props, 'modelValue');
</script>

<template>
    <label v-if="label" class="flex flex-row w-full">
        <div
            class="flex items-center px-2 border border-r-0 rounded-l-sm border-guest-light text-sm text-zinc-400 select-none"
        >
            <span v-if="typeof label === 'string'">{{ label }}</span>
            <component v-else :is="label" />
        </div>
        <select
            v-model="localValue"
            class="guest-control !rounded-l-none border-t border-b text-sm"
            :placeholder="placeholder"
        >
            <option
                v-for="[title, value] in entries"
                class="text-sm"
                :key="title"
                :value="value"
            >
                {{ title }}
            </option>
        </select>
    </label>
    <select
        v-else
        v-model="localValue"
        class="guest-control border text-sm"
        :placeholder="placeholder"
    >
        <option
            v-for="[title, value] in entries"
            class="text-sm"
            :key="title"
            :value="value"
        >
            {{ title }}
        </option>
    </select>
</template>
