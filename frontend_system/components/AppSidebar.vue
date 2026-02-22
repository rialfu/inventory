<template>
    <aside
        class="fixed inset-y-0 left-0 z-50 w-64 bg-white shadow-md transform transition duration-300 lg:static lg:translate-x-0"

        :class="open ? 'translate-x-0' : '-translate-x-full'"

    >
        <div class="p-6">
            <div class="text-2xl font-bold mb-6">
                Admin Panel
            </div>
            <nav>
                <NuxtLink
                    @click="$emit('close')"
                    to="/category"
                    class="block p-2 hover:bg-gray-100"
                >

                    Category
                </NuxtLink>
                <NuxtLink
                    @click="$emit('close')"
                    to="/merk"
                    class="block p-2 hover:bg-gray-100"
                >
                    Merk
                </NuxtLink>
                <NuxtLink
                    @click="$emit('close')"
                    to="/attribute"
                    class="block p-2 hover:bg-gray-100"
                >
                    Attribute
                </NuxtLink>
                <NuxtLink
                    @click="$emit('close')"
                    to="/product"
                    class="block p-2 hover:bg-gray-100"
                >
                    Product
                </NuxtLink>
            </nav>
        </div>
    </aside>
</template>



<script setup>

defineProps({

open:Boolean

})

</script>