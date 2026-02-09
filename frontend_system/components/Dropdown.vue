<script setup>
const props = defineProps(['modelValue', 'placeholder', "url", "idValue", "text"])
const emit = defineEmits(['update:modelValue'])

const { $api } = useNuxtApp();

const search = ref('')
const isOpen = ref(false)
const results = ref([])
const pending = ref(false)
const limit = 10

// State untuk Pagination
const page = ref(1)
const hasMore = ref(true) // Untuk cek apakah masih ada data selanjutnya

const fetchResults = async (isLoadMore = false) => {
    if (!search.value) {
        results.value = []
        return
    }

    // Jika bukan load more, reset ke page 1
    if (!isLoadMore) {
        page.value = 1
        results.value = []
        hasMore.value = true
    }

    pending.value = true
    try {
        const queryParams = {
            search: search.value,
            page: page.value,
        }
        
        const res = await $api('category/dropdown', {
            method: 'GET',
            query: queryParams,
        });

        // console.log(res['data'])
        const newData = res.data || [] // Sesuaikan struktur response Go kamu
        
        if (isLoadMore) {
            results.value = [...results.value, ...newData]
        } else {
            results.value = newData
        }

        // Jika data yang datang kurang dari limit, berarti sudah habis
        if (newData.length < limit) {
            hasMore.value = false
        }
    } catch (error) {
        console.error(error)
        results.value = []
    } finally {
        pending.value = false
    }
}

// Handler untuk Load More
const loadMore = () => {
    if (!pending.value && hasMore.value) {
        page.value++
        fetchResults(true)
    }
}

import { debounce } from 'lodash-es'
const debouncedSearch = debounce(() => fetchResults(false), 500)

watch(search, (newVal) => {
    if (isOpen.value) {
        debouncedSearch()
    }
})

const selectItem = (item) => {
    isOpen.value = false 
    search.value = item.name
    emit('update:modelValue', item.id)
}
</script>

<template>
    <div class="relative w-full max-w-xs">
        <div class="relative">
            <input
                v-model="search"
                @focus="isOpen = true"
                type="text"
                :placeholder="placeholder || 'Cari...'"
                class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 outline-none transition"
            />
            
            <div v-if="pending" class="absolute right-3 top-2.5">
                <div class="w-5 h-5 border-2 border-green-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
        </div>

        <ul v-if="isOpen && search.length > 0" 
            class="absolute z-50 w-full mt-2 bg-white border rounded-lg shadow-xl max-h-60 overflow-y-auto">
            
            <li v-if="results?.length === 0 && !pending" class="p-3 text-gray-500 text-sm">
                Data tidak ditemukan...
            </li>

            <li 
                v-for="item in results" 
                :key="item.id"
                @click="selectItem(item)"
                class="px-4 py-2 hover:bg-green-50 cursor-pointer border-b last:border-0 text-sm"
            >
                {{ item.name }}
            </li>

            <li v-if="hasMore && results.length > 0" class="p-2 text-center">
                <button 
                    @click.stop="loadMore"
                    :disabled="pending"
                    class="text-xs text-green-600 hover:text-green-800 font-semibold w-full py-1"
                >
                    {{ pending ? 'Loading...' : 'Muat Lebih Banyak' }}
                </button>
            </li>
        </ul>
    </div>
</template>