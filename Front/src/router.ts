// src/router.ts
import { createRouter, createWebHistory } from 'vue-router'
import Admin from './views/Admin.vue'
import Home from './views/Home.vue' // Assurez-vous d'avoir un composant Home.vue pour la route principale

const routes = [
  { path: '/', component: Home },
  { path: '/admin', component: Admin },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
