
<template>
    <div class="lg:flex h-screen bg-gray-100">
        <AppSidebar
            :open="isSidebarOpen"
            @close="isSidebarOpen=false"
        />
        <AppOverlay
            :open="isSidebarOpen"
            @close="isSidebarOpen=false"
        />
        <div class="flex-1 flex flex-col overflow-hidden">
            <AppHeader
                @toggleSidebar="isSidebarOpen=true"
            />
            <main class="flex-1 overflow-auto p-6">
                <slot></slot>
            </main>
        </div>
    </div>
</template>



<script setup>

const isSidebarOpen = ref(false)

</script>