<template>
    <div>
    <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
            {{ nameField }}
        </label>
        <div class="relative">
            <input
                type="text"
                placeholder="Search fruits..."
                class="w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                v-model="searchConfig['search']"
                @focus="dropdownOpen = true"
                @blur="handleBlur"
                @input="filterItems"
            >

            <div v-if="searchConfig['open'] && searchConfig['load']" class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 p-3 text-gray-500">
                Please Wait, still loading
            </div>
            <div
                v-else-if="searchConfig['open'] && searchConfig['data'].length >0 "
                class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 p-3 text-gray-500 max-h-60 overflow-y-auto"
            >
                <ul>
                    <li
                        v-for="item in searchConfig['data']"
                        :key="item"
                        @mousedown.prevent="selectItem(item)"
                        class="px-3 py-2 cursor-pointer hover:bg-blue-100 text-gray-800"
                    >
                        {{ item['name'] }}
                    </li>
                </ul>
            </div>
            <div v-else-if="searchConfig['open'] && searchConfig['data'].length == 0" class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 p-3 text-gray-500">
                No results found.
            </div>
        </div>
        <p>{{searchConfig['id'] != null? 'Choosen':'Not Choose'}}</p>
        <p class="text-red-500 text-xs italic" v-show=" searchConfig['error'] !== ''">{{searchConfig['error']}}</p>
    </div>
    </div>
</template>

<script setup>
    const props = defineProps({
        nameField: {
            type: String,
            required: true,
        },
        listData:{
            type: Array,
            required: true,
        },
        parentValue:{
            type: Object,
            required: true,
        },
        load:{
            type: Boolean,
            required: true,
        },
        // message:{
        //     type: String,
        //     required: true,
        // }

    });
    const {nameField,  } = props
    const searchConfig  = ref({
        open:false,
        id:null,
        search:'',
        cache:'',
        data:[],
        load:false,
        error:'',
        lastValue:{},

    })
    const emit = defineEmits(['findValue', 'setValue']);
    async function filterItems(){
        searchConfig.value['open'] = true
        if(searchConfig.value['load']){
            return
        }
        
        searchConfig.value['load'] = true
        emit('findValue', searchConfig.value['search'])
    }
    watch(() => props, (newVal, oldVal) => {
        console.log('jalan all')
        console.log(newVal)
    },{deep:true});
    function selectItem(value){
        searchConfig.value['id'] = value['id']
        searchConfig.value['open'] = false
        searchConfig.value['search'] = value['name']
        emit('setValue', {'id':value['id'], 'name':value['name']})
    }
    function handleBlur(){
        setTimeout(() => {
            try{
                searchConfig.value['open'] = false
                searchConfig.value['load'] = false;
            if(searchConfig['value']['search'] == ''){
                searchConfig.value['id'] = null
            }
            else if(searchConfig['value']['lastValue']['id'] !== undefined && searchConfig['value']['lastValue']['id'] !== null){
                searchConfig.value['search'] = searchConfig.value.lastValue['name']
                searchConfig.value['id'] = searchConfig.value.lastValue['id']
            }
            }catch(err){
                console.log(err)
            }
        }, 100);
    }
    watch(() => props.listData, (newVal, oldVal) => {
        if(searchConfig.value['load']){
            searchConfig.value['load'] = false;
        }
        searchConfig['value']['data'] = newVal
    },{deep:true});

    watch(() => props.load, (newVal, oldVal) => {
        searchConfig.value['load'] = newVal;
    },{deep:true});

    // watch(() => props.message, (newVal, oldVal) => {
    //     searchConfig.value['load'] = newVal;
    // },);

    watch(() => props.parentValue, (newVal, oldVal) => {
        searchConfig.value['lastValue'] = newVal
    },{deep:true});

</script>

<style  scoped>

</style>