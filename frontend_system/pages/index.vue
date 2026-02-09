<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Dashboard Overview</h1>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-2">Total Users</h3>
            <p class="text-4xl font-bold text-indigo-600 dark:text-indigo-400">1,234</p>
        </div>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-2">New Orders</h3>
            <p class="text-4xl font-bold text-green-600 dark:text-green-400">56</p>
        </div>
        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-2">Revenue</h3>
            <p class="text-4xl font-bold text-yellow-600 dark:text-yellow-400">$12,345</p>
        </div>
        </div>

        <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4">Recent Activity</h3>
            <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead>
                
                <tr>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Activity</th>
                    <th class="px-6 py-3 bg-gray-50 dark:bg-gray-700 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">Date</th>
                </tr>
                </thead>
                <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                <tr>
                    <td class="px-6 py-4 whitespace-nowrap">User John Doe created</td>
                    <td class="px-6 py-4 whitespace-nowrap">2025-06-07</td>
                </tr>
                <tr>
                    <td class="px-6 py-4 whitespace-nowrap">Product A updated</td>
                    <td class="px-6 py-4 whitespace-nowrap">2025-06-06</td>
                </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>


<script setup>
    definePageMeta({
        layout: 'base'
    });
</script>

<style lang="scss" scoped>

</style>