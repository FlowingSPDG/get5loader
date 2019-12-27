import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'
import VueAxios from 'vue-axios'

import Matches from '@/components/Matches'
import Match from '@/components/Match'
import MyMatches from '@/components/MyMatches'
import MyServers from '@/components/Myservers'
import Metrics from '@/components/Metrics'
import Team from '@/components/Team'
import Teams from '@/components/Teams'
import TeamCreate from '@/components/TeamCreate'
import User from '@/components/User'

Vue.use(VueAxios, axios)
Vue.use(Router)

export default new Router({
  mode: 'hash',
  routes: [
    {
      path: '/',
      redirect: '/matches'
    },
    {
      path: '/matches',
      name: 'Matches',
      component: Matches
    },
    {
      path: '/matches/:userid',
      name: 'Matches',
      component: Matches
    },
    {
      path: '/myservers',
      name: 'My Servers',
      component: MyServers
    },
    {
      path: '/mymatches',
      name: 'My matches',
      component: MyMatches
    },
    {
      path: '/match/create', // TODO
      name: 'Create Match',
      component: Matches
    },
    {
      path: '/match/:matchid',
      name: 'Match',
      component: Match
    },
    {
      path: '/metrics',
      name: 'Metrics',
      component: Metrics
    },
    {
      path: '/team/create', // TODO
      name: 'Create Team',
      component: TeamCreate
    },
    {
      path: '/team/:teamid',
      name: 'Team',
      component: Team
    },
    {
      path: '/teams/:userid',
      name: 'Teams',
      component: Teams
    },
    {
      path: '/user/:userid',
      name: 'User',
      component: User
    },
    {
      path: '/servers/create', // TODO
      name: 'Create Server',
      component: User
    }
  ]
})
