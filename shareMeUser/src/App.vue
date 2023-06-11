

<template>
    <Base/>
    <router-view></router-view>
</template>

<script>
import Base from "./components/Base.vue";

export default {
  components: {Base},
  methods: {
    redirect() {
      this.$router.push({ name: '/' }) // redirect to Home
    },
    back() {
      this.$router.go(-1) // go(negative) goes to history
    },
    forward() {
      this.$router.go(1)
    }
  }
}
</script>

<style>

</style>
