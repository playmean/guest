<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core';

import LayoutBlock from '@/components/Layout/LayoutBlock.vue';
import LayoutResizer from '@/components/Layout/LayoutResizer.vue';

import { useLayoutResize } from './compositions/use-layout-resize';

const { sideSize, startResize } = useLayoutResize();

const localSize = useLocalStorage('workspace-layout-size', sideSize);
</script>

<template>
    <div class="w-full h-[100vh] flex flex-col overflow-hidden font-roboto">
        <layout-block v-if="$slots.header" class="m-2 mb-0">
            <slot name="header" />
        </layout-block>
        <div
            class="grid gap-[0.2rem] mt-2 px-2 flex-grow"
            :style="{
                gridTemplateColumns: `${localSize.left}px 0.1rem auto 0.1rem ${localSize.right}px`,
            }"
        >
            <layout-block>
                <slot name="navigation" />
            </layout-block>
            <layout-resizer @resize-start="startResize('left', $event.pageX)" />
            <layout-block>
                <slot name="content" />
            </layout-block>
            <layout-resizer @resize-start="startResize('right', $event.pageX)" />
            <layout-block>
                <slot name="tools" />
            </layout-block>
        </div>
        <div class="my-2 px-2">
            <layout-block>
                <slot name="footer" />
            </layout-block>
        </div>
    </div>
</template>
