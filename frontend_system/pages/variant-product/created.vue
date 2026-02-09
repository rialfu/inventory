<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Create Generic Product</h1>
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
                    <li class="mb-1" >Success Create</li>
                    
                </ul>
            </div>
            <form class="">
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="sku">
                        SKU
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['sku'] !== undefined && messageError['sku'] !== ''? 'border-red-500 mb-3':'')]" id="sku" type="text" placeholder="SKU" @input="changeValue" v-model="form['sku']">
                    <p class="text-red-500 text-xs italic" v-show=" messageError['sku'] !== undefined && messageError['sku'] !== ''">{{messageError['sku']}}</p>
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Name
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['name'] !== undefined && messageError['name'] !== ''? 'border-red-500 mb-3':'')]" id="name" type="text" placeholder="Name" @input="changeValue" v-model="form['name']">
                    <p class="text-red-500 text-xs italic" v-show=" messageError['name'] !== undefined && messageError['name'] !== ''">{{messageError['name']}}</p>
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="unit">
                        Unit
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['unit'] !== undefined && messageError['unit'] !== ''? 'border-red-500 mb-3':'')]" id="unit" type="text" placeholder="Unit" @input="changeValue" v-model="form['unit']">
                    <p class="text-red-500 text-xs italic" v-show=" messageError['unit'] !== undefined && messageError['unit'] !== ''">{{messageError['unit']}}</p>
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Merk <input type="checkbox" v-model="form['withoutMerk']" id=""> Without Merk
                    </label>
                    <select v-show="form['withoutMerk'] == false" v-model="form['merk']" id="country" name="country" autocomplete="country-name" class="mb-3 col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pr-8 pl-3 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6">
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
                <DropdownCustom  @findValue="search_value" @setValue="setValue" nameField="Category" :listData="valueDropdownCategory['data']" :load="valueDropdownCategory['load']" :parentValue="form['category2']"/>
                
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
    import DropdownCustom from '~/components/dynamic_search_dropdown.vue'
    const { $api } = useNuxtApp();
    
    const form = ref({'sku':'', 'name':'','unit':'','withoutMerk':false,'category':null, 'naming_category':'', 'merk':null,  'category2':{'id':null, 'name':''}})
    const optionsMerk = ref([])
    const success = ref(false)
    const router = useRouter();
    const messageError = ref({'sku':'', 'name':'','unit':'', 'category':'', 'message':[]})
    definePageMeta({
        layout: 'base'
    });
    const searchConfig  = ref({
        open:false,
        id:null,
        search:'',
        cache:'',
        data:[],
        load:false,

    })
    const valueDropdownCategory = ref({
        data:[],
        field:'Category',
        load:false,
    })
    
    async function search_value(search){
        valueDropdownCategory.value.load = true
        try{
            const queryParams = new URLSearchParams({search });
            const data = await $api('category/search?'+queryParams,{
                method:'GET'
            });
            if(data['data'] != undefined){
                console.log(data['data'].map((e)=>({id:e['id'], name:e['full_path']})))
                valueDropdownCategory.value.data = data['data'].map((e)=>({id:e['id'], name:e['full_path']}))
            }else{
                valueDropdownCategory.value.data = []
            }
        }catch(err){
            valueDropdownCategory.value.data = []
            console.log(err)
        }
        valueDropdownCategory.value.load = false
    }
    function setValue(value){
        console.log('select parent',value)
        form.value.category2 = {
            ...value
        }
        // form['value']['category2']['id'] = value['id']
        // form['value']['category2']['name'] = value['name']
        console.log(form['value']['category2'])
    }
    watch(() => form.value.category2, (newVal, oldVal) => {
        console.log('watch new',newVal)
        console.log('watch old', oldVal)
    },{deep:true});
    async function filterItems(){
        if(searchConfig.value['load']){
            return
        }
        searchConfig.value['load'] = true
        // form.value['category'] = null;
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
        console.log(form.value)
        searchConfig.value['load'] = false
    }
    function selectItem(value){
        console.log(value)
        searchConfig.value['id'] = value['id']
        searchConfig.value['open'] = false
        searchConfig.value['search'] = value['full_path']
        // form.value['category'] = value['id'];
        // let res = searchConfig.value.data.find((e)=>e['id'] === value['id'])
        // form.value['naming_category'] = res === undefined ? '':res['full_path']
    }
    function handleBlur(){
        console.log(form.value)
        setTimeout(() => {
            searchConfig.value['open'] = false
            
            if(form.value['category'] != null){
                searchConfig.value['search'] = form.value['naming_category']
                searchConfig.value['id'] = form.value['category']
            }
        }, 100);
    }
    function changeValue(value){
        // messageError.value = {
        //     'name':'',
        //     'merk':'',
        //     'category':'',
        //     'message':messageError.value['message']
        // }
    }
    
    function goBack(){
        router.back();
    }

    function closeButtonSuccess(){
        success.value = false;
    }
    function closeButton(){
        messageError.value = {
            ...messageError.value,
            'message':[]
        }
    }
    async function save() {
        messageError.value = {
            ...messageError.value,
            'name':'',
            'merk':'',
            'sku':'',
            'unit':'',
            'category':'',
        }
        if(form.value['sku'] == ''){
            messageError.value = {
                ...messageError.value,
                'sku':'Please fill SKU',
            }
        }else if(form.value['sku'].length > 20){
            messageError.value = {
                ...messageError.value,
                'sku':'SKU must maximum 20 character',
            }
        }
        if(form.value['unit'] == ''){
            messageError.value = {
                ...messageError.value,
                'unit':'Please fill unit',
            }
        }else if(form.value['unit'].length > 10){
            messageError.value = {
                ...messageError.value,
                'unit':'Unit must maximum 10 character',
            }
        }
        if(form.value['name'] == ''){
            messageError.value = {
                ...messageError.value,
                'name':'Please fill name',
            }
        }else if(form.value['name'].length > 150){
            messageError.value = {
                ...messageError.value,
                'name':'Name must maximum 50 character',
            }
        }
        if(form.value['merk'] == null && form.value['withoutMerk'] == false){
            messageError.value = {
                ...messageError.value,
                'merk':'Please choose merk or set without merk',
            }
        }
        if(form.value['category'] == null){
            messageError.value = {
                ...messageError.value,
                'category':'Please choose category',
            }
        }
        if(messageError.value['name'] != '' || messageError.value['sku'] != '' || messageError.value['unit'] != '' || messageError.value['category'] != '' || messageError.value['merk'] != '') return;
        // if(messageError[]) ) return;
        try {
            
            const data1 = await $api('generic-product',{
                method:'POST',
                
                body:{ 
                    name: form.value['name'],
                    descriptive: form.value['descriptive'],
                    category: form.value['category'],
                    merk: form.value['withoutMerk'] ? null : form.value['merk']
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
            const data = await $api('merk',{
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
    onMounted(() => {
        fetchDataMerk()
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>

</style>