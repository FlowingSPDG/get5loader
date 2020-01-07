<template>
<div id="container">
    <div id="content">
        <el-form ref="form" :model="form" :rules="rules" label-position="left" label-width="120px">
            <el-form-item label="Server" style="width: 653px;" prop="display_name">
                <el-select v-model="form.server_id">
                    <el-option v-for="(server, index) in servers" :label="server.display_name" :key="index" :value="server.id"></el-option>
                </el-select>
                <router-link class="btn btn-primary" to="/server/create">Create a Server</router-link>
            </el-form-item>

            <el-form-item label="Team 1" style="width: 653px;" prop="team1_id">
                <el-select v-model="form.team1_id">
                    <el-option v-for="(team, index) in teams" :label="team.name" :key="index" :value="team.id"></el-option>
                </el-select>
            </el-form-item>

            <el-form-item v-if="match_text_option" label="Team 1 title text" style="width: 653px;" prop="team1_string">
                <el-input v-model="form.team1_string"></el-input>
            </el-form-item>

            <el-form-item label="Team2" style="width: 653px;" prop="team2_id">
                <el-select v-model="form.team2_id">
                    <el-option v-for="(team, index) in teams" :label="team.name" :key="index" :value="team.id"></el-option>
                </el-select>
            </el-form-item>

            <el-form-item v-if="match_text_option" label="Team 2 title text" style="width: 653px;" prop="team2_string">
                <el-input v-model="form.team2_string"></el-input>
            </el-form-item>

            <el-form-item v-if="match_text_option" label="Match title text" style="width: 653px;" prop="match_title">
                <el-input v-model="form.match_title"></el-input>
            </el-form-item>

            <el-form-item label="Series Type" style="width: 653px;" prop="series_type">
                <el-radio-group v-model="form.series_type">
                    <el-radio v-for="(option, index) in series_type" :label="option.type" :key="index"></el-radio>
                </el-radio-group>
            </el-form-item>

            <el-form-item label="Map Pool" style="width: 653px;" prop="veto_mappool">
                <el-checkbox-group v-model="form.veto_mappool">
                    <el-checkbox v-for="(map, index) in mappool" :label="map.system" :key="index"></el-checkbox>
                </el-checkbox-group>
            </el-form-item>

            <el-form-item style="width: 653px;">
                <el-button type="primary" @click="RegisterMatch">Create Match</el-button>
            </el-form-item>
        </el-form>
    </div>
</div>
</template>

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

<script>
export default {
  name: 'MatchCreate',
  data () {
    return {
      user: {
        id: 0,
        steam_id: '',
        name: '',
        admin: false,
        servers: null,
        teams: null,
        matches: null
      },
      servers: [],
      teams: [],
      mappool: [
        {
          system: 'de_dust2',
          formal: 'Dust II'
        },
        {
          system: 'de_mirage',
          formal: 'Mirage'
        },
        {
          system: 'de_inferno',
          formal: 'Inferno'
        },
        {
          system: 'de_overpass',
          formal: 'Overpass'
        },
        {
          system: 'de_nuke',
          formal: 'NUKE'
        },
        {
          system: 'de_train',
          formal: 'Train'
        },
        {
          system: 'de_vertigo',
          formal: 'Vertigo'
        },
        {
          system: 'de_cache',
          formal: 'Cache'
        }
      ],
      series_type: [
        {
          type: 'bo1-preset',
          desc: 'Bo1 with preset map'
        },
        {
          type: 'bo1',
          desc: 'Bo1 with map vetoes'
        },
        {
          type: 'bo2',
          desc: 'Bo2 with map vetoes'
        },
        {
          type: 'bo3',
          desc: 'Bo3 with map vetoes'
        },
        {
          type: 'bo5',
          desc: 'Bo5 with map vetoes'
        },
        {
          type: 'bo7',
          desc: 'Bo7 with map vetoes'
        }
      ],
      form: {
        server_id: 0,
        team1_id: 0,
        team2_id: 0,
        max_maps: 0,
        title: '',
        skip_veto: false,
        veto_mappool: [],
        series_type: 'bo1'
      },
      rules: {
        server_id: [{
          required: true,
          trigger: 'change',
          message: 'Please chose server'
        }],
        team1_id: [{
          required: true,
          trigger: 'change',
          message: 'Please chose team1 id'
        }],
        team2_id: [{
          required: true,
          trigger: 'change',
          message: 'Please chose team2 id'
        }],
        title: [{
          required: true,
          trigger: 'change',
          message: 'Please input title'
        }],
        skip_veto: [{
          required: false,
          trigger: 'change',
          message: 'Please chose skip veto option'
        }],
        veto_mappool: [{
          required: false,
          trigger: 'change',
          message: 'Please chose map(s)'
        }]
      },
      match_text_option: false // TODO
    }
  },
  async created () {
    this.user = await this.axios.get('/api/v1/CheckLoggedIn')
    this.servers = await this.GetServers()
    this.teams = await this.GetTeams()
  },
  methods: {
    async RegisterMatch () {
      this.form.team1_id = parseInt(this.form.team1_id, 10)
      this.form.team2_id = parseInt(this.form.team2_id, 10)
      this.form.server_id = parseInt(this.form.server_id, 10)
      if (this.form.team1_id === this.form.team2_id) {
        this.$message.error('Teams cannot be equal')
        return
      }
      switch (this.form.series_type) {
        case 'bo1-preset':
          this.form.max_maps = 1
          this.form.skip_veto = true
          if (this.form.veto_mappool.length !== 1) {
            this.$message.error('You must have exactly 1 map selected to do a bo1 with a preset map')
            return
          }
          break
        case 'bo1':
          this.form.max_maps = 1
          break
        case 'bo2':
          this.form.max_maps = 3
          break
        case 'bo3':
          this.form.max_maps = 3
          break
        case 'bo5':
          this.form.max_maps = 5
          break
        case 'bo7':
          this.form.max_maps = 7
          break
      }
      const json = JSON.stringify(this.form)
      this.$refs['form'].validate(async (valid) => {
        if (valid) {
          try {
            let res = await this.axios.post('/api/v1/match/create', json)
            this.form = {}
            this.$message({
              message: 'Successfully registered match.',
              type: 'success'
            })
            this.$router.push('/mymatches')
          } catch (err) {
            if (typeof err.response.data === 'string') {
              this.$message.error(err.response.data)
            } else if (typeof err.response.data === 'object') {
              this.$message.error(err.response.data.errormessage)
            }
          }
        } else {
          this.$message.error('Please fill form')
        }
      })
    },
    async GetTeams () {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get('/api/v1/GetTeamList')
        resolve(res.data)
      })
    },
    async GetServers () {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get('/api/v1/GetServerList')
        resolve(res.data)
      })
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

#container {
    padding-left: 500px;
    margin-right: auto;
    margin-left: auto;
    align-content: center;
}
</style>
