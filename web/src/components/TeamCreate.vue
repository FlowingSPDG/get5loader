<template>
<div id="content">
  <form role="form" action="" method="post" class="form-horizontal">

    <!--{{ form.csrf_token }} -->

    <div class="form-group">
    <label class="col-sm-2 control-label" for="name">Team Name</label>
      <div class="col-sm-7">
        <input class="form-control" id="name" name="name" placeholder="Team name" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="tag">Team Tag</label>
      <div class="col-sm-7">
        <input class="form-control" id="tag" name="tag" placeholder="Team tag" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="logo">Logo Name</label>
      <div class="col-sm-7">
        <select class="form-control input-sm" id="logo" name="logo"><option selected value="">None</option></select>
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2 control-label" for="country_flag">Country Flag</label>
        <div class="col-sm-7">
          <select class="form-control input-sm" id="country_flag" name="country_flag"><option selected value="">None</option><option value="dz">Algeria - DZ</option><option value="ar">Argentina - AR</option><option value="au">Australia - AU</option><option value="at">Austria - AT</option><option value="by">Belarus - BY</option><option value="be">Belgium - BE</option><option value="br">Brazil - BR</option><option value="bg">Bulgaria - BG</option><option value="ca">Canada - CA</option><option value="cl">Chile - CL</option><option value="cn">China - CN</option><option value="cc">Cocos Islands - CC</option><option value="hr">Croatia - HR</option><option value="cz">Czech Republic - CZ</option><option value="dk">Denmark - DK</option><option value="ee">Estonia - EE</option><option value="eu">European Union - EU</option><option value="fi">Finland - FI</option><option value="fr">France - FR</option><option value="de">Germany - DE</option><option value="gr">Greece - GR</option><option value="gp">Guadeloupe - GP</option><option value="hk">Hong Kong - HK</option><option value="hu">Hungary - HU</option><option value="is">Iceland - IS</option><option value="in">India - IN</option><option value="id">Indonesia - ID</option><option value="ir">Iran - IR</option><option value="ie">Ireland - IE</option><option value="il">Israel - IL</option><option value="it">Italy - IT</option><option value="jp">Japan - JP</option><option value="kz">Kazahkstan - KZ</option><option value="lv">Latvia - LV</option><option value="ly">Libya - LY</option><option value="lt">Liechtenstein - LT</option><option value="lu">Luxembourg - LU</option><option value="mo">Macao - MO</option><option value="mk">Macedonia - MK</option><option value="my">Malaysia - MY</option><option value="mx">Mexico - MX</option><option value="nl">Netherlands - NL</option><option value="nz">New Zealand - NZ</option><option value="no">Norway - NO</option><option value="pk">Pakistan - PK</option><option value="pe">Peru - PE</option><option value="ph">Phillippines - PH</option><option value="pl">Poland - PL</option><option value="pt">Portugal - PT</option><option value="re">Reunion - RE</option><option value="ro">Romania - RO</option><option value="ru">Russia - RU</option><option value="sa">Saudi Arabia - SA</option><option value="rs">Serbia - RS</option><option value="sg">Singapore - SG</option><option value="si">Slovenia - SI</option><option value="sk">Slovokia - SK</option><option value="za">South Africa - ZA</option><option value="kr">South Korea - KR</option><option value="es">Spain - ES</option><option value="se">Sweden - SE</option><option value="ch">Switzerland - CH</option><option value="tw">Taiwan - TW</option><option value="th">Thailand - TH</option><option value="tr">Turkey - TR</option><option value="ua">Ukraine - UA</option><option value="ae">United Arab Emirates - AE</option><option value="gb">United Kingdom - GB</option><option value="us">United States - US</option><option value="ve">Venezuela - VE</option><option value="vn">Vietnam - VN</option></select>
        </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth1">Player 1</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth1" name="auth1" placeholder="Any steamid format" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth2">Player 2</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth2" name="auth2" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth3">Player 3</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth3" name="auth3" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth4">Player 4</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth4" name="auth4" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth5">Player 5</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth5" name="auth5" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth6">Player 6</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth6" name="auth6" type="text" value="">
      </div>
    </div>

    <div class="form-group">
    <label class="col-sm-2 control-label" for="auth7">Player 7</label>
      <div class="col-sm-7">
        <input class="form-control" id="auth7" name="auth7" type="text" value="">
      </div>
    </div>

    <div class="form-group" v-if="user.adminaccess == true">
      <label class="col-sm-2 control-label" for="public_team">Public Team</label>
      <div class="col-sm-offset-1">
        <input id="public_team" name="public_team" type="checkbox" value="y">
      </div>
    </div>

    <div class="input submit col-sm-offset-1" v-if="edit">
      <input type="submit" class="btn btn-primary" value="Update Team">
    </div>

    <div class="input submit col-sm-offset-1" v-else>
      <input type="submit" class="btn btn-primary" value="Create Team">
    </div>

    <br>

  </form>

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
      edit: false // TODO
    }
  },
  async created () {
    this.user = await this.axios.get('/api/v1/CheckLoggedIn')
  },
  methods: {
    async RegisterTeam () {
      // TODO
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
