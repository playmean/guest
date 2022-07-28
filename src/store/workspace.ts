import { defineStore } from 'pinia';
import { ref } from 'vue';

import { ApiRequest } from '@/request';
import { ApiGetWorkspaceResponse } from '@/types/api';

const currentWorkspace = ref<ApiGetWorkspaceResponse>();
const isLoading = ref(false);

export const useWorkspace = defineStore('workspace', () => {
    async function load() {
        try {
            isLoading.value = true;

            currentWorkspace.value = await ApiRequest.Get<ApiGetWorkspaceResponse>(
                '/workspace'
            );
        } catch (err) {
            console.error(err);
        } finally {
            isLoading.value = false;
        }
    }

    async function open(path: string) {
        try {
            isLoading.value = true;

            currentWorkspace.value = await ApiRequest.Put<ApiGetWorkspaceResponse>(
                '/workspace',
                {
                    path,
                }
            );
        } catch (err) {
            console.error(err);
        } finally {
            isLoading.value = false;
        }
    }

    async function create() {
        try {
            isLoading.value = true;

            currentWorkspace.value = await ApiRequest.Post<ApiGetWorkspaceResponse>(
                '/workspace'
            );
        } catch (err) {
            console.error(err);
        } finally {
            isLoading.value = false;
        }
    }

    return {
        currentWorkspace,
        isLoading,
        load,
        open,
        create,
    };
});
