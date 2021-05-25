<template>
  <div>
    <form v-on:submit.prevent="onSubmitRolls()">
      <label for="sides">How many sides:</label>
      <input type="text" id="sides" name="sides" v-model="modelSides"><br><br>
      <label for="number">Number of rolls:</label>
      <input type="text" id="number" name="number" v-model="modelNumber"><br><br>
      <input type="submit" value="Submit">
    </form>
  </div>
  <br><br><br>
  <div v-for="res in results.slice().reverse()" v-bind:key="res">
    <table>
      <tr>
        {{ res.rawRolls }}
      </tr>
      <tr>
        {{ res.orderedRolls }}
      </tr>
    </table>
    <br>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      modelNumber: 1,
      modelSides: 6,
      results: []
    };
  },
  methods: {
    async onSubmitRolls() {
      const params = new URLSearchParams();
      params.append('number', this.modelNumber);
      params.append('sides', this.modelSides);
      const request = {
        params: params
      }
      const data = await axios.get('http://localhost:3000/battlelog/roll', request)
      if (data.status != 200) {
        return
      }
      let element = {rawRolls: [], orderedRolls: ''};
      element.rawRolls = data.data.results
      let ordered = ''
      for (let i = 0; i < data.data.sorted.length; i++) {
        const j = i + 1;
        ordered = ordered + '[' + j + ': ' + data.data.sorted[i] + ']  '
      }
      element.orderedRolls = ordered
      this.results.push(element)
    }
  }
}
</script>

<style scoped>

</style>