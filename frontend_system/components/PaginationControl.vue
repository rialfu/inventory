<script setup>
const props = defineProps({
    modelValue: { type: Number, required: true }, // Halaman aktif
    totalItems: { type: Number, required: true }, // Total data dari DB
    perPage:    { type: Number, default: 10 },    // Limit data per hal
    isLoading:  { type: Boolean, default: false },
    isEmpty:    { type: Boolean, default: false }  // Apakah data di hal ini kosong?
})
onMounted(() => {
  console.log("PaginationControl berhasil dipanggil!");
  console.log("Total data yang diterima:", props.totalItems);
});
const emit = defineEmits(['update:modelValue'])

// Hitung total halaman
const totalPages = computed(() => Math.ceil(props.totalItems / props.perPage))

// Logika Munculkan Pesan Error (Jika user di page > 1 tapi data kosong)
const showEmptyWarning = computed(() => {
    return props.isEmpty && props.modelValue > 1 && !props.isLoading
})

const setPage = (p) => {
    if (p >= 1 && p <= totalPages.value) {
        emit('update:modelValue', p)
    }
}
</script>

<template>
    <!-- <div>Kontol</div> -->
    <div class="mt-6 space-y-4">
        <div v-if="showEmptyWarning" class="flex flex-col items-center p-6 border-2 border-dashed border-amber-300 bg-amber-50 rounded-xl text-center">
            <UIcon name="i-heroicons-exclamation-triangle" class="w-10 h-10 text-amber-500 mb-2" />
            <p class="text-amber-800 font-medium">Halaman {{ modelValue }} tidak memiliki data.</p>
            <div class="mt-3 flex gap-2">
                <UButton color="amber" variant="soft" size="sm" @click="setPage(modelValue - 1)">
                Ke Halaman Sebelumnya
                </UButton>
                <UButton color="amber" variant="solid" size="sm" @click="setPage(1)">
                Kembali ke Awal
                </UButton>
            </div>
        </div>

        <div v-else-if="totalItems > 0" 
            class="flex items-center justify-between bg-white px-4 py-3 border border-gray-200 rounded-lg shadow-sm">
            
            <div class="hidden sm:block text-sm text-gray-600">
                Total <span class="font-bold text-gray-900">{{ totalItems }}</span> data 
                <span class="mx-1 text-gray-300">|</span> 
                Hal <span class="font-bold text-primary-600">{{ modelValue }}</span> dari {{ totalPages }}
            </div>

            <div class="flex items-center gap-1">
                <UButton
                icon="i-heroicons-chevron-left"
                color="gray"
                variant="ghost"
                :disabled="modelValue === 1 || isLoading"
                @click="setPage(modelValue - 1)"
                />

                <div class="flex items-center gap-1">
                <UButton
                    v-for="p in totalPages"
                    :key="p"
                    size="sm"
                    :variant="p === modelValue ? 'solid' : 'ghost'"
                    :color="p === modelValue ? 'primary' : 'gray'"
                    @click="setPage(p)"
                    :disabled="isLoading"
                >
                    {{ p }}
                </UButton>
                </div>

                <UButton
                icon="i-heroicons-chevron-right"
                color="gray"
                variant="ghost"
                :disabled="modelValue === totalPages || isLoading"
                @click="setPage(modelValue + 1)"
                />
            </div>
        </div>
    </div>
</template>