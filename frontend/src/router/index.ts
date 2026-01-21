import { getCurrentUser } from '@/firebase'
import { createRouter, createWebHistory } from 'vue-router'
const HomeView = () => import('../views/HomeView.vue')
const AppView = () => import('../views/AppView.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/app',
      name: 'app',
      component: AppView,
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: { name: 'home', replace: true },
    },
  ],
})

router.beforeEach(async (to) => {
  const user = await getCurrentUser()

  if (!user && to.name !== 'home') {
    return { name: 'home', replace: true }
  }

  if (user && to.name === 'home') {
    return { name: 'app', replace: true }
  }
})

export default router
