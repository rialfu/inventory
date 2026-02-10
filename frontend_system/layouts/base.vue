<template>
  <UApp>

  <div class="flex h-screen bg-gray-100 font-sans antialiased">
    <aside class="w-64 bg-white text-gray-800 p-6 shadow-md flex-shrink-0">
      <div class="text-2xl font-bold text-blue-700 mb-8">Admin Panel</div>
      <nav>
        <ul>
          <li class="mb-2">
            <button
              @click="toggleMenu('main1')"
              class="flex items-center justify-between w-full p-2 rounded-md hover:bg-gray-100 text-gray-600 hover:text-blue-500 transition-colors duration-200 focus:outline-none"
              :class="{ 'bg-gray-100 text-blue-500': openMenus.includes('main1') }"
            >
              <div class="flex items-center">
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2 2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path></svg>
                <span class="text-lg font-semibold">Main 1</span>
              </div>
              <svg
                class="w-4 h-4 transition-transform duration-200"
                :class="{ 'rotate-90': openMenus.includes('main1') }"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
            <ul v-show="openMenus.includes('main1')" class="ml-8 mt-2 space-y-2">
              <li class="mb-2">
                <NuxtLink to="/merk" class="flex items-center p-2 rounded-md hover:bg-gray-100 text-gray-600 hover:text-blue-500 transition-colors duration-200">
                  Dashboard
                </NuxtLink>
              </li>
              <li class="mb-2">
                <NuxtLink to="/merk" class="flex items-center p-2 rounded-md hover:bg-gray-100 text-gray-600 hover:text-blue-500 transition-colors duration-200">
                  Users
                </NuxtLink>
              </li>
            </ul>
          </li>

          <li class="mb-2">
            <button
              @click="toggleMenu('main2')"
              class="flex items-center justify-between w-full p-2 rounded-md hover:bg-gray-100 text-gray-600 hover:text-blue-500 transition-colors duration-200 focus:outline-none"
              :class="{ 'bg-gray-100 text-blue-500': openMenus.includes('main2') }"
            >
              <div class="flex items-center">
                <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path></svg>
                <span class="text-lg font-semibold">Main 2</span>
              </div>
              <svg
                class="w-4 h-4 transition-transform duration-200"
                :class="{ 'rotate-90': openMenus.includes('main2') }"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
            <ul v-show="openMenus.includes('main2')" class="ml-8 mt-2 space-y-2">
              <li class="mb-2">
                <NuxtLink to="/attribute" :class="['flex items-center p-2 rounded-md transition-colors duration-200', (route.path.startsWith('/attribute') ? ` bg-gray-100 text-blue-500`:` hover:bg-gray-100 text-gray-600 hover:text-blue-500`)]">
                  Attribute
                </NuxtLink>
              </li>
              <li class="mb-2">
                <NuxtLink to="/category" :class="['flex items-center p-2 rounded-md transition-colors duration-200', (route.path.startsWith('/category') ? ` bg-gray-100 text-blue-500`:` hover:bg-gray-100 text-gray-600 hover:text-blue-500`)]">
                  Category
                </NuxtLink>
              </li>
              <li class="mb-2">
                <NuxtLink to="/merk" :class="['flex items-center p-2 rounded-md transition-colors duration-200', (route.path.startsWith('/merk') ? ` bg-gray-100 text-blue-500`:` hover:bg-gray-100 text-gray-600 hover:text-blue-500`)]">
                  Merk
                </NuxtLink>
              </li>
              <li class="mb-2">
                <NuxtLink to="/variant-product" :class="['flex items-center p-2 rounded-md transition-colors duration-200', (route.path.startsWith('/variant-product') ? ` bg-gray-100 text-blue-500`:` hover:bg-gray-100 text-gray-600 hover:text-blue-500`)]">
                  General Product
                </NuxtLink>
              </li>
            </ul>
          </li>

          <li class="mb-2">
            <NuxtLink to="/merk" class="flex items-center p-2 rounded-md hover:bg-gray-100 text-gray-600 hover:text-blue-500 transition-colors duration-200">
              <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.942 3.333.999 2.405 2.503a1.724 1.724 0 001.066 2.573c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.942 1.543-.999 3.333-2.503 2.405a1.724 1.724 0 00-2.573 1.066c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.942-3.333-.999-2.405-2.503a1.724 1.724 0 00-1.066-2.573c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.942-1.543.999-3.333 2.503-2.405a1.724 1.724 0 002.573-1.066z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>
              App Settings
            </NuxtLink>
          </li>
        </ul>
      </nav>
    </aside>

    <main class="flex-1 p-6 overflow-auto">
      <header class="bg-white shadow-sm p-4 mb-6 rounded-lg flex justify-between items-center">
        <h1 class="text-2xl font-semibold text-gray-800">
          {{ $route.meta.title || 'Admin Dashboard' }}
        </h1>
        <div class="flex items-center space-x-4">
          <span class="text-gray-600">Welcome, Admin!</span>
          <button class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md transition-colors duration-200">
            Logout
          </button>
        </div>
      </header>

      <slot />

      <footer class="mt-8 text-center text-gray-500 text-sm">
        &copy; {{ new Date().getFullYear() }} My Admin App. All rights reserved.
      </footer>
    </main>
  </div>
  </UApp>

</template>

<script setup>
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router'; // Import useRoute dari vue-router

const openMenus = ref([]); // Array untuk menyimpan ID menu yang sedang terbuka

const toggleMenu = (menuId) => {
  const index = openMenus.value.indexOf(menuId);
  if (index > -1) {
    // Menu sudah terbuka, tutup
    openMenus.value.splice(index, 1);
  } else {
    // Menu tertutup, buka
    openMenus.value.push(menuId);
  }
};

// Logika untuk membuka menu secara otomatis jika salah satu sub-menunya aktif
const route = useRoute(); // Dapatkan objek route saat ini
console.log(route)
watch(() => route.path, (newPath) => {
  // Anda perlu memetakan path sub-menu ke ID menu utama yang sesuai
  // Ini adalah contoh sederhana, Anda mungkin ingin membuat mapping yang lebih canggih
  if (newPath.startsWith('/admin/dashboard') || newPath.startsWith('/admin/users')) {
    if (!openMenus.value.includes('main1')) {
      openMenus.value.push('main1');
    }
  } else {
    // Opsi: Tutup menu jika tidak ada sub-menu yang aktif di dalamnya
    const index = openMenus.value.indexOf('main1');
    if (index > -1 && !newPath.startsWith('/admin/dashboard') && !newPath.startsWith('/admin/users')) {
      // openMenus.value.splice(index, 1); // Uncomment jika ingin menutup otomatis
    }
  }

  if (newPath.startsWith('/category') || newPath.startsWith('/attribute') || newPath.startsWith('/variant-product') || newPath.startsWith('/merk')) {
    if (!openMenus.value.includes('main2')) {
      openMenus.value.push('main2');
    }
  } else {
    const index = openMenus.value.indexOf('main2');
    if (index > -1 && !newPath.startsWith('/category') && !newPath.startsWith('/attribute') && !newPath.startsWith('/variant-product') || newPath.startsWith('/merk')) {
      // openMenus.value.splice(index, 1); // Uncomment jika ingin menutup otomatis
    }
  }
}, { immediate: true }); // `immediate: true` agar watcher berjalan saat komponen pertama kali di-mount

</script>

<style scoped >
/*
  Jika Anda perlu menimpa atau menambahkan gaya spesifik untuk layout ini yang tidak bisa dihandle Tailwind
  atau jika Anda ingin mengisolasi gaya ke layout ini.
*/
/* .router-link-active {
  @apply bg-blue-500 text-white !important; 
} */
/* .router-link-active svg {
    @apply text-white; 
} */
</style>