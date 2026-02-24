<script setup>
    const props = defineProps({

        modelValue: {
            type: Array,
            default: () => []
        },

        placeholder: {
            type: String,
            default: 'Cari...'
        },

        url: {
            type: String,
            required: true
        },

        idValue: {
            type: String,
            default: 'id'
        },

        text: {
            type: String,
            default: 'name'
        },

        label: String,

        error: String,

        formId: String

    })

    const emit = defineEmits(['update:modelValue'])
    const { $api } = useNuxtApp()

    const search = ref('')
    const results = ref([])
    const selectedItems = ref([])
    const isOpen = ref(false)
    const pending = ref(false)
    const page = ref(1)
    const limit = 10
    const hasMore = ref(true)
    const dropdownRef = ref(null)

    watch(() => props.modelValue, (val) => {

        selectedItems.value = val || []

    }, { immediate: true })


    async function fetchResults(loadMore = false){
        if(!loadMore){
            page.value = 1
            results.value = []
            hasMore.value = true
        }
        if(!hasMore.value) return
        pending.value = true
        try{

            const res = await $api(props.url, {
                method:'GET',
                query:{
                    search: search.value,
                    page: page.value,
                    limit
                }
            })
            const data = res.data || []
            if(loadMore){
                results.value.push(...data)
            }else{
                results.value = data
            }

            if(data.length < limit){
                hasMore.value = false
            }

        }
        finally{
            pending.value = false

        }

    }
    // LOAD MORE
    function loadMore(){
        page.value++
        fetchResults(true)
    }
    // DEBOUNCE
    function debounce(fn, delay){
        let t
        return (...args)=>{
            clearTimeout(t)
            t = setTimeout(()=> fn(...args), delay)
        }
    }

    const debouncedSearch = debounce(fetchResults, 500)

    watch(search, ()=>{
        if(isOpen.value){
            debouncedSearch()
        }
    })


    // SELECT

    function selectItem(item){
        const id = item[props.idValue]
        const exists = selectedItems.value.find(
            v => v[props.idValue] === id
        )
        if(exists){
            selectedItems.value =
                selectedItems.value.filter(
                    v => v[props.idValue] !== id
                )
        }else{
            selectedItems.value.push(item)
        }
        emit('update:modelValue', selectedItems.value)
    }


    // REMOVE TAG

    function removeTag(id){

        selectedItems.value =
            selectedItems.value.filter(
                v => v[props.idValue] !== id
            )
        emit('update:modelValue', selectedItems.value)
    }


    // CHECK SELECTED

    function isSelected(id){
        return selectedItems.value.some(
            v => v[props.idValue] === id
        )
    }

    // FOCUS
    function onFocus(){
        isOpen.value = true
        fetchResults()
    }
    // CLICK OUTSIDE
    if(process.client){
        window.addEventListener('click', e=>{
            if(
                dropdownRef.value &&
                !dropdownRef.value.contains(e.target)
            ){
                isOpen.value = false
            }

        })

    }
</script>


<template>

    <div
        ref="dropdownRef"
        class="relative w-full"
    >
        <div class="flex">
            <label
                v-if="label"
                class="block text-gray-700 text-sm font-bold mb-2 me-2"
            >
                {{ label }}
            </label>
             <!-- TAG -->
            <div class="flex flex-wrap gap-1">
                <span
                    v-for="item in selectedItems"
                    :key="item[idValue]"
                    class="bg-blue-100 text-blue-700 text-xs px-2  rounded flex items-center"
                >

                    {{ item[text] }}
                    <button
                        type="button"
                        @click="removeTag(item[idValue])"
                        class="ml-1 text-red-500"
                    >
                        ×
                    </button>
                </span>
            </div>
        </div>
        


       

        <!-- INPUT -->
        <div class="relative">
            <input
                v-model="search"
                @focus="onFocus"
                :placeholder="placeholder"

                class="border rounded w-full py-2 px-3 text-sm focus:outline-none focus:ring-1 focus:ring-blue-500"
            />
            <div
                v-if="pending"
                class="absolute right-2 top-2"
            >
                <div class="w-4 h-4 border-2 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
        </div>

        <!-- DROPDOWN -->
        <ul
            v-if="isOpen"
            class="absolute z-50 w-full bg-white border rounded shadow max-h-60 overflow-auto mt-1"
        >
            <li
                v-for="item in results"
                :key="item[idValue]"
                @click="selectItem(item)"

                :class="[
                    'px-3 py-2 cursor-pointer text-sm',
                    isSelected(item[idValue])
                        ? 'bg-blue-100 text-blue-700'
                        : 'hover:bg-blue-50'

                ]"
            >{{ item[text] }}
            </li>

            <li
                v-if="hasMore"
                class="text-center p-2"
            >
                <button
                    type="button"
                    @click="loadMore"
                    class="text-blue-500 text-xs"
                >
                    Load more
                </button>

            </li>
        </ul>
        <p
            v-if="error"
            class="text-red-500 text-xs mt-1"
        >
            {{ error }}
        </p>
    </div>
</template>