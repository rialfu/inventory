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
        <h1 class="text-3xl font-bold mb-6"> Data Attribute</h1>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4"><span @click="goBack" class="cursor-pointer">←</span> Data Attribute</h3>
            <p class="mb-3">Name : {{ showTable.name }} </p>
            <span 
                v-if="showTable.isLoad" 
                class="bg-green-600 px-2 py-2 rounded text-gray-100 opacity-50 cursor-not-allowed"
            >
                Loading...
            </span>

            <NuxtLink 
                v-else
                :to="{ path: `/attribute/${id}/create` }" 
                class="bg-green-600 px-2 py-2 rounded text-gray-100 hover:bg-green-700"
            >
                Create
            </NuxtLink>
            <table class="mt-4 min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead>
                <tr>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">No. </th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Attribute Value</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Activity</th>
                </tr>
                </thead>
                <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-if="(showTable.attribute_value ?? []).length == 0">
                    <td class="px-6 py-4 whitespace-nowrap">Empty Data</td>
                </tr>
                <tr v-else v-for="(data, index) in showTable['attribute_value'] ?? []" :key="index">
                    <td class="px-6 py-4 whitespace-nowrap">{{ index + 1 }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                        <NuxtLink :to="{ path: `/attribute/${id}/update`,query: { id: data['id'], name:data['name'] } }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
                        
                    </td>
                </tr>
                
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
    const { $api } = useNuxtApp();
    const router = useRouter();
    const route = useRoute()
    const { id } = route.params
    const notFound = ref(false)
    const showTable = ref({
        name:'',
        attribute_value:[],
        isLoad:true

    })
    definePageMeta({
        layout: 'base'
    });
    function goBack(){
        router.back();
    }
    function redirect(){
        router.push('/attribute')
    }
    async function fetchData() {
        try {
           
            
            const res = await $api(`attribute/${id}/attribute-value`,{
                method:'GET'
            });
            console.log(res)
            if(res['data'] !== undefined){
                if(res['data']['name'] !== undefined){
                    showTable.value.name = res['data']['name']
                }
                console.log(res['data'])
                // showTable.value = data['data']
            }
        } catch (error) {
            if(error.statusCode == 404){
                notFound.value = true
            }
            console.log(error.statusCode)
            console.error('Failed to fetch data:', error);
        }finally{
            showTable.value.isLoad = false
        }
    }

    onMounted(() => {
        fetchData()
        console.log(route.path.startsWith("/attribute"))
    })
</script>

<style lang="scss" scoped>

</style>