<template>
    <b-container id="groupList" class="col-6">
        <b-row class="groups">
            <router-link class="col-3" to="/allgroups/group1">Group1</router-link>
            <p class="col-3" v-if="$groups.group1 == 0" id="group1" @click="joinGroup($event)">Join Group</p>
            <p class="col-3" v-else id="group1" @click="LeaveGroup($event)">Leave Group</p>
        </b-row>
        <b-row class="groups">
            <router-link class="col-3" to="/allgroups/group2">Group2</router-link>
            <p class="col-3" v-if="$groups.group2 == 0" id="group2" @click="joinGroup($event)">Join Group</p>
            <p class="col-3" v-else id="group2" @click="LeaveGroup($event)">Leave Group</p>
        </b-row>
    </b-container>
</template>
<script>
export default {
    props: ["socket"],
    methods: {
        joinGroup(event) {
            var targetId = event.currentTarget.id;
            if (targetId == "group1") {
                this.$groups.group1 = 1;
            } else {
                this.$groups.group2 = 1;
            }
            console.log(targetId);
            this.socket.emit("joingroup", targetId);
        },
        LeaveGroup(event) {
            var targetId = event.currentTarget.id;
            if (targetId == "group1") {
                this.$groups.group1 = 0;
            } else {
                this.$groups.group2 = 0;
            }
            console.log(targetId);
            this.socket.emit("LeaveGroup", {
                room: targetId,
                name: this.$clientname.name,
            });
        },
    },
};
</script>
