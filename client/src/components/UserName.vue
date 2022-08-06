<template>
  <b-input-group id="set-user"  class="mb-2 mr-sm-2 mb-sm-0 col-4">
    <b-form-input id="input-user"
      placeholder="Enter your name" maxlength="10"
      v-model="userName"
      @keyup.enter="setUsername"
    ></b-form-input>
    <b-button
      size="md"
      variant="primary"
      class="mb-2"
      @click="setUsername"
      :disabled="!userName"
      >Set</b-button
    >
  </b-input-group>
</template>

<script>
export default {
  props: ["socket"],
  data() {
    return {
      userName: "",
    };
  },
  methods: {
    setUsername() {
      if (this.userName) {
        // this.$clientname.name = this.userName;
        console.log("set user name inside");
        // console.log(this.socket);
        this.socket.emit("setUsername", this.userName);
      }
    },
  },
};
</script>
