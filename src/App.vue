<script setup lang="ts">
import { onBeforeMount } from 'vue';

import { toRefs } from '@vueuse/core';

import { useWorkspace } from './store/workspace';
import NoWorkspaceView from './views/NoWorkspaceView.vue';
import WorkspaceLoadingView from './views/WorkspaceLoadingView.vue';
import WorkspaceView from './views/WorkspaceView.vue';

const workspace = useWorkspace();
const { isLoading, currentWorkspace } = toRefs(workspace);

onBeforeMount(async () => {
    await workspace.load();
});
</script>

<template>
    <workspace-loading-view v-if="isLoading" />
    <template v-else>
        <no-workspace-view v-if="!currentWorkspace" />
        <workspace-view v-else />
    </template>
</template>
