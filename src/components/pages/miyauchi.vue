<template>
  <b-table :data="data" :columns="columns"></b-table>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      data: [],
      columns: [
        { field: "id", label: "ID", width: '40', numeric: true },
        { field: "title", label: "タイトル" },
        { field: "author", label: "作者" },
        { field: "votes", label: "投票数" }
      ]
    }
  },
  created() {
    axios.get("/api/miyauchi").then((res) => {
      this.data = res.data
    })
  }
}
</script>
