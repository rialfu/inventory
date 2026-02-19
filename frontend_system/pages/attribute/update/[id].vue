<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Update Attribute</h1>
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
                    <li class="mb-1" >Success Update</li>
                </ul>
            </div>
            <form class="" @submit.prevent>
                <div class="mb-4">
                    <CustomInput
                        v-model="state.form.name"
                        label="Name"
                        id="name"
                        placeholder="Please fill category name"
                        :error="state.message.name"
                    />
                </div>
                
                <CustomButtonForm
                    :is-load="state.load"
                    @save="save"
                />
                
            </form>
        </div>
    </div>
</template>

<script setup>
    import * as v from 'valibot'

            // 1. Definisikan aturan (Schema)
    const schema = v.object({
        name: v.pipe(v.string(), v.minLength(1, 'Name must fill')),
    })

    definePageMeta({
        layout: 'base'
    });

    const { $api } = useNuxtApp();
    
    const success = ref(false)
    const route = useRoute();
    const router = useRouter()
    const id = route.params.id
    
    const state = reactive({
        form: {
            id:null,
            name: '',
            parent_id: null, 
            selected_label:'',
        },
        message: {
            name: '',
            parent_id: '',
            
        },
        listMessage:[],
        load:false,
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
    const validateForm = () => {
        const result = v.safeParse(schema, state.form);

        if (!result.success) {
            const flattenedErrors = v.flatten(result.issues).nested;

            if (flattenedErrors.name) state.message.name = flattenedErrors.name[0];

            state.listMessage = v.flatten(result.issues).root || Object.values(flattenedErrors).flat();
            
            return false;
        }

        return true; // Validasi berhasil
    };
    async function save() {
        resetMessageAll()
        try {
            if(validateForm()){
                let body = { name: state.form.name }
                const data = await $api('attribute/'+state.form.id,{
                    method:'PUT',
                    body,
                });
                
                success.value = true
            }
        } catch (error) {
            console.log(error.data)
            if(error.data['error'] !== undefined && error.data['error']['Name'] != undefined){
                // err = error.data['error']
                if(error.data['error']['Name'] != undefined){
                    state.message.name = error.data['error']['Name']
                    state.listMessage = [...state.listMessage, "Name "+error.data['error']['Name']]
                }
                

                
            
            }else if(error.data['error']){
                state.listMessage = [...state.listMessage, `${error.data['error']}`]
            }else if(error.data['message'] !=undefined){
                state.listMessage = [...state.listMessage, error.data['message']]
            }
            
            success.value = false;
        }
    }
    async function fetch(){
        try{
            if (state.load) return
            state.load = true
            const res = await $api('attribute/'+id,{
                method:'GET',
            });
            const resData = res['data']
            console.log(resData)
            if (resData['id'] != undefined){
                state.form.id = resData['id']
                
            }
            if (resData['name'] != undefined){
                state.form.name = resData['name']
                
            }
            console.log(res)
        }catch(err){
            console.log(err)
        }finally{
            // await delay(5000);
           state.load = false
        }
        
    }
    onMounted(async() => {
        await nextTick();
        fetch()
        
    })
</script>

<style lang="scss" scoped>

</style>