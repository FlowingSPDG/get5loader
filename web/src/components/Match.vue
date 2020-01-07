<template>
<div>

<div id="content">

    <div class="container" v-loading="loading" v-cloak v-if="matchdata">
        <h1>
            <img :src="get_logo_or_flag_link(team1,team2).team1" /> <router-link :to="'/team/'+team1.id"> {{team1.name}}</router-link>
            {{ matchdata.team1_score }}
            {{ score_symbol(matchdata.team1_score, matchdata.team2_score) }}
            {{ matchdata.team2_score }}
            <img :src="get_logo_or_flag_link(team1,team2).team2" /> <router-link :to="'/team/'+team2.id"> {{team2.name}}</router-link>
              <el-dropdown v-if="AdminToolsAvailable()" @command="handleCommand">
                <el-button type="primary">Admin tools<i class="el-icon-arrow-down el-icon--right"></i></el-button>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item v-if="matchdata.live">Pause match</el-dropdown-item><br>
                  <el-dropdown-item v-if="matchdata.live">Unpause match</el-dropdown-item><br>
                  <el-dropdown-item command="AddPlayerToTeam1">Add player to team1</el-dropdown-item><br>
                  <el-dropdown-item command="AddPlayerToTeam2">Add player to team2</el-dropdown-item><br>
                  <el-dropdown-item command="AddPlayerToSpec">Add player to specator list</el-dropdown-item><br>
                  <el-dropdown-item command="rcon_command">Send rcon command</el-dropdown-item><br>
                  <el-dropdown-item devided command="backup_manager">Load a backup file</el-dropdown-item><br>
                  <el-dropdown-item devided command="cancelmatch">Cancel match</el-dropdown-item><br>
                </el-dropdown-menu>
              </el-dropdown>
        </h1>

        <el-dialog title="Select Backup file" :visible.sync="chose_backup" width="30%">
          <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="Backup Files">
              <el-select v-model="chosed_backup">
                <el-option
                  v-for="(backup,index) in backups"
                  :key="index"
                  :label="backup"
                  :value="backup">
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>

          <span slot="footer" class="dialog-footer">
            <el-button @click="chose_backup = !chose_backup">Cancel</el-button>
            <el-button type="primary" @click="SendBackup">Confirm</el-button>
          </span>
        </el-dialog>

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

        <div class="panel panel-default" role="alert" v-if="matchdata.start_time == '0001-01-01T00:00:00Z'">
            <div class="panel-body">
                This match is pending start.
            </div>
        </div>

        <el-timeline>
          <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :timestamp="activity.timestamp"
            :icon="activity.icon"
            :color="activity.color"
            >
            {{activity.content}}
          </el-timeline-item>
        </el-timeline>

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
                    <tbody v-if="matchdata.team1_player_stats && matchdata.team2_player_stats">
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
      loading: true,
      backups: [],
      activities: [], // {timestamp:"",content:""}
      chosed_backup: '',
      chose_backup: false,
      matchdata: {
        id: 0,
        user_id: 0,
        team1: {
          'id': 0,
          'user_id': 0,
          'name': 'LOADING...',
          'tag': '',
          'flag': '',
          'logo': '',
          'steamids': [],
          'public_team': false
        },
        team2: {
          'id': 0,
          'user_id': 0,
          'name': 'LOADING...',
          'tag': '',
          'flag': '',
          'logo': '',
          'steamids': [],
          'public_team': false
        },
        winner: 0,
        cancelled: false,
        start_time: '',
        end_time: '',
        max_maps: 0,
        title: '',
        skip_veto: false,
        veto_mappool: [],
        team1_score: 0,
        team2_score: 0,
        team1_string: '',
        team2_string: '',
        forfeit: false,
        map_stats: [],
        team1_player_stats: [],
        team2_player_stats: [],
        server: {
          id: 0,
          user_id: 0,
          in_use: false,
          ip_string: '',
          port: 0,
          display: '',
          public_server: false
        },
        user: {
          id: 0,
          steam_id: '',
          name: '',
          admin: false,
          servers: null,
          teams: null,
          matches: null
        },
        pending: false,
        live: false,
        status: ''
      },
      user: {
        isLoggedIn: false,
        adminaccess: false,
        steamid: '',
        userid: ''
      },
      team1: {
        'id': 0,
        'user_id': 0,
        'name': 'LOADING...',
        'tag': '',
        'flag': '',
        'logo': '',
        'steamids': [],
        'public_team': false
      },
      team2: {
        'id': 0,
        'user_id': 0,
        'name': 'LOADING...',
        'tag': '',
        'flag': '',
        'logo': '',
        'steamids': [],
        'public_team': false
      }
    }
  },
  async created () {
    this.matchdata = await this.GetMatchData(this.$route.params.matchid)
    if (this.matchdata.start_time !== '0001-01-01T00:00:00Z') {
      this.activities.push({ timestamp: this.matchdata.start_time, content: 'Match Started', icon: 'el-icon-plus', color: '#0bbd87' })
    }
    for (let i = 0; i < this.matchdata.map_stats.length; i++) {
      if (this.matchdata.map_stats[i].start_time !== '0001-01-01T00:00:00Z') {
        this.activities.push({ timestamp: this.matchdata.map_stats[i].start_time, content: `Map ${i + 1} Started`, icon: 'el-icon-circle-plus-outline', color: '#0bbd87' })
      }
      if (this.matchdata.map_stats[i].end_time !== '0001-01-01T00:00:00Z') {
        this.activities.push({ timestamp: this.matchdata.map_stats[i].end_time, content: `Map ${i + 1} Finished`, icon: 'el-icon-circle-check', color: '#0bbd87' })
      }
    }
    if (this.matchdata.end_time !== '0001-01-01T00:00:00Z') {
      this.activities.push({ timestamp: this.matchdata.end_time, content: 'Match Finished', icon: 'el-icon-success', color: '#0bbd87' })
    }
    for (let i = 0; i < this.matchdata.map_stats.length; i++) {
      this.GetPlayerStats(this.matchdata.id, this.matchdata.map_stats[i].id)
    }
    let team1Promise = this.GetTeam1Data(this.matchdata.team1.id)
    let team2Promise = this.GetTeam2Data(this.matchdata.team2.id)
    this.team1 = await team1Promise
    this.team2 = await team2Promise
    this.loading = false
    let res = await this.axios.get('/api/v1/CheckLoggedIn')
    this.user = res.data
    // this.Editable = this.CheckTeamEditable(this.$route.params.teamid,this.user.userid) // TODO
  },
  methods: {
    async GetTeam1Data (team1id) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/team/${team1id}/GetTeamInfo`)
        resolve(res.data)
      })
    },
    async GetTeam2Data (team2id) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/team/${team2id}/GetTeamInfo`)
        resolve(res.data)
      })
    },
    async GetMatchData (matchid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`)
        resolve(res.data)
      })
    },
    async GetMapStat (matchid) { // TODO?
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/match/${matchid}/GetMatchInfo`)
        resolve(res.data)
      })
    },
    async GetPlayerStats (matchid, mapid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/match/${matchid}/GetPlayerStatInfo?mapID=${mapid}`)
        if (!this.matchdata.team1_player_stats) {
          this.matchdata.team1_player_stats = {}
          this.matchdata.team1_player_stats[mapid] = []
        }
        if (!this.matchdata.team2_player_stats) {
          this.matchdata.team2_player_stats = {}
          this.matchdata.team2_player_stats[mapid] = []
        }

        let team1stats = res.data.filter(player => player.team_id === this.matchdata.team1.id)
        let team2stats = res.data.filter(player => player.team_id === this.matchdata.team2.id)

        for (let i = 0; i < team1stats.length; i++) {
          this.$set(this.matchdata.team1_player_stats, mapid, team1stats)
        }
        for (let i = 0; i < team2stats.length; i++) {
          this.$set(this.matchdata.team2_player_stats, mapid, team2stats)
        }
        resolve(res.data)
      })
    },
    GetSteamURL: function (steamid) {
      return `https://steamcommunity.com/profiles/${steamid}`
    },
    get_logo_or_flag_link: function (team1, team2) { // get_logo_or_flag_link(team1)
      if (team1.logo && team2.logo) {
        return {
          // team1: get_logo_link(team1),
          // team2: get_logo_link(team2)
        }
      } else {
        return {
          team1: this.get_flag_link(team1),
          team2: this.get_flag_link(team2)
        }
      }
    },
    get_logo_html: function (team) {
      // TODO...
    },
    get_flag_link: function (team) {
      if (team.flag === '') {
        return `/img/_unknown.png`
      }
      // return `<img src="/img/valve_flags/${team.flag}"  width="24" height="16">`
      return `/img/valve_flags/${team.flag}.png`
    },
    score_symbol: function (score1, score2) {
      if (score1 > score2) {
        return '>'
      } else {

      } if (score1 < score2) {
        return '<'
      } else {
        return '=='
      }
    },
    get_loser: function (matchdata) { // returns loser's teamname
      if (matchdata.team1_score > matchdata.team2_score) {
        return matchdata.team2.name
      } else if (matchdata.team1_score < matchdata.team2_score) {
        return matchdata.team1.name
      } else {
        return ''
      }
    },
    GetKDR: function (playerstat) {
      if (playerstat.deaths === 0) {
        return playerstat.kills
      }
      return playerstat.kills / playerstat.deaths
    },
    GetHSP: function (playerstat) {
      if (playerstat.deaths === 0) {
        return playerstat.kills
      }
      return playerstat.headshot_kills / playerstat.kills * 100
    },
    GetADR: function (playerstat) {
      if (playerstat.roundsplayed === 0) {
        return 0.0
      }
      return playerstat.damage / playerstat.roundsplayed
    },
    GetFPR: function (playerstat) {
      if (playerstat.roundsplayed === 0) {
        return 0.0
      }
      return playerstat.kills / playerstat.roundsplayed
    },
    handleCommand: function (command) {
      switch (command) {
        case 'cancelmatch':
          this.CancelMatch(this.matchdata.id)
          break
        case 'AddPlayerToTeam1':
          this.AddPlayerToTeam1()
          break
        case 'AddPlayerToTeam2':
          this.AddPlayerToTeam2()
          break
        case 'AddPlayerToSpec':
          this.AddPlayerToSpec()
          break
        case 'rcon_command':
          this.SendRCON()
          break
        case 'backup_manager':
          this.GetBackupList()
          break
        default:
          this.$message.error('Unknown command occured!')
      }
    },
    async CancelMatch (matchid) {
      try {
        await this.$confirm('This will cancel the match. Continue?', 'Warning', {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning'
        })
        const res = await this.axios.post(`/api/v1/match/${matchid}/cancel`)
        this.$message({
          message: 'Successfully cancelled match.',
          type: 'success'
        })
        this.$router.push('/mymatches')
      } catch (err) {
        if (typeof err === 'object' && err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        } else if (typeof err === 'string') {
          this.$message.error(err)
        }
      }
    },
    AdminToolsAvailable: function () {
      if (this.user.isAdmin && this.matchdata.live) {
        return true
      } else if (this.user.isAdmin && this.matchdata.pending) {
        return true
      }
      return false
    },
    async AddPlayerToTeam1 () {
      let steamid = await this.$prompt(`Please enter a SteamID to add to ${this.team1.name}`, 'Tip', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: 'Invalid SteamID'
      })
      try {
        const res = await this.axios.post(`/api/v1/match/${this.matchdata.id}/adduser?team=team1&auth=${steamid.value}`)
        this.$message({
          message: 'Successfully added player.',
          type: 'success'
        })
        this.$router.push(`/match/${this.matchdata.id}`)
      } catch (err) {
        if (typeof err === 'object' && err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        } else if (typeof err === 'string') {
          this.$message.error(err)
        }
      }
    },
    async AddPlayerToTeam2 () {
      let steamid = await this.$prompt(`Please enter a SteamID to add to ${this.team2.name}`, 'Tip', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: 'Invalid SteamID'
      })
      try {
        const res = await this.axios.post(`/api/v1/match/${this.matchdata.id}/adduser?team=team2&auth=${steamid.value}`)
        this.$message({
          message: 'Successfully added player.',
          type: 'success'
        })
        this.$router.push(`/match/${this.matchdata.id}`)
      } catch (err) {
        if (typeof err === 'object' && err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        } else if (typeof err === 'string') {
          this.$message.error(err)
        }
      }
    },
    async AddPlayerToSpec () {
      let steamid = await this.$prompt(`Please enter a SteamID to add to Spectators`, 'Tip', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: 'Invalid SteamID'
      })
      try {
        const res = await this.axios.post(`/api/v1/match/${this.matchdata.id}/adduser?team=spec&auth=${steamid.value}`)
        this.$message({
          message: 'Successfully added player.',
          type: 'success'
        })
        this.$router.push(`/match/${this.matchdata.id}`)
      } catch (err) {
        if (typeof err === 'object' && err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        } else if (typeof err === 'string') {
          this.$message.error(err)
        }
      }
    },
    async SendRCON () {
      let command = await this.$prompt(`Enter a command to send`, 'Tip', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: 'Invalid command'
      })
      try {
        const res = await this.axios.post(`/api/v1/match/${this.matchdata.id}/rcon?command=${command.value}`)
        this.$message({
          message: 'Successfully sent command.',
          type: 'success'
        })
        this.$router.push(`/match/${this.matchdata.id}`)
      } catch (err) {
        if (err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        }
      }
    },
    async GetBackupList () {
      try {
        let backups = await this.axios.get(`/api/v1/match/${this.matchdata.id}/backup`)
        this.backups = backups.data.files
        this.chose_backup = true
      } catch (err) {
        if (err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        }
      }
    },
    async SendBackup () {
      try {
        await this.axios.post(`/api/v1/match/${this.matchdata.id}/backup?file=${this.chosed_backup}`)
        this.chose_backup = false
      } catch (err) {
        if (err.response) {
          if (typeof err.response.data === 'string') {
            this.$message.error(err.response.data)
          } else if (typeof err.response.data === 'object') {
            this.$message.error(err.response.data.errormessage)
          }
        }
      }
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
    display: block;
    margin: 0 10px;
}

a {
    color: #42b983;
}

.panel-body {
  overflow: scroll
}
</style>
