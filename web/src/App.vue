<template>
<div id="app">
    <div id="header">
        <nav class="navbar navbar-default">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="/">Get5 Web Panel</a>
                </div>
                <div class="collapse navbar-collapse" id="myNavbar">
                    <ul class="nav navbar-nav">
                        <li><router-link id="matches" to="/matches" v-if="user.isLoggedIn">All Matches</router-link></li>
                        <li><a id="mymatches" href="/mymatches" v-if="user.isLoggedIn">My Matches</a></li>
                        <li><a id="match_create" href="/match/create" v-if="user.isLoggedIn">Create a Match</a></li>
                        <li><a id="myteams" href="/myteams" v-if="user.isLoggedIn">My Teams</a></li>
                        <li><a id="team_create" href="/team/create" v-if="user.isLoggedIn">Create a Team</a></li>
                        <li><a id="myservers" href="/myservers" v-if="user.isLoggedIn">My Servers</a></li>
                        <li><a id="server_create" href="/server/create" v-if="user.isLoggedIn">Add a Server</a></li>
                        <li><a href="/logout" v-if="user">Logout</a></li>
                        <li><a href="/login" v-if="!user.isLoggedIn"> <img src="/static/img/login_small.png" height="18" /></a></li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>
    <router-view />
    <div class="panel-footer text-muted">
        <div>
            Powered by <a href="http://steampowered.com">Steam</a> -
            <a href="/metrics">Stats</a>
            <div v-if="COMMIT_STRING">- Version <a href="https://github.com/FlowingSPDG/get5-web-go">{{ COMMIT_STRING }}+</a></div>
        </div>
    </div>
</div>
</template>

<script>
import axios from 'axios'
export default {
    name: 'App',
    data() {
        return {
            user: {
                isLoggedIn:false,
                steamid:undefined,
                userid:undefined
            }, // should be object from JSON response
            COMMIT_STRING: "COMMIT NUMBER HERE"
        }
    },
    mounted () {
        axios
            .get('/api/v1/CheckLoggedIn')
            .then((res) => {
                console.log(res.data)
                this.user = res.data
            })
  }
}
</script>

<style>
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
