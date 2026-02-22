<template>
    <div>
        <h1 class="text-3xl font-bold mb-6 text-slate-600">Category</h1>
        <div class="bg-white  p-6 rounded-lg shadow-md">
            <BaseToast
                :show="errorMessage !== '' && errorMessage !== null"
                :message="errorMessage"
            />
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
            <div class="mt-5 overflow-x-auto">
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
                        <td class="px-6 py-4 whitespace-nowrap">{{ (index + 1) + (currentPage - 1) * perPage }}</td>
                        <td class="px-6 py-4 whitespace-nowrap">{{ data.name }}</td>
                        <td class="px-6 py-4 whitespace-nowrap">
                            <NuxtLink :to="{ path: `/category/${data['id']}`,  }" class="bg-blue-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border hover:border-yellow-600 rounded cursor-pointer mx-2">Open</NuxtLink>
                            <NuxtLink :to="{ path: `/category/update/${data['id']}`,  }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
                        </td>
                    </tr>
                    
                    </tbody>
                </table>
            </div>
            
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
    const route = useRoute()
    const id = route.params.id

    definePageMeta({
        layout: 'base'
    });
    const query = {}

    if (id && id.length > 0)
        query.parent = id[0]
    
    const {
        data:showTable,
        loading,
        currentPage,
        perPage,
        totalData:totalItems,
        allData,
        errorMessage,
        
    } = useFetchCustom('category/', query)
    const parent = computed(()=>{
        if(allData['path'] !== undefined){
            let path = allData['path']
            if(path != null){
                let names = path.split("/")
                names = names.map((data)=>{
                    let spl = data.split("-")
                    return {'id':spl[0], "name":spl[1]}
                })
                return ["Main", ...names]
            }else{
                return ["Main"]
            }
        }
        
        return []
    })
    
    onMounted(() => {
        
    })
</script>

<style lang="scss" scoped>

</style>