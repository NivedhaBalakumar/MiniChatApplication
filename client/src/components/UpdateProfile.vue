<template>
  <transition name="modal-fade">
    <div class="modal-overlay" @click="$emit('close-modal')">
      <div
        class="modalUpdate"
        @click.stop
      >
        <header class="modal-header">
          <slot name="header">
            Upload Profile Picture
          </slot>
        </header>

        <section class="modal-body">
          <slot name="body">
            <input
              class="chooseFile"
              ref="fileInput"
              type="file"
              @input="pickFile"
            />
            <div
              class="imagePreviewWrapper"
              :style="{ 'background-image': `url(${previewImage})` }"
              @click="selectImage"
            ></div>
          </slot>
        </section>
        <button id="save" v-if="value1" @click="save">Save</button>

        <!-- <SavedModal v-show="showModal" @close-modal="showModal = false" /> -->
      </div>
      <div class="closeSave" @click="$emit('close-modal')">
        <img class="close-img" src="../assests/close-icon.svg" alt="" />
      </div>
    </div>
  </transition>
</template>

<script>
// import SavedModal from "./SavedModal.vue";
export default {
  components: {
    // SavedModal,
  },
  data() {
    return {
      previewImage: null,
      dp: this.$displayPicture.url,
      value1: false,
      showModal: false,
    };
  },
  methods: {
    //image upload and preview methods
    selectImage() {
      this.$refs.fileInput.click();
    },
    pickFile() {
      let input = this.$refs.fileInput;
      console.log(input);
      let file = input.files;
      console.log(file);
      if (file && file[0]) {
        let reader = new FileReader();
        reader.onload = (e) => {
          this.previewImage = e.target.result;
          this.$displayPicture.url = e.target.result;
          this.value1 = true;
        };
        reader.readAsDataURL(file[0]);
      }
    },
    save() {
      this.showModal = true;
      this.value1 = false;
      if (document.getElementById("dp")) {
        document.getElementById("dp").remove();
      }
      if (document.getElementById("dp1")) {
        document.getElementById("dp1").remove();
      }
      document.getElementById("dp-pic").innerHTML +=
        '<span id="dp1" style="background-image: url(' +
        this.$displayPicture.url +
        ')"></span>';
       this.$emit('close-modal')
    },
  },
};
</script>
<style scoped>
.imagePreviewWrapper {
  background-repeat: no-repeat;
  width: 150px;
  height: 150px;
  display: block;
  cursor: pointer;
  margin: auto auto auto;
  background-size: contain;
  background-position: center center;
}
.chooseFile {
  margin-bottom: 5%;
}

.modal-overlay {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  background-color: #000000da;
}
.modalUpdate {
  text-align: center;
  background-color: white;
  height: auto;
  width: 400px;
  margin-top: 5%;
  margin-bottom: 5%;
  padding: 10px 0;
  /* border-radius: 20px; */
  display: block;
  overflow-y: initial !important;
}
.modal-body {
  height: 250px;
  overflow-y: auto;
}
.closeSave {
  margin: 4% 0 0 16px;
  cursor: pointer;
}
.close-img {
  width: 25px;
}
.check {
  width: 150px;
}
h6 {
  font-weight: 500;
  font-size: 28px;
  margin: 20px 0;
}
p {
  /* font-weight: 500; */
  font-size: 16px;
  margin: 20px 0;
}
button {
  background-color: #ac003e;
  width: 150px;
  height: 40px;
  color: white;
  font-size: 14px;
  border-radius: 16px;
  margin-top: 50px;
}
.modal-fade-enter,
.modal-fade-leave-to {
  opacity: 0;
}
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.5s ease;
}
.modal-header {
  display: flex;
  align-items: flex-start;
  justify-content: center;
  border: unset;
}
</style>
