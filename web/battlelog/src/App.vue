<template>
  <div id="app">
    <div>
      <button type="button" v-on:click="onClickRoller()">Roller</button>
      <button type="button" v-on:click="onClickStrats()">Stratagems</button>
    </div>
    <Roller v-if="roller"></Roller>
    <Player v-if="strats"></Player>
    <Player v-if="strats"></Player>
  </div>
</template>

<script>
import Player from './components/Player';
import Roller from './components/Roller'

export default {
  name: 'App',
  components: {
    Player,
    Roller
  },
  data() {
    return {
      roller: true,
      strats: false,
    };
  },
  methods: {
    onClickRoller() {
      this.roller = true;
      this.strats = false;
    },
    onClickStrats() {
      this.strats = true;
      this.roller = false;
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
