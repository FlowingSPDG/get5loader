<template>

  <div class="container">
    <h1 v-cloak>
      <img :src="get_flag_link(team)" /> {{ team.name }} {{ team.logo }}
      <div class="pull-right" v-if="Editable == true">
        <router-link :to="'/team/'+team.id+'/edit'" class="btn btn-primary btn-xs">Edit</router-link>
      </div>
    </h1>

    <br>

    <div class="panel panel-default">
      <div class="panel-heading">Players</div>
      <div class="panel-body" v-cloak>
        <el-table :data="players">
          <el-table-column label="SteamID" prop="steamid" width="400">
            <template slot-scope="scope">
              <a :href="'https://steamcommunity.com/profiles/'+scope.row.steamid">{{scope.row.steamid}}</a>
            </template>
          </el-table-column>
          <el-table-column
            prop="name"
            label="Name"
            width="180">
          </el-table-column>
        </el-table>
      </div>
    </div>

    <div class="panel panel-default">
      <div class="panel-heading">Recent Matches</div>
        <div class="panel-body"  v-if="team">
          <div v-for="(match, index) in matches" :key="index" >
            <router-link :to="'/match/'+match.id">#{{match.id}}</router-link>: {{ matchdata[index] }}
            <br>
          </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Team',
  data () {
    return {
      team: {
        flag: '',
        name: '',
        logo: ''
      },
      matches: [],
      matchdata: [],
      players: [],
      teamdatas: {},
      user: {
        isLoggedIn: false,
        steamid: '',
        userid: ''
      },
      Editable: false
    }
  },
  async created () {
    const teamdataPromise = this.GetTeamData(this.$route.params.teamid)
    const matchesPromise = this.GetRecentMatches(this.$route.params.teamid)
    this.team = await teamdataPromise
    this.matches = await matchesPromise
    for (let i = 0; i < this.matches.length; i++) {
      if (!this.matchdata) {
        this.matchdata = []
      }
      let res = await this.get_vs_match_result(this.matches[i])
      this.matchdata.push(res)
    }
    for (let i = 0; i < this.team.steamids.length; i++) {
      this.GetSteamName(this.team.steamids[i])
    }
    const loggedin = await this.axios.get('/api/v1/CheckLoggedIn')
    this.user = loggedin.data
    this.Editable = this.CheckTeamEditable(this.user.userid)
  },
  methods: {
    async GetTeamData (teamid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`)
        resolve(res.data)
      })
    },
    async GetRecentMatches (teamid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/team/${teamid}/GetRecentMatches`)
        this.matches = res.data
        resolve(res.data)
      })
    },
    async GetSteamName (steamid) {
      let self = this
      if (steamid === '') {
        return
      }
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/GetSteamName?steamID=${steamid}`)
        self.players.push({ steamid: steamid, name: res.data })
        resolve(res.data)
      })
    },
    CheckTeamEditable: function (userid) {
      return this.team.user_id === userid
    },
    get_flag_link: function (team) {
      if (team.flag === '') {
        return `/img/_unknown.png`
      }
      // return `<img src="/img/valve_flags/${team.flag}"  width="24" height="16">`
      return `/img/valve_flags/${team.flag}.png`
    },
    async get_vs_match_result (match) {
      return new Promise(async (resolve, reject) => {
        let MyScore
        let OtherTeamScore
        let OtherTeam
        if (match.team1.id === this.$route.params.teamid) {
          MyScore = match.team1_score
          OtherTeamScore = match.team2_score
          OtherTeam = await this.GetTeamData(match.team2.id)
          // for a bo1 replace series score with the map score
          if (match.max_maps === 1) {
            if (match.map_stats.length === 1) {
              if (match.team1_id === self.id) {
                MyScore = match.map_stats[0].team1_score
                OtherTeamScore = match.map_stats[0].team2_score
              } else {
                MyScore = match.map_stats[0].team2_score
                OtherTeamScore = match.map_stats[0].team1_score
              }
            }
          }
          if (match.live) {
            let r = `Live, ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          }
          if (MyScore < OtherTeamScore) {
            let r = `Lost ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          } else if (MyScore > OtherTeamScore) {
            let r = `Won ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          } else {
            let r = `Tied ${OtherTeamScore}:${MyScore} vs ${OtherTeam.name}`
            resolve(r)
          }
        } else {
          MyScore = match.team2_score
          OtherTeamScore = match.team1_score
          OtherTeam = await this.GetTeamData(match.team1.id)
          // for a bo1 replace series score with the map score
          if (match.max_maps === 1) {
            if (match.map_stats.length === 1) {
              if (match.team1_id === self.id) {
                MyScore = match.map_stats[0].team1_score
                OtherTeamScore = match.map_stats[0].team2_score
              } else {
                MyScore = match.map_stats[0].team2_score
                OtherTeamScore = match.map_stats[0].team1_score
              }
            }
          }
          if (match.live) {
            let r = `Live, ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          }
          if (MyScore < OtherTeamScore) {
            let r = `Lost ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          } else if (MyScore > OtherTeamScore) {
            let r = `Won ${MyScore}:${OtherTeamScore} vs ${OtherTeam.name}`
            resolve(r)
          } else {
            let r = `Tied ${OtherTeamScore}:${MyScore} vs ${OtherTeam.name}`
            resolve(r)
          }
        }
      })
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
