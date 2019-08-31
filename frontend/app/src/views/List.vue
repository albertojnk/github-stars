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
          <input
            type="search"
            class="search"
            placeholder="search by tag"
            v-model="searchValue"
            @keypress="searchRepository"
          />
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
            <td class="repository-name">
              <a v-bind:href="repository.html_url" target="_blank">
                {{ repository.name }}
              </a>
            </td>
            <td class="repository-description">{{ repository.description }}</td>
            <td class="repository-language">{{ repository.language }}</td>
            <td class="repository-tags">
              <template v-for="(tag, idx) in repository.tags">
                {{ tag | tagNormalize(idx, repository.tags.length) }}
              </template>
            </td>
            <td class="repository-edit">
              <a
                href="javascript:void(0)"
                class="edit-tags"
                @click="show(repository)"
              >
                edit
              </a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="modal" v-bind:class="{ 'edit-modal': modal }">
      <div class="content" v-click-outside="outsideClose">
        <div class="text-input-wrapper">
          <p class="content-p">edit tags for {{ currentRepo.name }}</p>
          <input
            type="text"
            name="tags"
            id="tags"
            v-bind:placeholder="currentRepo.tag_suggester"
            v-model="currentTags"
            @keypress="sendMonitor"
          />
        </div>
        <div class="btnWrapper">
          <button class="btn" id="btn-submit" type="submit" @click="Save">
            save
          </button>
          <button class="btn" id="btn-close" type="button" @click="Close">
            close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapMutations } from "vuex";

export default {
  name: "list",
  components: {},
  data() {
    return {
      APIURL: "http://localhost:8090",
      modal: false,
      currentRepo: null,
      currentTags: null,
      searchValue: null,
      response: null
    };
  },
  mounted() {
    if (this.id == undefined && this.loading != true) {
      this.setLoading({ loading: true, loaded: false });
      let user = JSON.parse(localStorage.getItem("user"));

      this.response = {
        id: user.id,
        repositories: user.repositories
      };

      localStorage.setItem("user", JSON.stringify(this.response));
      this.setNewUser(this.response);
      this.setLoaded({ loading: false, loaded: true });
    }
  },
  computed: {
    ...mapState(["id", "repositories", "loading", "loaded"])
  },
  methods: {
    ...mapMutations(["setNewUser", "setLoading", "setLoaded"]),
    searchRepository(event) {
      if (event.key == "Enter") {
        this.$http
          .post(this.APIURL + "/search", {
            id: this.id,
            search: this.searchValue
          })
          .then(resp => {
            this.setNewUser({
              id: this.id,
              repositories: resp.data
            });
            this.response = {
              id: this.id,
              repositories: resp.data
            };
            localStorage.setItem("user", JSON.stringify(this.response));
          });
      }
    },
    Save() {
      this.modal = false;
      this.setLoading({ loading: true, loaded: false });
      if (this.currentTags == null || this.currentTags.length == 0) {
        this.currentTags = this.currentRepo.tag_suggester
      }
      this.$http
        .patch(this.APIURL + "/update", {
          username: this.id,
          repo_id: this.currentRepo.id,
          tags: this.transformInArray(this.currentTags)
        })
        .then(resp => {
          this.response = {
            id: this.id,
            repositories: resp.data
          };

          this.modal = false;
          this.setNewUser(this.response);
          localStorage.setItem("user", JSON.stringify(this.response));
          this.setLoaded({ loading: false, loaded: true });
        })
        .catch(this.$router.push("/"))
        .finally(this.$router.push("/repositories"));
    },
    sendMonitor(event) {
      if (event.key == "Enter") {
        this.Save();
      }
    },
    transformInArray(value) {
      var splited = value.split(/[ ,]+/);
      var result = [];
      splited.forEach(tag => {
        if (tag.trim() !== "") {
          result.push(tag.trim());
        }
      });
      return result;
    },
    outsideClose(event) {
      if (event.target.className !== "edit-tags" || event.key === "Escape") {
        this.modal = false;
      }
    },
    Close() {
      this.modal = false;
    },
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
  padding-left: 10px;
}
.repository-description {
  border-right: solid 1px #000000;
  padding-left: 10px;
}
.repository-language {
  border-right: solid 1px #000000;
  padding: 0 5px;
  padding-left: 10px;
}
.repository-tags {
  border-right: solid 1px #000000;
  padding: 0 5px;
  padding-left: 10px;
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

.tableBody {
  text-align: left;
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
  height: 120px;
  margin: 0 auto;
  position: fixed;
  z-index: 1;
  top: calc(50% - 120px / 2);
  left: calc(50% - 400px / 2);
}

.edit-modal:before {
  background: rgba(124, 127, 129, 0.7);
  content: "";
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  z-index: -1;
}

.content {
  height: 100%;
  width: 100%;
  font-family: "Comic Sans MS", cursive, sans-serif;
  font-size: 15px;
  font-weight: 500;
  background: white;
  border-radius: 10px;
}

.text-input-wrapper {
  height: 65%;
  width: 100%;
}

.content-p {
  width: 100%;
  margin: auto auto 1% 7%;
  text-align: left;
  padding-top: 10px;
}

#tags {
  width: 90%;
  border: solid 2px #000000;
}

.btnWrapper {
  height: 35%;
  width: 100%;
  align-items: center;
  justify-content: center;
}
.btn {
  width: 80px;
  height: 25px;
  font-family: "Comic Sans MS", cursive, sans-serif;
  font-size: 14px;
  text-transform: lowercase;
  font-weight: bolder;
  color: #000;
  background-color: #fff;
  border: 2px solid #000000 !important;
  border-radius: 0px !important;
  box-shadow: 3px 3px 0px rgba(0, 0, 0, 1);
  cursor: pointer;
  outline: none;
  line-height: 0.5 !important;
}

[type="submit"]:not(:disabled):hover {
  cursor: pointer;
  background: lightblue;
}

[type="button"]:not(:disabled):hover {
  cursor: pointer;
  background: lightgray;
}
</style>
