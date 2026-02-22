<template>
    <label 
      v-if="label" 
      :for="inputId" 
      class="block text-gray-700 text-sm font-bold mb-2"
    >
      {{ label }}
    </label>
    
    <input
        :id="inputId"
        :type="type"
        :value="modelValue"
        @input="handleInput"
        :readonly="readOnly"
        :placeholder="placeholder"
        :class="[
            'border rounded w-full py-2 px-3 leading-tight focus:outline-none transition-all',
            // 'border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-1 focus:ring-blue-500 transition-all',
            error ? 'border-red-500' : 'border-gray-300',
            readOnly
                ? 'bg-gray-100 text-gray-500 cursor-not-allowed pointer-events-none'
                : 'bg-white text-gray-700 focus:ring-1 focus:ring-blue-500'

        ]"
    />
    
    <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="transform -translate-y-2 opacity-0"
        enter-to-class="transform translate-y-0 opacity-100"
    >
        <p 
            v-if="error" 
            class="text-red-500 text-xs italic mt-1"
        >
            {{ error }}
        </p>
    </Transition>
    
</template>

<script setup>
const props = defineProps({
    modelValue: [String, Number],
    label: { type: String, default: 'Form Input' },
    // Kita beri nilai default null, lalu kita fallback ke useId()
    id: { type: String, default: null }, 
    type: { type: String, default: 'text' },
    placeholder: String,
    error: String,
    readOnly: { type:Boolean, default:false }
})

const emit = defineEmits(['update:modelValue', 'change'])

// Otomatis generate ID unik jika prop 'id' tidak diisi
const inputId = props.id || useId() 

const handleInput = (event) => {
    emit('update:modelValue', event.target.value)
}

</script>

