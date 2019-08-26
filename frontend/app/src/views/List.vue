<template>
  <div class="list">
    <div
      class="loading-div"
      v-if="this.loading == true && this.loaded == false"
    >
      <div id="noTrespassingOuterBarG">
        <div id="noTrespassingFrontBarG" class="noTrespassingAnimationG">
          <div class="noTrespassingBarLineG"></div>
          <div class="noTrespassingBarLineG"></div>
          <div class="noTrespassingBarLineG"></div>
          <div class="noTrespassingBarLineG"></div>
          <div class="noTrespassingBarLineG"></div>
          <div class="noTrespassingBarLineG"></div>
        </div>
      </div>
      <p>Getting the repositories list from Github...</p>
    </div>
    <div
      class="listContainer"
      v-if="this.loading == false && this.loaded == true"
    >
      <div class="lighter">
        <span class="searchContainer">
          <input type="search" class="search" placeholder="search by tag" />
        </span>
      </div>
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
            <td class="repository-tags">
              <template v-for="(tag, idx) in repository.tags">
                {{ tag | tagNormalize(idx, repository.tags.length) }}
              </template>
            </td>
            <td class="repository-edit">
              <a href="javascript:void(0)" @click="show(repository)">edit</a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="modal" v-bind:class="{ 'edit-modal': modal }">
      <div class="content">
        <p>edit tags for {{ currentRepo.name }}</p>
        <input type="text" name="tags" id="tags" v-bind:value="currentTags" />
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "list",
  components: {},
  data() {
    return {
      modal: false,
      currentRepo: null,
      currentTags: null
    };
  },
  computed: {
    ...mapState(["id", "repositories", "loading", "loaded"])
  },
  methods: {
    show(repository) {
      this.modal = true;
      this.currentRepo = repository;
      this.currentTags = this.currentTagsNormalizer(repository.tags);
    },
    hide() {
      this.modal = false;
    },
    currentTagsNormalizer(tags) {
      var len = tags.length;
      var results = "";
      if (len == 0) {
        return "";
      }
      tags.forEach(tag => {
        results += tag + ", ";
      });

      return results;
    }
  },
  filters: {
    tagNormalize(tag, index, len) {
      if (index == len - 1) {
        return tag;
      }
      return tag + ", ";
    }
  }
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
.loading-div {
  height: 100%;
  margin: 15%;
}
#noTrespassingOuterBarG {
  height: 19px;
  width: 156px;
  border: 1px solid rgb(0, 0, 0);
  overflow: hidden;
  background-color: rgb(255, 255, 255);
  margin: auto;
}

.noTrespassingBarLineG {
  background-color: rgb(0, 0, 0);
  float: left;
  width: 14px;
  height: 117px;
  margin-right: 23px;
  margin-top: -27px;
  transform: rotate(45deg);
  -o-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  -webkit-transform: rotate(45deg);
  -moz-transform: rotate(45deg);
}

.noTrespassingAnimationG {
  width: 230px;
  animation-name: noTrespassingAnimationG;
  -o-animation-name: noTrespassingAnimationG;
  -ms-animation-name: noTrespassingAnimationG;
  -webkit-animation-name: noTrespassingAnimationG;
  -moz-animation-name: noTrespassingAnimationG;
  animation-duration: 0.515s;
  -o-animation-duration: 0.515s;
  -ms-animation-duration: 0.515s;
  -webkit-animation-duration: 0.515s;
  -moz-animation-duration: 0.515s;
  animation-iteration-count: infinite;
  -o-animation-iteration-count: infinite;
  -ms-animation-iteration-count: infinite;
  -webkit-animation-iteration-count: infinite;
  -moz-animation-iteration-count: infinite;
  animation-timing-function: linear;
  -o-animation-timing-function: linear;
  -ms-animation-timing-function: linear;
  -webkit-animation-timing-function: linear;
  -moz-animation-timing-function: linear;
}

@keyframes noTrespassingAnimationG {
  0% {
    margin-left: 0px;
  }

  100% {
    margin-left: -37px;
  }
}

@-o-keyframes noTrespassingAnimationG {
  0% {
    margin-left: 0px;
  }

  100% {
    margin-left: -37px;
  }
}

@-ms-keyframes noTrespassingAnimationG {
  0% {
    margin-left: 0px;
  }

  100% {
    margin-left: -37px;
  }
}

@-webkit-keyframes noTrespassingAnimationG {
  0% {
    margin-left: 0px;
  }

  100% {
    margin-left: -37px;
  }
}

@-moz-keyframes noTrespassingAnimationG {
  0% {
    margin-left: 0px;
  }

  100% {
    margin-left: -37px;
  }
}
.search {
  padding: 6px 15px 6px 30px;
  margin: 3px;
  background: url("../assets/search.png") no-repeat 7px 10px;
  border-radius: 20px;
  -moz-border-radius: 20px;
  -webkit-border-radius: 20px;
  border-style: groove;
}
.lighter {
  width: 90%;
  height: 50px;
  padding: 0 25px;
}
.searchContainer {
  float: left;
  margin: 5% auto auto 5%;
}
.edit-modal {
  width: 400px;
  height: 200px;
  background-color: green;
  margin: 0 auto;
  position: absolute;
  z-index: 1;
  top: calc(50% - 200px / 2);
  left: calc(50% - 400px / 2);
}

.edit-modal:before {
  background: rgba(189, 195, 199, 0.6);
  content: "";
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  z-index: -1;
}

.content {
  background: yellow;
  height: 100%;
  width: 100%;
}

/* .bg-modal-active {
  position: absolute;
  height: 100%;
  width: 100%;
  background: rgba(189, 195, 199, 0.6);
} */

/* .listContainer{
  z-index: -2;
} */
/* body.app {
  position: absolute;
  height: 100%;
  width: 100%;
  background: rgba(189, 195, 199, 0.6);  
} */
</style>
