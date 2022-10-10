const infoRoutes = [
  {
    path: '/article',
    name: 'Article',
    component: () => import('@/views/Article.vue')
  }, {
    path: '/list',
    name: 'List',
    component: () => import('@/views/List.vue')
  }, {
    path: '/search',
    name: 'Search',
    component: () => import('@/views/Search.vue')
  }
]

export default infoRoutes
