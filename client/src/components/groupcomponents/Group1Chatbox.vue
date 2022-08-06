<template>
  <b-container>
    <div id="grouponecomponent" class>
      <b-row
        class="Clientzero"
        :class="{ Clientone: chat.name == $clientname.name }"
        align-h="start"
        v-for="(chat, inx) in chatHistory"
        :key="inx"
      >
        <b-col v-if="chat.name == $clientname.name" id="chat-width" class="no-gutters" cols="9">
          <div class="chat-align">
            <p class="received-chat row">
              <span class="chat-name">{{ chat.name }}</span>:
              <span class="chat-msg">{{ chat.msg }}</span>
              <span v-if="chat.dp === './assests/bg.jpg'" id="dp" class="profilePic"></span>
              <span
                v-else
                id="dp1"
                class="profilePic"
                :style="{ 'background-image': `url(${chat.dp})` }"
              ></span>
            </p>
          </div>
        </b-col>
        <b-col v-else id="chat-width" class="no-gutters" cols="9">
          <div class="chat-align">
            <p class="received-chat row">
              <span v-if="chat.dp === './assests/bg.jpg'" id="dp" class="profilePic"></span>
              <span
                v-else
                id="dp1"
                class="profilePic"
                :style="{ 'background-image': `url(${chat.dp})` }"
              ></span>

              <span class="chat-name">{{ chat.name }}</span>:
              <span class="chat-msg">{{ chat.msg }}</span>
            </p>
          </div>
        </b-col>
      </b-row>
    </div>
  </b-container>
</template>

<script>
export default {
  props: ["socket"],
  data() {
    return {
      chatHistory: [],
      currentUrl: "",
    };
  },
  created() {
    this.currentUrl = window.location.hash;
    this.currentUrl = this.currentUrl.substring(
      this.currentUrl.lastIndexOf("/") + 1
    );
  },
  mounted() {
    if (this.socket) {
      this.socket.on("group1", (data) => {
        var count = ++this.$notifications.group1;
        this.$notifications.group1 =count;
        var audio = new Audio(
          "http://soundbible.com/mp3/Elevator Ding-SoundBible.com-685385892.mp3"
        );
        audio.play();
        console.log("group data", JSON.stringify(data));
        if (data) {
          this.chatHistory.push({
            name: data.username,
            msg: data.msg.trim(),
            dp: data.dp,
            clientno: "zero",
          });
        }
      });
      this.socket.on("group1userjoined", (data) => {
        if (data) {
          var audio = new Audio(
            "http://soundbible.com/mp3/Elevator Ding-SoundBible.com-685385892.mp3"
          );
          audio.play();
          if (document.getElementById("grouponecomponent")) {
            if(document.getElementById("client-left")){
              document.getElementById("client-left").remove();
            }
            document.getElementById("grouponecomponent").innerHTML +=
              '<div id="client-left" class="row no-gutters justify-content-center">' +
              data.name +
              "   Joined !</div>";
          }
        }
      });
    }
  },
};
</script>

<style scoped>
.container {
  height: calc(100% - 100px);
}

.sender-name {
  margin: 0px;
}
.chat {
  background-color: lightgreen;
  padding: 10px;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
}
.received-chat {
  /*  background-color: #d39bbe;
  padding: 10px;
  border-top-right-radius: 10px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px; */
  margin: 1rem;
  padding: 0.5rem;
}
.chat-align {
  margin-right: -15px;
  margin-left: -15px;
  border-top-right-radius: 10px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
}
.Clientzero .chat-align {
  background-color: #d39bbe;
}
.Clientone .chat-align {
  background-color: #d39bd2;
}
.invisible {
  visibility: hidden;
}

.Clientone .received-chat {
  justify-content: end;
}
.Clientone {
  justify-content: flex-end !important;
}
</style>
