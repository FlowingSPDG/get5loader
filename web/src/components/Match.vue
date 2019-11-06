<template>
<div>

<div id="content">

    <div class="container" v-loading="loading" v-cloak>

        <h1>
            <img :src="get_logo_or_flag_link(team1,team2).team1" /> <a :href="'/team?teamid='+team1.id"> {{team1.name}}</a>
            {{ matchdata.team1_score }}
            {{ score_symbol(matchdata.team1_score, matchdata.team2_score) }}
            {{ matchdata.team2_score }}
            <img :src="get_logo_or_flag_link(team1,team2).team2" /> <a :href="'/team?teamid='+team2.id"> {{team2.name}}</a>

            <div class="dropdown dropdown-header pull-right" v-if="user.adminaccess == true && matchdata.live && match.pending">
                <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                    Admin tools
                    <span class="caret"></span>
                </button>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
                    <li v-if="matchdata.live"><a id="pause" :href="this.$route.path+'/pause'">Pause match</a></li>
                    <li v-if="matchdata.live"><a id="unpause" :href="this.$route.path+'/unpause'">Unpause match</a></li>
                    <li><a id="addplayer_team1" href="#">Add player to team1</a></li>
                    <li><a id="addplayer_team2" href="#">Add player to team2</a></li>
                    <li><a id="addplayer_spec" href="#">Add player to specator list</a></li>
                    <li><a id="rcon_command" href="#">Send rcon command</a></li>
                    <li role="separator" class="divider"></li>
                    <li><a id="backup_manager" :href="this.$route.path+'/backup'">Load a backup file</a></li>
                    <li><a :href="this.$route.path+'/cancel'">Cancel match</a></li>
                </ul>
            </div>

        </h1>

        <br>
        <div class="alert alert-danger" role="alert" v-if="matchdata.cancelled">
            <span class="glyphicon glyphicon-exclamation-sign" aria-hidden="true"></span>
            <span class="sr-only">Error:</span>
            This match has been cancelled.
        </div>

        <div class="alert alert-warning" role="alert" v-if="matchdata.forfeit">
            <span class="glyphicon glyphicon-exclamation-sign" aria-hidden="true"></span>
            <span class="sr-only">Error:</span>
            This match was forfeit by {{get_loser(matchdata)}}.
        </div>

        <p v-if="matchdata.start_time != '0001-01-01T00:00:00Z'">Started at {{ matchdata.start_time }}</p>
        <div class="panel panel-default" role="alert" v-else>
            <div class="panel-body">
                This match is pending start.
            </div>
        </div>

        <p v-if="matchdata.end_time != '0001-01-01T00:00:00Z'">Ended at {{ matchdata.end_time }}</p>

        <div v-for="map_stats in matchdata.map_stats" :key="map_stats.id">
        <br>
        <div class="panel panel-primary">
            <div class="panel-heading">
                Map {{map_stats.map_number + 1}}: {{ map_stats.map_name }},
                {{team1.name}} {{ score_symbol(map_stats.team1_score, map_stats.team2_score) }} {{team2.name}},
                {{map_stats.team1_score}}:{{map_stats.team2_score}}
            </div>

            <div class="panel-body">
                <p>Started at {{ map_stats.start_time }}</p>

                <p v-if="map_stats.end_time != '0001-01-01T00:00:00Z'">Ended at {{ map_stats.end_time }}</p>

                <table class="table table-hover">
                    <thead>
                        <tr>
                            <th>Player</th>
                            <th class="text-center">Kills</th>
                            <th class="text-center">Deaths</th>
                            <th class="text-center">Assists</th>
                            <th class="text-center">Flash assists</th>
                            <th class="text-center">1v1</th>
                            <th class="text-center">1v2</th>
                            <th class="text-center">1v3</th>
                            <th class="text-center">Rating</th>
                            <th class="text-center"><acronym title="Frags per round">FPR</acronym></th>
                            <th class="text-center"><acronym title="Average damage per round">ADR</acronym></th>
                            <th class="text-center"><acronym title="Headshot percentage">HSP</acronym></th>
                        </tr>
                    </thead>
                    <tbody>
                        <td> <b>{{ team1.name }}</b> </td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        
                        <tr v-for="player in matchdata.team1_player_stats[map_stats.id]" :key="player.id">
                            <td v-if="player.roundsplayed"> <a :href="GetSteamURL(player.steam_id)"> {{ player.name }} </a></td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.kills }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.deaths }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.assists }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.flashbang_assists }} </td>
                        
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v1 }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v2 }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v3 }} </td>
                        
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.rating }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.fpr }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.adr }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.hsp }} </td>
                        </tr>
                        
                        
                        <td> <b>{{ team2.name }}</b> </td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        <td></td>
                        
                        <tr v-for="player in matchdata.team2_player_stats[map_stats.id]" :key="player.id">
                            <td v-if="player.roundsplayed"> <a :href="GetSteamURL(player.steam_id)"> {{ player.name }} </a></td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.kills }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.deaths }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.assists }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.flashbang_assists }} </td>
                        
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v1 }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v2 }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.v3 }} </td>
                        
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.rating }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.fpr }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.adr }} </td>
                            <td v-if="player.roundsplayed" class="text-center"> {{ player.hsp }} </td>
                        </tr>
                    </tbody>

                </table>
            </div>

        </div>
        </div>

    </div>

    <br>
</div>

<!--script>
    jQuery("#addplayer_team1").click(function (e) {
        var input = prompt("Please enter a steamid to add to {{team1.name}}", "");
        if (input != null) {
            window.location.href = "{{request.path}}/adduser?team=team1&auth=" + encodeURIComponent(input);
        }
    });

    jQuery("#addplayer_team2").click(function (e) {
        var input = prompt("Please enter a steamid to add to {{team2.name}}", "");
        if (input != null) {
            window.location.href = "{{request.path}}/adduser?team=team2&auth=" + encodeURIComponent(input);
        }
    });

    jQuery("#addplayer_spec").click(function (e) {
        var input = prompt("Please enter a steamid to add to the spectators list", "");
        if (input != null) {
            window.location.href = "{{request.path}}/adduser?team=spec&auth=" + encodeURIComponent(input);
        }
    });

    jQuery("#rcon_command").click(function (e) {
        var input = prompt("Enter a command to send", "");
        if (input != null) {
            window.location.href = "{{request.path}}/rcon?command=" + encodeURIComponent(input);
        }
    });
</script>-->
</div>
</template>

<script>
export default {
  name: 'Match',
  data () {
    return {
        loading:true,
        matchdata:{
            id:0,
            user_id:0,
            team1: {
                "id":0,
                "user_id":0,
                "name":"LOADING...",
                "tag":"",
                "flag":"",
                "logo":"",
                "steamids":[],
                "public_team":false
                },
            team2: {
                "id":0,
                "user_id":0,
                "name":"LOADING...",
                "tag":"",
                "flag":"",
                "logo":"",
                "steamids":[],
                "public_team":false
                },
            winner:0,
            cancelled:false,
            start_time:"",
            end_time:"",
            max_maps:0,
            title:"",
            skip_veto:false,
            veto_mappool:[],
            team1_score:0,
            team2_score:0,
            team1_string:"",
            team2_string:"",
            forfeit:false,
            map_stats:{},
            team1_player_stats:[],
            team2_player_stats:[],
            server:{
                id: 0,
		        user_id: 0,
		        in_use: false,
		        ip_string: "",
		        port: 0,
		        display: "",
		        public_server: false
            },
            user:{
                id: 0,
		        steam_id: "",
		        name: "",
		        admin: false,
		        servers: null,
		        teams: null,
                matches: null
            },
            pending:false,
            live:false,
            status:""
        },
        user: {
          isLoggedIn:false,
          adminaccess:false,
          steamid:"",
          userid:""
        },
        team1:{
            "id": 0,
	        "user_id": 0,
	        "name":"LOADING...",
	        "tag": "",
	        "flag": "",
	        "logo": "",
	        "steamids": [],
	        "public_team": false
        },
        team2:{
            "id": 0,
	        "user_id": 0,
	        "name":"LOADING...",
	        "tag": "",
	        "flag": "",
	        "logo": "",
	        "steamids": [],
	        "public_team": false
        },
    }
  },
  created () {
    this.GetMatchData(this.$route.query.matchid).then((res)=>{
        console.log("GetMatchData")
        console.log(res)
        for (let i=0;i<res.map_stats.length;i++){
            console.log("GetPlayerStats")
            this.GetPlayerStats(res.id,res.map_stats[i].id)
        }
        console.log("GetTeam1Data")
        this.GetTeam1Data(res.team1.id)
        console.log("GetTeam2Data")
        this.GetTeam2Data(res.team2.id)
        this.loading = false;
    })
    this.axios
      .get('/api/v1/CheckLoggedIn')
      .then((res) => {
          console.log(res.data)
          this.user = res.data
          //this.Editable = this.CheckTeamEditable(this.$route.query.teamid,this.user.userid) // TODO
      })
  },
  methods: {
    GetTeam1Data: function(team1id){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/team/${team1id}/GetTeamInfo`).then((res) => {
        this.team1 = res.data
        resolve(res.data)
      })
     })
    },
    GetTeam2Data: function(team2id){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/team/${team2id}/GetTeamInfo`).then((res) => {
        this.team2 = res.data
        resolve(res.data)
      })
     })
    },
    GetMatchData: function(matchid){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`).then((res) => {
        this.matchdata = res.data
        resolve(res.data)
      })
     })
    },
    GetMapStat: function(matchid){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`).then((res) => {
        this.matchdata.map_stats.push(res.data)
        console.log(res.data)
        resolve(res.data)
      })
     })
    },
    GetPlayerStats: function(matchid,mapid){
     return new Promise((resolve, reject) => {
      this.axios.get(`/api/v1/match/${matchid}/GetPlayerStatInfo?mapID=${mapid}`).then((res) => {
        console.log(res.data)
        if(!this.matchdata.team1_player_stats){
          this.matchdata.team1_player_stats = {}
          this.matchdata.team1_player_stats[mapid] = []
        }
        if(!this.matchdata.team2_player_stats){
          this.matchdata.team2_player_stats = {}
          this.matchdata.team2_player_stats[mapid] = []
        }

        let team1stats = res.data.filter(player => player.team_id == this.matchdata.team1.id)
        let team2stats = res.data.filter(player => player.team_id == this.matchdata.team2.id)

        for(let i=0;i<team1stats.length;i++){
            this.$set(this.matchdata.team1_player_stats, mapid, team1stats);
        }
        for(let i=0;i<team2stats.length;i++){
            this.$set(this.matchdata.team2_player_stats, mapid, team2stats);
        }
        resolve(res.data)
      })
     })
    },
    GetSteamURL: function(steamid){
        return `https://steamcommunity.com/profiles/${steamid}`
    },
    get_logo_or_flag_link: function(team1,team2){ // get_logo_or_flag_link(team1)
        if (team1.logo && team2.logo){
            return {
                team1:get_logo_link(team1),
                team2:get_logo_link(team2)
            }
        } else {
            return {
                team1:this.get_flag_link(team1),
                team2:this.get_flag_link(team2)
            }
        }
    },
    get_logo_html : function(team){
        // TODO...
    },
    get_flag_link : function(team){
        if(team.flag == ""){
          return `/static/img/_unknown.png`  
        }
        //return `<img src="/static/img/valve_flags/${team.flag}"  width="24" height="16">`
        return `/static/img/valve_flags/${team.flag}.png`
    },
    score_symbol: function(score1,score2){
        if ( score1 > score2){
            return ">"
        } else {

        }if ( score1 < score2){
            return "<"
        } else{
            return "=="
        }
    },
    get_loser : function(matchdata){ // returns loser's teamname
        if (matchdata.team1_score > matchdata.team2_score){
            return matchdata.team2.name
        } else if (matchdata.team1_score < matchdata.team2_score){
            return matchdata.team1.name
        } else {
            return ""
        }
    },
    GetKDR : function(playerstat){
        if (playerstat.deaths == 0) {
	    	return playerstat.kills
	    }
	    return playerstat.kills / playerstat.deaths
    },
    GetHSP : function(playerstat){
        if (playerstat.deaths == 0) {
		    return playerstat.kills
	    }
	    return playerstat.headshot_kills / playerstat.kills * 100
    },
    GetADR : function(playerstat){
        if (playerstat.roundsplayed == 0) {
    		return 0.0
    	}
    	return playerstat.damage / playerstat.roundsplayed
    },
    GetFPR : function(playerstat){
        if (playerstat.roundsplayed == 0) {
    		return 0.0
    	}
    	return playerstat.kills / playerstat.roundsplayed
    },
    SendRCON: function(command){
        //TODO
        /*
        this.$notify.info({
          title: 'Info',
          message: 'This is an info message'
        });
      */
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>
h1,
h2 {
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
