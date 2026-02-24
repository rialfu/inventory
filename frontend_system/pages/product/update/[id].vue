<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Update Product</h1>
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
            <div>
                <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
                    <div class="mb-4 md:col-span-2 lg:col-span-1">
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
                            :default-name="state.form.category_selected_label"
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
                            :default-name="state.form.merk_selected_label"
                            :error="state.message.merk_id"
                        />
                    </div>
                </div>
                
                
                <CustomButtonForm
                    :is-load="state.load"
                    
                    @save="save"
                />
            </div>
            <div class="my-4">
                <CustomButtonForm
                    :is-load="false"
                    @save="AddVariant"
                    button-name="Add Variant"
                />
            </div>
            
            <div v-for="(variant, index) in stateVariants" :key="index" class="border rounded py-2 px-3 mb-2">
                <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">

                    <div class="mb-2">
                        <CustomInput
                            v-model="stateVariants[index].form.sku"
                            label="SKU"
                            placeholder="Please fill sku"
                            :error="stateVariants[index].message.sku"
                        />
                    </div>
                    <div class="mb-2">
                        <CustomInput
                            v-model="stateVariants[index].form.name"
                            label="Name"
                            placeholder="Please fill name"
                            :error="stateVariants[index].message.name"
                        />
                    </div>
                    <div class="mb-2">
                        <CustomInput
                            v-model="stateVariants[index].form.description"
                            label="Description"
                            placeholder=""
                            :error="stateVariants[index].message.description"
                        />
                    </div>
                    
                </div>
                <div class="mb-2">
                    <CustomButtonForm
                        :is-load="variant.isLoad"
                        color-button="success"
                        @save="()=>AddAttribute(index)"
                        button-name="Add Attribute"
                    />
                </div>
                
                <div class="grid md:grid-cols-2 gap-4 mb-3" v-for="(attribute, j) in variant.form.attributes">
                    <!-- {{ attribute }} -->
                    <div>
                        <div class="flex">
                            <label class="block text-gray-700 text-sm font-bold mb-2"> 
                                Attribute 
                                
                            </label>
                            <button 
                                @click="()=>RemoveAttribute(index, j)"
                                class="ms-2 font-bold text-xs px-2 py-1 rounded focus:outline-none transition-colors duration-200 cursor-pointer bg-red-500 hover:bg-red-700 text-white shadow-md">
                                X
                            </button> 
                        </div>
                        
                        <Dropdown 
                            v-model="variant.form.attributes[j].id" 
                            placeholder="Find Attribute" 
                            url="/item/option/attribute-name"
                            id-value="id",
                            text="name"
                            :default-name="variant.form.attributes[j].selected_label"
                        />
                    </div>
                    
                    <MultiSelect v-if="attribute.id !== null"
                        v-model="variant.form.attributes[j].values" 
                        placeholder="Find Merk" 
                        :url="`/item/option/attribute-value/${attribute.id}`"
                        id-value="id",
                        text="name"
                        label="Attribute Value"
                        
                    />
                </div>
                <div class="flex gap-3 mt-2">
                    <CustomButtonForm
                        :is-load="variant.isLoad"
                        @save="()=>saveVariant(index)"
                        button-name="Save"
                    />
                    <template v-if="variant.form.id === null">
                        <CustomButtonForm
                            :is-load="variant.isLoad"
                            color-button="danger"
                            @save="()=>DeleteVariant(index)"
                            button-name="Delete"
                        />
                    </template>
                    
                </div>
            </div>
                
        </div>
    </div>
</template>

<script setup>
    import * as v from 'valibot'
    import { validateFormCustom, isAssociativeArray } from '~/utils/helpers'
    import { useToastStore } from '~/stores/useToastStore'
    
    definePageMeta({
        layout: 'base'
    });
    const { $api } = useNuxtApp();

    const toastStore = useToastStore()
    const success = ref(false)
    const router = useRouter();
    const route = useRoute();
    const id = route.params.id

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
        load:false
    })
    const createDefaultFormatVariant = ()=> {
       return {
            form:{
                id:null,
                sku:'',
                name:'',
                description:'',
                image_url:'',
                stock:0,
                values:[],
                attributes:[]
            },
            message:{
                sku:'',
                name:'',
                description:'',
                values:''
            },
            isLoad:false,
       }
        
    }
    const createDefaultFormatAttribute = ()=>{
        return {
            id:null,
            selected_label:'',
            values:[]
        }
       
    }
    const stateVariants = ref([])
    
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
        return
        try {
            if(validateFormCustom(v, schema, state, ['name', 'merk_id', 'category_id'])){
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
    function AddVariant(){
        stateVariants.value.push(createDefaultFormatVariant())
    }
    function DeleteVariant(index){
        const res = stateVariants.value[index]
        if(res == undefined || res['id'] != null) return
        stateVariants.value.splice(index, 1)
    }

    function AddAttribute(index){
        stateVariants.value[index].form.attributes.push(createDefaultFormatAttribute())
    }
    function RemoveAttribute(i, j){
        stateVariants.value[i].form.attributes.splice(j, 1)
    }
    async function saveVariant(index){
        const val = stateVariants.value[index]
        if(val === undefined || val === null){
            return
        }
        stateVariants.value[index].isLoad = true
        console.log(val)
        stateVariants.value[index].isLoad = false
    }
    async function fetchItem(){
        try{
            const res = await $api(`item/${id}/item`,{
                method:'GET',    
            });
            if(res['data'] === undefined){
                toastStore.addToast('Data not found')
                return
            }
            if(res['data'] !== undefined){
                if(res['data']['merk_name'] !== undefined){
                    state.form.merk_selected_label = res['data']['merk_name'] ?? ''
                }
                if(res['data']['merk_id'] !== undefined){
                    state.form.merk_id = res['data']['merk_id'] ?? ''
                }
                if(res['data']['category_id'] !== undefined){
                    state.form.category_id = res['data']['category_id'] ?? ''
                }
                if(res['data']['category_name'] !== undefined){
                    state.form.category_selected_label = res['data']['category_name'] ?? ''
                }
                if(res['data']['name'] !== undefined){
                    state.form.name = res['data']['name'] ?? ''
                }
            }
            console.log(res)
    
        }catch(err){
            // err.s
        }
       
    }
    async function fetchVariant(){
        try{
            const res = await $api(`item/${id}/variant`,{
                method:'GET',    
            });
            if(res['data'] === undefined){
                toastStore.addToast('Data not found')
                return
            }
            
            console.log(res)
    
        }catch(err){
            // err.s
        }
       
    }
    onMounted(() => {
        fetchItem()
        fetchVariant()
    })
</script>

<style lang="scss" scoped>

</style>