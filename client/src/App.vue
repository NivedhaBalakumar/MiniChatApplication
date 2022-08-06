<template>
  <div id="app" class="chatApp">
    <div class="row">
      <chat-navbar
        class="col-sm-10 col-md-10 col-lg-10"
        :socket="socket"
      ></chat-navbar>
      <div id="welcomeMessage" class="col-sm-2 col-md-2 col-lg-2">
        <span id="clientName">{{ $clientname.name }} !</span>
        <div id="dp-pic" @click="editProfile = true"></div>
        <span id="dp" @click="editProfile = true"></span>
      </div>
    </div>
    <div class="row">
      <div id="panel" class="col-sm-2 col-md-2 col-lg-1">
        <sidebar-menu
          :collapsed="collapsed"
          :menu="menu"
          :show-one-child="true"
          @collapse="onCollapse"
          @itemClick="onItemClick"
        ></sidebar-menu>
      </div>
      <div class="col-sm-10 col-md-10 col-lg-11">
        <keep-alive>
          <router-view :socket="socket"></router-view>
        </keep-alive>
      </div>
    </div>
    <UpdateProfile v-show="editProfile" @close-modal="editProfile = false" />
  </div>
</template>

<script>
import io from "socket.io-client";
import ChatNavbar from "./components/ChatNavbar.vue";
import UpdateProfile from "./components/UpdateProfile.vue";
import "../node_modules/@syncfusion/ej2-vue-notifications/styles/material.css";
export default {
  name: "App",
  components: {
    ChatNavbar,
    UpdateProfile,
  },
  data() {
    return {
      socket: null,
      serverUrl: process.env.VUE_APP_SOCKET_URL || "ws://localhost:8080",
      collapsed: true,
      editProfile: false,
      menu: [
        {
          title: "Home",
          icon: "fa-home fa",
          href: "/",
          active: true,
        },
        {
          title: "Available Groups",
          icon: "fa fa-users",
          child: [
            {
              href: "/group1",
              title: "Group1",
            },
            {
              href: "/group2",
              title: "Group2",
            },
          ],
        },
      ],
    };
  },
  created() {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
      this.socket = io(this.serverUrl, {
        transports: ["websocket"],
      });
    },
    onItemClick(e, i) {
      console.log("onItemClick " + e + "--" + i);
    },
    onCollapse(c) {
      console.log("onCollapse");
      this.collapsed = c;
    },
  },
  props: {
    relative: {
      type: Boolean,
      default: true,
    },
  },
};
window.onclick = function (event) {
  if (!event.target.matches(".dropbtn")) {
    var dropdowns = document.getElementsByClassName("dropdown-content");
    var i;
    for (i = 0; i < dropdowns.length; i++) {
      var openDropdown = dropdowns[i];
      if (openDropdown.classList.contains("show")) {
        openDropdown.classList.remove("show");
      }
    }
  }
};
</script>

<style>
html,
body,
#app {
  height: 100%;
  background-image: url("./assests/chat4.jpg");
  /*   background-repeat: no-repeat;
    background-size: 100%; */
}
.card {
  height: 100%;
  margin: 15px;
}
.chatbox {
  padding: 10px;
  height: calc(100% - 35px);
  /*overflow-y: auto; */
  overflow-y: auto;
}
nav.navbar.navbar-dark {
  /* background-color: #cf71af; */
  background-color: #f5f5f5 !important;
  opacity: 0.96;
}
.navbar-dark .navbar-brand {
  color: #532c5e !important;
}
.navbar {
  justify-content: center !important;
}
.v-sidebar-menu .vsm--header,
.v-sidebar-menu .vsm--link,
.v-sidebar-menu .vsm--toggle-btn {
  /* color: rgba(255,255,255,0.7); */
  color: #532c5e !important;
}
button {
  background-color: #532c5e !important;
}
.fade-animation-enter-active {
  transition: transform 0.15s, opacity 0.45s !important;
}
.fade-animation-enter,
.fade-animation-leave-to {
  transform: translateX(-60%) !important;
  opacity: 0 !important;
}
.v-sidebar-menu.vsm_expanded {
  max-width: unset !important;
  width: unset !important;
}
.v-sidebar-menu.vsm_collapsed {
  width: 5% !important;
}
.v-sidebar-menu,
.v-sidebar-menu .vsm--toggle-btn {
  /* background-color: #cf71af !important; */
  background-color: #f5f5f5 !important;
  opacity: 0.7;
}
.v-sidebar-menu .vsm--toggle-btn:after {
  content: "\f0c9" !important;
}

.v-sidebar-menu .vsm--link_level-1 .vsm--icon {
  background-color: transparent !important;
}

.v-sidebar-menu.vsm_collapsed .vsm--link_level-1.vsm--link_hover,
.v-sidebar-menu.vsm_collapsed .vsm--link_level-1:hover {
  background-color: #532c5e !important;
}

.v-sidebar-menu .vsm--link_level-1.vsm--link_exact-active,
.v-sidebar-menu .vsm--link_level-1.vsm--link_active {
  -webkit-box-shadow: 3px 0px 0px 0px #532c5e inset !important;
  box-shadow: 3px 0px 0px 0px #532c5e inset !important;
}
.v-sidebar-menu .vsm--mobile-bg {
  background-color: #532c5e !important;
  position: absolute;
  top: unset !important;
  left: unset !important;
  padding-left: unset !important;
  z-index: unset !important;
  width: auto !important;
  max-width: auto !important;
}

.v-sidebar-menu .vsm--item {
  position: relative;
  display: block;
  width: 100%;
  white-space: nowrap;
}
div#client-left {
  margin-bottom: 10px;
}
div#send-email {
  margin-bottom: 5px;
  margin-top: 15px;
}
div#welcome-msg {
  font-style: oblique;
  font-weight: 600;
}
div#set-user {
  margin-top: 10px !important;
  float: right;
}
#input-user {
  border-top-right-radius: 0;
  border-width: medium;
  border-bottom-right-radius: 0;
  border-color: #532c5e;
}
div#chat-box,
#grouponecomponent,
#grouptwocomponent {
  margin-bottom: 20px;
}
button#save {
  color: white;
}
#dp {
  width: 30px;
  height: 30px;
  display: inline-block;
  cursor: pointer;
  border-radius: 50%;
  background-size: contain;
  background-image: url("./assests/dp.jpg");
}
#dp1 {
  width: 30px;
  height: 30px;
  display: inline-block;
  cursor: pointer;
  border-radius: 50%;
  background-size: contain;
}
#nav-bar-pic,
#welcomeMessage {
  background-color: #f5f5f5 !important;
  opacity: 0.96;
  display: flex;
  justify-content: right;
  /* padding: 0.5rem  2rem 0.5rem 0rem; */
  padding: unset;
  padding-top: 0.5rem;
  padding-right: 1rem;
}
svg.svg-inline--fa.fa-arrow-circle-left.fa-w-16 {
  margin-bottom: 5px;
  color: #532c5e;
}
/*chat pic*/
.profilePic {
  margin-right: 2%;
}
.chat-name {
  margin-right: 2%;
}
.chat-msg {
  margin-left: 2%;
  margin-right: 2%;
}
.container.chat-container {
  height: 400px;
}
.v-sidebar-menu .vsm--dropdown .vsm--list {
  background-color: #98fafc !important;
}
.v-sidebar-menu.vsm_expanded .vsm--item_open .vsm--link_level-1 {
  background-color: #e5ccff !important;
}
/** 
List groups css
*/
div#groupList {
  margin-top: 2%;
}
.row.gorups {
  padding-top: 5%;
}
.emailerror {
  margin-left: 2%;
  color: red;
}
button#group2,
button#group1 {
  color: white;
}
/** Notifiaction **/

span.notification-badges {
  margin-right: 20%;
}
button .e-badge {
    line-height: unset;
    position: absolute;
    top: unset !important; 
}
.e-badge {
     background: unset; 
    border-color: unset; 
    border-radius:50%;
    box-shadow: unset;
   box-sizing: unset; 
     color: unset; 
     display: unset;
    font-family:unset;
    font-size: 8px;
    font-weight: 500; 
   line-height: 1; 
    overflow: hidden;
    padding: 0.25em 0.4em 0.25em 0.4em;
    text-align: center;
    text-decoration: none;
    text-indent: 0;
    vertical-align: middle;
}
/**  bell icon */
button#n-dropdown__BV_toggle_ {
  background: unset !important;
  color: #532c5e !important;
  padding: unset;
}
/**  username */
span#clientName {
    margin-right: 2%;
}
</style>
