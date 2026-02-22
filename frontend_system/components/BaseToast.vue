<template>

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

</template>


<script setup>

const props = defineProps({

    message:String,

    type:{
        type:String,
        default:"error"
    },

    show:Boolean

})


const colorClass = computed(()=>{

    if(props.type === "success")
        return "bg-green-500"

    if(props.type === "warning")
        return "bg-yellow-500"

    return "bg-red-500"

})

</script>