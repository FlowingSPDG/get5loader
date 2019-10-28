<template>
  <div id="metrics">
    <h1>Get5 Metrics</h1>
    <ul class="list-group">
        <li class="list-group-item" v-for="(value,key, index) in metrics" :key="value.key">
            {{key}}: {{value}}
        </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Metrics',
  data () {
    return {
      metrics:{}
    }
  },
  created () {
    this.axios.get('/api/v1/GetMetrics').then(res => {
      this.metrics = res.data
    })
  },
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
