<template>
    <header class="bg-white shadow-sm p-4 flex items-center justify-between">
        <div class="flex items-center gap-3">
            <button
            @click="$emit('toggleSidebar')"

            class="lg:hidden"

            >

            ☰

            </button>


            <h1 class="text-xl font-bold">
            Admin Panel
            </h1>
        </div>
    </header>
</template>