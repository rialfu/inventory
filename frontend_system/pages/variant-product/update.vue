<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Update Generic Product</h1>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
           
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4"><span @click="goBack" class="cursor-pointer">←</span> Form</h3>
            <div class="relative bg-red-600 text-white p-4 rounded-lg shadow-md max-w-sm" v-show="(messageError['message'] ?? []).length > 0">
                <button class="absolute top-2 right-2 text-white hover:text-gray-200 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50" @click="closeButton">
                    <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                </button>

                <h3 class="font-bold text-lg mb-2">Validation Error!</h3>
                <ul class="list-disc list-inside p-0 m-0 text-sm">
                    <li class="mb-1" v-for="(data, index) in (messageError['message'] ?? []) " :key="index">{{data}}</li>
                    
                </ul>
            </div>
            <div class="relative bg-green-600 text-white p-4 rounded-lg shadow-md max-w-sm" v-show="success">
                <button class="absolute top-2 right-2 text-white hover:text-gray-200 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50" @click="closeButtonSuccess">
                    <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                </button>

                <h3 class="font-bold text-lg mb-2">Notification</h3>
                <ul class="list-disc list-inside p-0 m-0 text-sm">
                    <li class="mb-1" >Success Update</li>
                    
                </ul>
            </div>
            <form class="">
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Name
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['name'] !== undefined && messageError['name'] !== ''? 'border-red-500 mb-3':'')]" id="name" type="text" placeholder="Name" @input="changeValue" v-model="name['name']">
                    <p class="text-red-500 text-xs italic" v-show=" messageError['name'] !== undefined && messageError['name'] !== ''">{{messageError['name']}}</p>
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="descriptive">
                        Descriptive
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['descriptive'] !== undefined && messageError['descriptive'] !== ''? 'border-red-500 mb-3':'')]" id="descriptive" type="text" placeholder="Descriptive" @input="changeValue" v-model="name['descriptive']">
                    
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Merk <input type="checkbox" v-model="name['withoutMerk']" id=""> Without Merk
                    </label>
                    <select v-show="name['withoutMerk'] == false" v-model="name['merk']" id="country" name="country" autocomplete="country-name" class="mb-3 col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pr-8 pl-3 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6">
                        <option v-for="(data, index) in optionsMerk" :key="index" :value="data['id']">{{data['name']}}</option>
                        
                    </select>
                    <!-- <a href="#"  @click="refreshValue"  class="px-4 py-2 rounded bg-green-400 text-white text-sm">Refresh Category</a> -->
                    
                    <p class="text-red-500 text-xs italic" v-show=" messageError['merk'] !== undefined && messageError['merk'] !== ''">{{messageError['merk']}}</p>
                </div>
                <!-- start -->
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Category
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
                                    {{ item['full_path'] }}
                                </li>
                            </ul>
                        </div>
                        <div v-else-if="searchConfig['open'] && searchConfig['data'].length == 0" class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 p-3 text-gray-500">
                            No results found.
                        </div>
                    </div>
                    <p>{{searchConfig['id'] != null? 'Choosen':'Not Choose'}}</p>
                    <p class="text-red-500 text-xs italic" v-show=" messageError['category'] !== undefined && messageError['category'] !== ''">{{messageError['category']}}</p>
                </div>
                
                <!-- finish -->
                <div class="flex items-center justify-between">
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" @click="save"  type="button">
                        Save
                    </button>
                
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { useRoute } from "vue-router";


    const { $api } = useNuxtApp();
    
    const name = ref({'name':'', 'descriptive':'', 'withoutMerk':false,'category':null, 'merk':null,})
    const optionsMerk = ref([])
    const success = ref(false)
    const router = useRouter();
    const messageError = ref({'name':'','merk':'', 'category':'', 'message':[]})
    const route = useRoute()
    const {id, name:nameParam } = route.query

    definePageMeta({
        layout: 'base'
    });
    const searchConfig  = ref({
        open:false,
        id:null,
        search:'',
        data:[],
        load:false,

    })
    async function filterItems(){
        if(searchConfig.value['load']){
            return
        }
        searchConfig.value['load'] = true
        name.value['category'] = null;
        let search = searchConfig.value['search']
        searchConfig.value['id'] = null;
        try{
            const queryParams = new URLSearchParams({search });
            const data = await $api('category/search?'+queryParams,{
                method:'GET'
            });
            if(data['data'] != undefined){
                console.log(data['data'])
                searchConfig.value['data'] = data['data']
            }else{
                searchConfig.value['data'] = []
            }
            searchConfig.value['open'] =true
            // console.log(data)
        }catch(err){
            console.log(err)
        }
        searchConfig.value['load'] = false
    }
    function selectItem(value){
        searchConfig.value['id'] = value['id']
        searchConfig.value['open'] = false
        searchConfig.value['search'] = value['full_path']
        name.value['category'] = value['id'];
    }
    function handleBlur(){
        setTimeout(() => {
            searchConfig.value['open'] = false
        }, 100);
    }
    function changeValue(value){
        messageError.value = {
            'name':'',
            'merk':'',
            'category':'',
            'message':messageError.value['message']
        }
    }
    function goBack(){
        router.back();
    }

    function closeButtonSuccess(){
        success.value = false;
    }
    function closeButton(){
        messageError.value = {
            'name':messageError.value['name'],
            'merk':messageError.value['merk'],
            'category':messageError.value['category'],
            'message':[]
        }
    }
    async function save() {
        messageError.value = {
            'name':'',
            'merk':'',
            'category':'',
            'message':messageError.value['message']
        }
        if(name.value['name'] == ''){
            messageError.value = {
                'name':'Please fill name',
                'merk':messageError.value['merk'],
                'category':messageError.value['category'],
                'message':messageError.value['message']
            }
            
        }
        if(name.value['merk'] == null && name.value['withoutMerk'] == false){
            messageError.value = {
                'name':messageError.value['name'],
                'merk':'Please choose merk or set without merk',
                'category':messageError.value['category'],
                'message':messageError.value['message']
            }
        }
        if(name.value['category'] == null){
            messageError.value = {
                'name':messageError.value['name'],
                'merk':messageError.value['merk'],
                'category':'Please choose',
                'message':messageError.value['message']
            }
        }
        console.log(name.value['category'])
        if(name.value['name'] == '' || (name.value['merk'] == null && name.value['withoutMerk'] == false) ) return;
        try {
            
            const data1 = await $api('generic-product',{
                method:'PUT',
                
                body:{ 
                    id: id,
                    name: name.value['name'],
                    descriptive: name.value['descriptive'],
                    category: name.value['category'],
                    merk: name.value['withoutMerk'] ? null : name.value['merk']
                }
            });
            messageError.value = {
                'name':'',
                'message':[]
            }
            success.value = true
        } catch (error) {
             if(error.data['message'] !== undefined){
                messageError.value = {
                    'name':'',
                    'message':error.data['message']
                }
            
            }else{
                messageError.value = {
                    'name':'',
                    'message':['Something wrong from server']
                }
            }
            
            success.value = false;
        }
    }
    async function fetchDataMerk(){
        try{
            const data = await $api('generic-product/merk',{
                method:'GET',
            });
            if(data['data'] != undefined){
                console.log(data['data'])
                optionsMerk.value = data['data']
            }
            console.log(data)
        }catch(err){
            console.log(err)
        }
    }
    async function fetchData(){
        try{
            console.log(id)
            const data = await $api(`generic-product/${id ?? 'lol'}`,{
                method:'GET',
            });
            if(data['data'] != undefined){
                name.value ={
                    'name':data['data']['name'], 
                    'descriptive':data['data']['descriptive'], 
                    'withoutMerk': data['data']['merk'] == null, 
                    'category':data['data']['category']['id'], 
                    'merk':data['data']['merk']== null ? null:data['data']['merk']['id'],
                }
            }
            console.log(data)
        }catch(err){
            console.log(err)
        }
    }
    // async 
    onMounted(() => {
        fetchDataMerk()
        fetchData()
        name.value['name']
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>

</style>