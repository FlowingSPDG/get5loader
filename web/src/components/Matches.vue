<template>
  <div id="content" class="matches">
    <h1 v-if="my_matches">Your matches</h1>
    <h1 v-else-if="all_matches">All matches</h1>
    <h1 v-else>Matches for <a :href="'/user/'+match_owner.id"> {{ match_owner.name }}</a></h1>

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
        <td v-if="match"><a :href="'/match/'+match.id">{{match.id}}</a></td>

        <td v-if="matchinfo[match.id]">
          {{matchinfo[match.id].team1.flag}}
          <router-link :to="'/team?teamid='+match.team1_id">{{matchinfo[match.id].team1.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          {{matchinfo[match.id].team2.flag}}
          <router-link :to="'/team?teamid='+match.team2_id">{{matchinfo[match.id].team2.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          {{ matchinfo[match.id].status }}
        </td>

        <td v-if="my_matches && matchinfo[match.id].server">{{ matchinfo[match.id].server.display }} </td>
        <td v-if="my_matches && matchinfo[match.id]">
          <a v-if="(match.pending || match.live)" :href="'/match/'+match.id+'cancel'" class="btn btn-danger btn-xs align-right">Cancel</a>
        </td>
        <td v-if="!my_matches && matchinfo[match.id]"> <a :href="'/user/'+matchinfo[match.id].user.id"> {{ matchinfo[match.id].user.name }} </a> </td>

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
      my_matches:false,
      all_matches:false,
      matches:[],
      matchinfo:{},
      match_owner:{
        id:1,
        name:"hoge"
      },
      teamdatas:{},
      userdatas:{},
      serverdatas:{},
    }
  },
  created () {
    this.GetMatches().then((res) => {
      console.log(res)
      for(let i=0;i<res.length;i++){
        this.GetMatchInfo(res[i].id)
      }
    })
  },
  methods: {
    GetMatches: function(){
      return new Promise((resolve, reject) => {
        this.axios.get('/api/v1/GetMatches').then(res => {
          this.matches = res.data
          /*for(let i=0;i<res.data.length;i++){
            this.$set(this.matches, i, res.data[i])
          }*/
          resolve(res.data)
        })
      })
    },
    GetTeamData: function(teamid){
      return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`).then((res) => {
          this.$set(this.teamdatas,teamid,res.data)
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    GetUserData: function(userid){
      return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/user/${userid}/GetUserInfo`).then((res) => {
          this.$set(this.userdatas,userid,res.data)
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    GetServerData: function(serverid){
      return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/server/${serverid}/GetServerInfo`).then((res) => {
          this.$set(this.serverdatas,serverid,res.data)
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    GetMatchInfo: function(matchid){
      this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`).then((res) => {
        this.$set(this.matchinfo,matchid,res.data)
        console.log(res.data)
        return(res.data)
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
