<script setup>
import { debounce } from 'lodash-es'

const props = defineProps({
    modelValue: [String, Number],
    placeholder: { type: String, default: 'Cari...' },
    url: { type: String, required: true },
    // Mapping key dari API (misal: 'id', 'category_id', dll)
    idValue: { type: String, default: 'id' },
    // Mapping display text dari API (misal: 'name', 'title', dll)
    text: { type: String, default: 'name' },
    // Nama awal yang muncul di input (untuk Edit Mode)
    defaultName: { type: String, default: '' },
    error: String,
    label: String,
    inputId: {type:String, default:useId()},
})

const emit = defineEmits(['update:modelValue'])
const { $api } = useNuxtApp()

// --- STATES ---
const search = ref(props.defaultName)
const isOpen = ref(false)
const results = ref([])
const pending = ref(false)
const limit = 10
const page = ref(1)
const hasMore = ref(true)

// --- LOGIC FETCH ---
const fetchResults = async (isLoadMore = false) => {
    if (!search.value && !isLoadMore) {
        results.value = []
        return
    }

    if (!isLoadMore) {
        page.value = 1
        results.value = []
        hasMore.value = true
    }

    pending.value = true
    try {
        console.log("search value", search.value)
        const res = await $api(props.url, {
        method: 'GET',
        query: {
            search: search.value,
            page: page.value,
            limit: limit
        },
        })

        const newData = res.data || []
        
        if (isLoadMore) {
        results.value = [...results.value, ...newData]
        } else {
        results.value = newData
        }

        if (newData.length < limit) {
        hasMore.value = false
        }
    } catch (error) {
        console.error('Dropdown Error:', error)
        results.value = []
    } finally {
        pending.value = false
    }
}

// --- HANDLERS ---
const loadMore = () => {
    if (!pending.value && hasMore.value) {
        page.value++
        fetchResults(true)
    }
}

const debouncedSearch = debounce(() => fetchResults(false), 500)

watch(search, (newVal) => {
  // Jika user menghapus teks sampai kosong, kosongkan juga ID di parent
    if (newVal === '') {
        emit('update:modelValue', null)
        results.value = []
    }
    
    if (isOpen.value) {
        debouncedSearch()
    }
})

// Sinkronisasi jika defaultName berubah dari parent
watch(() => props.defaultName, (newVal) => {
    search.value = newVal
})

const selectItem = (item) => {
    isOpen.value = false 
    search.value = item[props.text] 
    console.log(item[props.idValue])
    emit('update:modelValue', item[props.idValue]) 
}
function onFocus(){
    isOpen.value=true
    debouncedSearch()
}
// Close dropdown saat klik di luar (Optional Logic)
const dropdownRef = ref(null)
if (process.client) {
    window.addEventListener('click', (e) => {
        if (dropdownRef.value && !dropdownRef.value.contains(e.target)) {
            isOpen.value = false
        }
    })
}
</script>

<template>
    <div class="relative w-full" ref="dropdownRef">
        <label 
            v-if="label" 
            :for="inputId" 
            class="block text-gray-700 text-sm font-bold mb-2"
        >
        {{ label }}
        </label>
        <div class="relative">
            <input
                v-model="search"
                @focus="onFocus"
                type="text"
                :id="inputId"
                :placeholder="placeholder"
                :class="[
                'border rounded w-full py-2 px-3 text-sm text-gray-700 leading-tight focus:outline-none focus:ring-1 focus:ring-blue-500 transition-all',
                error ? 'border-red-500' : 'border-gray-300'
                ]"
            />
        
            <div v-if="pending" class="absolute right-3 top-2.5">
                <div class="w-4 h-4 border-2 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
        </div>

        <ul v-if="isOpen && (results.length > 0 || (search.length > 0 && !pending))" 
            class="absolute z-50 w-full mt-1 bg-white border border-gray-200 rounded-md shadow-lg max-h-60 overflow-y-auto">
        
            <li v-if="results.length === 0 && !pending" class="p-3 text-gray-500 text-xs italic">
                Data tidak ditemukan...
            </li>

            <li 
                v-for="item in results" 
                :key="item[props.idValue]"
                @click="selectItem(item)"
                class="px-4 py-2 hover:bg-blue-50 cursor-pointer border-b border-gray-50 last:border-0 text-sm text-gray-700"
            >
                {{ item[props.text] }}
            </li>

            <li v-if="hasMore && results.length > 0" class="p-2 text-center bg-gray-50">
                <button 
                @click.stop="loadMore"
                :disabled="pending"
                type="button"
                class="text-xs text-blue-600 hover:text-blue-800 font-semibold w-full py-1"
                >
                {{ pending ? 'Loading...' : 'Muat Lebih Banyak' }}
                </button>
            </li>
        </ul>

        <p v-if="error" class="text-red-500 text-xs italic mt-1">{{ error }}</p>
    </div>
</template>