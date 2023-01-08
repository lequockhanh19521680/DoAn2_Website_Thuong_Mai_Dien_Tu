import { transformFilters } from "../stogare_function/listActions";
import axiosClient from "./Client";
export const ProductApi = {
    GetComment:(id,filters) =>{
        const url = (filters) ? `/comments/product/${id}?${filters}` : `/comments/product/${id} `
        return axiosClient.get(url)
    },
    AddNewComment:(idProduct,idUser,body)=>{
        const url=`/comments/product/${idProduct}/user/${idUser}`
        return axiosClient.post(url,body)
    },
    GetBanners:()=>{
        const url='/banners'
        return axiosClient.get(url)
    },
    GetDetailBanner:(id)=>{
        const url=`/banners/${id}/product/preview`
        return axiosClient.get(url)
    },
    GetCategoriesRoof:()=>{
        const url=`/categories/roof`
        return axiosClient.get(url)
    },
    GetProductPreview:(filters)=>{
        console.log(filters)
        const url=`/products/preview?${transformFilters(filters)}`
        return axiosClient.get(url)
    },
    GetDetailProduct:(id)=>{
        const url=`/products/${id}`
        return axiosClient.get(url)
    },
    GetCategoriesTree:()=>{
        const url='/categories'
        return axiosClient.get(url)
    },
    GetCategoryChildren:(id)=>{
        const url=`/categories/children/${id}`
        return  axiosClient.get(url)
    },
    GetDescriptionFromProduct:(id)=>{
        const url=`/products/${id}/description`
        return axiosClient.get(url)
    },
    GetSpecification:(id)=>{
        const url=`/products/${id}/specification`
        return axiosClient.get(url)
    },
    GetMedia:(id)=>{
        const url=`/products/${id}/media`
        return axiosClient.get(url) 
    },
    GetProductFromCategorySelected:(id,filters)=>{
        const url=`/products/category/${id}/preview?${transformFilters(filters)}`
        return axiosClient.get(url)
    },
    //for add product
    AddNewProduct:(idProvider,idUser,body)=>{
        const url= `/products/provider/${idProvider}/user/${idUser}`
        return axiosClient.post(url,body)
    },
    AddSpecificationTreeInProduct:(idProduct,body)=>{
        const url=`/products/${idProduct}/specification`
        return axiosClient.post(url,body)
    },
    DeleteProduct: (idProduct) =>{
        const url = `/products/${idProduct}`
        return axiosClient.delete(url)
    }
}