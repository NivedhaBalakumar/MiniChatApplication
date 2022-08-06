<template>
  <b-input-group class="mb-2 mr-sm-2 mb-sm-0">
    <b-form-input
      id="send-msg"
      placeholder="type here.."
      v-model.trim="newMsg"
      @keyup.enter="sendMessage"
    ></b-form-input>
    <b-button class="mb-2" size="md" variant="primary" @click="sendMessage" :disabled="!newMsg">Send</b-button>
  </b-input-group>
</template>

<script>
export default {
  props: ["socket"],
  data() {
    return {
      newMsg: "",
      currentUrl: "",
    };
  },
  methods: {
    sendMessage() {
      if (this.newMsg) {
          this.currentUrl = window.location.hash;
          this.currentUrl = this.currentUrl.substring(
          this.currentUrl.lastIndexOf("/") + 1
        );
        console.log(this.currentUrl, "---group newmsg");
        this.socket.emit("/groupchat", {
          username: "",
          msg: this.newMsg,
          dp: this.$displayPicture.url,
          current_user_name: this.$clientname.name,
          groupid: this.currentUrl,
        });
        this.newMsg = "";
      }
      var $cont = document.getElementById("chat-container-scroll");
      if ($cont !== null) {
        $cont.scrollTop = $cont.scrollHeight;
      }
    },
  },
};
</script>
