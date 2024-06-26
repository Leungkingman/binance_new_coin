import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)
const baseRouters = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: () => import ('@/view/login/login.vue')
  },
  {
    path: '/console/workplatform',
    name: 'workplatform',
    component: () => import ('@/view/console/workplatform.vue')
  },
  {
    path: '/console/logplatform',
    name: 'logplatform',
    component: () => import ('@/view/console/logplatform.vue')
  }
]

const createRouter = () => new Router({
    routes: baseRouters
})

const router = createRouter()
export default router
