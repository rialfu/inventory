export default defineNuxtPlugin((nuxtApp) => {
    const config = useRuntimeConfig();

    // Inisialisasi `$api` kustom Anda
    const $api = $fetch.create({
        baseURL: config.public.apiBase,
        onRequest({ request, options, error }) {
        // console.log()
        // Tambahkan header otorisasi jika token ada
            const token = useCookie('auth_token').value; // Contoh mendapatkan token dari cookie
            // options.headers ={
            //     ...options.headers,
            //     'Content-Type': 'application/json'
            // }
            // options.headers.append( 'Content-Type', 'application/json')
            if (token) {
                options.headers.append('Authorization', `Bearer ${token}`)
            }
            // console.log('Request:', request, options);
        },
        onResponse({ request, response, options }) {
            // console.log('Response:', request, response);
        // Anda bisa memproses respons di sini
        },
        onRequestError({ request, options, error }) {
            console.error('Request Error:', request, error);
        // Tangani kesalahan permintaan di sini
        // Misalnya, tampilkan notifikasi error
        },
        onResponseError({ request, response, options }) {
            
            console.error('Response Error:', request, response.status, response._data);
            // Tangani kesalahan respons (misalnya, status 4xx, 5xx)
            // Misalnya, redirect ke halaman login jika 401
            if (response.status === 401) {
                // useNuxtApp().$router.push('/login'); // Hanya jika di sisi client
                console.log('Unauthorized, redirecting to login...');
            }
        }
    });

    // Daftarkan $api sebagai helper global di NuxtApp
    nuxtApp.provide('api', $api);
});