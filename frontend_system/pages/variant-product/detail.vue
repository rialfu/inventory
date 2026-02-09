<template>
    <div>
        <h1 class="text-3xl font-bold mb-6">Detail Variant Product</h1>
        
        <div class="grid grid-cols-3 gap-3">
            <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md col-span-3 md:col-span-1 min-h-100">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4">Information data</h3>
                <div class="grid grid-cols-[auto_1fr] gap-x-2">
                    <div class="text-gray-700 font-semibold text-left">Nama</div>
                    <div class="text-gray-900">: Bla bla</div>

                    <div class="text-gray-700 font-semibold text-left">Umur</div>
                    <div class="text-gray-900">: 12</div>

                    <div class="text-gray-700 font-semibold text-left">Tinggal</div>
                    <div class="text-gray-900">: jln swaday 1 rt 12/10 no 9b pejaten timur, pasar mingu jakarta selatan jakarta indonesia, asia tengagara asia bumi tata surya</div>
                </div>
            
            </div>
            <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md col-span-3 md:col-span-2 min-h-100">
            
            </div>
        </div>
        
    </div>
</template>

<script setup>
    const {$api } = useNuxtApp()
    const route = useRoute()
    const { id = ''} = route.query;
    definePageMeta({
        layout: 'base'
    });
    const data = ref({
        'load':false,
        'sku':'',
        'name':'',
        'unit':'',
        'merk':'',
        'category':'',
        'active': false,
        'attributes':[],
    })
    const atributes = ref({})
    async function fetchData() {
        if(data.value['load']) return false;
        data.value = {...data.value, 'load':true}
        try {
            // const queryParams = new URLSearchParams({page:setting.value['page'], limit:setting.value['limit'] });
            const dataServer = await $api(`variant-product/search/${id}`,{
                method:'GET'
            });
            
            if(dataServer['data'] !== undefined){
                data.value = {...data.value, 
                    sku:dataServer['data']['sku'], 
                    name:dataServer['data']['name'], 
                    unit:dataServer['data']['unit'],
                    merk:dataServer['data']['merk'] === null ? 'No Merk': dataServer['data']['merk']['merk'],
                    category: dataServer['data']['category_name'],
                    active: dataServer['data']['active']
                }
                let cache = {}
                if(dataServer['data']['variant_product_attribute'] instanceof Array){
                    let attrData = dataServer['data']['variant_product_attribute'];
                    let value = null
                    let key = null;
                    for(let i=0; i<attrData.length; i++){
                        value = {...attrData[i]['attribute_value']}
                        key = {...attrData[i]['attribute_value']['attribute']}
                        console.log(key, value)
                        if(key !== undefined){
                            if(cache[key['id']] === undefined){
                                delete value['attribute']
                                cache[key['id']] = { name:key['name'], data:[{...value}] }
                            }
                            
                        }else{
                            cache[key['id']]['data'] =  [...cache[key['id']]['data'], {...value}]
                        }
                    }
                }
                console.log(cache)
            }
            console.log(data)
            
        } catch (error) {
            console.error('Failed to fetch data:', error);
        }
        data.value = {...data.value, 'load':false}
    }
    onMounted(() => {
        fetchData()
        console.log(`the component is now mounted.`)
    })
</script>

<style lang="scss" scoped>
    
</style>