<template>
  
  <div class="container">
    <h1 v-cloak>
      <img :src="get_flag_link(team)" /> {{ team.name }} {{ team.logo }}
      <div class="pull-right" v-if="Editable == true">
        <a :href="'/team/'+team.id+'/edit'" class="btn btn-primary btn-xs">Edit</a>
      </div>
    </h1>

    <br>

    <div class="panel panel-default">
      <div class="panel-heading">Players</div>
      <div class="panel-body" v-cloak>
        <div v-for="player in players" :key="player.steamid">
          <a :href="'http://steamcommunity.com/profiles/'+player.steamid" class="col-sm-offset-0"> {{player.steamid}}</a>
          <div>{{player.name}}</div>
          <br>
        </div>
      </div>
    </div>


    <div class="panel panel-default">
      <div class="panel-heading">Recent Matches</div>
        <div class="panel-body"  v-if="team">
          <!--{% for match in team.get_recent_matches() %}
            <a :href="'/match/'+match.id">#{{match.id}}</a>: {{ team.get_vs_match_result(match.id) }}
            <br>
          {% endfor %}-->
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
    this.GetTeamData(this.$route.query.teamid).then(()=>{
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
        this.team = res.data
        console.log(res.data)
        resolve(res.data)
      })
    })
  },
  GetRecentMatches: function(teamid) {
      
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
        //return `<img src="/static/img/valve_flags/${team.flag}"  width="24" height="16">`
        return `/static/img/valve_flags/${team.flag}.png`
    },
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
