const manageRoutes = [
  {
    path: '/secret/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  }, {
    path: '/secret/manage',
    name: 'Manage',
    meta: {
      auth: true
    },
    component: () => import('@/views/Manage.vue')
  }, {
    path: '/secret/edit',
    name: 'Edit',
    meta: {
      auth: true
    },
    component: () => import('@/views/Edit.vue')
  }
]

export default manageRoutes
