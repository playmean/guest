import { Component } from 'vue';

export declare type Shortcuts = 'return' | 'ctrl+return' | 'ctrl+p';

export declare interface ShortcutInfo {
    label: string;
    icon?: Component;
}
