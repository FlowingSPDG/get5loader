<template>
  
  <div class="container">
    <h1 v-cloak>
      <img :src="get_flag_link(team)" /> {{ team.name }} {{ team.logo }}
      <div class="pull-right" v-if="Editable == true">
        <router-link :to="'/team?teamid='+team.id+'/edit'" "btn btn-primary btn-xs">Edit</router-link>
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
            <router-link :to="'/match?matchid='+match.id">#{{match.id}}</router-link>: {{ matchdata[index] }}
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
        flag:"",
        name:"",
        logo:"",
      },
      matches:[],
      matchdata:[],
      players:[],
      teamdatas: {},
      user: {
        isLoggedIn:false,
        steamid:"",
        userid:""
      },
      Editable:false
    }
  },
  created () {
    this.GetTeamData(this.$route.query.teamid).then((teamdata)=>{
      this.team = teamdata
      this.GetRecentMatches(this.$route.query.teamid).then((matches) => {
        this.matches = matches
        for(let i=0;i<this.matches.length;i++){
          if (!this.matchdata){
            this.matchdata = new Array
          }
          this.get_vs_match_result(this.matches[i]).then((res) => {
            this.matchdata.push(res)
          })
        }
      })
      for(let i=0;i<this.team.steamids.length;i++){
        this.GetSteamName(this.team.steamids[i])
      }
    })
    this.axios
      .get('/api/v1/CheckLoggedIn')
      .then((res) => {
          console.log(res.data)
          this.user = res.data
          this.Editable = this.CheckTeamEditable(this.user.userid)
      })
  },
  methods: {
    GetTeamData: function(teamid){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`).then((res) => {
        console.log(res.data)
        resolve(res.data)
      })
    })
  },
  GetRecentMatches: function(teamid) {
    return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/team/${teamid}/GetRecentMatches`).then(res => {
        this.matches = res.data
          resolve(res.data)
      })
    })
  },
  GetSteamName: function(steamid){
    var self = this
    if(steamid == ""){
      return
    }
    return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/GetSteamName?steamID=${steamid}`).then((res) => {
        console.log(res.data)
        console.log(self.team)
        self.players.push({steamid:steamid,name:res.data})
        resolve(res.data)
      })
    })
  },
  CheckTeamEditable: function(userid){
    return this.team.user_id == userid
  },
  get_flag_link : function(team){
        if(team.flag == ""){
          return `/static/img/_unknown.png`  
        }
        //return `<img src="/static/img/valve_flags/${team.flag}"  width="24" height="16">`
        return `/static/img/valve_flags/${team.flag}.png`
  },
  get_vs_match_result: function(match){
    return new Promise((resolve, reject) => {
    console.log("get_vs_match_result")
    console.log(match)
    let my_score
    let other_team_score
    let other_team;
    if (match.team1.id == this.$route.query.teamid){
        my_score = match.team1_score
        other_team_score = match.team2_score
        this.GetTeamData(match.team2.id).then((res) => {
          other_team = res;
          //for a bo1 replace series score with the map score
    if (match.max_maps == 1){
      mapstat = match.map_stats[0]
      if (mapstat){
        if (match.team1_id == self.id){
          my_score = mapstat.team1_score
          other_team_score = mapstat.team2_score
        }
        else{
          my_score = mapstat.team2_score
          other_team_score = mapstat.team1_score      
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
      mapstat = match.map_stats[0]
      if (mapstat){
        if (match.team1_id == self.id){
          my_score = mapstat.team1_score
          other_team_score = mapstat.team2_score
        }
        else{
          my_score = mapstat.team2_score
          other_team_score = mapstat.team1_score      
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
