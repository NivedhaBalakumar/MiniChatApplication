<template>
  <b-navbar toggleable type="dark">
    <br />
    <b-navbar-brand href="#">NvedaKumar's Mini Chat App</b-navbar-brand>
  </b-navbar>
</template>
<script>
export default {
  props: ["socket"],
  data() {
    return {
      dp: this.$displayPicture.url,
    };
  }
};
</script>