<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Create Merk</h1>
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
                    <input :class="[
                            'border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-1 focus:ring-blue-500 transition-all', 
                            (messageError['name'] ? 'border-red-500' : 'border-gray-300')
                        ]" 
                        id="name" type="text" placeholder="Name" @input="changeValue" v-model="name">
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
    const router = useRouter();
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

    function closeButtonSuccess(){
        success.value = false;
    }
    function closeButton(){
        messageError.value = {
            'name':messageError.value['name'],
            'message':[]
        }
    }
    async function save() {
        // if(name.value == ''){
        //     messageError.value = {
        //         'name':'Please fill name',
        //         'message':messageError.value['message']
        //     }
        //     return
        // }
        
        try {
            
            const data1 = await $api('merk/',{
                method:'POST',
                
                body:{ name: name.value }
            });
            messageError.value = {
                'name':'',
                'message':[]
            }
            success.value = true
        } catch (error) {
            console.log(error.data)
             if(error.data['error'] !== undefined && error.data['error']['Name'] != undefined){
                // err = error.data['error']

                messageError.value = {
                    'name':error.data['error']['Name'],
                    'message':[`Name ${error.data['error']['Name']}`]
                }
            
            }else if(error.data['message'] !=undefined){
                messageError.value = {
                    'name':'',
                    'message':error.data['message']
                }
            }
            
            success.value = false;
        }
    }
    onMounted(() => {
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>

</style>