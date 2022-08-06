<template>
    <b-input-group id="emailCont" class="mb-2 mr-sm-2 mb-sm-0 col-8">
      <b-form-input placeholder="Enter your email-id" v-model="emailId" @keyup.enter="sendemail"></b-form-input>
      <b-button
        size="md"
        variant="primary"
        class="mb-2"
        v-on:click="sendemail"
         :disabled="disable"
      >Export your chat to email</b-button>
      <span v-if="msg.email" class="emailerror">{{ msg.email }}</span>
    </b-input-group>
</template>

<script>
export default {
  props: ["socket"],
  data() {
    return {
      emailId: "",
      msg: [],
      disable:true,
    };
  },
  watch: {
    emailId(value) {
      // binding this to the data value in the email input
      this.emailId = value;
      this.validateEmail(value);
    }
  },

  methods: {
    validateEmail(value) {
     // eslint-disable-next-line
     if ((/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/).test(value))
     {
        this.msg['email'] = '';
        this.disable = false;
      } else {
        this.msg['email'] = 'Invalid Email Address';
        this.disable = true;
      }
    },
    sendemail() {
      if (this.emailId) {
        console.log("setemailId insides");
        console.log(this.socket);
        this.socket.emit("send_email", this.emailId);
      }
    },
  },
};
</script>
