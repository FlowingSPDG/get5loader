<template>
  <div id="content" class="matches">
    <h1 v-if="my_matches">Your matches</h1>
    <h1 v-else-if="all_matches">All matches</h1>
    <h1 v-else>Matches for <router-link :to="'/user/'+match_owner.id">{{match_owner.name}}</router-link></h1>

  <table class="table table-striped">
    <thead>
      <tr>
        <th>Match ID</th>
        <th>Team 1</th>
        <th>Team 2</th>
        <th>Status</th>
        <th v-if="my_matches">Server</th>
        <th v-if="my_matches"></th>
        <th v-else>Owner</th>
      </tr>
    </thead>
    <tbody>

      <tr v-for="(match, index) in matches" :key="index" align="left">
        <td v-if="match"><router-link :to="'/match/'+match.id">{{match.id}}</router-link></td>

        <td v-if="matchinfo[match.id]">
          <img :src="get_flag_link(matchinfo[match.id].team1)" />
          <router-link :to="'/team/'+match.team1_id">{{matchinfo[match.id].team1.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          <img :src="get_flag_link(matchinfo[match.id].team2)" />
          <router-link :to="'/team/'+match.team2_id">{{matchinfo[match.id].team2.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          {{ matchinfo[match.id].status }}
        </td>

        <td v-if="my_matches && matchinfo[match.id]">{{ matchinfo[match.id].server.display }} </td>
        <td v-if="my_matches && matchinfo[match.id]">
          <a v-if="(match.pending || match.live)" :href="'/match/'+match.id+'cancel'" class="btn btn-danger btn-xs align-right">Cancel</a>
        </td>
        <td v-if="!my_matches && matchinfo[match.id]">
          <router-link :to="'/user/='+matchinfo[match.id].user.id">{{ matchinfo[match.id].user.name }}</router-link>
        </td>

      </tr>

    </tbody>
  </table>

  <!--{{ pagination_buttons(matches) }}-->

  </div>
</template>

<script>
export default {
  name: 'matches',
  data () {
    return {
      user: {
        isLoggedIn: false,
        steamid: '',
        userid: ''
      },
      my_matches: true,
      all_matches: false, // TODO
      matches: [],
      matchinfo: {},
      match_owner: { // TODO
        id: 1,
        name: 'hoge'
      },
      teamdatas: {},
      userdatas: {},
      serverdatas: {}
    }
  },
  async created () {
    const res = await this.axios.get('/api/v1/CheckLoggedIn')
    this.user = res.data
    const matches = await this.GetMatches(this.user.userid)
    console.log(matches)
    for (let i = 0; i < matches.length; i++) {
      this.GetMatchInfo(matches[i].id)
    }
  },
  methods: {
    async GetMatches (userid) {
      return new Promise(async (resolve, reject) => {
        if (!userid) {
          const res = await this.axios.get('/api/v1/GetMatches')
          this.matches = res.data
          resolve(res.data)
        } else {
          const res = await this.axios.get(`/api/v1/GetMatches?userID=${userid}`)
          this.matches = res.data
          resolve(res.data)
        }
      })
    },
    async GetTeamData (teamid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`)
        this.$set(this.teamdatas, teamid, res.data)
        console.log(res.data)
        resolve(res.data)
      })
    },
    async GetUserData (userid) {
      return new Promise(async (resolve, reject) => {
        const res = this.axios.get(`/api/v1/user/${userid}/GetUserInfo`)
        this.$set(this.userdatas, userid, res.data)
        console.log(res.data)
        resolve(res.data)
      })
    },
    async GetServerData (serverid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/server/${serverid}/GetServerInfo`)
        this.$set(this.serverdatas, serverid, res.data)
        console.log(res.data)
        resolve(res.data)
      })
    },
    async GetMatchInfo (matchid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`)
        this.$set(this.matchinfo, matchid, res.data)
        console.log(res.data)
        resolve(res.data)
      })
    },
    get_flag_link: function (team) {
      if (team.flag === '') {
        return `/static/img/_unknown.png`
      }
      // return `<img src="/static/img/valve_flags/${team.flag}"  width="24" height="16">`
      return `/static/img/valve_flags/${team.flag}.png`
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
