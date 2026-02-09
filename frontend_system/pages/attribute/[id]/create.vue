<template>
    <div>
        <div v-if="notFound">
            <button class="rounded-md bg-gray-950/5 px-2.5 py-1.5 text-sm font-semibold text-gray-900 hover:bg-gray-950/10">Open dialog</button>

            <div class="relative z-10" aria-labelledby="dialog-title" role="dialog" aria-modal="true">
                <!--
                Background backdrop, show/hide based on dialog state.

                Entering: "ease-out duration-300"
                    From: "opacity-0"
                    To: "opacity-100"
                Leaving: "ease-in duration-200"
                    From: "opacity-100"
                    To: "opacity-0"
                -->
                <div class="fixed inset-0 bg-gray-500/75 transition-opacity" aria-hidden="true"></div>

                <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
                <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <!--
                    Dialog panel, show/hide based on dialog state.

                    Entering: "ease-out duration-300"
                        From: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        To: "opacity-100 translate-y-0 sm:scale-100"
                    Leaving: "ease-in duration-200"
                        From: "opacity-100 translate-y-0 sm:scale-100"
                        To: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    -->
                    <div class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                    <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                        <div class="sm:flex sm:items-start">
                            <div class="mx-auto flex size-12 shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:size-10">
                                <svg class="size-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true" data-slot="icon">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
                                </svg>
                            </div>
                            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                                <h3 class="text-base font-semibold text-gray-900" id="dialog-title">Notification</h3>
                                <div class="mt-2">
                                <p class="text-sm text-gray-500">Data Not Found</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                        <button type="button" @click="redirect" class="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-red-500 sm:ml-3 sm:w-auto cursor-pointer">Go Back</button>
                    </div>
                    </div>
                </div>
                </div>
            </div>
        </div>
        <h1 class="text-3xl font-bold mb-6">Create Attribute Value</h1>
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
                    <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
                        Name
                    </label>
                    <input :class="['shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ', (messageError['name'] !== undefined && messageError['name'] !== ''? 'border-red-500 mb-3':'')]" id="name" type="text" placeholder="Name" @input="changeValue" v-model="name">
                    <p class="text-red-500 text-xs italic" v-show=" messageError['name'] !== undefined && messageError['name'] !== ''">{{messageError['name']}}</p>
                </div>
                
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

    const { $api } = useNuxtApp();
    
    const name = ref('')
    const success = ref(false)
    const notFound = ref(false)
    const router = useRouter();
    const route = useRoute()
    const { id } = route.params
    const messageError = ref({'name':'', 'message':[]})
    definePageMeta({
        layout: 'base'
    });
    function changeValue(value){
        messageError.value = {
            'name':'',
            'message':messageError.value['message']
        }
    }
    function goBack(){
        console.log('test')
        router.back();
    }
    function redirect(){
        router.push('/attribute')
    }
    function closeButton(){
        messageError.value = {
            'name':messageError.value['name'],
            'message':[]
        }
    }
    function closeButtonSuccess(){
        success.value = false;
    }
    async function save() {
        if(name.value == ''){
            messageError.value = {
                'name':'Please fill name',
                'message':messageError.value['message']
            }
            return
        }
        
        try {
            if(isNaN(parseInt(id))) return;

            const data1 = await $api('attribute/attribute-value',{
                method:'POST',
                
                body:{ name: name.value, attributeId:parseInt(id) }
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
            // console.error('Failed to fetch data:', error);
        }
    }
    async function fetchData() {
        try {
            const data = await $api(`attribute/${id}`,{
                method:'GET'
            });
            
        } catch (error) {
            if(error.statusCode == 404){
                notFound.value = true
            }
        }
    }
    onMounted(() => {
        fetchData()
    })
</script>

<style lang="scss" scoped>

</style>