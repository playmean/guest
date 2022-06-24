<script setup lang="ts">
import { shortcuts } from './definitions';
import type { Shortcuts } from './types';

defineProps<{
    shortcut: Shortcuts;
}>();
</script>

<template>
    <template v-if="shortcuts[shortcut]">
        {{ shortcuts[shortcut].label }}
        <component :is="shortcuts[shortcut].icon" class="ml-1 mt-0.5 text-[0.6rem]" />
    </template>
</template>
