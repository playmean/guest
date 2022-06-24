<script setup lang="ts">
import { ref } from 'vue';

import VInput from './components/VInput.vue';
import VLink from './components/VLink.vue';
import InitialLayout from './layouts/InitialLayout.vue';

const repoPath = ref('');
</script>

<template>
    <initial-layout>
        <div class="font-semibold select-none">Open workspace</div>
        <v-input
            v-model="repoPath"
            class="mt-4"
            label="git"
            shortcut="return"
            placeholder="..."
        />

        <template #or>
            <v-link class="text-sm" href="new">Create new</v-link>
        </template>
    </initial-layout>
</template>
