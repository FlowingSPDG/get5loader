<template>
  
  <div class="container">
    <h1 v-if="team">
      {{ team.flag }} {{ team.name }} {{ team.logo }}
      {% if team.can_edit(user) %}
      <div class="pull-right">
        <a :href="'/team/'+team.id+'/edit'" class="btn btn-primary btn-xs">Edit</a>
      </div>
      {% endif %}
    </h1>

    <br>

    <div class="panel panel-default">
      <div class="panel-heading">Players</div>
      <div class="panel-body"  v-if="team">
          {% for auth,name in team.get_players() %}
          <a :href="'http://steamcommunity.com/profiles/'+auth" class="col-sm-offset-0"> {{auth}}</a>
          {% if name %}
          {{name}}
          {% endif %}
          <br>
          {% endfor %}
      </div>
    </div>


    <div class="panel panel-default">
      <div class="panel-heading">Recent Matches</div>
        <div class="panel-body"  v-if="team">
          {% for match in team.get_recent_matches() %}
            <a :href="'/match/'+match.id">#{{match.id}}</a>: {{ team.get_vs_match_result(match.id) }}
            <br>
          {% endfor %}
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      team: {},
      teamdatas: {}
    }
  },
  created () {
      this.GetTeamData(this.$route.query.teamid)
  },
  methods: {
    GetTeamData: function(teamid){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/team/${teamid}/GetTeamInfo`).then((res) => {
        this.$set(this.teamdatas,teamid,res.data)
        this.team = this.teamdatas[0]
        console.log(res.data)
        resolve(res.data)
      })
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
