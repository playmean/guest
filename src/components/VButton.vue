<script setup lang="ts">
import type { Component } from 'vue';

defineProps<{
    outline?: boolean;
    icon?: Component;
}>();
</script>

<template>
    <button
        type="button"
        class="relative flex justify-center items-center px-2 h-8 text-white text-sm rounded-sm outline-none border border-transparent focus:!border-white select-none transition-colors"
        :class="{
            'text-guest-accent border border-guest-accent hover:bg-guest-accent hover:text-white':
                outline,
            'bg-guest-accent text-white/90 hover:bg-guest-accent/90': !outline,
        }"
    >
        <component
            v-if="icon"
            :class="{
                'absolute left-2': $slots.default,
            }"
            :is="icon"
        />
        <slot />
    </button>
</template>
