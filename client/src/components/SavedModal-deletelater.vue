<template>
  <transition name="modal-fade">
    <div class="modal-overlay" @click="$emit('close-modal')">
      <div class="modalSave" @click.stop>
        <img class="check" src="../assests/check-icon.png" alt="" />
        <h6>Saved!</h6>
      </div>
      <div class="closeSave" @click="$emit('close-modal')">
        <img class="close-img" src="../assests/close-icon.svg" alt="" />
      </div>
    </div>
  </transition>
</template>

<script>
export default {}
</script>

<style scoped>
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
.modalSave {
text-align: center;
    background-color: white;
    height: 250px;
    width: 300px;
    margin-top: 5%;
    padding: 10px 0;
    /* border-radius: 20px; */
    display: block;
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

</style>