import { ShortcutInfo, Shortcuts } from './types';

import IconOutlineKeyboardReturn from '~icons/ic/outline-keyboard-return';

export const shortcuts: Record<Shortcuts, ShortcutInfo> = {
    return: {
        label: 'Return',
        icon: IconOutlineKeyboardReturn,
    },
};
