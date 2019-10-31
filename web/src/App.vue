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
                    <router-link to="/" class="navbar-brand" >Get5 Web Panel</router-link>
                </div>
                <div class="collapse navbar-collapse" id="myNavbar">
                        <el-menu
                            class="nav navbar-nav"
                            :default-active="activeIndex"
                            mode="horizontal"
                            router
                        >
                        <el-menu-item index="Matches" id="matches" :route="{ path:'/matches' }" v-if="user.isLoggedIn">All Matches</el-menu-item>
                        <el-menu-item index="My matches" id="mymatches" :route="{ path:'/mymatches' }" v-if="user.isLoggedIn">My Matches</el-menu-item>
                        <el-menu-item index="match_create" id="match_create" :route="{ path:'/match/create' }" v-if="user.isLoggedIn">Create a Match</el-menu-item>
                        <el-menu-item index="myteams" id="myteams" :route="{ path:'/myteams' }" v-if="user.isLoggedIn">My Teams</el-menu-item>
                        <el-menu-item index="team_create" id="matches" :route="{ path:'/team/create' }" v-if="user.isLoggedIn">Create a Team</el-menu-item>
                        <el-menu-item index="myservers" id="myservers" :route="{ path:'/myservers' }" v-if="user.isLoggedIn">My Servers</el-menu-item>
                        <el-menu-item index="server_create" id="server_create" :route="{ path:'/server/create' }" v-if="user.isLoggedIn">Add a Server</el-menu-item>
                        <el-menu-item index="Logout" :route="{ path:'/logout' }" v-if="user.isLoggedIn">Logout</el-menu-item>
                        <el-menu-item index="login" id="login" :route="{ path:'/login' }" v-if="!user.isLoggedIn"> <img src="/static/img/login_small.png" height="18" /></el-menu-item>
                    </el-menu>
                </div>
            </div>
        </nav>
    </div>
    <router-view />
    <div class="panel-footer text-muted">
        <div>
            Powered by <a href="http://steampowered.com">Steam</a> -
            <router-link id="metrics" to="/metrics" >Stats</router-link>
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
            activeIndex: '',
            user: {
                isLoggedIn:false,
                adminaccess:false,
                steamid:"",
                userid:""
            }, // should be object from JSON response
            COMMIT_STRING: "COMMIT NUMBER HERE"
        }
    },
    mounted () {
        this.activeIndex = this.$route.name;
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
