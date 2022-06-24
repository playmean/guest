<script setup lang="ts">
import { onBeforeMount } from 'vue';

import { toRefs } from '@vueuse/core';

import { useWorkspace } from './store/workspace';
import NoWorkspaceView from './views/NoWorkspaceView.vue';

const workspace = useWorkspace();
const { workspaceLoaded, currentWorkspace } = toRefs(workspace);

onBeforeMount(async () => {
    await workspace.load();
});
</script>

<template>
    <no-workspace-view v-if="workspaceLoaded && !currentWorkspace" />
    <div v-else-if="!workspaceLoaded">loading...</div>
</template>
