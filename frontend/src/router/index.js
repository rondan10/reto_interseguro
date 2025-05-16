import { createRouter, createWebHistory } from 'vue-router'
import MatrixCalculator from '../components/MatrixCalculator.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: MatrixCalculator
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
