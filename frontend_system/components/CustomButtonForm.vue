<template>
    <div class="flex items-center justify-between">
        
        <button 
            type="button"
            :disabled="isLoad"
            :class="[
                'font-bold py-2 px-4 rounded focus:outline-none transition-colors duration-200 cursor-pointer',
                buttonClass
                // isLoad
                // ? 'bg-blue-100 text-blue-400 cursor-not-allowed' 
                // : 'bg-blue-500 hover:bg-blue-700 text-white shadow-md'
            ]"
            @click="save"
        >
            <div class="flex items-center justify-center gap-2">
                
                <svg 
                    v-if="isLoad" 
                    class="animate-spin h-4 w-4 text-blue-400" 
                    xmlns="http://www.w3.org/2000/svg" 
                    fill="none" 
                    viewBox="0 0 24 24"
                    >
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
            
                <span v-if="isLoad">Load / Process ...</span>
                <span v-else>{{ buttonName }}</span>
            </div>
        </button>
        
    </div>
</template>

<script setup>
    const props = defineProps({
        isLoad: { type: Boolean, default: false },
        buttonName: {type:String, default:'Save'},
        colorButton: {type:String, default:'primary'}
    })
    const emit = defineEmits(['save'])
    function save(){
        emit('save', '')
    }
    const buttonClass = computed(() => {
        if(props.colorButton == 'success' && props.isLoad){
            return 'bg-green-100 text-green-400 cursor-not-allowed'
        }else if(props.colorButton == 'success'){
            return 'bg-green-500 hover:bg-green-700 text-white shadow-md'
        }else if(props.colorButton == 'danger' && props.isLoad){
            return 'bg-red-100 text-red-400 cursor-not-allowed'
        }else if(props.colorButton == 'danger'){
            return 'bg-red-500 hover:bg-red-700 text-white shadow-md'
        }else if (props.isLoad) {
            return 'bg-blue-100 text-blue-400 cursor-not-allowed'
        }
         
        return 'bg-blue-500 hover:bg-blue-700 text-white shadow-md'

    })
</script>