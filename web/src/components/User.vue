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
      displaying_user:{
          id:"",
          name:"",
          steam_id:"",
          matches:[],
          teams:[],
          servers:[]
      },
      matches:[],
      matchdata:[]
    }
  },
  created () {
    let self = this
      this.GetUserData(this.$route.params.userid).then((user) => {
        self.displaying_user = user;
        self.matches = user.matches
        for(let i=0;i<self.matches.length;i++){
          self.get_vs_match_result(self.matches[i]).then((res) => {
            self.matchdata.push(res)
          })
        }
      })
  },
  methods: {
    GetSteamURL: function(steamid){
      return `https://steamcommunity.com/profiles/${steamid}`
    },
    GetUserData: function(userid){
      return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/user/${userid}/GetUserInfo`).then((res) => {
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    GetTeamData: function(teamid){
     return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`).then((res) => {
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    get_vs_match_result: function(match){
    return new Promise((resolve, reject) => {
    console.log("get_vs_match_result")
    console.log(match)
    let my_score
    let other_team_score
    let other_team;
    if (match.team1.id == this.$route.params.teamid){
        my_score = match.team1_score
        other_team_score = match.team2_score
        this.GetTeamData(match.team2.id).then((res) => {
          other_team = res;
          //for a bo1 replace series score with the map score
    if (match.max_maps == 1){
      if (match.map_stats.length == 1){
        if (match.team1_id == self.id){
          my_score = match.map_stats[0].team1_score
          other_team_score = match.map_stats[0].team2_score
        }
        else{
          my_score = match.map_stats[0].team2_score
          other_team_score = match.map_stats[0].team1_score      
        }
      }
    }
    console.log("other_team")
    console.log(other_team)
    if (match.live){
      let r = `Live, ${my_score}:${other_team_score} vs ${other_team.name}` 
      console.log(r)
      resolve(r)
    }
    if (my_score < other_team_score){
      let r = `Lost ${my_score}:${other_team_score} vs ${other_team.name}`
      console.log(r)
      resolve(r)
    }
    else if(my_score > other_team_score){
       let r = `Won ${my_score}:${other_team_score} vs ${other_team.name}`
       console.log(r)
      resolve(r)
    }
    else{
      let r = `Tied ${other_team_score}:${my_score} vs ${other_team.name}`
      console.log(r)
      resolve(r)
    }
        })
    }
    else{
      my_score = match.team2_score
      other_team_score = match.team1_score
      this.GetTeamData(match.team1.id).then((res) => {
        other_team = res;
        //for a bo1 replace series score with the map score
    if (match.max_maps == 1){
      if (match.map_stats.length == 1){
        if (match.team1_id == self.id){
          my_score = match.map_stats[0].team1_score
          other_team_score = match.map_stats[0].team2_score
        }
        else{
          my_score = match.map_stats[0].team2_score
          other_team_score = match.map_stats[0].team1_score      
        }
      }
    }
    console.log("other_team")
    console.log(other_team)
    if (match.live){
      let r = `Live, ${my_score}:${other_team_score} vs ${other_team.name}` 
      console.log(r)
      resolve(r)
    }
    if (my_score < other_team_score){
      let r = `Lost ${my_score}:${other_team_score} vs ${other_team.name}`
      console.log(r)
      resolve(r)
    }
    else if(my_score > other_team_score){
       let r = `Won ${my_score}:${other_team_score} vs ${other_team.name}`
       console.log(r)
      resolve(r)
    }
    else{
      let r = `Tied ${other_team_score}:${my_score} vs ${other_team.name}`
      console.log(r)
      resolve(r)
    }
      })
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
