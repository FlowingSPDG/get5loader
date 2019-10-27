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

      <tr v-for="(match, index) in matches" :key="index">
        <td><a :href="'/match/'+match.id">{{match.id}}</a></td>

        <td>
          {{GetTeamData(match.team1_id)}}
          <a :href="'/team/'+match.team1_id">{{GetTeamData(match.team1_id)}}</a>
        </td>

        <td>
          {{GetTeamData(match.team2_id)}}
          <a :href="'/team/'+match.team2_id">{{GetTeamData(match.team2_id)}}</a>
        </td>

        <td>
          {{ GetMatchStatusString(match.id) }}
        </td>

        {% if my_matches %}
        <td>{% if match.get_server() is not none   %} {{ match.get_server().get_display() }} {% endif %}</td>
        <td>
          {% if match.pending() or match.live() %}
          <a :href="'/match/'+match.id+'cancel'" class="btn btn-danger btn-xs align-right">Cancel</a>
          {% endif %}
        </td>
        {% else %}
        <td> <a :href="match.get_user().get_url()"> {{ match.get_user().name }} </a> </td>
        {% endif %}

      </tr>

    </tbody>
  </table>

  <!--{{ pagination_buttons(matches) }}-->

  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'matches',
  data () {
    return {
      my_matches:false,
      all_matches:false,
      matches:[],
      match_owner:{
        id:1,
        name:"hoge"
      }
    }
  },
  mounted () {
    axios.get('/api/v1/GetMatches').then((res) => {
      this.matches = res.data
      console.log(res.data)
    })
  },
  methods: {
    GetTeamData: function(teamid){
      axios.get(`/api/v1/team/${teamid}/GetTeamInfo`).then((res) => {
        console.log(res.data)
        return res.data
      })
    },
    GetMatchStatusString: function(matchid){
      axios.get(`/api/v1/match/${matchid}/GetStatusString`).then((res) => {
        console.log(res.data)
        return res.data
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
