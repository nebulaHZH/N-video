import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const route = [{
        path:'/',
        component:()=>import('./page/HomePage.vue'),
        children:[
            {
                path:'/my',
                component:()=>import('./page/MyDetail.vue')
            }
        ]
}]
const router =  createRouter({
    routes:route as RouteRecordRaw[],
    history:createWebHistory()
})

export default router
