<template>
  <div class="home">
    <div v-if="loading == false && loaded == true">
      <p>
        https://github.com/
        <input
          type="text"
          name="username"
          class="usernameTxt"
          placeholder=" username"
          v-model="username"
        />
      </p>

      <div class="wrap">
        <button class="button" v-on:click="submit(username)">
          get repositories
          <span class="arrow-right"></span>
        </button>
      </div>
    </div>

    <div class="loading-div" v-if="loading == true && loaded == false">
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
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  name: "home",
  components: {},
  data() {
    return {
      response: {
        id: null,
        repositories: null
      },
      username: "",
      APIURL: "http://localhost:8090",
      loading: false,
      loaded: true
    };
  },
  methods: {
    ...mapMutations(["setNewUser"]),
    submit: function() {
      this.$http
        .post(this.APIURL + "/create", {
          username: this.username
        })
        .then(resp => {
          this.loading = true;
          this.loaded = false;
          this.response = {
            id: resp.data[0]._id,
            repositories: resp.data[0].repositories
          };
          this.setNewUser(this.response);
          // the test requested to set on localstorage but I doubt it is necessary ?
          localStorage.setItem("user", JSON.stringify(this.response));
        })
        .catch(err => {
          console.log(err);
        })
        .finally(
          (this.loaded = true),
          (this.loading = false),
          this.$router.push("list")
        );
    }
  }
};
</script>

<style>
.usernameTxt {
  border: solid 0.15em;
  height: 20px;
}
.arrow-right {
  display: inline-block;
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 5px 0 5px 5px;
  border-color: transparent transparent transparent #000000;
  vertical-align: bottom;
}
.wrap {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.button {
  width: 140px;
  height: 25px;
  font-family: "Comic Sans MS", cursive, sans-serif;
  font-size: 8px;
  text-transform: uppercase;
  letter-spacing: 1.5px;
  font-weight: bolder;
  color: #000;
  background-color: #fff;
  border: solid 0.2em;
  border-radius: 0px;
  box-shadow: 4px 4px 0px rgba(0, 0, 0, 1);
  cursor: pointer;
  outline: none;
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
</style>
