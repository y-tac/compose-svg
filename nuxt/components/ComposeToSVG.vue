<template>
  <div>
    <textarea id='compose' v-model='input' debounce='100'></textarea>
    <div id='svg' v-html='input | marked'></div>
  </div>
</template>

<style>
</style>
<script>
import marked from 'marked';

export default {
  data () {
    input: '# Hello, Vue.js\n\n- hoge\n- foo\n- bar'
  },
  filters: {
     marked: marked,
  },

  created () {
  },
  mounted: function () {
  }

}
</script>
