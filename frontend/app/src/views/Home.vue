<template>
  <div class="home">
    <div class="container">
      <p>
        https://github.com/
        <input
          type="text"
          name="username"
          class="usernameTxt"
          placeholder=" username"
          v-model="username"
          @keypress="sendMonitor"
        />
      </p>

      <div class="wrap">
        <button class="button" @click="submit()">
          get repositories
          <span class="arrow-right"></span>
        </button>
      </div>
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
      APIURL: "http://localhost:8090"
    };
  },
  methods: {
    ...mapMutations(["setNewUser", "setLoading", "setLoaded"]),
    sendMonitor(event) {
      if (event.key == "Enter") {
        this.submit();
      }
    },
    submit() {
      this.setLoading({ loading: true, loaded: false });
      this.$http
        .post(this.APIURL + "/create", {
          username: this.username
        })
        .then(resp => {
          this.response = {
            id: resp.data._id,
            repositories: resp.data.repositories
          };
          localStorage.setItem("user", JSON.stringify(this.response));
          this.setNewUser(this.response);
          this.setLoaded({ loading: false, loaded: true });
        })
        .catch(err => {
          console.log(err);
          this.$router.push("/");
        })
        .finally(this.$router.push("list"));
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
.container {
  height: 100%;
  margin: 10% auto auto auto;
}
</style>
