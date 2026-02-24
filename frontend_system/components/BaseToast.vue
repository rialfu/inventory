<!-- <template>

    <div
        v-if="show"
        class="fixed top-5 right-5 z-50"
    >

        <div
            class="px-4 py-3 rounded shadow-lg text-white"
            :class="colorClass"
        >

            {{ message }}

        </div>

    </div>

</template> -->
<template>
    <div class="fixed top-5 right-5 z-50 flex flex-col gap-3">
        <TransitionGroup name="list">
            <div
                v-for="toast in toasts"
                :key="toast.id"
                class="px-4 py-3 rounded shadow-lg text-white transition-all duration-300"
                :class="getColorClass(toast.type)"
            >
                {{ toast.message }}
            </div>
        </TransitionGroup>
    </div>
</template>

<script setup>
    import { useToastStore } from '~/stores/useToastStore'

    const toastStore = useToastStore()
    const props = defineProps({

        // message:String,

        // type:{
        //     type:String,
        //     default:"error"
        // },

        // show:Boolean
        toasts: {
            type: Array,
            default: () => []
        }
    })


const getColorClass = (type) => {
    if (type === "success") return "bg-green-500"
    if (type === "warning") return "bg-yellow-500"
    return "bg-red-500"
}
</script>

<style scoped>
    /* Animasi List */
    .list-enter-from { opacity: 0; transform: translateX(30px); }
    .list-leave-to { opacity: 0; transform: scale(0.9); }

</style>