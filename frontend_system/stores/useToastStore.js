import { defineStore } from 'pinia'

export const useToastStore = defineStore('toast', () => {
    const toasts = ref([])
    const addToast = (message, type = 'error') => {
        const id = Date.now()
        toasts.value.push({ id, message, type })
        setTimeout(() => {
            removeToast(id)
        }, 3000)
    }

    const removeToast = (id) => {
        toasts.value = toasts.value.filter(t => t.id !== id)
    }
    return { toasts, addToast, removeToast }
})