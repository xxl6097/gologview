<template>
  <div>
    <LogView></LogView>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import LogView from "@/views/LogView.vue"; // @ is an alias to /src

export default defineComponent({
  name: "HomeView",
  components: {
    LogView,
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
