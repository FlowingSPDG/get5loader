<template>
<div id="container">
<div id="content">
  <el-form ref="form" :model="form" :rules="rules" label-position="left" label-width="120px">
  <el-form-item label="Display Name" style="width: 653px;" prop="display_name">
    <el-input v-model="form.display_name"></el-input>
  </el-form-item>

  <el-form-item label="Server IP" style="width: 653px;" prop="ip_string">
    <el-input v-model="form.ip_string"></el-input>
  </el-form-item>

  <el-form-item label="Server Port" style="width: 653px;" prop="port">
    <el-input v-model="form.port" placeholder=27015></el-input>
  </el-form-item>

  <el-form-item label="RCON Password" style="width: 653px;" prop="rcon_password">
    <el-input type="password"  v-model="form.rcon_password"></el-input>
    <p class="help-block">Your server information will not be exposed to other users.</p>
  </el-form-item>

  <el-form-item label="Public Team" style="width: 653px;" prop="public_server" v-if="user.admin">
    <el-switch v-model="form.public_server"></el-switch>
  </el-form-item>

  <el-form-item style="width: 653px;" v-if="edit">
    <el-button type="primary" @click="UpdateServer">UpdateServer</el-button>
  </el-form-item>

  <el-form-item style="width: 653px;" v-else>
    <el-button type="primary" @click="RegisterServer">Create</el-button>
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
  name: 'ServerCreate',
  data () {
    var ValidateIPaddress = (rule, value, callback) => {
      if (/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/.test(value)) {
        callback()
      } else {
        callback(new Error('Please input valid IP'))
      }
    }
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
      form: {
        ip_string: '',
        port: 27015,
        rcon_password: '',
        display_name: '',
        public_server: false
      },
      rules: {
        ip_string: [
          { required: true, trigger: 'change', validator: ValidateIPaddress, message: 'Please enter valid IP' }
        ],
        rcon_password: [
          { required: true, message: 'Please input rcon password', trigger: 'change' }
        ],
        display_name: [
          { required: true, message: 'Please input server display name', trigger: 'change' }
        ],
        public_server: [
          { required: false, message: 'Please check if server is public', trigger: 'change' }
        ]
      },
      edit: false // TODO
    }
  },
  async created () {
    this.$message({
      dangerouslyUseHTMLString: true,
      message: `Make sure your server is running and has the get5 server plugins installed first.<br>See <a href="https://github.com/splewis/get5/wiki/Step-by-step-installation-guide">the get5 wiki</a> for help installing the get5 and get5_apistats plugins.`,
      type: 'info',
      duration: 1000 * 10,
      showClose: true
    })
    this.user = await this.axios.get('/api/v1/CheckLoggedIn')
  },
  methods: {
    async RegisterServer () {
      const json = JSON.stringify(this.form)
      this.$refs['form'].validate(async (valid) => {
        if (valid) {
          try {
            let res = await this.axios.post('/api/v1/server/create', json)
            this.form = {}
            this.$message({
              message: 'Successfully registered server.',
              type: 'success'
            })
            this.$router.push('/matches')
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
    }
  },
  async UpdateServer () { // TODO

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

#container {
  padding-left: 500px;
  margin-right: auto;
  margin-left: auto;
  align-content: center;
}
</style>
