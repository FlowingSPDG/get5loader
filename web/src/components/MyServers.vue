<template>
  <ul class="list-group" v-if="servers">
    <li class="list-group-item" v-if="servers.length == 0">
      No servers found.
    </li>

    <el-table :data="servers" style="width: 100%" v-else>
      <el-table-column prop="id" label="Server ID" width="180"></el-table-column>
      <el-table-column prop="display_name" label="Display Name" width="180"></el-table-column>
      <el-table-column prop="ip_string" label="IP Address"></el-table-column>
      <el-table-column prop="port" label="Port"></el-table-column>
      <el-table-column label="Status">
        <template slot-scope="scope">
          <template v-if="scope.row.in_use"> In use </template>
          <template v-else> Free  </template>
        </template>
      </el-table-column>
      <el-table-column>
        <template slot-scope="scope">
          <a :href="'/server/'+scope.row.id+'/edit'" class="btn btn-primary btn-xs">Edit</a>
          <a :href="'/server/'+scope.row.id+'/delete'" class="btn btn-danger btn-xs" v-if="!scope.row.in_use">Delete</a>
        </template>
      </el-table-column>
    </el-table>
  </ul>
</template>

<script>
export default {
  name: 'MyServers',
  data () {
    return {
      servers: []
    }
  },
  async created () {
    const res = await this.axios.get('/api/v1/CheckLoggedIn')
    const user = await this.GetUserData(res.data.userid)
    this.servers = user.servers
  },
  methods: {
    async GetUserData (userid) {
      return new Promise(async (resolve, reject) => {
        const res = await this.axios.get(`/api/v1/user/${userid}/GetUserInfo`)
        resolve(res.data)
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
