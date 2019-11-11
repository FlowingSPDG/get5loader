<template>
<div id="content">

  <h1 v-if="my_matches">Your teams</h1>
  <h1 v-else-if="owner">Teams for <a :href="'/user/'+owner.id"> {{ owner.name }}</a></h1>

  <ul class="list-group">
    <li class="list-group-item" v-if="teams.length == 0">
    No teams found.
    </li>

    <li class="list-group-item" v-else v-for="(team,index) in teams" :key="index">

      <img :src="get_flag_link(team)" />
      <router-link :to="'/team/'+team.id" class="col-sm-offset-1">{{team.name}}</router-link>

      <div class="pull-right" v-if="CheckTeamDeletable(team)">
        <a :href="'/team'+team.id+'/delete'" class="btn btn-danger btn-xs">Delete</a>
      </div>

      <div class="pull-right" v-if="CheckTeamEditable(team)">
        <a :href="'/team/'+team.id+'/edit'" class="btn btn-primary btn-xs">Edit</a>
      </div>

    </li>
  </ul>

</div>
</template>

<script>
export default {
  name: 'Teams',
  data () {
    return {
      user:{},
      my_matches:false,
      teams:[],
      owner:{}
    }
  },
  created () {
    this.axios
      .get('/api/v1/CheckLoggedIn')
      .then((res) => {
        this.user = res;
        this.GetUserData(this.$route.params.userid).then((res) => {
          this.owner = res
          this.my_matches = this.$route.params.userid == res.id
          this.teams = res.teams
        })
      })
  },
  methods : {
    GetUserData: function(userid){
      return new Promise((resolve, reject) => {
        this.axios.get(`/api/v1/user/${userid}/GetUserInfo`).then((res) => {
          console.log(res.data)
          resolve(res.data)
        })
      })
    },
    get_flag_link : function(team){
      if(team.flag == ""){
        return `/static/img/_unknown.png`  
      }
      //return `<img src="/static/img/valve_flags/${team.flag}"  width="24" height="16">`
      return `/static/img/valve_flags/${team.flag}.png`
    },
    CheckTeamEditable: function(team){
        return team.user_id == this.user.id
    },
    CheckTeamDeletable: function(team){
        return team.user_id == this.user.id
    },

  }
}
</script>
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

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s
}
.fade-enter, .fade-leave {
  opacity: 0
}
</style>
