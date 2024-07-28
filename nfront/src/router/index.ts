import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import VideoList from '../components/VideoList.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: VideoList
  },
  {
    path: '/my',
    name: 'my',
    component: () => import('../views/MyDetail.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
