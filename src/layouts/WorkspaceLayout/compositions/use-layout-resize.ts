import { ref } from 'vue';

import { useEventListener } from '@vueuse/core';

type SideDirection = 'left' | 'right';
type SizesTable = Record<SideDirection, number>;

const defaultSizes: SizesTable = {
    left: 200,
    right: 200,
};

export function useLayoutResize(
    // TODO 29.07.2022 calculate from window size
    minimalSize = 200,
    maximalSize = 400,
    initialSizes: Record<SideDirection, number> = defaultSizes
) {
    const sideSize = ref<Record<SideDirection, number>>({
        left: initialSizes.left,
        right: initialSizes.right,
    });
    const lastX = ref<Record<SideDirection, number>>({
        left: 0,
        right: 0,
    });

    const isMoving = ref(false);
    const resizeSide = ref<SideDirection>('left');

    useEventListener(
        'mousemove',
        (event) => {
            if (!isMoving.value) return;

            triggerResize(event.pageX);
        },
        { passive: true }
    );

    useEventListener(
        'mouseup',
        () => {
            if (!isMoving.value) return;

            isMoving.value = false;
        },
        { passive: true }
    );

    function triggerResize(pageX: number) {
        if (resizeSide.value === 'left') {
            sideSize.value.left += pageX - lastX.value.left;

            if (sideSize.value.left < minimalSize) {
                sideSize.value.left = minimalSize;

                return;
            }

            if (sideSize.value.left > maximalSize) {
                sideSize.value.left = maximalSize;

                return;
            }
        } else {
            sideSize.value.right += lastX.value.right - pageX;

            if (sideSize.value.right < minimalSize) {
                sideSize.value.right = minimalSize;

                return;
            }

            if (sideSize.value.right > maximalSize) {
                sideSize.value.right = maximalSize;

                return;
            }
        }

        lastX.value[resizeSide.value] = pageX;
    }

    function startResize(side: SideDirection, pageX: number) {
        lastX.value[side] = pageX;
        resizeSide.value = side;
        isMoving.value = true;
    }

    return {
        sideSize,

        startResize,
    };
}
