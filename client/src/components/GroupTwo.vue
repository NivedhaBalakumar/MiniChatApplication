<template>
    <div id="group-1" class="container">
        <div v-if="$clientname.set == 1">
            <div v-if="$groups.group2 == 0"  class="col-8" style="margin-top: 10%;margin-left: 35%;">
                <button class="col-3" id="group2" @click="joinGroup($event)">Join Group</button>
            </div>
            <div v-else>
                <button class="col-3" id="group2" @click="LeaveGroup($event)">Leave Group</button>
                <group-container :socket="socket"></group-container>
            </div>
        </div>
        <div v-else class="col-8" style="margin-top: 10%;margin-left: 30%;">
            <p>
                Please 
                <router-link to="/">Go back</router-link> and Set your userName to Proceed !
            </p>
        </div>
    </div>
</template>
<script>
import GroupContainer from "./groupcomponents/GroupContainer.vue";
export default {
    props: ["socket"],
    components: {
        GroupContainer,
    },
    created() {
        this.$notifications.group2 = 0
    },
    methods: {
        joinGroup(event) {
            var targetId = event.currentTarget.id;
            this.$groups.group2 = 1;
            console.log(targetId);
            this.socket.emit("joingroup", {
                room: targetId,
                name: this.$clientname.name,
            });
        },
        LeaveGroup(event) {
            var targetId = event.currentTarget.id;
            this.$groups.group2 = 0;
            console.log(targetId);
            this.socket.emit("LeaveGroup", {
                room: targetId,
                name: this.$clientname.name,
            });
        },
    },
    mounted() {
    if (this.socket) {
      this.socket.on("group2userleft", function (data) {
        if (data) {
          var audio = new Audio(
            "https://soundbible.com/mp3/Air%20Plane%20Ding-SoundBible.com-496729130.mp3"
          );
          audio.play();
          if (document.getElementById("client-left")) {
            document.getElementById("client-left").remove();
          }
          document.getElementById("grouptwocomponent").innerHTML +=
            '<div id="client-left" class="row no-gutters justify-content-center">' + data.name + '  left</div>';
        }
      });
    }
  },
};
</script>
