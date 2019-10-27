<template>
  <div class="matches">
    <h1>Matches Page</h1>
    <p>{{ msg }}</p>
    <p>{{ info }}</p>
    <ul>
      <li>
        <router-link to="/">/</router-link>
        <router-link to="/matches">Matches</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'matches',
  data () {
    return {
      msg: 'Matches page',
      info: ''
    }
  },
  mounted () {
    
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
