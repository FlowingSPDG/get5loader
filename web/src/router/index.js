import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'
import VueAxios from 'vue-axios'

import Matches from '@/components/Matches'
import Match from '@/components/Match'
import MatchCreate from '@/components/MatchCreate'
import MyMatches from '@/components/MyMatches'
import MyServers from '@/components/MyServers'
import Metrics from '@/components/Metrics'
import Team from '@/components/Team'
import Teams from '@/components/Teams'
import TeamCreate from '@/components/TeamCreate'
import ServerCreate from '@/components/ServerCreate'
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
      path: '/match/create',
      name: 'Create Match',
      component: MatchCreate
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
      path: '/team/create',
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
      path: '/server/create',
      name: 'Create Server',
      component: ServerCreate
    }
  ]
})
