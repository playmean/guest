<script setup lang="ts">
defineEmits<{
    (e: 'resizeStart', event: MouseEvent): void;
}>();
</script>

<template>
    <div
        class="hover:bg-guest-accent rounded-full cursor-col-resize select-none transition-colors"
        @mousedown="$emit('resizeStart', $event)"
    ></div>
</template>
