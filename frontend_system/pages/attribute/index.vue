<template>
    <div>
        <h1 class="text-3xl font-bold mb-6 text-slate-600">Attribute Name</h1>
        <div class="bg-white  p-6 rounded-lg shadow-md">
            <!-- <BaseToast
                :show="errorMessage !== '' && errorMessage !== null"
                :message="errorMessage"
            /> -->
            <h3 class="text-lg font-semibold text-gray-700 mb-4">List data</h3>
            <NuxtLink to="/attribute/create" class="bg-green-600 px-2 py-2 rounded text-gray-100">Create</NuxtLink>
            <div class="mt-5 overflow-x-auto">
                <table class=" min-w-full divide-y divide-gray-200 ">
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
                            <NuxtLink :to="{ path: `/attribute/${data['id']}`,  }" class="bg-blue-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border hover:border-yellow-600 rounded cursor-pointer mx-2">Open</NuxtLink>
                            <NuxtLink :to="{ path: `/attribute/update/${data['id']}`,  }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
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
    import { useToastStore } from '~/stores/useToastStore'

    definePageMeta({
        layout: 'base'
    });
    const {
        data:showTable,
        loading,
        currentPage,
        perPage,
        totalData:totalItems,
        errorMessage,
        
    } = useFetchCustom('attribute/', {})

    const toastStore = useToastStore()
    
    watch(errorMessage, (newError) => {
        if (newError) {
            // Panggil Pinia untuk memunculkan toast
            toastStore.addToast(newError, 'error')
        }
    })
</script>

<style lang="scss" scoped>

</style>