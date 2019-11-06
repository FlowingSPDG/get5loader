<template>
<div id="content">

  <div class="panel panel-default">
    <div class="panel-heading">User information</div>
    <div class="panel-body">
      Name: {{displaying_user.name}}<br>
      Steam account: <a :href="GetSteamURL(displaying_user.steam_id)"> {{displaying_user.steam_id}}</a> <br>
      Teams saved: <a :href="'/teams/'+displaying_user.id"> {{displaying_user.teams.length}}</a> <br>
      Matches created: <a :href="'/matches/'+displaying_user.id"> {{displaying_user.matches.length}}</a> <br>
    </div>
  </div>

  <div class="panel panel-default">
    <div class="panel-heading">Recent Matches</div>
    <div class="panel-body">
      {% for match in displaying_user.get_recent_matches() %}
        <a :href="'/match/'+match.id">#{{match.id}}</a>: {{ match.get_vs_string() }}
        <br>
      {% endfor %}
    </div>
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
          matches:[]        
      }
    }
  },
  created () {
      this.GetUserData(this.$route.query.userid)
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
