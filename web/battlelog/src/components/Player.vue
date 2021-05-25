<template>
  <div class="column">
    <form v-on:submit.prevent="onSubmit">
      <label for="player1">Choose a faction: </label>
      <select name="player1" v-model="selectedFaction" @change="changeSubs()">
        <option v-for="faction in factions" :key="faction">
          {{ faction.name }}
        </option>
      </select>
      <br><br>
      <label>Choose a sub-faction: </label>
      <div v-for="(sub, index) in subs" :key="sub">
        <input v-model="selectedSubs[index]" type="checkbox" id="sub" :true-value="sub">
        <label for="sub">{{ sub }}</label><br>
      </div>
      <br><br>
      <label for="playerPhases">Choose a phase: </label>
      <select name="playerPhases" v-model="selectedPhase">
        <option v-for="phase in phases" :value="phase" :key="phase">
          {{ phase }}
        </option>
      </select>
      <br><br>
      <input type="submit" value="Submit">
    </form>
    <div v-for="strat in strats" :key="strat" style="position: relative;margin: 2px 8px 2px 8px">
      <div class="breakInsideAvoid" style="padding: 6px;max-width: 720px;">
        <table class="collapse" border="0" cellpadding="0" cellspacing="0">
          <tbody>
          <tr></tr>
          <tr>
            <td style="width: 24px"></td>
            <td bgcolor="#d3d3d3">
              <p class="stratName">{{ strat.name }}</p>
              <p class="stratFaction">{{ strat.type }}</p>
              <p class="stratLegend"> {{ strat.fluff }}</p>
              {{ strat.description }}
              <p class="stratName">{{ strat.cp_cost }}</p>
            </td>
            <td style="width: 24px"></td>
          </tr>
          <tr></tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios';
import { config } from '../config.js';

export default {
  data() {
    return {
      strats: null,
      subs: null,
      factions: null,
      phases: [],
      selectedPhase: '',
      selectedFaction: '',
      selectedSubs: [],
      url: config.url
    };
  },
  methods: {
    changeSubs () {
      this.selectedSubs = []
      for (var i = 0; i < this.factions.length; i++) {
        if (this.factions[i].name == this.selectedFaction) {
          console.log(this.factions[i].subs)
          this.subs = this.factions[i].subs
        }
      }
    },
    async onSubmit() {
      const params = new URLSearchParams();
      if (this.selectedPhase != '') {
        params.append('phase', this.selectedPhase)
      }
      if (this.selectedSubs.length > 0) {
        for (const s in this.selectedSubs) {
          console.log(this.selectedSubs[s])
          params.append('sub', this.selectedSubs[s])
        }
      }
      const request = {
        params: params
      }
      const responseData = await axios.get(this.url + '/battlelog/stratagems/' + this.selectedFaction, request)
      this.strats = responseData.data.stratagems
    }
  },
  created() {
    axios.get(this.url + '/battlelog/phases').then(response => {this.phases = response.data})
    axios.get(this.url + '/battlelog/stratagems').then(response => {this.factions = response.data})
  }
}
</script>
<style>
.column {
  float: left;
  width: 50%;
}
</style>