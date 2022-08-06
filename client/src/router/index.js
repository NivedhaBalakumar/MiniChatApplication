import Vue from 'vue'
import Router from 'vue-router'
import ChatContainer from './../components/ChatContainer'
import GroupOne from './../components/GroupOne'
import GroupTwo from './../components/GroupTwo'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'ChatContainer',
      component: ChatContainer
    },
    {
      path: '/group1',
      name: 'GroupOne',
      component: GroupOne
    },
    {
      path: '/group2',
      name: 'GroupTwo',
      component: GroupTwo
    },
  ]
})
