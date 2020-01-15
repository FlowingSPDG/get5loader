<template>
<div id="content">

  <div class="panel panel-default">
    <div class="panel-heading">User information</div>
    <div class="panel-body">
      Name: {{displaying_user.name}}<br>
      Steam account: <a :href="GetSteamURL(displaying_user.steam_id)"> {{displaying_user.steam_id}}</a> <br>
      Teams saved: <router-link :to="'/teams/'+displaying_user.id"> {{displaying_user.teams.length}}</router-link> <br>
      Matches created: <router-link :to="'/matches/'+displaying_user.id">{{displaying_user.matches.length}}</router-link> <br>
    </div>
  </div>

  <div class="panel panel-default" v-if="displaying_user">
    <div class="panel-heading">Recent Matches</div>
    <span class="panel-body" v-for="(match, index) in matches" :key=match.id>
        <router-link :to="'/match/'+match.id">#{{match.id}}</router-link> {{ matchdata[index] }}
    </span>
  </div>

</div>
</template>

<script>
export default {
  name: 'User',
  data () {
    return {
      displaying_user: {
        id: '',
        name: '',
        steam_id: '',
        matches: [],
        teams: [],
        servers: []
      },
      matches: [],
      matchdata: []
    }
  },
  async created () {
    let self = this
    const user = await this.GetUserData(this.$route.params.userid)
    self.displaying_user = user
    self.matches = user.matches
    for (let i = 0; i < self.matches.length; i++) {
      const res = await self.get_vs_match_result(self.matches[i])
      self.matchdata.push(res)
    }
  },
  methods: {
    GetSteamURL: function (steamid) {
      return `https://steamcommunity.com/profiles/${steamid}`
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
