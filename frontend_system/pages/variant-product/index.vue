<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Variant Product</h1>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4">List data</h3>
            <NuxtLink to="/variant-product/create" class="bg-green-600 px-2 py-2 rounded text-gray-100">Create</NuxtLink>
            <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead>
                <tr>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">No. </th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">SKU</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Name</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Unit</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Merk</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Category</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Activity</th>
                </tr>
                </thead>
                <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-for="(data, index) in resp['data']" :key="index">
                    <td class="px-6 py-4 whitespace-nowrap">{{ (setting['page'] - 1) * setting['limit'] + (index + 1) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data['sku'] }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data['name'] }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data['unit'] }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data['merk'] == null ? 'No Merk': (data['merk']['name'] ?? '')}}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data['name_category'] }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                        <NuxtLink :to="{ path: '/variant-product/detail', query: { id: data['id'],  } }" class="bg-green-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Detailt</NuxtLink>
                        <NuxtLink :to="{ path: '/variant-product/update', query: { id: data['id'],  } }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
                    </td>
                </tr>
                
                </tbody>
                
            </table>
            <div class="flex justify-center items-center">
                
                <UPagination v-model:page="setting['page']" :sibling-count="2" active-color="neutral" show-edges :items-per-page="setting['limit']"  :total="resp['total']" />
                
            </div>
            <div class="mx-auto ">
                
            </div>
            
        </div>
    </div>
</template>

<script setup>
    const { $api } = useNuxtApp();
    // const {} = 'nuxt/'
    const resp = ref({
        'data':[],
        'total':0,
    })
    const setting = ref({
        'page':1,
        'limit':10,
        'search':''
    })
    definePageMeta({
        layout: 'base'
    });
    async function fetchData() {
        try {
            const queryParams = new URLSearchParams({page:setting.value['page'], limit:setting.value['limit'] });
            const data = await $api(`variant-product?${queryParams}`,{
                method:'GET'
            });
            
            if(data['data'] !== undefined && data['count_all'] !== undefined){
                resp.value = {
                    'data':data['data'],
                    'total':data['count_all']
                }
            }
            
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }
    }
    watch(()=>setting.value,  (newQuestion, oldQuestion) => {
        fetchData()
    },{deep:true})
    onMounted(() => {
        fetchData()
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>

</style>