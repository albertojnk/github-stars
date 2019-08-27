<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">
        <header>
          <p>githubstarts</p>
        </header>
      </router-link>
      <!-- <router-link to="/about">About</router-link>  -->
    </div>
    <router-view />
  </div>
</template>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #000000;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #000000;
}

header p {
  font-family: "Comic Sans MS", cursive, sans-serif;
  text-align: left;
  font-weight: bolder;
  font-size: larger;
}
</style>
