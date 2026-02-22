<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Create Category</h1>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
           
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4"><span @click="goBack" class="cursor-pointer">←</span> Form</h3>
            <div class="relative bg-red-600 text-white p-4 rounded-lg shadow-md max-w-sm" v-show="(state.listMessage ?? []).length > 0">
                <button class="absolute top-2 right-2 text-white hover:text-gray-200 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50" @click="closeButton">
                    <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                </button>

                <h3 class="font-bold text-lg mb-2">Validation Error!</h3>
                <ul class="list-disc list-inside p-0 m-0 text-sm">
                    <li class="mb-1" v-for="(data, index) in (state.listMessage ?? []) " :key="index">{{data}}</li>
                    
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
            <div class="grid grid-cols-3 gap-4">
                <div class="mb-4">
                    <CustomInput
                        v-model="state.form.name"
                        label="Name"
                        id="name"
                        placeholder="Please fill category name"
                        :error="state.message.name"
                    />
                </div>
                <div class="mb-4">
                    <Dropdown
                        v-model="state.form.category_id" 
                        placeholder="Find Category " 
                        url="/item/option/category"
                        id-value="id",
                        text="name"
                        label="Category"
                        :input-id="state.message.category_id"
                        :error="state.message.category_id"
                    />
                </div>
                <div class="mb-4">
                    <Dropdown
                        v-model="state.form.merk_id" 
                        placeholder="Find Merk" 
                        url="/item/option/merk"
                        id-value="id",
                        text="name"
                        label="Merk"
                        :input-id="state.message.merk_id"
                        :error="state.message.merk_id"
                    />
                </div>
            </div>
            
            
            <CustomButtonForm
                :is-load="state.load"
                @save="save"
            />
                
        </div>
    </div>
</template>

<script setup>
    import * as v from 'valibot'
    import { validateFormCustom, isAssociativeArray } from '~/utils/helpers'
    definePageMeta({
        layout: 'base'
    });
    const { $api } = useNuxtApp();
    
    const success = ref(false)
    const router = useRouter();
    
    const schema = v.object({
        name: v.pipe(v.string(), v.minLength(1, 'Name must fill')),
        category_id: v.pipe(
            v.number('Category must fill'),
            v.minValue(1, 'Category must fill')
        ),
        merk_id: v.pipe(
            v.number('Merk must fill'),
            v.minValue(1, 'Merk must fill')
        ),
    })
    
    const state = reactive({
        form: {
            name: '',
            category_id: null, 
            category_selected_label:'',
            merk_id:null,
            merk_selected_label:'',
        },
        message: {
            name: '',
            category_id: '',
            merk_id:'',
            
        },
        listMessage:[],
        load:true
    })
    
    function goBack(){
        
        router.back();
    }

    function closeButtonSuccess(){
        success.value = false;
    }
    function resetMessageAll(){
        state.message = {
            name: '',
            parent_id: '',
        }
        state.listMessage = []
    }
    function closeButton(){
        resetMessageAll() 
    }
    function selectedValue(item){

    }
    async function save() {
        resetMessageAll()
        try {
            // validateFormCustom(v, schema, state, ['name', 'merk_id', 'category_id'])
            if(true){
                let body = { name:state.form.name, category:state.form.category_id,
                    merk:state.form.merk_id 
                }
                const data = await $api('item/',{
                    method:'POST',
                    body,
                });
                success.value = true
            }
        } catch (error) {
            console.log(error.data)
            if(error.data['error'] !== undefined && (
                error.data['error']['Name'] !== undefined || 
                error.data['error']['Category'] !== undefined ||
                error.data['error']['Merk'] !== undefined
            )){
                if(error.data['error']['Name'] !== undefined){
                    state.message.name = error.data['error']['Name']
                    state.listMessage = [...state.listMessage, "Name "+error.data['error']['Name']]
                }
                if(error.data['error']['Category'] !== undefined){
                    state.message.category_id = error.data['error']['Category']
                    state.listMessage = [...state.listMessage, "Category "+error.data['error']['Category']]
                }
                if(error.data['error']['Merk'] !== undefined){
                    state.message.merk_id = error.data['error']['Merk']
                    state.listMessage = [...state.listMessage, "Merk "+error.data['error']['Merk']]
                }
            
            }else if(error.data['error']){
                if(Array.isArray(error.data['error'])){
                    state.listMessage = [...state.listMessage, ...(error.data['error'])]
                }else if(isAssociativeArray(error.data['error'])){
                    const values = Object.entries(error.data['error']).map(([key, value]) => `${key}: ${value}`)
                    state.listMessage = [...state.listMessage, ...values]
                }else{
                    state.listMessage = [...state.listMessage, `${error.data['error']}`]
                }
            }else if(error.data['message'] !=undefined){
                state.listMessage = [...state.listMessage, error.data['message']]
            }
            
            success.value = false;
        }
    }
    onMounted(() => {
        state.load = false
    })
</script>

<style lang="scss" scoped>

</style>