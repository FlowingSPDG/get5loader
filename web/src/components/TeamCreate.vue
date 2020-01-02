<template>
<div id="container">
<div id="content">
  <el-form ref="form" :model="form" :rules="rules" label-position="left" label-width="120px">
  <el-form-item label="Team Name" style="width: 653px;" prop="name">
    <el-input v-model="form.name"></el-input>
  </el-form-item>

  <el-form-item label="Team Tag" style="width: 653px;" prop="tag">
    <el-input v-model="form.tag"></el-input>
  </el-form-item>

  <el-form-item label="Logo Name" style="width: 653px;" prop="logo_name">
    <el-select v-model="form.logo_name" placeholder="None" style="width: 653px;">
      <el-option selected value="">None</el-option>
    </el-select>
  </el-form-item>

  <el-form-item label="Country Flag" style="width: 653px;" prop="flag">
    <el-select v-model="form.flag" placeholder="None" style="width: 653px;">
      <el-option selected value="">None</el-option><br>
      <el-option label="Algeria" value="dz">Algeria - DZ</el-option><br>
      <el-option label="Argentina" value="ar">Argentina - AR</el-option><br>
      <el-option label="Australia" value="au">Australia - AU</el-option><br>
      <el-option label="Austria" value="at">Austria - AT</el-option><br>
      <el-option label="Belarus" value="by">Belarus - BY</el-option><br>
      <el-option label="Belgium" value="be">Belgium - BE</el-option><br>
      <el-option label="Brazil" value="br">Brazil - BR</el-option><br>
      <el-option label="Bulgaria" value="bg">Bulgaria - BG</el-option><br>
      <el-option label="Canada" value="ca">Canada - CA</el-option><br>
      <el-option label="Chile" value="cl">Chile - CL</el-option><br>
      <el-option label="China" value="cn">China - CN</el-option><br>
      <el-option label="Cocos Islands" value="cc">Cocos Islands - CC</el-option><br>
      <el-option label="Croatia" value="hr">Croatia - HR</el-option><br>
      <el-option label="Czech Republic" value="cz">Czech Republic - CZ</el-option><br>
      <el-option label="Denmark" value="dk">Denmark - DK</el-option><br>
      <el-option label="Estonia" value="ee">Estonia - EE</el-option><br>
      <el-option label="European Union" value="eu">European Union - EU</el-option><br>
      <el-option label="Finland" value="fi">Finland - FI</el-option><br>
      <el-option label="France" value="fr">France - FR</el-option><br>
      <el-option label="Germany" value="de">Germany - DE</el-option><br>
      <el-option label="Greece" value="gr">Greece - GR</el-option><br>
      <el-option label="Guadeloupe" value="gp">Guadeloupe - GP</el-option><br>
      <el-option label="Hong Kong" value="hk">Hong Kong - HK</el-option><br>
      <el-option label="Hungary" value="hu">Hungary - HU</el-option><br>
      <el-option label="Iceland" value="is">Iceland - IS</el-option><br>
      <el-option label="India" value="in">India - IN</el-option><br>
      <el-option label="Indonesia" value="id">Indonesia - ID</el-option><br>
      <el-option label="Iran" value="ir">Iran - IR</el-option><br>
      <el-option label="Ireland" value="ie">Ireland - IE</el-option><br>
      <el-option label="Israel" value="il">Israel - IL</el-option><br>
      <el-option label="Italy" value="it">Italy - IT</el-option><br>
      <el-option label="Japan" value="jp">Japan - JP</el-option><br>
      <el-option label="Kazahkstan" value="kz">Kazahkstan - KZ</el-option><br>
      <el-option label="Latvia" value="lv">Latvia - LV</el-option><br>
      <el-option label="Libya" value="ly">Libya - LY</el-option><br>
      <el-option label="Liechtenstein" value="lt">Liechtenstein - LT</el-option><br>
      <el-option label="Luxembourg" value="lu">Luxembourg - LU</el-option><br>
      <el-option label="Macao" value="mo">Macao - MO</el-option><br>
      <el-option label="Macedonia" value="mk">Macedonia - MK</el-option><br>
      <el-option label="Malaysia" value="my">Malaysia - MY</el-option><br>
      <el-option label="Mexico" value="mx">Mexico - MX</el-option><br>
      <el-option label="Netherlands" value="nl">Netherlands - NL</el-option><br>
      <el-option label="New Zealand" value="nz">New Zealand - NZ</el-option><br>
      <el-option label="Norway" value="no">Norway - NO</el-option><br>
      <el-option label="Pakistan" value="pk">Pakistan - PK</el-option><br>
      <el-option label="Peru" value="pe">Peru - PE</el-option><br>
      <el-option label="Phillippines" value="ph">Phillippines - PH</el-option><br>
      <el-option label="Poland" value="pl">Poland - PL</el-option><br>
      <el-option label="Portugal" value="pt">Portugal - PT</el-option><br>
      <el-option label="Reunion" value="re">Reunion - RE</el-option><br>
      <el-option label="Romania" value="ro">Romania - RO</el-option><br>
      <el-option label="Russia" value="ru">Russia - RU</el-option><br>
      <el-option label="Saudi Arabia" value="sa">Saudi Arabia - SA</el-option><br>
      <el-option label="Serbia" value="rs">Serbia - RS</el-option><br>
      <el-option label="Singapore" value="sg">Singapore - SG</el-option><br>
      <el-option label="Slovenia" value="si">Slovenia - SI</el-option><br>
      <el-option label="Slovokia" value="sk">Slovokia - SK</el-option><br>
      <el-option label="South Africa" value="za">South Africa - ZA</el-option><br>
      <el-option label="South Korea" value="kr">South Korea - KR</el-option><br>
      <el-option label="Spain" value="es">Spain - ES</el-option><br>
      <el-option label="Sweden" value="se">Sweden - SE</el-option><br>
      <el-option label="Switzerland" value="ch">Switzerland - CH</el-option><br>
      <el-option label="Taiwan" value="tw">Taiwan - TW</el-option><br>
      <el-option label="Thailand" value="th">Thailand - TH</el-option><br>
      <el-option label="Turkey" value="tr">Turkey - TR</el-option><br>
      <el-option label="Ukraine" value="ua">Ukraine - UA</el-option><br>
      <el-option label="United Arab Emirates" value="ae">United Arab Emirates - AE</el-option><br>
      <el-option label="United Kingdom" value="gb">United Kingdom - GB</el-option><br>
      <el-option label="United States" value="us">United States - US</el-option><br>
      <el-option label="Venezuela" value="ve">Venezuela - VE</el-option><br>
      <el-option label="Vietnam" value="vn">Vietnam - VN</el-option><br>
    </el-select>
  </el-form-item>

  <el-form-item label="Player 1" style="width: 653px;" prop="auths[0]">
    <el-input v-model="form.auths[0]"></el-input>
  </el-form-item>
  <el-form-item label="Player 2" style="width: 653px;" prop="auths[1]">
    <el-input v-model="form.auths[1]"></el-input>
  </el-form-item>
  <el-form-item label="Player 3" style="width: 653px;" prop="auths[2]">
    <el-input v-model="form.auths[2]"></el-input>
  </el-form-item>
  <el-form-item label="Player 4" style="width: 653px;" prop="auths[3]">
    <el-input v-model="form.auths[3]"></el-input>
  </el-form-item>
  <el-form-item label="Player 5" style="width: 653px;" prop="auths[4]">
    <el-input v-model="form.auths[4]"></el-input>
  </el-form-item>
  <el-form-item label="Player 6" style="width: 653px;" prop="auths[5]">
    <el-input v-model="form.auths[5]"></el-input>
  </el-form-item>
  <el-form-item label="Player 7" style="width: 653px;" prop="auths[6]">
    <el-input v-model="form.auths[6]"></el-input>
  </el-form-item>

  <el-form-item label="Public Team" style="width: 653px;">
    <el-switch v-model="form.public_team"></el-switch>
  </el-form-item>

  <el-form-item style="width: 653px;">
    <el-button type="primary" @click="RegisterTeam">Create</el-button>
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
  name: 'TeamCreate',
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
      form: {
        name: '',
        tag: '',
        flag: '',
        logo: '',
        auths: [],
        public_team: false
      },
      rules: {
        name: [
          { required: true, trigger: 'change', min: 5, max: 30, message: 'Length should be 5 to 30' }
        ],
        tag: [
          { required: false, message: 'Please input team tag', trigger: 'change' }
        ],
        flag: [
          { required: false, message: 'Please select country flag', trigger: 'change' }
        ],
        logo: [
          { required: true, message: 'Please pick a logo', trigger: 'change' }
        ],
        auths: [
          { required: true, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' },
          { required: false, message: 'Please input your team members steamid', trigger: 'blur' }
        ],
        public_team: [
          { required: false, message: 'Please select team publicity', trigger: 'change' }
        ]
      },
      edit: false // TODO
    }
  },
  async created () {
    this.user = await this.axios.get('/api/v1/CheckLoggedIn')
  },
  methods: {
    async RegisterTeam () {
      const json = JSON.stringify(this.form)
      this.$refs['form'].validate(async (valid) => {
        if (valid) {
          try {
            let res = await this.axios.post('/api/v1/team/create', json)
            this.form = {}
            this.$message({
              message: 'Successfully registered team.',
              type: 'success'
            })
            this.$router.push('/myteams')
          } catch (err) {
            this.$message.error(err.response.data.errormessage)
          }
        } else {
          this.$message.error('Please fill form')
        }
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
