<template>
    <div>
        <h1 class="text-3xl font-bold mb-6 text-slate-600">Category</h1>
        <div class="bg-white  p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 mb-4">List data</h3>
            <NuxtLink to="/category/create" class="bg-green-600 px-2 py-2 rounded text-gray-100">Create</NuxtLink>
            
            <div v-if="parent.length > 0 && id!=''" class="flex items-center mt-5 flex-wrap">
  
                <NuxtLink
                    to="/category"
                    class="hover:text-blue-600 transition-colors"
                >
                    Main
                </NuxtLink>
                <template v-if="parent.length > 1">

                </template>
                

                <template v-for="(item, index) in parent" :key="index">
                    <template v-if="index != 0">
                        <span class="mx-2 text-gray-400">/</span>
                        <NuxtLink
                  
                        :to="`/category/${item.id}`"
                        class="hover:text-blue-600 transition-colors"
                        >
                        {{ item.name }}
                        </NuxtLink>
                    </template>
                    
                </template>
            </div>
            <table class="mt-5 min-w-full divide-y divide-gray-200 ">
                <thead class="divide-y divide-gray-700 border-b border-gray-200">
                    <tr class="border-b border-gray-200 divide-x divide-gray-200">
                        <th class="w-[80px] px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider whitespace-nowrap">
                            No.
                        </th>
                        
                        <th class="px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Name
                        </th>
                        
                        <th class="px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Activity
                        </th>
                    </tr>
                </thead>
                <tbody class="bg-white  divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-if="showTable.length == 0">
                    <td colspan="3" class="px-6 py-4 text-slate-600 whitespace-nowrap text-center">
                        No Data
                    </td>
                </tr>
                <tr v-else v-for="(data, index) in showTable" :key="index">
                    <td class="px-6 py-4 whitespace-nowrap">{{ index + 1 }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                        <NuxtLink :to="{ path: `/category/${data['id']}`,  }" class="bg-blue-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border hover:border-yellow-600 rounded cursor-pointer mx-2">Open</NuxtLink>
                        <NuxtLink :to="{ path: `/category/update/${data['id']}`,  }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
                    </td>
                </tr>
                
                </tbody>
            </table>
            <PaginationControl
                v-model="currentPage"
                :total-items="totalItems"
                :per-page="perPage"
                :is-loading="loading"
                :is-empty="showTable.length === 0"
            />
        </div>
        
    </div>
</template>

<script setup>
    const { $api } = useNuxtApp();
    const showTable = ref([])
    const currentPage = ref(1)
    const perPage = ref(10)
    const totalItems = ref(0)
    const loading = ref(false)
    const parent = ref([])
    
    const route = useRoute()
    const id = route.params.id

    definePageMeta({
        layout: 'base'
    });

    async function fetchData() {
        try {
            loading.value = true
            const query = {
                page:currentPage.value,
                limit:perPage.value,
            }
            if (id.length > 0){
                query['parent'] = id[0]
            }
            const res = await $api('category/',{
                method:'GET',
                query,
            });
            let data = []
            if(res['data'] != undefined){
                console.log(res['data'])
                if(res['data']['data'] !== undefined){
                    data = res['data']['data']
                }
                if(res['data']['limit'] !==undefined){
                    perPage.value = res['data']['limit']
                }
                if(res['data']['total'] !==undefined){
                    totalItems.value = res['data']['total']
                }
                if(res['data']['path'] !==undefined){
                    if(res['data']['path'] != null){
                        let names = res['data']['path'].split("/")
                        names = names.map((data)=>{
                            let spl = data.split("-")
                            return {'id':spl[0], "name":spl[1]}
                        })
                        parent.value = ["Main", ...names]
                    }else{
                        parent.value = ["Main"]
                    }
                    
                }
               
            }
            showTable.value = data
            console.log(res);
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }finally{
            loading.value = false
        }
    }
    onMounted(() => {
        fetchData()
        
    })
</script>

<style lang="scss" scoped>

</style>