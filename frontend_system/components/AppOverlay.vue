<template>
    <div
        v-if="open"
        @click="$emit('close')"
        class="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
    >
    </div>
</template>



<script setup>
    defineProps({
        open:Boolean
    })

</script>