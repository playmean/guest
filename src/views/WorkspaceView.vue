<script setup lang="ts">
import { ref } from 'vue';

import VButton from '@/components/VButton.vue';
import VInput from '@/components/VInput.vue';
import VSelect from '@/components/VSelect.vue';

import WorkspaceLayout from '@/layouts/WorkspaceLayout/WorkspaceLayout.vue';

import IconOutlineLink from '~icons/ic/outline-link';
import IconOutlineSearch from '~icons/ic/outline-search';
import IconOutlineWorkspaces from '~icons/ic/outline-workspaces';
import IconSharpRocketLaunch from '~icons/ic/sharp-rocket-launch';

const requestMethods: [string, string][] = [
    ['GET', 'get'],
    ['POST', 'post'],
    ['PUT', 'put'],
    ['DELETE', 'delete'],
];

const searchQuery = ref('');

const requestMethod = ref('get');
const requestUrl = ref('');

// TODO 29.07.2022 divide into hand-specific views
</script>

<template>
    <workspace-layout>
        <template #header>
            <div class="flex justify-between items-center">
                <div class="flex items-center gap-1 text-xl">
                    <icon-outline-workspaces />
                    Workspace
                </div>
                <v-button>create</v-button>
            </div>
        </template>
        <template #navigation>
            <div>
                <v-input
                    v-model="searchQuery"
                    class="w-full"
                    :label="IconOutlineSearch"
                    placeholder="search..."
                    shortcut="ctrl+p"
                />
            </div>
        </template>
        <template #content>
            <div class="grid grid-cols-[7rem_auto_8rem] gap-2">
                <v-select v-model="requestMethod" :entries="requestMethods" />
                <v-input
                    v-model="requestUrl"
                    :label="IconOutlineLink"
                    placeholder="https://..."
                    shortcut="ctrl+return"
                />
                <v-button :icon="IconSharpRocketLaunch">Send</v-button>
            </div>
        </template>
        <template #tools></template>
    </workspace-layout>
</template>
