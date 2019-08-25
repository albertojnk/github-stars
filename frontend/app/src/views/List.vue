<template>
  <div class="list">
    <table class="table-repositories">
      <thead>
        <tr class="tableHeader">
          <th>Repository</th>
          <th>Description</th>
          <th>Language</th>
          <th>Tags</th>
          <th></th>
        </tr>
      </thead>
      <tbody class="tableBody">
        <tr v-for="(repository, index) in repositories" v-bind:key="index">
          <td class="repository-name">{{ repository.name }}</td>
          <td class="repository-description">{{ repository.description }}</td>
          <td class="repository-language">{{ repository.language }}</td>
          <td class="repository-tags">{{ repository.tags }}</td>
          <td class="repository-edit">edit</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "list",
  components: {},
  data() {
    return {
      tags: ""
    };
  },
  computed: {
    ...mapState(["id", "repositories"])
  },
  methods: {}
};
</script>
<style>
tr:nth-child(odd) {
  background-color: #f2f2f2;
}
tr:nth-child(odd) {
  border-right: solid 1px #000000;
}
.table-repositories {
  border-collapse: collapse;
  width: 90%;
  margin: 5% 5%;
  border: solid 1px #000000;
}
.repository-name {
  border-right: solid 1px #000000;
}
.repository-description {
  border-right: solid 1px #000000;
}
.repository-language {
  border-right: solid 1px #000000;
  padding: 0 5px;
}
.repository-tags {
  border-right: solid 1px #000000;
  padding: 0 5px;
}
.repository-edit {
  padding: 0 7px;
}
.tableHeader {
  background-color: #cfcfcf;
  border-bottom: solid 1px #000000;
}
</style>
