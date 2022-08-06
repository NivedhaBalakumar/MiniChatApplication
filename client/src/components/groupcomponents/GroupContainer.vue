<template>
  <div class="container chat-container">
    <div id="chat-container-scroll" class="chatbox">
      <group1-chatbox id="1" v-if="currentUrl == 'group1'" :socket="socket"></group1-chatbox>
      <group2-chatbox id="2" v-if="currentUrl == 'group2'" :socket="socket"></group2-chatbox>
    </div>
    <group-message id="message-box" class="message-box" :socket="socket"></group-message>
  </div>
</template>
<script>
import Group1Chatbox from "./Group1Chatbox.vue";
import Group2Chatbox from "./Group2Chatbox.vue";
import GroupMessage from "./GroupMessage.vue";
export default {
  name: "chat",
  components: {
    Group1Chatbox,
    Group2Chatbox,
    GroupMessage,
  },
  props: ["socket"],
  data() {
    return {
      // socket: null,
      // serverUrl: process.env.VUE_APP_SOCKET_URL || "ws://localhost:8080",
      userMessage: "",
      currentUrl:""
    };
  },
  created() {
    
    this.currentUrl = window.location.hash;
    this.currentUrl = this.currentUrl.substring(
    this.currentUrl.lastIndexOf("/") + 1
    );
  },
  methods: {
    // connectToWebsocket() {
    //   this.socket = io(this.serverUrl, {
    //     transports: ["websocket"],
    //   });
    // },
  },
  mounted() {
    if (this.socket) {
      this.socket.on("clientLeft", function (data) {
        if (data) {
          var audio = new Audio(
            "https://soundbible.com/mp3/Air%20Plane%20Ding-SoundBible.com-496729130.mp3"
          );
          audio.play();
          if (document.getElementById("client-left")) {
            document.getElementById("client-left").remove();
          }
          document.getElementById("chat-box").innerHTML +=
            '<div id="client-left" class="row no-gutters justify-content-center"></div>';
        }
      });
    }
  },
};
</script>
