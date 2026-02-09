<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Attribute</h1>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4">List data</h3>
            <NuxtLink to="/attribute/create" class="bg-green-600 px-2 py-2 rounded text-gray-100">Create</NuxtLink>
            <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead>
                <tr>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">No. </th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Name</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Activity</th>
                </tr>
                </thead>
                <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-for="(data, index) in showTable" :key="index">
                    <td class="px-6 py-4 whitespace-nowrap">{{ index + 1 }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ data.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap">
                        <NuxtLink :to="{ path: `/attribute/${data['id']}`, }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Detail</NuxtLink>
                        <NuxtLink :to="{ path: '/attribute/update', query: { id: data['id'], name:data['name'] } }" class="bg-red-400 hover:bg-yellow-600 text-white font-bold py-2 px-4 border border-yellow-600 rounded cursor-pointer mx-2">Edit</NuxtLink>
                    </td>
                </tr>
                
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
    const { $api } = useNuxtApp();
    const showTable = ref([])
    definePageMeta({
        layout: 'base'
    });
    async function fetchData() {
        try {
            const data = await $api('attribute',{
                method:'GET'
            });
            if(data['data'] !== undefined){
                showTable.value = data['data']
            }
             // Ini akan memanggil https://api.example.com/v1/users
            console.log(data);
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }
    }
    onMounted(() => {
        fetchData()
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>

</style>