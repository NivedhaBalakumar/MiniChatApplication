<template>
  <div class="container chat-container">
    <user-name :socket="socket"></user-name>
    <send-email id="send-email" class="invisible" :socket="socket"></send-email>
    <div id="chat-container-scroll" class="chatbox">
      <chat-chatbox :socket="socket"></chat-chatbox>
    </div>
    <chat-message-box id="message-box" class="invisible message-box" :socket="socket"></chat-message-box>
  </div>
</template>
<script>
import UserName from "./UserName.vue";
import SendEmail from "./SendEmail.vue";
import ChatChatbox from "./ChatChatbox.vue";
import ChatMessageBox from "./ChatMessagebox.vue";
export default {
  name: "chat",
  components: {
    UserName,
    SendEmail,
    ChatChatbox,
    ChatMessageBox,
  },
  props: ["socket"],
  data() {
    return {
      // socket: null,
      // serverUrl: process.env.VUE_APP_SOCKET_URL || "ws://localhost:8080",
      userMessage: "",
    };
  },
  created() {
    this.$notifications.inbox = 0
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
            '<div id="client-left" class="row no-gutters justify-content-center">User left please wait </div>';
        }
      });
      this.socket.on("userExists", (data) => {
        if (data.msg) {
          // welcome-msg
          if (document.getElementById("welcome-msg")) {
            document.getElementById("welcome-msg").remove();
          }
          var audio = new Audio(
            "https://soundbible.com/mp3/Air%20Plane%20Ding-SoundBible.com-496729130.mp3"
          );
          audio.play();
          if (document.getElementById("user-exists")) {
            document.getElementById("user-exists").remove();
          }
          document.getElementById("chat-box").innerHTML +=
            '<div id="user-exists" class="row no-gutters justify-content-center">Username taken already , Try with new one</div>';
        }
      });
      this.socket.on("userSet", (data) => {
        if (data) {
          // console.log(document.getElementById("set-user"))
          document.getElementById("message-box").classList.remove("invisible");
          document.getElementById("set-user").classList.add("invisible");
          document.getElementById("set-user").remove();
          if (document.getElementById("user-exists")) {
            document.getElementById("user-exists").remove();
          }
          this.$clientname.name = data.username;
          this.$clientname.set = 1;
          this.socket.emit("check-prev-msg");
        }
      });
    }
  },
};
</script>
