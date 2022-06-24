import { defineStore } from 'pinia';
import { ref } from 'vue';

import { ApiRequest } from '@/request';
import { ApiGetWorkspaceResponse } from '@/types/api';

const currentWorkspace = ref<ApiGetWorkspaceResponse>();
const workspaceLoaded = ref(false);

export const useWorkspace = defineStore('workspace', () => {
    async function load() {
        try {
            currentWorkspace.value = await ApiRequest.Get<ApiGetWorkspaceResponse>(
                '/workspace'
            );
        } catch (err) {
            console.error(err);
        } finally {
            workspaceLoaded.value = true;
        }
    }

    return {
        currentWorkspace,
        workspaceLoaded,
        load,
    };
});
