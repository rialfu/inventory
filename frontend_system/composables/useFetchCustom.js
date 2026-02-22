export function useFetchCustom(url, initialQuery = {}, ) {

    const { $api } = useNuxtApp()

    const data = ref([])
    const allData = reactive({})
    const loading = ref(false)
    const errorMessage = ref(null)
    const currentPage = ref(1)
    const perPage = ref(10)
    const search = ref("")
    const isError401 = ref(false)
    const totalData = ref(0)
    let timeOutToast = null

    function setTimeoutToast(){
        if(timeOutToast != null){
            clearTimeout(timeOutToast)
        }
        setTimeout(()=>{
            timeOutToast = null
            errorMessage.value = ""
        }, 3000)
    }
    const fetchData = async () => {
        errorMessage.value = null
        loading.value = true

        try {
            const query = {...initialQuery}
            query['page'] = currentPage.value
            query['limit'] = perPage.value

            const res = await $api(url, {
                query
            })
            
            data.value = res?.data?.data || []
            totalData.value = res?.data?.total || 0
            Object.assign(allData, res?.data ?? {})
            
        }catch(err){
            const status = err.response.status
            if(status === 401){
                isError401.value = true
            }
            const errData = err.data
            if(errData['error'] !== undefined && errData['error'] !== null){
                errorMessage.value = `${errData['error']}`
            }else if(errData['message'] !== undefined && errData['message'] !== null){
                errorMessage.value = `${errData['message']}`
            }else{
                errorMessage.value = "Something is wrong"
            }
            setTimeoutToast()
        }finally {

            loading.value = false

        }

    }


    // AUTO FETCH ketika berubah

    watch(
        [currentPage, perPage, search],
        fetchData
    )
    onMounted(fetchData)

    return {

        data,
        allData,
        loading,
        isError401,
        errorMessage,

        currentPage,
        perPage,
        search,
        totalData,


        fetchData

    }

}